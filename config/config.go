// Copyright 2018 panigrahi.kiran@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package config

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"../utils"
)

type folder struct {
	In   string `json:"input"`
	Wrkg string `json:"working"`
	Out  string `json:"output"`
}

type environment struct {
	HTTP     string `json:"http"`
	Name     string `json:"name"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type header struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type event struct {
	Path         string        `json:"path"`
	Headers      []header      `json:"headers"`
	Environments []environment `json:"environments"`
}

// Config - structure to hold the entire configuration
type Config struct {
	Folders         folder `json:"folders"`
	ChunkSize       int    `json:"chunkSize"`
	Threads         int    `json:"threads"`
	NoOfPrllEvnts   int    `json:"noOfParallelEvents"`
	TimeEvent       event  `json:"timeEvent"`
	TimeRecordEvent event  `json:"timeRecordEvent"`

	UUID        string // not part of the config.json
	RunDttmStmp string // not part of the config.json
}

// Load - function to load the configuration from the file and gets a
// Config object
func Load(configFileName string) Config {
	var config Config

	configFile, err := os.Open(configFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer configFile.Close()

	err = json.NewDecoder(configFile).Decode(&config)
	if err != nil {
		log.Fatal(err)
	}
	// get the current time stamp
	sCurDttmStmp := time.Now().Format("20060102150405")
	// set the UUID
	config.UUID = sCurDttmStmp
	// for different format please refer to the below URL
	// https://medium.com/@Martynas/formatting-date-and-time-in-golang-5816112bf098
	// set the run date time
	config.RunDttmStmp = time.Now().Format("2006-Jan-02 15:04:05.000")

	currWrkgDir, _ := os.Getwd()
	// set the fully qualified path if user has provided relative path
	config.Folders.In = utils.GetFullQualifiedPath(currWrkgDir, config.Folders.In)
	config.Folders.Wrkg = utils.GetFullQualifiedPath(currWrkgDir, config.Folders.Wrkg)
	config.Folders.Out = utils.GetFullQualifiedPath(currWrkgDir, config.Folders.Out)

	return config
}
