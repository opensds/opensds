// Copyright 2018 The OpenSDS Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package multicloud

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/opensds/opensds/contrib/backup/multicloud/credentials/keystonecredentials"
	"github.com/opensds/opensds/contrib/backup/multicloud/signer"

	"github.com/astaxie/beego/httplib"
	log "github.com/golang/glog"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	creds "github.com/gophercloud/gophercloud/openstack/identity/v3/credentials"
	"github.com/gophercloud/gophercloud/openstack/identity/v3/tokens"
	"github.com/opensds/opensds/contrib/backup/multicloud/credentials"
	"github.com/opensds/opensds/pkg/utils/constants"
	"github.com/opensds/opensds/pkg/utils/pwd"
)

const (
	DefaultTenantId      = "adminTenantId"
	DefaultTimeout       = 60 // in Seconds
	DefaultUploadTimeout = 30 // in Seconds
	ApiVersion           = "v1"
)

type Client struct {
	endpoint      string
	tenantId      string
	version       string
	baseURL       string
	userId        string
	Identity      *gophercloud.ServiceClient
	auth          *AuthOptions
	token         *tokens.Token
	timeout       time.Duration
	uploadTimeout time.Duration
}

type Blob struct {
	Access string `json:"access"`
	Secret string `json:"secret"`
}

type getCredentialsOutput struct {
	AccessKeyID     string
	SecretAccessKey string
}

type Signature struct {
	Service string
	Region  string
	Request *http.Request
	Body    string
	Query   url.Values

	SignedHeaderValues http.Header

	credValues       credentials.Value
	requestDateTime  string
	requestDate      string
	requestPayload   string
	signedHeaders    string
	canonicalHeaders string
	canonicalString  string
	credentialString string
	stringToSign     string
	signature        string
	authorization    string
}

func NewClient(endpooint string, opt *AuthOptions, uploadTimeout int64) (*Client, error) {
	u, err := url.Parse(endpooint)
	if err != nil {
		return nil, err
	}
	u.Path = path.Join(u.Path, ApiVersion)
	baseURL := u.String() + "/"

	client := &Client{
		endpoint:      endpooint,
		tenantId:      DefaultTenantId,
		version:       ApiVersion,
		baseURL:       baseURL,
		timeout:       time.Duration(DefaultTimeout) * time.Minute,
		uploadTimeout: time.Duration(uploadTimeout) * time.Minute,
		auth:          opt,
	}

	if opt.Strategy == "keystone" || opt.Strategy == "AK/SK" {
		if err := client.UpdateToken(); err != nil {
			return nil, err
		}
	}
	return client, nil
}

type ReqSettingCB func(req *httplib.BeegoHTTPRequest) error

func (c *Client) getToken(opt *AuthOptions) (*tokens.CreateResult, error) {
	var pwdCiphertext = opt.Password

	if opt.EnableEncrypted {
		// Decrypte the password
		pwdTool := pwd.NewPwdEncrypter(opt.PwdEncrypter)
		password, err := pwdTool.Decrypter(pwdCiphertext)
		if err != nil {
			return nil, err
		}
		pwdCiphertext = password
	}

	auth := gophercloud.AuthOptions{
		IdentityEndpoint: opt.AuthUrl,
		DomainName:       opt.DomainName,
		Username:         opt.UserName,
		Password:         pwdCiphertext,
		TenantName:       opt.TenantName,
	}

	provider, err := openstack.AuthenticatedClient(auth)
	if err != nil {
		log.Error("When get auth client:", err)
		return nil, err
	}

	// Only support keystone v3
	identity, err := openstack.NewIdentityV3(provider, gophercloud.EndpointOpts{})
	c.Identity = identity

	if err != nil {
		log.Error("When get identity session:", err)
		return nil, err
	}
	r := tokens.Create(identity, &auth)
	return &r, nil
}

func (c *Client) UpdateToken() error {
	t, err := c.getToken(c.auth)
	if err != nil {
		log.Errorf("Get token failed, %v", err)
		return err
	}
	user, err := t.ExtractUser()
	if err != nil {
		log.Errorf("extract user failed, %v", user)
		return err
	}
	c.userId = user.ID
	log.V(5).Infof("userId:%v", user.ID)
	project, err := t.ExtractProject()
	if err != nil {
		log.Errorf("extract project failed, %v", err)
		return err
	}
	c.tenantId = project.ID
	token, err := t.ExtractToken()
	if err != nil {
		log.Errorf("extract token failed, %v", err)
		return err
	}
	c.token = token
	log.V(5).Infof("TokenId:%s, ExpiresAt:%v", token.ID, token.ExpiresAt)
	return nil
}

