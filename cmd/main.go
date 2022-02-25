package main

import (
	"github.com/crazyacking/gwik/cmd/app"
	"github.com/crazyacking/gwik/pkg/common/log"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	defer func() {
		log.Sync()
	}()

	if err := app.RootCmd.Execute(); err != nil {
		log.NewLogger("RootCMD").Errorf("RootCmd execute failed: %v", err)
		os.Exit(-1)
	}
}
