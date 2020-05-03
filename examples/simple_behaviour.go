package main

import (
	"fmt"

	"github.com/JusticeN/chaos"
)

// you can declare an Action to later add as behaviour like this
func simpleShootBehaviour(agent *chaos.Agent) error {
	for i := 0; i < 5; i++ {
		fmt.Printf("[%v] Bang %v ... 🃏", agent.Name, i)
	}
	return nil
}

func simpleBehaviour() {

	chaos.NewAgentWithSetup("my Agent", func(agent *chaos.Agent) {
		agent.AddOneShotBehaviour(simpleShootBehaviour)
	})
	// start agent
	chaos.Start()

}
