import { Component, OnInit, ViewContainerRef, ViewChild, Directive, ElementRef, HostBinding, HostListener, AfterViewInit } from '@angular/core';
import { Http } from '@angular/http';
import { Router } from '@angular/router';
import { I18NService, Consts, ParamStorService } from 'app/shared/api';
// import { AppService } from 'app/app.service';
import { I18nPluralPipe } from '@angular/common';
import { MenuItem, SelectItem } from './components/common/api';

@Component({
    selector: 'app-root',
    templateUrl: './app.component.html',
    styleUrls: []
})
export class AppComponent implements OnInit, AfterViewInit{
    chromeBrowser: boolean = false;

    isLogin: boolean;

    hideLoginForm: boolean = false;

    linkUrl = "";

    username: string;

    password: string;

    dropMenuItems: MenuItem[];

    currentTenant: string="";

    showLoginAnimation: boolean=false;

    showLogoutAnimation: boolean=false;
    currentTime=new Date().getTime();
    lastTime=new Date().getTime();
    minExpireTime = 2 * 60 * 1000;
    advanceRefreshTime = 1*60*1000;
    defaultExpireTime = 10*60*1000;
    interval:any;
    intervalRefreshToken:any;
    showErrorMsg:boolean=false;
    errorMsg:string="";

    tenantItems = [];

    menuItems = [];

    menuItems_tenant = [
        {
            "title": "Home",
            "description": "Resources statistics",
            "routerLink": "/home"
        },
        {
            "title": "Volume",
            "description": "Block storage resources",
            "routerLink": "/block"
        }
    ]

    menuItems_admin = [
        {
            "title": "Home",
            "description": "Resources statistics",
            "routerLink": "/home"
        },
        {
            "title": "Volume",
            "description": "Block storage resources",
            "routerLink": "/block"
        },
        // {
        //     "title": "Multi-Cloud Service",
        //     "description": "5 replications, 1 migrations",
        //     "routerLink": "/cloud"
        // },
        {
            "title": "Profile",
            "description": "Block profiles",
            "routerLink": "/profile"
        },
        {
            "title": "Resource",
            "description": "Regions, availability zones and storages",
            "routerLink": "/resource"
        },
        {
            "title": "Identity",
            "description": "Managing tenants and users",
            "routerLink": "/identity"
        }
    ];

    activeItem: any;

    private msgs: any = [{ severity: 'warn', summary: 'Warn Message', detail: 'There are unsaved changes'}];

    constructor(
        private el: ElementRef,
        private viewContainerRef: ViewContainerRef,
        private http: Http,
        private router: Router,
        private paramStor: ParamStorService,
        public I18N: I18NService
    ){}
    
    ngOnInit() {
        let currentUserInfo = this.paramStor.CURRENT_USER();
        if(currentUserInfo != undefined && currentUserInfo != ""){
            this.hideLoginForm = true;

            let [username, userid, tenantname, tenantid] = [
                    this.paramStor.CURRENT_USER().split("|")[0],
                    this.paramStor.CURRENT_USER().split("|")[1],
                    this.paramStor.CURRENT_TENANT().split("|")[0],
                    this.paramStor.CURRENT_TENANT().split("|")[1] ];
            this.AuthWithTokenScoped({'name': username, 'id': userid});
        }else{
            this.isLogin = false;
            this.hideLoginForm = false;
        }
    }
    checkTimeOut(){
        this.currentTime = new Date().getTime(); //update current time
        let timeout = this.paramStor.TOKEN_PERIOD() ? this.paramStor.TOKEN_PERIOD() : this.defaultExpireTime;
        if(this.currentTime - this.lastTime > timeout){ //check time out
            this.logout();
        }
    }
    refreshLastTime(){
        this.lastTime = new Date().getTime(); 
    }
    /**
     * this function must be called after AuthWithTokenScoped succeed ,because it need username and password
     */
    refreshToken(){
        let request: any = { auth: {} };
        request.auth = {
            "identity": {
                "methods": [
                    "password"
                ],
                "password":{
                    "user": {
                        "name": this.paramStor.CURRENT_USER().split("|")[0],
                        "domain": {
                            "name": "Default"
                        },
                        "password": this.paramStor.PASSWORD()
                    }
                }
            }
        }

        this.http.post("/v3/auth/tokens", request).subscribe((res)=>{
            let token_id = res.headers.get('x-subject-token');
            let projectName = this.paramStor.CURRENT_TENANT().split("|")[0];
            let req: any = { auth: {} };
            req.auth = {
                "identity": {
                    "methods": [
                        "token"
                    ],
                    "token": {
                        "id": token_id
                    }
                },
                "scope": {
                "project": {
                    "name": projectName,
                    "domain": { "id": "default" }
                }
                }
            }

            this.http.post("/v3/auth/tokens", req).subscribe((r)=>{
                this.paramStor.AUTH_TOKEN( r.headers.get('x-subject-token') );
            });
        },
        error=>{
            console.log("Username or password incorrect.")
        });
    }
    ngAfterViewInit(){
        this.loginBgAnimation();
    }

