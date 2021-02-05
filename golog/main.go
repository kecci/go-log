package main

import (
	"github.com/kataras/golog"
)

func main() {
	golog.SetLevel("debug")

	golog.Print("Print")
	// No Trace
	golog.Debug("This is a debug message")
	golog.Info("This is an info message, with colors (if the output is terminal)")
	golog.Warn("This is a warning message")
	golog.Error("This is an error message")
	golog.Fatal(`Fatal will exit no matter what,
    but it will also print the log message if logger's Level is >=FatalLevel`)
	// No Panic

}
