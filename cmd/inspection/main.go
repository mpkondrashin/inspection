/*

Inspection (c) by Mikhail Kondrashin (mkondrashin@gmail.com)

Code is released under CC BY license: https://creativecommons.org/licenses/by/4.0/

main.go - main file

*/

package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

const (
	EnvPrefix      = "C1NS"
	ConfigFileName = "config"
	ConfigFileType = "yaml"
)

const (
	cmdStatus = "status"
	cmdOn     = "on"
	cmdOff    = "off"
)

const (
	flagApiKey    = "api_key"
	flagRegion    = "region"
	flagAccountID = "account_id"
	flagAWSRegion = "aws_region"
)

var commands = []command{
	newStatusCommand(),
	newOnCommand(),
	newOffCommand(),
}

func usage() {
	var commandNames []string
	for _, c := range commands {
		commandNames = append(commandNames, c.Name())
	}
	fmt.Fprintf(os.Stderr, "Inspection — control Trend Micro Cloud One Network Security Hosted Infrastructure inspection\nUsage: %s%s {%s} [options]\n",
		name(), exe(), strings.Join(commandNames, "|"))
	fmt.Fprintf(os.Stderr, "Commands available:\n")
	for _, c := range commands {
		fmt.Fprintf(os.Stderr, "\t%s - %s\n", c.Name(), c.Description())
	}
	fmt.Fprintf(os.Stderr, "For more details, run vone <command> --help\n")
	os.Exit(2)
}

func pickCommand(args []string) error {
	subcommand := args[0]
	for _, cmd := range commands {
		if cmd.Name() == subcommand {
			err := cmd.Init(args[1:])
			if err != nil {
				return fmt.Errorf("Init error: %w", err)
			}
			log.Printf("Command: %s\n", cmd.Name())
			return cmd.Execute()
		}
	}
	return fmt.Errorf("unknown command: %s", subcommand)
}

func main() {
	if len(os.Args) < 2 {
		usage()
	}
	err := pickCommand(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Done")
	}
}

func exe() string {
	if runtime.GOOS == "windows" {
		return ".exe"
	}
	return ""
}

func name() string {
	_, path, _, ok := runtime.Caller(0)
	if !ok {
		panic(fmt.Errorf("runtime.Caller() error"))
	}
	dir := filepath.Dir(path)
	folder := filepath.Base(dir)
	return folder
}
