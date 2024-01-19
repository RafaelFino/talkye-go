package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	rotatelogs "github.com/lestrrat/go-file-rotatelogs"

	config "talkye/internal/config"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Print("Usage: talkye <config_file>\n")
		os.Exit(1)
	}

	configFile := os.Args[1]

	cfg, err := config.LoadConfigFromFile(configFile)
	if err != nil {
		fmt.Printf("Error loading config file %s: %s\n", configFile, err)
		os.Exit(1)
	}

	err = initLogger(cfg.LogPath)
	if err != nil {
		fmt.Printf("Error opening log file: %s, using stdout", err)
		log.SetOutput(os.Stdout)
	}

	fmt.Printf("%s\nStarting...", getLogo())

	//service := procspy.NewSpy(cfg)
	//go spy.Start(onUpdate)
	//defer spy.Stop()

	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)
	<-quitChannel

	log.Print("Stopping...\n")
}

func initLogger(path string) error {
	if err := os.Mkdir(path, 0755); !os.IsExist(err) {
		fmt.Printf("Error creating directory %s: %s", path, err)
		return err
	}

	writer, err := rotatelogs.New(
		fmt.Sprintf("%s/%s.log", path, "%Y%m%d"),
		rotatelogs.WithMaxAge(24*time.Hour),
		rotatelogs.WithRotationTime(time.Hour),
		rotatelogs.WithRotationCount(30), //30 days
	)
	if err != nil {
		log.Fatalf("Failed to Initialize Log File %s", err)
	}
	log.SetOutput(writer)

	return nil
}

func getLogo() string {
	return `
	_ __  _ __ ___   ___ ___ _ __  _   _ 
	| '_ \| '__/ _ \ / __/ __| '_ \| | | |
	| |_) | | | (_) | (__\__ \ |_) | |_| |
	| .__/|_|  \___/ \___|___/ .__/ \__, |
	| |                      | |     __/ |
	|_|                      |_|    |___/ 

`
}
