// Copyright © 2018 Robert Attard <robert.attard@mail.mcgill.ca>
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
	Long: `Find a fragrance from the BaseNotes database and add it (along with all of its relevant information) to your collection.
the add command needs to be used with flags followed by the necessary information.
The flags are: --url/-u, --name/-n , or a combination of --name/-n and the --brand/-b flags, 
please see the readme for more info.`,
	Run: func(cmd *cobra.Command, args []string) {
		if url != "" {
			mfi.AddToCollection(url)
		} else if brand != "" && name != "" {
			mfi.SearchByHouse(brand, name)
		} else if name != "" {
			mfi.SearchByName(name)
		} else {
			fmt.Println("Please use either --name/-n, --url/-u or the combination of --name/-n and --brand/-b together")
		}

	},
}

var url string
var name string
var brand string

func init() {
	rootCmd.AddCommand(addCmd)
	// Here you will define your flags and configuration settings.
	addCmd.Flags().StringVarP(&url, "url", "u", "", "URL of the item on BaseNotes")
	addCmd.Flags().StringVarP(&name, "name", "n", "", "Name of fragrance to search on BaseNotes")
	addCmd.Flags().StringVarP(&brand, "brand", "b", "", "Fragrance house to start search with")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
