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

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a fragrance from the collection",
	Long:  `Select a fragrance from your collection to remove.`,
	Run: func(cmd *cobra.Command, args []string) {
		mfi.RemoveFromCollection()
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
