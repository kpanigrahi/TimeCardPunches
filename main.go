// Copyright 2018 panigrahi kiran@gmail com  All rights reserved
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file

package main

import (
	"fmt"
	"log"

	"./config"
	"./process"
	"./utils"
)

func main() {
	// printBanner()
	// fmt.Println()

	configFileLoc := "./config/config.json"
	config := config.Load(configFileLoc)

	fileList := utils.GetFileList(config.Folders.In)
	if len(fileList) == 0 {
		// no file(s) found for processing
		log.Fatal(fmt.Sprintf("no file(s) %s: The system cannot find any file(s) for processing.", config.Folders.In))
	} else {
		process.Context.Config = config
		process.PreProcess()
		process.Process()
		process.PostProcess()
	}
}
