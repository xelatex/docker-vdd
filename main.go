// Copyright 2015-2018 Arthur Chunqi Li. All rights reserved.

package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/xelatex/docker-vdd/daemon"
)

func main() {
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
	log.Info("Hello World!")
	daemon.Start("/run/docker/plugins/vdd.sock")
}
