/*

Inspection (c) by Mikhail Kondrashin (mkondrashin@gmail.com)

Code is released under CC BY license: https://creativecommons.org/licenses/by/4.0/

on_command.go - command to turn on inspection

*/

package main

import (
	"context"
	"inspection/pkg/cone"
)

type onCommand struct {
	baseCommand
}

func (c *onCommand) Execute() error {
	return c.cOne.SetInspectionBypass_(context.TODO(), c.awsRegion, cone.ActionInspect)
}

func newOnCommand() *onCommand {
	c := &onCommand{}
	c.Setup(cmdOn, "turn inspection on")
	return c
}
