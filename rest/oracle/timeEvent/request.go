// Copyright 2018 panigrahi kiran@gmail com  All rights reserved
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file

package timeEvent

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// MetaHDR structure to store the header information
type MetaHDR struct {
	REQUESTNUMBER       int8
	SOURCEID            int8
	REQUESTTIMESTAMP    int8
	TIMEEVENTS          int8
	DEVICEID            int8
	EVENTDATETIME       int8
	SUPPLIERDEVICEEVENT int8
	REPORTERID          int8
	REPORTERIDTYPE      int8
	TIMEEVENTATTRIBUTES int8
	NAME1               int8
	VALUE1              int8
}

// CreateRequest - function to create the request file(s)
func CreateRequest(sFileName string) []string {
	var sAFileNames []string
	sFileNameNameOnly := filepath.Base(sFileName)[:len(filepath.Base(sFileName))-len(filepath.Ext(sFileName))]

	var tmEvntRqsts Request
	var tmEvnts []TimeEvent
	intNumberOfFiles := 1
	intNumberOfLines := 0
	intChunkLines := 0

	f, error := os.Open(sFileName)
	if error != nil {
		log.Fatal(error)
	}
	defer f.Close()
	r := csv.NewReader(bufio.NewReader(f))
	hdrData, error := r.Read()
	if error != nil {
		log.Fatal(error)
	}
	metaHdrTmEnvt := getMetaTimeEvent(sFileName, hdrData)
	curTime := time.Now().Format("20060102150405")
	for {
		fileData, error := r.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		tmEvnts = append(tmEvnts, populateTimeEvent(fileData, metaHdrTmEnvt))
		strRequestNumber := curTime + "_" + strconv.Itoa(intNumberOfFiles)
		// retrieve all the values from line and populate the structure
		tmEvntRqsts = Request{
			RequestNumber:    strRequestNumber,
			SourceID:         fileData[metaHdrTmEnvt.SOURCEID],
			RequestTimeStamp: fileData[metaHdrTmEnvt.REQUESTTIMESTAMP],
			TimeEvents:       tmEvnts,
		}
		intNumberOfLines++
		intChunkLines++
		if intChunkLines == Context.Config.ChunkSize {
			sFileName = filepath.Join(Context.Config.Folders.Wrkg, Context.Config.UUID, "requests", "timeEvents", sFileNameNameOnly+"_"+strconv.Itoa(intNumberOfFiles)+".json")
			sAFileNames = append(sAFileNames, sFileName)
			fmt.Println(intNumberOfFiles, "->", filepath.Base(writeTimeEventsToFile(tmEvntRqsts, sFileName)), "->", intChunkLines)
			intNumberOfFiles++
			intChunkLines = 0
			tmEvnts = make([]TimeEvent, 0)
		}
	}
	if intChunkLines < Context.Config.ChunkSize && intChunkLines != 0 {
		sFileName = filepath.Join(Context.Config.Folders.Wrkg, Context.Config.UUID, "requests", "timeEvents", sFileNameNameOnly+"_"+strconv.Itoa(intNumberOfFiles)+".json")
		sAFileNames = append(sAFileNames, sFileName)
		fmt.Println(intNumberOfFiles, "->", filepath.Base(writeTimeEventsToFile(tmEvntRqsts, sFileName)), "->", intChunkLines)
	}
	return sAFileNames
}

// getMetaTimeEvent is a function which is used to read the
// header and based on the header it populates the Metadata
// structure for Time Event. Once the structure is populated
// then using these values we can then populate the actual
// structure, which then can be used to print the JSON.
func getMetaTimeEvent(inputFile string, hdrData []string) MetaHDR {
	var metaHdrTmEnvt MetaHDR
	// hdrData := GetHeader(inputFile)
	for index, hdrName := range hdrData {
		metaTmEvntStructVal := reflect.ValueOf(&metaHdrTmEnvt)
		structField := metaTmEvntStructVal.Elem().FieldByName(strings.ToUpper(hdrName))
		if structField.IsValid() {
			structField.SetInt(int64(index))
		}
	}
	return metaHdrTmEnvt
}

// populateTimeEventAttribute - populates the time event
// attributes from the input file data
func populateTimeEventAttribute(line []string, metaHdrTmEnvt MetaHDR) []Attribute {
	var tmEvntAttrs []Attribute
	var name, value string
	for index, lineValue := range line[metaHdrTmEnvt.NAME1:] {
		if strings.TrimSpace(lineValue) != "" {
			if index%2 == 0 {
				name = lineValue
			} else {
				value = lineValue
				tmEvntAttrs = append(tmEvntAttrs, Attribute{
					Name:  name,
					Value: value,
				})
				name, value = "", ""
			}
		}
	}
	return tmEvntAttrs
}

// populateTimeEvent - populate the time event details from
// the input file data
func populateTimeEvent(line []string, metaHdrTmEnvt MetaHDR) TimeEvent {
	tmEvntAttrs := populateTimeEventAttribute(line, metaHdrTmEnvt)
	intReportID, _ := strconv.Atoi(line[metaHdrTmEnvt.REPORTERID])
	tmEvnt := TimeEvent{
		DeviceID:            line[metaHdrTmEnvt.DEVICEID],
		EventDateTime:       line[metaHdrTmEnvt.EVENTDATETIME],
		SupplierDeviceEvent: line[metaHdrTmEnvt.SUPPLIERDEVICEEVENT],
		ReporterID:          intReportID,
		ReporterIDType:      line[metaHdrTmEnvt.REPORTERIDTYPE],
		TimeEventAttributes: tmEvntAttrs,
	}
	return tmEvnt
}

// writeTimeEventsToFile - writes the time event structure
// into a JSON file
func writeTimeEventsToFile(tmEvntRqsts Request, strOutputFileName string) string {
	punchJSON, _ := json.MarshalIndent(tmEvntRqsts, "", "  ")
	// create the output file
	outFile, outFileError := os.Create(strOutputFileName)
	if outFileError != nil {
		log.Fatal(outFileError)
	}
	defer outFile.Close()
	outFile.Write(punchJSON)
	return strOutputFileName
}
