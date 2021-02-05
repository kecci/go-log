package main

import (
	"fmt"

	"go.uber.org/zap"
)

func main() {
	fmt.Println("ZAP")
	log, _ := zap.NewProduction()
	// No Print
	// No Trace
	log.Debug("Debug")
	log.Info("Info")
	log.Warn("Warn")
	log.Error("Error")
	// log.Fatal("Fatal")
	// log.Panic("Panic")

	fmt.Println("ZAP SUGAR")
	// Zap Sugar
	sugar, _ := zap.NewProduction()
	// No Print
	// No Trace
	sugar.Debug("Debug")
	sugar.Info("Info")
	sugar.Warn("Warn")
	sugar.Error("Error")
	// sugar.Fatal("Fatal")
	// sugar.Panic("Panic")
}
