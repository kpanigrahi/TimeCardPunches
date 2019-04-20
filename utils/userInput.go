// Copyright 2018 panigrahi kiran@gmail com  All rights reserved
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file

package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// GetUserInput - ask user the specific question and returns
// the answer back to from where it is being called.
func GetUserInput(strQuestion string) string {
	fmt.Print(strQuestion)
	strAnswer, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	// convert CRLF to LF
	strAnswer = strings.Replace(strAnswer, "\r\n", "", -1)
	return strAnswer
}
