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

// alphaCmd represents the alpha command
var alphaCmd = &cobra.Command{
	Use:   "alpha",
	Short: "Display your collection in alphabetical order by name",
	Long:  `Visualize your collection with a numbered list of your fragrances ordered alphabetically by their names.`,
	Run: func(cmd *cobra.Command, args []string) {
		display.CollectionAlphabetical()
	},
}

func init() {
	mycollectionCmd.AddCommand(alphaCmd)
}
