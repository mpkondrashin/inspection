/*

Inspection (c) by Mikhail Kondrashin (mkondrashin@gmail.com)

Code is released under CC BY license: https://creativecommons.org/licenses/by/4.0/

base_command.go - base struct for all commands

*/

package main

import (
	"context"
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
			panic(err)
		} else {
			log.Printf("%s: loaded", ConfigFileName)
		}
	}
	c.cOne = cOne()
	return nil
}

func cOne() *cone.CloudOneNS {
	var err error
	accountID := GetNotEmpty(flagAccountID)
	apiKey := GetNotEmpty(flagApiKey)
	awsRegion := viper.GetString(flagAWSRegion)
	c1Region := viper.GetString(flagRegion)

	if c1Region == "" {
		if awsRegion == "" {
			log.Println("Cloud One and AWS regions are missing. Detecting...")
			_, c1Region, err = cone.DetectBothRegions(context.TODO(), accountID, apiKey)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			log.Println("Cloud One region is missing. Detecting...")
			c1Region, err = cone.DetectCloudOneRegion(context.TODO(), accountID, apiKey, awsRegion)
			if err != nil {
				log.Fatal(err)
			}
		}
		log.Println("Detected Cloud One region:", c1Region)
	}

	if awsRegion == "" {
		log.Println("AWS region region is missing. Detecting...")
		awsRegions := cone.DetectAWSRegions(context.TODO(), accountID, apiKey, c1Region)
		switch len(awsRegions) {
		case 0:
			log.Fatal("No AWS regions with Network Security Hosted Infrastructure detected")
		case 1:
			log.Println("Detected NSHI in following AWS region:", c1Region)
			awsRegion = awsRegions[0]
		default:
			for _, r := range awsRegions {
				log.Println("Detected NSHI in AWS region:", r)
			}
			log.Fatal("Provide AWS region value and run Inspection again")
		}
	}

	return cone.NewCloudOneNS(
		apiKey,
		c1Region,
		accountID,
		awsRegion,
	)
}

func GetNotEmpty(key string) string {
	value := viper.GetString(key)
	if value == "" {
		log.Fatalf("Missing %s", key)
	}
	return value
}
