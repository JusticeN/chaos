package main

import (
	"chaos"
	"fmt"
)

// you can declare an Action to later add as behaviour like this
func shootAction(agent *chaos.Agent) error {
	for i := 0; i < 5; i++ {
		fmt.Printf("[%v] Bang %v ... 🃏 \n", agent.Name, i)
	}
	return nil
}

func behaviourExample() {

	smith := &chaos.Agent{
		Name: "Smith",
		Setup: func(agent *chaos.Agent) {
			agent.AddOneShotBehaviour(shootAction)
		},
	}
	// register new created agent
	chaos.RegisterAgent(smith)
	// start agent
	chaos.Start()

}
