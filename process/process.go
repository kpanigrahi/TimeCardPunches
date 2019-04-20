// Copyright 2018 panigrahi kiran@gmail com  All rights reserved
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file

package process

import (
	"path/filepath"

	"../rest/oracle/timeEvent"
	"../rest/oracle/timeRecordEvent"
	"../utils"
)

// Process - the core function which actually processes the file(s)
func Process() {
	// get the list of file(s) from working folder
	sWrkgFldr := filepath.Join(Context.Config.Folders.Wrkg, Context.Config.UUID, filepath.Base(Context.Config.Folders.In))
	fileList := utils.GetFileList(sWrkgFldr)
	for _, file := range fileList {
		createRequest(file)
	}
}

func createRequest(sFileName string) []string {
	// check the file for which we have to create the request
	if utils.IsTimeEventInput(sFileName) {
		timeEvent.Context.Config = Context.Config
		return timeEvent.CreateRequest(sFileName)
	} else if utils.IsTimeRecordEvent(sFileName) {
		timeRecordEvent.Context.Config = Context.Config
		return timeRecordEvent.CreateRequest(sFileName)
	}
	return nil
}
