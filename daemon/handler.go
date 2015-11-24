// Copyright 2015-2018 Arthur Chunqi Li. All rights reserved.

package daemon

import (
	log "github.com/Sirupsen/logrus"
	"net/http"
)

type pluginInfo struct {
	Implements []string
}

type pluginRequest struct {
	Name string
	Opts map[string]string
}

type pluginResponse struct {
	Mountpoint string `json:",omitempty"`
	Err        string `json:",omitempty"`
}

func (s *daemon) pluginActivate(w http.ResponseWriter, r *http.Request) {
	log.Infof("Handle plugin activate: %v %v", r.Method, r.RequestURI)
	info := pluginInfo{
		Implements: []string{"VolumeDriver"},
	}
	sendResponse(w, info)
}

func (s *daemon) pluginCreateVolume(w http.ResponseWriter, r *http.Request) {
	log.Infof("Handle plugin create volume: %v %v", r.Method, r.RequestURI)
	request := readRequest(r)
	log.Infof("Request volume name: %s", request.Name)
	log.Infof("Request opts: %v", request.Opts)
	info := pluginResponse{}
	sendResponse(w, info)
}

func (s *daemon) pluginRemoveVolume(w http.ResponseWriter, r *http.Request) {
	log.Infof("Handle plugin remove volume: %v %v", r.Method, r.RequestURI)
	request := readRequest(r)
	log.Infof("Request volume name: %s", request.Name)
	log.Infof("Request opts: %v", request.Opts)
	info := pluginResponse{}
	sendResponse(w, info)
}

func (s *daemon) pluginMountVolume(w http.ResponseWriter, r *http.Request) {
	log.Infof("Handle plugin mount volume: %v %v", r.Method, r.RequestURI)
	request := readRequest(r)
	log.Infof("Request volume name: %s", request.Name)
	log.Infof("Request opts: %v", request.Opts)
	info := pluginResponse{
		Mountpoint: "/tmp/test",
	}
	sendResponse(w, info)
}

func (s *daemon) pluginUnmountVolume(w http.ResponseWriter, r *http.Request) {
	log.Infof("Handle plugin unmount volume: %v %v", r.Method, r.RequestURI)
	request := readRequest(r)
	log.Infof("Request volume name: %s", request.Name)
	log.Infof("Request opts: %v", request.Opts)
	info := pluginResponse{}
	sendResponse(w, info)
}

func (s *daemon) pluginVolumePath(w http.ResponseWriter, r *http.Request) {
	log.Infof("Handle plugin volume path: %v %v", r.Method, r.RequestURI)
	request := readRequest(r)
	log.Infof("Request volume name: %s", request.Name)
	log.Infof("Request opts: %v", request.Opts)
	info := pluginResponse{
		Mountpoint: "/tmp/test",
	}
	sendResponse(w, info)
}
