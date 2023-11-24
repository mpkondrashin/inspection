package main

import (
	"context"
)

type onCommand struct {
	baseCommand
}

func (c *onCommand) Execute() error {
	return c.cOne.SetInspectionBypass(context.TODO(), ActionInspect)
}

func newOnCommand() *onCommand {
	c := &onCommand{}
	c.Setup(cmdOn, "turn inspection on")
	return c
}