    loginBgAnimation(){
        let obj =this.el.nativeElement.querySelector(".login-bg");
        if(obj){
            let obj_w = obj.clientWidth;
            let obj_h = obj.clientHeight;
            let dis = 50;
            obj.addEventListener("mousemove", (e)=>{
                let MX = e.clientX;
                let MY = e.clientY;
                let offsetX = (obj_w - 2258)*0.5 + (obj_w-MX)*dis / obj_w;
                let offsetY = (obj_h - 1363)*0.5 + (obj_h-MY)*dis / obj_h;
                obj.style.backgroundPositionX = offsetX +"px";
                obj.style.backgroundPositionY = offsetY +"px";
            })
        }
    }
    
    login() {
        let request: any = { auth: {} };
        request.auth = {
            "identity": {
                "methods": [
                    "password"
                ],
                "password":{
                    "user": {
                        "name": this.username,
                        "domain": {
                            "name": "Default"
                        },
                        "password": this.password
                    }
                }
            }
        }

        this.http.post("/v3/auth/tokens", request).subscribe((res)=>{
            //set token period start
            let token = res.json().token;
            let expires_at = token.expires_at;
            let issued_at = token.issued_at;
            let tokenPeriod = Date.parse(expires_at) - Date.parse(issued_at);
            if(tokenPeriod >= this.minExpireTime){
                this.paramStor.TOKEN_PERIOD(tokenPeriod);
            }
            //set token period end
            this.paramStor.AUTH_TOKEN(res.headers.get('x-subject-token'));
            this.paramStor.PASSWORD(this.password);
            let user = res.json().token.user;
            this.AuthWithTokenScoped(user);
            this.showErrorMsg = false;
        },
        error=>{
            switch(error.status){
                case 401:
                    this.errorMsg = this.I18N.keyID['sds_login_error_msg_401'];
                    break;
                case 503:
                    this.errorMsg = this.I18N.keyID['sds_login_error_msg_503'];
                    break;
                default:
                    this.errorMsg = this.I18N.keyID['sds_login_error_msg_default'];                
            }
            this.showErrorMsg = true;
        });
    }

