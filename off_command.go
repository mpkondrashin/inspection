package main

import (
	"context"
)

type offCommand struct {
	baseCommand
}

func (c *offCommand) Execute() error {
	return c.cOne.SetInspectionBypass(context.TODO(), ActionBypass)
}

func newOffCommand() *offCommand {
	c := &offCommand{}
	c.Setup(cmdOff, "turn inspection off")
	return c
}
