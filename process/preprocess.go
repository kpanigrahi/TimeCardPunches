// Copyright 2018 panigrahi kiran@gmail com  All rights reserved
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file

package process

import (
	"log"
	"os"
	"path/filepath"

	"../utils"
)

// PreProcess - function to prepare the environment before all the
// processing begins
func PreProcess() {
	createFldrs()
	moveFiles()
}

func createFldrs() {
	// create all the required folders
	utils.MkdirAll(Context.Config.Folders.Out)
	sFileName := filepath.Base(Context.Config.Folders.In)
	sWrkgFldr := filepath.Join(Context.Config.Folders.Wrkg, Context.Config.UUID, sFileName)
	utils.MkdirAll(sWrkgFldr)
	sRequest := filepath.Join(Context.Config.Folders.Wrkg, Context.Config.UUID, "requests", "timeEvents")
	utils.MkdirAll(sRequest)
	sRequest = filepath.Join(Context.Config.Folders.Wrkg, Context.Config.UUID, "requests", "timeRecordEvents")
	utils.MkdirAll(sRequest)
}

func moveFiles() {
	// get the list of file(s) from input folder
	fileList := utils.GetFileList(Context.Config.Folders.In)
	// move all the file(s) from input folder to working/uuid/input folder
	for _, file := range fileList {
		sWrkgFldr := filepath.Join(Context.Config.Folders.Wrkg, Context.Config.UUID, filepath.Base(Context.Config.Folders.In), filepath.Base(file))
		err := os.Rename(file, sWrkgFldr)
		if err != nil {
			log.Fatal(err)
		}
	}
}
