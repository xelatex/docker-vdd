// Copyright 2015-2018 Arthur Chunqi Li. All rights reserved.

package daemon

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"net/http"
)

func sendResponse(w http.ResponseWriter, v interface{}) error {
	output, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		info := fmt.Sprintf("Convert to json failed, %v", err)
		log.Errorf(info)
		return err
	}
	log.Infof("Send response:\n %v", string(output))
	_, err = w.Write(output)
	return err
}

func readRequest(r *http.Request) *pluginRequest {
	request := &pluginRequest{}
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		log.Errorf("Convert request from json failed, %v", err)
	}
	return request
}
