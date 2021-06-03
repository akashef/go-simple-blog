package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"time"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
func processBody(Body io.ReadCloser) map[string]interface{} {
	body, err := ioutil.ReadAll(Body)
	checkErr(err)
	bodyMap := make(map[string]interface{})
	err = json.Unmarshal(body, &bodyMap)
	checkErr(err)
	return bodyMap
}
func CurrentTime() string {
	return time.Now().UTC().Format("2006-01-02 15:04:05")
}
