package main

import (
	"chaos"
)

func simple() {
	neo := &chaos.Agent{
		Name:  "Neo",
		Setup: func(agent *chaos.Agent) {},
	}
	chaos.RegisterAgent(neo)
	chaos.Start()
}
