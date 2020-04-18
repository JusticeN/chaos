// Copyright 2020 Justice Nanhou. All rights reserved.
// license that can be found in the LICENSE file.

package chaos

import (
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/google/uuid"
)

// Context chaos Context
// type Context context.Context

// RegisterAgent register a new Agent
func RegisterAgent(a *Agent) {
	// AgentSotre()[a.Name] = a
	a.ID = genID()
	GetAgentStoreInstance().Add(a)
}
func genID() string {
	id := uuid.New().String()
	res := strings.Split(id, "-")
	return res[len(res)-1]
}

// RegisterAgentAndScale register a new Agent and scale
func RegisterAgentAndScale(a *Agent, scale int) {
	for i := 0; i < scale; i++ {
		agt := *a
		name := fmt.Sprintf("%v_%v", a.Name, i)
		agt.Name = name
		GetAgentStoreInstance().Add(&agt)
	}
}

// StartAgent ...
func StartAgent(wg *sync.WaitGroup, agent *Agent) {
	go func(wg *sync.WaitGroup, agent *Agent) {
		defer wg.Done()
		agent.start()
		agent.doTakeDown()
	}(wg, agent)
}

// Start start all agents
func Start() {
	var wg sync.WaitGroup
	log.Println("--------------------------------")
	log.Println("--------- Start chaos ----------")
	log.Println("--------------------------------")
	setUpAllAgent()
	// fmt.Printf("agentStore count: %v\n", GetAgentStoreInstance().Count())
	wg.Add(GetAgentStoreInstance().Count())
	for _, agent := range GetAgentStoreInstance().List() {
		StartAgent(&wg, agent)
	}
	wg.Wait()
	// acl.GetDispatcherInstance().Close()
	log.Println("--------------------------------")
	log.Println("--------- End chaos   ----------")
	log.Println("--------------------------------")

}

func setUpAllAgent() {
	for _, agent := range GetAgentStoreInstance().List() {
		agent.doSetup()
	}
}
