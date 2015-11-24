// Copyright 2015-2018 Arthur Chunqi Li. All rights reserved.

package daemon

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"net"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

type daemon struct {
	Router *mux.Router
}

func (s *daemon) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	info := fmt.Sprintf("Handler not found: %v %v", r.Method, r.RequestURI)
	log.Errorf(info)
	w.Write([]byte(info))
}

func Start(sockFile string) error {
	log.Info("Prepare for sock file: ", sockFile)
	s := &daemon{}

	s.Router = createRouter(s)

	path := filepath.Dir(sockFile)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.MkdirAll(path, os.ModeDir|0755); err != nil {
			return err
		}
	}
	if _, err := os.Stat(sockFile); err == nil {
		log.Warnf("Remove previous sockfile: %v", sockFile)
		if err := os.Remove(sockFile); err != nil {
			return err
		}
	}

	conn, err := net.Listen("unix", sockFile)
	if err != nil {
		log.Errorf("Listen sockfile error: %v %v", sockFile, err)
		return err
	}
	defer conn.Close()

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, os.Interrupt, os.Kill, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		log.Infof("Caught signal %s: shutting down.", sig)
		done <- true
	}()

	go func() {
		log.Info("Server start")
		err = http.Serve(conn, s.Router)
		if err != nil {
			log.Errorf("http server error, %v", err)
		}
		done <- true
	}()

	<-done
	return nil
}
