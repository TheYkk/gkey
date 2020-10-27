/*
Copyright 2020 Kaan Karakaya

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package gkey

import (
	"crypto/sha256"
	"errors"
	"strings"
)

var Chars = []string{
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
	"1", "2", "3", "4", "5", "6", "7", "8", "9", "0",
	"`", "~", "!", "@", "#", "$", "%", "^", "&", "*", "(", ")", "-", "_", "=", "+",
	"[", "{", "]", "}", "|", ";", ":", ",", "<", ".", ">", "?",
}

func GenPass(password, realm string, length int) (string, error) {
	if password == "" {
		return "", errors.New("password can not be empty")
	}

	if realm == "" {
		return "", errors.New("realm can not be empty")
	}

	if length == 0 {
		return "", errors.New("length can not be empty")
	}
	if length > 32 {
		return "", errors.New("length can not be bigger than 32")
	}
	// ? Create sha256sum from string
	h := sha256.New()
	_ , _ = h.Write([]byte(password + "-" + realm))

	// ? Translate hex to base 94 password
	var str []string
	for _, hexa := range h.Sum(nil) {
		pos := HexToGivenBase(hexa, byte(len(Chars)))
		str = append(str, Chars[pos])
	}

	// ? Shorten to a certain length
	return strings.Join(str[0:length], ""), nil
}

// ? Translate base 255 code to base 94
func HexToGivenBase(base byte, max byte) byte {
	return base % max
}