// from Keystone And error will be returned if the retrieval fails.
func (c *Client) getCredentials() (*getCredentialsOutput, error) {
	credsOptList := creds.ListOpts{
		UserID: c.userId,
	}
	allPages, err := creds.List(c.Identity, credsOptList).AllPages()

	credentials, err := creds.ExtractCredentials(allPages)
	log.V(4).Infof("Credentials: %s", credentials)

	if err != nil {
		return nil, err
	}

	blob, err := getBlob(credentials)

	if blob != nil {
		return &getCredentialsOutput{
			AccessKeyID:     blob.Access,
			SecretAccessKey: blob.Secret,
		}, nil
	}
	return nil, err
}

// Returns a credential Blob for getting access and secret
// And error will be returned if it fails.
func getBlob(credentials []creds.Credential) (*Blob, error) {
	blob := &Blob{}
	credential := credentials[0]

	var blobStr = credential.Blob
	b := strings.Replace(blobStr, "\\", "", -1)
	err := json.Unmarshal([]byte(b), blob)

	if err != nil {
		return nil, err
	}
	return blob, nil
}

func (c *Client) doRequest(method, u string, in interface{}, cb ReqSettingCB) ([]byte, http.Header, error) {
	req := httplib.NewBeegoRequest(u, method)
	req.Header("Content-Type", "application/xml")
	beforeExpires := c.token.ExpiresAt.Add(time.Minute)
	if time.Now().After(beforeExpires) {
		log.Warning("token is about to expire, update it")
		if err := c.UpdateToken(); err != nil {
			return nil, nil, err
		}
	}
	if c.auth.Strategy == "keystone" {
		req.Header("X-Auth-Token", c.token.ID)
	} else if c.auth.Strategy == "AK/SK" {
		getCredentialsOutput, err := c.getCredentials()
		if err != nil {
			return nil, nil, err
		}
		AK := getCredentialsOutput.AccessKeyID
		SK := getCredentialsOutput.SecretAccessKey
		log.V(5).Infof("AK:%v, SK:%v", AK, SK)
		//Create a keystone credentials Provider client for retrieving credentials
		credentials := keystonecredentials.NewCredentialsClient(AK)
		//Create a Signer and the calculate the signature based on the Header parameters passed in request
		signer := signer.NewSigner(credentials)
		request, err := http.NewRequest(method, u, nil)

		if err != nil {
			return nil, nil, err
		}

		requestDate := time.Now().UTC().Format("20060102")
		requestDateTime := time.Now().UTC().Format("20060102T150405Z")
		request.Header.Add(constants.SignDateHeader, requestDateTime)
		credentialStr := AK + "/" + requestDate + "/" + constants.Region + "/" + constants.Service + "/" + "sign_request"
		calculatedSignature, err := signer.Sign(request, "", constants.Service, constants.Region, requestDateTime, requestDate, credentialStr)
		log.V(5).Infof("req.Request=%v,service=%v,region=%v,requestDateTime=%v,requestDate=%v,credentialStr=%v", request, constants.Service, constants.Region, requestDateTime, requestDate, credentialStr)
		log.V(5).Infof("calculatedSignature:%v", calculatedSignature)
		Authorization := "OPENSDS-HMAC-SHA256 Credential=" + credentialStr + ",SignedHeaders=host;x-auth-date,Signature=" + calculatedSignature
		req.Header(constants.AuthorizationHeader, Authorization)
		req.Header(constants.SignDateHeader, requestDateTime)
	}

	req.SetTimeout(c.timeout, c.timeout)
	if cb != nil {
		if err := cb(req); err != nil {
			return nil, nil, err
		}
	}

	if in != nil {
		var body interface{}
		switch in.(type) {
		case string, []byte:
			body = in
		default:
			body, _ = xml.Marshal(in)
		}
		req.Body(body)
	}

	resp, err := req.Response()
	if err != nil {
		log.Errorf("Do http request failed, method: %s\n url: %s\n error: %v", method, u, err)
		return nil, nil, err
	}

	log.V(5).Infof("%s: %s OK\n", method, u)
	rbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("Get byte[] from response failed, method: %s\n url: %s\n error: %v", method, u, err)
		return nil, nil, err
	}
	return rbody, resp.Header, nil
}

func (c *Client) request(method, p string, in, out interface{}, cb ReqSettingCB) error {
	u, err := url.Parse(p)
	if err != nil {
		return err
	}
	base, err := url.Parse(c.baseURL)
	if err != nil {
		return err
	}

	fullUrl := base.ResolveReference(u)
	b, _, err := c.doRequest(method, fullUrl.String(), in, cb)
	if err != nil {
		return err
	}

	if out != nil {
		log.V(5).Infof("Response:\n%s\n", string(b))
		err := xml.Unmarshal(b, out)
		if err != nil {
			log.Errorf("unmarshal error, reason:%v", err)
			return err
		}
	}
	return nil
}