    AuthWithTokenScoped(user, tenant?){
        if(this.interval){
            clearInterval(this.interval);
        }
        this.lastTime=new Date().getTime();
        this.interval = window.setInterval(()=>{
            this.checkTimeOut()
        }, 10000);
        // Get user owned tenants
        let reqUser: any = { params:{} };
        this.http.get("/v3/users/"+ user.id +"/projects", reqUser).subscribe((objRES) => {
            let projects = objRES.json().projects;
            let defaultProject = user.name != 'admin' ? projects[0] : projects.filter((project) => { return project.name == 'admin'})[0]; 
            let project = tenant===undefined ? defaultProject : tenant;

            this.tenantItems = [];
            projects.map(item => {
                let tenantItemObj = {};
                tenantItemObj["label"] = item.name;
                tenantItemObj["command"] = ()=>{
                    let username =  this.paramStor.CURRENT_USER().split("|")[0];
                    let userid =  this.paramStor.CURRENT_USER().split("|")[1];
                    this.AuthWithTokenScoped({'name': username, 'id': userid}, item);
                };
                this.tenantItems.push(tenantItemObj);
            })
 
            // Get token authentication with scoped
            let token_id = this.paramStor.AUTH_TOKEN(); 
            let req: any = { auth: {} };
            req.auth = {
                "identity": {
                    "methods": [
                        "token"
                    ],
                    "token": {
                        "id": token_id
                    }
                },
                "scope": {
                "project": {
                    "name": project.name,
                    "domain": { "id": "default" }
                }
                }
            }

            this.http.post("/v3/auth/tokens", req).subscribe((r)=>{
                this.paramStor.AUTH_TOKEN( r.headers.get('x-subject-token') );
                this.paramStor.CURRENT_TENANT(project.name + "|" + project.id);
                this.paramStor.CURRENT_USER(user.name + "|"+ user.id);

                this.username = this.paramStor.CURRENT_USER().split("|")[0];
                this.currentTenant = this.paramStor.CURRENT_TENANT().split("|")[0];

                if(this.username == "admin"){
                    this.menuItems = this.menuItems_admin;
                    this.dropMenuItems = [
                        { 
                            label: "Switch Region", 
                            items: [{ label: "default_region", command:()=>{} }]
                        },
                        { 
                            label: "Logout", 
                            command:()=>{ this.logout() }
                        }
                    ];
                }else{
                    this.menuItems = this.menuItems_tenant;
                    this.dropMenuItems = [
                        { 
                            label: "Switch Region", 
                            items: [{ label: "default_region", command:()=>{} }]
                        },
                        { 
                            label: "Switch Tenant", 
                            items: this.tenantItems
                        },
                        { 
                            label: "Logout", 
                            command:()=>{ this.logout() }
                        }
                    ];
                }

                this.isLogin = true;
                this.router.navigateByUrl("home");
                this.activeItem = this.menuItems[0];

                // annimation for after login
                this.showLoginAnimation = true;
                setTimeout(() => {
                    this.showLoginAnimation = false;
                    this.hideLoginForm = true;
                }, 500);
                if(this.intervalRefreshToken){
                    clearInterval(this.intervalRefreshToken);
                }
                let tokenPeriod = this.paramStor.TOKEN_PERIOD();
                let refreshTime = tokenPeriod ? (Number(tokenPeriod) - this.advanceRefreshTime) : this.defaultExpireTime;
                this.intervalRefreshToken = window.setInterval(()=>{
                    this.refreshToken()
                },refreshTime);
            })
        },
        error => {
            this.logout();
        })
    }

    logout() {
        this.paramStor.AUTH_TOKEN("");
        this.paramStor.CURRENT_USER("");
        this.paramStor.CURRENT_TENANT("");
        this.paramStor.PASSWORD("");
        this.paramStor.TOKEN_PERIOD("");
        if(this.interval){
            clearInterval(this.interval);
        }
        if(this.intervalRefreshToken){
            clearInterval(this.intervalRefreshToken);
        }
        // annimation for after logout
        this.hideLoginForm = false;
        this.showLogoutAnimation = true;
        setTimeout(() => {
            this.showLogoutAnimation = false;
            this.username = "";
            this.password = "";
            this.isLogin = false;
        }, 500);

    }

    onKeyDown(e) {
        let keycode = window.event ? e.keyCode : e.which;
        if(keycode == 13){
            this.login();
        }
    }

    menuItemClick(event, item) {
        this.activeItem = item;
    }

    supportCurrentBrowser(){
        let ie,
            firefox,
            safari,
            chrome,
            cIE = 11,
            cFirefox = 40,
            cChrome = 40;
        let ua = navigator.userAgent.toLowerCase();
        let isLinux = (ua.indexOf('linux') >= 0);

        if(this.isIE()) {
            if(ua.indexOf('msie') >= 0) {
                ie = this.getSys(ua.match(/msie ([\d]+)/));
            } else {
                ie = this.getSys(ua.match(/trident.*rv:([\d]+)/));
            }
        }else if(navigator.userAgent.indexOf("Firefox") > 0){
            firefox = this.getSys(ua.match(/firefox\/([\d]+)/));
        }else if(ua.indexOf("safari") != -1 && !(ua.indexOf("chrome") != -1)) {
            safari = this.getSys(ua.match(/version\/([\d]+)/));
        }else if(ua.indexOf("chrome") != -1) {
            chrome = this.getSys(ua.match(/chrome\/([\d]+)/));
        }

        if ((firefox) / 1 < cFirefox || (chrome) / 1 < cChrome || (ie) / 1 < cIE) {
            return true;
        }

        return false;
    }

    isIE() {
        return navigator.userAgent.toLowerCase().indexOf('trident') >= 0;
    }

    getSys (browserVersionArr) {
        if( !browserVersionArr) {
            return 0;
        } else if( browserVersionArr.length < 2) {
            return 0;
        } else {
            return browserVersionArr[1];
        }
    }
}
