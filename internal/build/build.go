package build

import (
	"time"
)

var Version = "DEV"

var Date = "" // YYYY-MM-DD

func init() {
	Date = time.Now().String()
}
