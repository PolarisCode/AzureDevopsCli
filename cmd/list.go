/*
Copyright © 2021 IMRAN ALIYEV <imran.aliyev@gmail.com>

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
	"log"
	"os"

	"github.com/microsoft/azure-devops-go-api/azuredevops/core"
	"github.com/polariscode/AzureDevopsCli/azureclient"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Returns list of projects",
	Long:  `Return full list of projects from AzureDevops`,
	Run: func(cmd *cobra.Command, args []string) {

		url := os.Getenv("azure_project_url")
		token := os.Getenv("azure_token")

		client, ctx := azureclient.CreateClient(url, token)

		// Get first page of the list of team projects for your organization
		responseValue, err := client.GetProjects(ctx, core.GetProjectsArgs{})
		if err != nil {
			log.Fatal(err)
		}

		index := 0
		for responseValue != nil {
			// Log the page of team project names
			for _, teamProjectReference := range (*responseValue).Value {
				fmt.Printf("%v. %v\n", index, *teamProjectReference.Name)
				index++
			}

			// if continuationToken has a value, then there is at least one more page of projects to get
			if responseValue.ContinuationToken != "" {
				// Get next page of team projects
				projectArgs := core.GetProjectsArgs{
					ContinuationToken: &responseValue.ContinuationToken,
				}
				responseValue, err = client.GetProjects(ctx, projectArgs)
				if err != nil {
					log.Fatal(err)
				}
			} else {
				responseValue = nil
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
