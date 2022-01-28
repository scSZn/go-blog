package global

import (
	"github.com/scSZn/blog/conf"
	"io"
	"log"
	"os"
)

var LogFileWriter io.Writer

func init() {
	var err error
	LogFileWriter, err = os.OpenFile(conf.GetLogPath(), os.O_APPEND|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatalf("init logger writer fail, err: %v", err)
	}
}
