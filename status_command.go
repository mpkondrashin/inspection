package main

import (
	"context"
	"log"
)

type statusCommand struct {
	baseCommand
}

func (c *statusCommand) Execute() error {
	status, err := c.cOne.GetInspectionBypassStatus(context.TODO())
	if err != nil {
		return err
	}
	log.Printf("Action: %v", status.Action)
	log.Printf("Status: %v", status.Status)
	log.Printf("Last change: %v", status.UpdateTime.Local())
	return nil
}

func newStatusCommand() *statusCommand {
	c := &statusCommand{}
	c.Setup(cmdStatus, "get current inspection status")

	return c
}
