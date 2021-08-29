/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"fmt"
	"strings"

	"github.com/lirlia/go-mypj/voyage/kurl/client"
	"github.com/spf13/cobra"
)

type options struct {
	headers client.Headers
}

var (
	o       = &options{}
	headers = []string{}
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kurl",
	Short: "Simple http client by Golang",
	// 	Long: `A longer description that spans multiple lines and likely contains
	// examples and usage of using your application. For example:

	// Cobra is a CLI library for Go thacobt empowers applications.
	// This application is a tool to generate the needed files
	// to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())

}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.kurl.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().StringP("toggle", "t", false, "Help message for toggle")

	rootCmd.PersistentFlags().StringArrayVarP(&headers, "header", "H", nil, "HTTP Header for http request. Header must be set key:value style. (ex: Content-Type:application/json)")
	// rootCmd.MarkFlagRequired("region")

}

func setHeader() error {
	var headerSplit []string
	for _, header := range headers {

		if !strings.Contains(header, ":") {
			return fmt.Errorf("Header has no colon(:). Header must be set key:value style")
		}
		headerSplit = strings.Split(header, ":")
		if len(headerSplit) != 2 {
			return fmt.Errorf("Header has multiple colon. Header must be set key:value style")
		}
		o.headers = append(o.headers, client.Header{Key: headerSplit[0], Value: headerSplit[1]})
	}
	return nil
}
