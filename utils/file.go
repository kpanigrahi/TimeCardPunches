// Copyright 2018 panigrahi kiran@gmail com  All rights reserved
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file

package utils

import (
	"bufio"
	"encoding/csv"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// GetFullQualifiedPath - returns the fully qualified path of the
// specified folder name
func GetFullQualifiedPath(qualifiedPath, folder string) string {
	if folder[0:1] == "." {
		return filepath.Join(qualifiedPath, folder)
	}
	return filepath.Join(folder)
}

// GetFileList - returns the file(s) and folder(s) under a specified folder
func GetFileList(folder string) []string {
	files, err := ioutil.ReadDir(folder)
	if err != nil {
		log.Fatal(err)
	}
	var fileList []string
	for _, file := range files {
		fileList = append(fileList, filepath.Join(folder, file.Name()))
	}
	return fileList
}

// MkdirAll - creates a folder
func MkdirAll(folderName string) {
	err := os.MkdirAll(folderName, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}

// GetHeader - gets only the header from the file name provided
func GetHeader(strInputFileName string) []string {
	inFile, error := os.Open(strInputFileName)
	if error != nil {
		log.Fatal(error)
	}
	defer inFile.Close()
	reader := csv.NewReader(bufio.NewReader(inFile))
	hdrData, error := reader.Read()
	if error != nil {
		log.Fatal(error)
	}
	return hdrData
}

// IsTimeEventInput - from the input file looking at the header
// data it tries to identify whether it is a Time Event or not
func IsTimeEventInput(strInputFileName string) bool {
	var timeEvent bool
	hdrData := GetHeader(strInputFileName)
	for _, data := range hdrData {
		if strings.EqualFold(data, "timeEvents") {
			timeEvent = true
			break
		}
	}
	return timeEvent
}

// IsTimeRecordEvent - from the input file looking at the header
// data it tries to identify whether it is a Time Record Event or not
func IsTimeRecordEvent(strInputFileName string) bool {
	var timeRecordEvent bool
	hdrData := GetHeader(strInputFileName)
	for _, data := range hdrData {
		if strings.EqualFold(data, "timeRecordEvent") {
			timeRecordEvent = true
			break
		}
	}
	return timeRecordEvent
}
