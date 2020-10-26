// Copyright 2020 Kaan Karakaya
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
	"flag"
	"fmt"
	"github.com/theykk/gkey"
	"os"
)

func main() {

	password := flag.String("p", "", "-p Master Password")
	realm := flag.String("r", "", "-r Realm example github.com")
	length := flag.Int("l", 16, "-l length of password")

	// ? Parse flags
	flag.Parse()

	fancyPassword, err := gkey.GenPass(*password, *realm, *length)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// ? Print password and print new line to stderr
	// ? because we don't want to grap new line on pipeline

	fmt.Print(fancyPassword)
	_, _ = fmt.Fprint(os.Stderr, "\n")
}
