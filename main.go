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

package main

import (
	"FragCollector/cmd"
	"os"
)

//pyr
const adgp string = "http://www.basenotes.net/ID26145389.html"
const shalimar string = "http://www.basenotes.net/ID10211632.html"

//flat
const btf string = "http://www.basenotes.net/ID26147432.html"
const pradaCandy string = "http://www.basenotes.net/ID26132465.html"

func main() {
	if _, err := os.Stat("/path/to/whatever"); os.IsNotExist(err) {
		f, err := os.Create("./Collection.json")
		if err != nil {
			panic(err)
		}
		defer f.Close()
	}

	cmd.Execute()
}
