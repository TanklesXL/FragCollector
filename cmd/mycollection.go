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
	"FragCollector/display"

	"github.com/spf13/cobra"
)

// mycollectionCmd represents the mycollection command
var mycollectionCmd = &cobra.Command{
	Use:   "mycollection",
	Short: "Visualize your fragrance collection",
	Long: `This command by default displays the collection in alphabetical order by fragrance house.
Use the subcommands alpha, and notes to see the collection differently`,
	Run: func(cmd *cobra.Command, args []string) {
		display.CollectionAlphabeticalByBrand()
	},
}
var alphabetical bool

func init() {
	rootCmd.AddCommand(mycollectionCmd)
}
