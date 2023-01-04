package build

import (
	"os"
	"time"
)

var Version = "DEV"

var Date = "" // YYYY-MM-DD

func init() {
	if version := os.Getenv("OMF_VERSION"); version != "" {
		Version = version
	}
	Date = time.Now().String()
}
