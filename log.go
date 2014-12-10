package ws

import (
	"fmt"
	"github.com/stvp/rollbar"
	"os"
)

func InitLogger() {
	rollbar.Token = os.Getenv("ROLLBAR_TOKEN")
	rollbar.Environment = os.Getenv("ROLLBAR_ENV")
}

func LogError(err error) {
	if rollbar.Token == "" {
		fmt.Println("ERROR: " + err.Error())
	} else {
		rollbar.Error(rollbar.ERR, err)
	}
}

func LogInfo(msg string) {
	if rollbar.Token == "" {
		fmt.Println("INFO: " + msg)
	} else {
		rollbar.Message(rollbar.INFO, msg)
	}
}

func LogDebug(msg string) {
	if rollbar.Token == "" {
		fmt.Println("DEBUG: " + msg)
	} else {
		rollbar.Message(rollbar.DEBUG, msg)
	}
}
