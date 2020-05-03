package main

import (
	"fmt"
	"time"

	"github.com/JusticeN/chaos"
)

const timeSnapshot string = "timeSnapshot"

// func actionSetAge(agent *chaos.Agent) error {
// 	time.Sleep(5 * time.Second)
// 	agent.SetData("age", "50")
// 	return nil
// }

// func actionDoSomethongWithAge(agent *chaos.Agent) error {
// 	age := agent.GetData("age")
// 	fmt.Println("age: ", age)
// 	return nil
// }

func globalExample() {
	agent1 := &chaos.Agent{
		Name: "agent1",
		Setup: func(agent *chaos.Agent) {
			agent.SetData("timeSnapshot", "0")
			// update snapshot time every 3 seconds
			agent.AddTickerBehaviour(chaos.DelayedBehaviour(func(agent *chaos.Agent, delay int) error {
				time := time.Now().Format(time.StampMilli)
				agent.SetData(timeSnapshot, time)
				// fmt.Println(agent.GetData("age"))
				return nil
			}), 3)

			// read snapshot variable every seconds
			agent.AddTickerBehaviour(chaos.DelayedBehaviour(func(agent *chaos.Agent, delay int) error {
				fmt.Println(agent.GetData("timeSnapshot"))
				return nil
			}), 1)

			// agent.AddOneShotBehaviour(actionSetAge)
			// agent.AddOneShotBehaviour(actionDoSomethongWithAge)
		},
	}

	chaos.RegisterAgent(agent1)
	chaos.Start()
}
