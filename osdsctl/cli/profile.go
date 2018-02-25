// Copyright (c) 2017 Huawei Technologies Co., Ltd. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
This module implements a entry into the OpenSDS service.

*/

package cli

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/opensds/opensds/pkg/model"
	"github.com/spf13/cobra"
)

var profileCommand = &cobra.Command{
	Use:   "profile",
	Short: "manage OpenSDS profile resources",
	Run:   profileAction,
}

var profileCreateCommand = &cobra.Command{
	Use:   "create <profile info>",
	Short: "create a new profile resource",
	Run:   profileCreateAction,
}

var profileShowCommand = &cobra.Command{
	Use:   "show <profile id>",
	Short: "show information of specified profile",
	Run:   profileShowAction,
}

var profileListCommand = &cobra.Command{
	Use:   "list",
	Short: "get all profile resources",
	Run:   profileListAction,
}

var profileDeleteCommand = &cobra.Command{
	Use:   "delete <profile id>",
	Short: "delete a specified profile resource",
	Run:   profileDeleteAction,
}

func init() {
	profileCommand.AddCommand(profileCreateCommand)
	profileCommand.AddCommand(profileShowCommand)
	profileCommand.AddCommand(profileListCommand)
	profileCommand.AddCommand(profileDeleteCommand)
}

func profileAction(cmd *cobra.Command, args []string) {
	cmd.Usage()
	os.Exit(1)
}

func profileCreateAction(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		fmt.Fprintln(os.Stderr, "The number of args is not correct!")
		cmd.Usage()
		os.Exit(1)
	}

	prf := &model.ProfileSpec{}
	if err := json.Unmarshal([]byte(args[0]), prf); err != nil {
		fmt.Fprintln(os.Stderr, err)
		cmd.Usage()
		os.Exit(1)
	}

	resp, err := client.CreateProfile(prf)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	keys := KeyList{"Id", "CreatedAt", "UpdatedAt", "Name", "Description", "Extra"}
	PrintDict(resp, keys, FormatterList{})
}

func profileShowAction(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		fmt.Fprintln(os.Stderr, "The number of args is not correct!")
		cmd.Usage()
		os.Exit(1)
	}

	resp, err := client.GetProfile(args[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	keys := KeyList{"Id", "CreatedAt", "UpdatedAt", "Name", "Description", "Extra"}
	PrintDict(resp, keys, FormatterList{})
}

func profileListAction(cmd *cobra.Command, args []string) {
	if len(args) != 0 {
		fmt.Fprintln(os.Stderr, "The number of args is not correct!")
		cmd.Usage()
		os.Exit(1)
	}

	resp, err := client.ListProfiles()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	keys := KeyList{"Id", "Name", "Description", "Extra"}
	PrintList(resp, keys, FormatterList{})
}

func profileDeleteAction(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		fmt.Fprintln(os.Stderr, "The number of args is not correct!")
		cmd.Usage()
		os.Exit(1)
	}

	err := client.DeleteProfile(args[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("Delete profile(%s) success.\n", args[0])
}
