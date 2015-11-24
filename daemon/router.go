// Copyright 2015-2018 Arthur Chunqi Li. All rights reserved.

package daemon

import (
	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"net/http"
)

func createRouter(s *daemon) *mux.Router {
	router := mux.NewRouter()
	router.NotFoundHandler = s
	pluginMap := map[string]map[string]http.HandlerFunc{
		"POST": {
			"/Plugin.Activate":      s.pluginActivate,
			"/VolumeDriver.Create":  s.pluginCreateVolume,
			"/VolumeDriver.Remove":  s.pluginRemoveVolume,
			"/VolumeDriver.Mount":   s.pluginMountVolume,
			"/VolumeDriver.Unmount": s.pluginUnmountVolume,
			"/VolumeDriver.Path":    s.pluginVolumePath,
		},
	}
	for method, routes := range pluginMap {
		for uri, f := range routes {
			log.Debugf("Register plugin handler %s, %s", method, uri)
			router.Path(uri).Methods(method).HandlerFunc(f)
		}
	}
	return router
}
