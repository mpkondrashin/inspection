/*

Inspection (c) by Mikhail Kondrashin (mkondrashin@gmail.com)

Code is released under CC BY license: https://creativecommons.org/licenses/by/4.0/

status_command.go - command to get current inspection status

*/

package main

import (
	"context"
	"log"
)

type statusCommand struct {
	baseCommand
}

func (c *statusCommand) Execute() error {
	status, err := c.cOne.GetInspectionBypassStatus(context.TODO(), c.awsRegion)
	if err != nil {
		return err
	}
	log.Printf("Action: %v", status.Action)
	log.Printf("Status: %v", status.Status)
	if status.Error != "" {
		log.Printf("Error: %v", status.Error)
	}
	log.Printf("Last change: %v", status.UpdateTime.Local())
	return nil
}

func newStatusCommand() *statusCommand {
	c := &statusCommand{}
	c.Setup(cmdStatus, "get current inspection status")

	return c
}
