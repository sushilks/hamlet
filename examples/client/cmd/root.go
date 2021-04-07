// Copyright 2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vmware/hamlet/examples/client/cmd/start"
	"github.com/vmware/hamlet/examples/client/cmd/version"
)

// RootCmd represents the base command to be executed when called without any
// command line arguments.
var RootCmd = &cobra.Command{
	Use:   "client",
	Short: "Client Example",
	Long:  "Client Example",
}

// init initializes the root command instance.
func init() {
	RootCmd.AddCommand(version.NewCommand())
	RootCmd.AddCommand(start.NewCommand())
}
