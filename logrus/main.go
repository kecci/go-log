package main

import (
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Print("Print")
	logrus.Trace("Trace")
	logrus.Debug("Debug")
	logrus.Info("Info")
	logrus.Warn("Warn")
	logrus.Error("Error")
	logrus.Fatal("Fatal")
	logrus.Panic("Panic")
}