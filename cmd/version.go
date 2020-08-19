/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"fmt"
	"github.com/kmdkuk/cricli/log"
	"google.golang.org/grpc"

	"github.com/kmdkuk/cricli/version"
	"github.com/spf13/cobra"
	runtimeapi "k8s.io/cri-api/pkg/apis/runtime/v1alpha2"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: Version,
}

func init() {
	rootCmd.AddCommand(versionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func Version(_ *cobra.Command, _ []string) {
	fmt.Printf("version: %s-%s (%s)\n", version.Version, version.Revision, version.BuildDate)
	conn, err := grpc.Dial("127.0.0.1:19003", grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Error("err: ", err)
		}
	}()
	client := runtimeapi.NewRuntimeServiceClient(conn)
	message := &runtimeapi.VersionRequest{}
	res, err := client.Version(context.TODO(), message)
	if err != nil {
		log.Fatalf("error:%#v \n", err)
	}
	fmt.Printf("result:%#v \n", res)
}
