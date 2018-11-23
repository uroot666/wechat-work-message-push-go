package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Request struct {
}

func (e Request) Get(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Print("err:")
		fmt.Print(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body
}

func (e Request) PostJson(url string, jsonStr []byte) []byte {

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var objmap map[string]*json.RawMessage
	err = json.Unmarshal(body, &objmap)
	if err != nil {
		logger.Infof("PostJson error")
		os.Exit(0)
	}

	errcode := string(*objmap["errcode"])
	if errcode != "0" {
		fmt.Print("PostJson errmsg:" + string(*objmap["errmsg"]))
	}
	return body
}