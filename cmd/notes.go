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
	"github.com/TanklesXL/FragCollector/display"

	"github.com/spf13/cobra"
)

// notesCmd represents the notes command
var notesCmd = &cobra.Command{
	Use:   "notes",
	Short: "See the list of fragrance notes in your collection and the fragrances containing them",
	Long: `Get a list of the fragrance notes contained in the collection, and select one for which to expand.
Using the --list/-n flag (without a parameter) will display the fully expanded list, with the fragrances containing each scent note. `,
	RunE: func(cmd *cobra.Command, args []string) error {
		if asList, err := cmd.Flags().GetBool("list"); asList {
			if err != nil {
				return err
			}
			display.CollectionNotes()
		} else {
			display.SingleNote()
		}
		return nil
	},
}

func init() {
	mycollectionCmd.AddCommand(notesCmd)
	notesCmd.Flags().BoolP("list", "l", false, "Display all scent notes and the fragrances that contain them.")
}
