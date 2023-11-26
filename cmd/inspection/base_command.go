/*

Inspection (c) by Mikhail Kondrashin (mkondrashin@gmail.com)

Code is released under CC BY license: https://creativecommons.org/licenses/by/4.0/

base_command.go - base struct for all commands

*/

package main

import (
	"fmt"
	"inspection/pkg/cone"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type command interface {
	Name() string
	Description() string
	Init(args []string) error
	Execute() error
}

type baseCommand struct {
	name        string
	description string
	cOne        *cone.CloudOneNS
	fs          *pflag.FlagSet
}

func (c *baseCommand) Setup(name, description string) {
	c.name = name
	c.description = description
	c.fs = pflag.NewFlagSet(name, pflag.ExitOnError)
	c.fs.String(flagApiKey, "", "CloudOne API key (On CloudOne console go to Administration->API Keys->New)")
	c.fs.String(flagRegion, "", "CloudOne region (On CloudOne console go to Administration-Account Settings->Region)")
	c.fs.String(flagAccountID, "", "CloudOne account ID (On CloudOne console go to Administration-Account Settings->ID)")
	c.fs.String(flagAWSRegion, "", "AWS region, i.e. us-east-1")
	c.fs.Usage = func() {
		fmt.Fprintf(os.Stderr, "%s\nAvailable options:\n", c.description)
		c.fs.PrintDefaults()
	}
}

func (c *baseCommand) Name() string {
	return c.name
}

func (c *baseCommand) Description() string {
	return c.description
}

func (c *baseCommand) String() string {
	return c.name
}

func (c *baseCommand) Init(args []string) error {
	err := c.fs.Parse(os.Args[2:])
	if err != nil {
		return err
	}
	if err := viper.BindPFlags(c.fs); err != nil {
		panic(err)
	}
	viper.SetEnvPrefix(EnvPrefix)
	viper.AutomaticEnv()

	viper.SetConfigName(ConfigFileName)
	viper.SetConfigType(ConfigFileType)
	path, err := os.Executable()
	if err == nil {
		dir := filepath.Dir(path)
		viper.AddConfigPath(dir)
	}
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		notFoundErr, ok := err.(viper.ConfigFileNotFoundError)
		_ = notFoundErr
		if !ok {
			panic(err) //Fatal(RCConfigReadError, "ReadInConfig: %v", err)
		} else {
			log.Printf("%s: loaded", ConfigFileName)
		}
		//LogIt(Debug, "ReadInConfig: %v", notFoundErr)
	}

	c.cOne = cone.NewCloudOneNS(
		GetNotEmpty(flagApiKey),
		GetNotEmpty(flagRegion),
		GetNotEmpty(flagAccountID),
		GetNotEmpty(flagAWSRegion),
	)
	return nil
}

func GetNotEmpty(key string) string {
	value := viper.GetString(key)
	if value == "" {
		log.Fatalf("Missing %s", key)
	}
	return value
}
