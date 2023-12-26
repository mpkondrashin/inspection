/*

Inspection (c) by Mikhail Kondrashin (mkondrashin@gmail.com)

Code is released under CC BY license: https://creativecommons.org/licenses/by/4.0/

off_command.go - command to turn off inspection

*/

package main

import (
	"context"
	"inspection/pkg/cone"
)

type offCommand struct {
	baseCommand
}

func (c *offCommand) Execute() error {
	return c.cOne.SetInspectionBypass_(context.TODO(), c.awsRegion, cone.ActionBypass)
}

func newOffCommand() *offCommand {
	c := &offCommand{}
	c.Setup(cmdOff, "turn inspection off")
	return c
}
