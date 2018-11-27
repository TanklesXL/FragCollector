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

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Refresh the json file to correspond to the main collection",
	Long: `In the case of a manual change to the json file, this command can be used to load
the new changes to the other parts (alphabetical, alphabetical by brand, and scent notes).`,
	Run: func(cmd *cobra.Command, args []string) {
		mfi.ManualUpdate()
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
