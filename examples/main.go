package main

import (
	"chaos"
	"chaos/message"
	"fmt"
	"math/rand"
	"time"
)

func shootOnce(agent *chaos.Agent) error {
	fmt.Printf(">>>> %v..... bang only once \n", agent.Name)
	return nil
}
func shootWaked(agent *chaos.Agent, delay int) error {
	fmt.Printf(">>>> %v..... bang after %v seconds \n", agent.Name, delay)
	return nil
}
func shoot(agent *chaos.Agent) error {
	fmt.Printf(">>>> %v..... bang \n", agent.Name)
	// time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
	return nil
}
func receiveMsg(msg message.Message) {
	fmt.Printf("%v received message :> %v", msg.Receiver, msg.Content)
}
func shootOnceandSendMsg(agent *chaos.Agent) error {
	agent.SendMessage("Trinity", "Hello from "+agent.Name)
	fmt.Printf(">>>> %v..... bang only once \n", agent.Name)
	return nil
}
func sendMsgRegulary(agent *chaos.Agent) error {
	time.Sleep(5 * time.Second)
	agent.SendMessage("Trinity", "Hello from "+agent.Name)
	return nil
}
func sendMsgRegulary1(agent *chaos.Agent) error {
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	agent.SendMessage("Neo", "cool guy "+agent.Name)
	return nil
}

func main() {
	// go run ./examples/
	// trinity := &chaos.Agent{
	// 	Name: "Trinity",
	// 	Setup: func(agent *chaos.Agent) {
	// 		agent.OnMessage(receiveMsg)
	// 		agent.AddOneShotBehaviour(shootOnceandSendMsg)
	// 		agent.AddCyclicBehaviour(sendMsgRegulary1)
	// 	},
	// }
	// neo := &chaos.Agent{
	// 	Name: "Neo",
	// 	Setup: func(agent *chaos.Agent) {
	// 		agent.OnMessage(receiveMsg)
	// 		// agent.AddOneShotBehaviour(shootOnceandSendMsg)
	// 		agent.AddCyclicBehaviour(sendMsgRegulary)
	// 	},
	// }
	// chaos.RegisterAgent(neo)
	// chaos.RegisterAgent(trinity)
	// chaos.Start()

	// simple()
	behaviourExample()
	// globalExample()

}
