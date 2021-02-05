package main

import (
	"fmt"
	"os"

	stdlog "log"

	"github.com/go-kit/kit/log"
)

func main() {
	w := log.NewSyncWriter(os.Stderr)
	loggerFmt := log.NewLogfmtLogger(w)
	err := loggerFmt.Log("question", "what is the meaning of life?", "answer", 42)
	if err != nil {
		fmt.Println(err.Error())
	}

	logger := log.NewJSONLogger(log.NewSyncWriter(os.Stdout))
	stdlog.SetOutput(log.NewStdlibAdapter(logger))
	stdlog.Print("I sure like pie")
}
