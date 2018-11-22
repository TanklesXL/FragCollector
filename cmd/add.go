// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
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

package cmd

import (
	mfi "FragCollector/manipulatefragranceitems"
	"fmt"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a fragrance to your collection",
	Long:  `Find a fragrance from the BaseNotes database and add it (along with all of its relevant information) to your collection`,
	Run: func(cmd *cobra.Command, args []string) {
		if url != "" {
			mfi.AddToCollection(url)
		} else if house != "" && name != "" {
			mfi.SearchByHouse(house, name)
		} else if name != "" {
			mfi.SearchByName(name)
		} else {
			fmt.Println("Please use either --name/-n, --url/-u or the combination of --name/-n and --house/-n together")
		}

	},
}
var url string
var name string
var house string

func init() {
	rootCmd.AddCommand(addCmd)
	// Here you will define your flags and configuration settings.
	addCmd.Flags().StringVarP(&url, "url", "u", "", "URL of item on BaseNotes")
	addCmd.Flags().StringVarP(&name, "name", "n", "", "Name of fragrance to search on BaseNotes")
	addCmd.Flags().StringVar(&house, "house", "", "Fragrance house to start search with")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