type Object struct {
	ObjectKey  string `xml:"ObjectKey"`
	BucketName string `xml:"BucketName"`
	Size       uint64 `xml:"Size"`
}

type ListObjectResponse struct {
	ListObjects []Object `xml:"ListObjects"`
}

type InitiateMultipartUploadResult struct {
	Xmlns    string `xml:"xmlns,attr"`
	Bucket   string `xml:"Bucket"`
	Key      string `xml:"Key"`
	UploadId string `xml:"UploadId"`
}

type UploadPartResult struct {
	Xmlns      string `xml:"xmlns,attr"`
	PartNumber int64  `xml:"PartNumber"`
	ETag       string `xml:"ETag"`
}

type Part struct {
	PartNumber int64  `xml:"PartNumber"`
	ETag       string `xml:"ETag"`
}

type CompleteMultipartUpload struct {
	Xmlns string `xml:"xmlns,attr"`
	Part  []Part `xml:"Part"`
}

type CompleteMultipartUploadResult struct {
	Xmlns    string `xml:"xmlns,attr"`
	Location string `xml:"Location"`
	Bucket   string `xml:"Bucket"`
	Key      string `xml:"Key"`
	ETag     string `xml:"ETag"`
}

func (c *Client) UploadObject(bucketName, objectKey string, data []byte) error {
	p := path.Join("s3", bucketName, objectKey)
	err := c.request("PUT", p, data, nil, nil)
	return err
}

func (c *Client) ListObject(bucketName string) (*ListObjectResponse, error) {
	p := path.Join("s3", bucketName)
	object := &ListObjectResponse{}
	if err := c.request("GET", p, nil, object, nil); err != nil {
		return nil, err
	}
	return object, nil
}

func (c *Client) RemoveObject(bucketName, objectKey string) error {
	p := path.Join("s3", bucketName, objectKey)
	err := c.request("DELETE", p, nil, nil, nil)
	return err
}

func (c *Client) InitMultiPartUpload(bucketName, objectKey string) (*InitiateMultipartUploadResult, error) {
	p := path.Join("s3", bucketName, objectKey)
	p += "?uploads"
	out := &InitiateMultipartUploadResult{}
	if err := c.request("PUT", p, nil, out, nil); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *Client) UploadPart(bucketName, objectKey string, partNum int64, uploadId string, data []byte, size int64) (*UploadPartResult, error) {
	log.Infof("upload part buf size:%d", len(data))
	p := path.Join("s3", bucketName, objectKey)
	p += fmt.Sprintf("?partNumber=%d&uploadId=%s", partNum, uploadId)
	out := &UploadPartResult{}
	reqSettingCB := func(req *httplib.BeegoHTTPRequest) error {
		req.Header("Content-Length", strconv.FormatInt(size, 10))
		req.SetTimeout(c.uploadTimeout, c.uploadTimeout)
		return nil
	}
	if err := c.request("PUT", p, data, out, reqSettingCB); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *Client) CompleteMultipartUpload(
	bucketName string,
	objectKey string,
	uploadId string,
	input *CompleteMultipartUpload) (*CompleteMultipartUploadResult, error) {

	p := path.Join("s3", bucketName, objectKey)
	p += fmt.Sprintf("?uploadId=%s", uploadId)
	out := &CompleteMultipartUploadResult{}
	if err := c.request("PUT", p, input, nil, nil); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *Client) AbortMultipartUpload(bucketName, objectKey string) error {
	// TODO: multi-cloud has not implemented it yet. so just comment it.
	//p := path.Join("s3", "AbortMultipartUpload", bucketName, objectKey)
	//if err := c.request("DELETE", p, nil, nil); err != nil {
	//	return err
	//}
	return nil
}

func (c *Client) DownloadPart(bucketName, objectKey string, offset, size int64) ([]byte, error) {
	p := path.Join("s3", bucketName, objectKey)

	reqSettingCB := func(req *httplib.BeegoHTTPRequest) error {
		rangeStr := fmt.Sprintf("bytes:%d-%d", offset, offset+size-1)
		req.Header("Range", rangeStr)
		req.SetTimeout(c.uploadTimeout, c.uploadTimeout)
		return nil
	}

	u, err := url.Parse(p)
	if err != nil {
		return nil, err
	}
	base, err := url.Parse(c.baseURL)
	if err != nil {
		return nil, err
	}

	fullUrl := base.ResolveReference(u)
	body, _, err := c.doRequest("GET", fullUrl.String(), nil, reqSettingCB)
	if err != nil {
		return nil, err
	}
	return body, nil
}
