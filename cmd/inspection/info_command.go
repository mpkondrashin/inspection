/*

Inspection (c) by Mikhail Kondrashin (mkondrashin@gmail.com)

Code is released under CC BY license: https://creativecommons.org/licenses/by/4.0/

on_command.go - command to turn on inspection

*/

package main

import (
	"context"
	"log"
)

type infoCommand struct {
	baseCommand
}

func (c *infoCommand) Execute() error {
	info, err := c.cOne.GetAccountInfo_(context.TODO())
	if err != nil {
		return err
	}
	log.Printf("ID: %s", info.ID)
	log.Printf("Alias: %s", info.Alias)
	log.Printf("Locale: %s", info.Locale)
	log.Printf("Timezone: %s", info.Timezone)
	log.Printf("Region: %s", info.Region)
	log.Printf("State: %s", info.State)
	log.Printf("Created: %v", info.Created)
	log.Printf("LastModified: %v", info.LastModified)
	log.Printf("URN: %s", info.Urn)
	log.Printf("MFA Required: %v", info.MfaRequired)
	for n, link := range info.Links {
		log.Printf("Link #%d: Rel: %s", n, link.Rel)
		log.Printf("Link #%d: Href: %s", n, link.Href)
		log.Printf("Link #%d: Method: %s", n, link.Method)
	}
	return nil
}

func newInfoCommand() *infoCommand {
	c := &infoCommand{}
	c.Setup(cmdInfo, "get account info")
	return c
}
