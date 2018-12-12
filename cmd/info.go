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

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "See the info about a fragrance in your collection",
	Long: `Get all info about the fragrance of your choice stored in the system, 
right now you can view its name, fragrance house, release year and scent note breakdown.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if asList, err := cmd.Flags().GetBool("list"); asList {
			if err != nil {
				return err
			}
			display.FragranceListInfo()
		} else {
			display.FragranceInfo()
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
	infoCmd.Flags().BoolP("list", "l", false, "Display the list of fragrances as an alpohabetical list.")
}
