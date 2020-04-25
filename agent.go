// Copyright 2020 Justice Nanhou. All rights reserved.
// license that can be found in the LICENSE file.

package chaos

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/JusticeN/chaos/globals"
	"github.com/JusticeN/chaos/message"
)

// Behaviour ...
type Behaviour func(*Agent) error

//DelayedBehaviour ...
type DelayedBehaviour func(*Agent, int) error

type timeBaseBehaviour struct {
	delay     int
	duration  time.Duration
	behaviour DelayedBehaviour
}
type setupFunc func(*Agent)
type takeDownFunc func(*Agent)

// Agent ...
type Agent struct {
	Name              string
	ID                string
	Setup             setupFunc
	TakeDown          takeDownFunc
	oneShotBehaviours []Behaviour
	cyclicBehaviours  []Behaviour
	tickerBehaviour   []timeBaseBehaviour
	wakerBehaviour    []timeBaseBehaviour
	cancelChan        chan struct{}
	dispatcher        *message.Dispatcher
	messagesChan      chan message.Message
	onMessage         func(message.Message)
}

// OnMessage add a callback function to receive message send to this agent
func (a *Agent) OnMessage(om func(message.Message)) {
	if a.onMessage != nil {
		return
	}
	a.onMessage = om
}

// Global state management

// SetData save key value pair
func (a *Agent) SetData(key string, data string) {
	globals.GetInstance().Set(a.ID, key, data)
}

// GetData get value for the given key
func (a *Agent) GetData(key string) string {
	if value := globals.GetInstance().Get(a.ID, key); value != "" {
		return value
	}
	return ""
}

// Send send message to another agent
func (a *Agent) Send(msg message.Message) {
	a.dispatcher.Publish(msg)
}

// SendMessage send Inform Message
func (a *Agent) SendMessage(receiver string, msgBody string) {
	msg := message.Message{
		Type:     message.TypeInform,
		Sender:   a.Name,
		Receiver: receiver,
		Content:  msgBody,
	}
	a.Send(msg)
}

func (a *Agent) doSetup() {
	log.Printf("--------- Setup agent        : [%v] ----------", a.Name)
	// configure message dispatcher
	a.dispatcher = message.GetDispatcherInstance()
	a.messagesChan = message.GetDispatcherInstance().Subscribe(a.Name)
	//
	if a.Setup != nil {
		a.Setup(a)
	}
	// fmt.Printf("%v a.cyclicBehaviours count: %v\n", a.Name, len(a.cyclicBehaviours))
	return
}

func (a *Agent) doTakeDown() {
	log.Printf("--------- Takedown agent     : [%v] ----------", a.Name)
	if a.TakeDown != nil {
		a.TakeDown(a)
	}
	return
}

//AddOneShotBehaviour add new behaviour to the agent
func (a *Agent) AddOneShotBehaviour(b Behaviour) {
	a.oneShotBehaviours = append(a.oneShotBehaviours, b)
}

//AddTickerBehaviour add new ticker behaviour to the agent
// waker behaviours execute repetitively after a given period of time
func (a *Agent) AddTickerBehaviour(b DelayedBehaviour, tickerSecond int) {
	tb := timeBaseBehaviour{
		delay:     tickerSecond,
		duration:  time.Duration(tickerSecond) * time.Second,
		behaviour: b,
	}
	a.tickerBehaviour = append(a.tickerBehaviour, tb)
}

//AddWakerBehaviour add new waker behaviour to the agent
// waker behaviours execute after a given timeout
func (a *Agent) AddWakerBehaviour(b DelayedBehaviour, timeoutSecond int) {
	tb := timeBaseBehaviour{
		delay:     timeoutSecond,
		duration:  time.Duration(timeoutSecond) * time.Second,
		behaviour: b,
	}
	a.wakerBehaviour = append(a.wakerBehaviour, tb)
}

//AddCyclicBehaviour add new cyclic behaviour to the agent
// cyclic behaviours never complete and run forever
func (a *Agent) AddCyclicBehaviour(b Behaviour) {
	a.cyclicBehaviours = append(a.cyclicBehaviours, b)
}

// Start start running the agent
func (a *Agent) stop() {
	a.cancelChan <- struct{}{}
}

// Start start running the agent
func (a *Agent) start() {
	log.Printf("--------- Start agent        : [%v] ----------", a.Name)
	var wg sync.WaitGroup

	// wakerBehaviour
	if len(a.wakerBehaviour) > 0 {
		wg.Add(len(a.wakerBehaviour))
		for _, action := range a.wakerBehaviour {
			go func(ch <-chan struct{}, actionBehaviour timeBaseBehaviour) {
				defer func() { wg.Done() }()
				select {
				case <-ch:
					return
				case <-time.After(actionBehaviour.duration):
					// fmt.Print(extime)
					err := actionBehaviour.behaviour(a, actionBehaviour.delay)
					if err != nil {
						log.Printf("somthing went wrong with Agent %v", a.Name)
						log.Println(err)
					}
					return
				}
			}(a.cancelChan, action)
		}
	}

	// tickerBehaviour
	if len(a.tickerBehaviour) > 0 {
		wg.Add(len(a.tickerBehaviour))
		for _, action := range a.tickerBehaviour {
			go func(ch <-chan struct{}, actionBehaviour timeBaseBehaviour) {
				defer func() { wg.Done() }()
				for {
					select {
					case <-ch:
						return
					case <-time.Tick(actionBehaviour.duration):
						// fmt.Print(extime)
						err := actionBehaviour.behaviour(a, actionBehaviour.delay)
						if err != nil {
							log.Printf("somthing went wrong with Agent %v", a.Name)
							log.Println(err)
						}
					}
				}
			}(a.cancelChan, action)
		}
	}
	// cyclicBehaviour
	if len(a.cyclicBehaviours) > 0 {
		wg.Add(len(a.cyclicBehaviours))
		for _, action := range a.cyclicBehaviours {
			go func(ch <-chan struct{}, actionBehaviour Behaviour) {
				defer func() { wg.Done() }()
				for {
					select {
					case <-ch:
						return
					default:
						err := actionBehaviour(a)
						if err != nil {
							log.Printf("somthing went wrong with Agent %v", a.Name)
							log.Println(err)
						}
					}
				}
			}(a.cancelChan, action)
		}
	}
	// mesage
	if a.onMessage != nil {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case msg := <-a.messagesChan:
					fmt.Printf("#%v\n", msg)
				case <-a.cancelChan:
					return
				default:
				}
			}
		}()
	}

	// one shot Behaviour
	if len(a.oneShotBehaviours) > 0 {
		for _, action := range a.oneShotBehaviours {
			err := action(a)
			if err != nil {
				log.Printf("somthing went wrong with Agent %v", a.Name)
				log.Println(err)
				a.stop()
			}
		}
	}

	wg.Wait()
	// a.doTakeDown()
}

//NewAgent create a new Simple Agent
func NewAgent(name string) {
	NewAgentWithSetupAndTakeDown(name, func(arg1 *Agent) {}, func(arg1 *Agent) {})
}

// NewAgentWithSetup create a new Agent with a setUp function.
// the setup function will be executed before starting the agent
func NewAgentWithSetup(name string, setUp setupFunc) {
	NewAgentWithSetupAndTakeDown(name, setUp, func(arg1 *Agent) {})
}

// NewAgentWithSetupAndTakeDown create a new Agent with a setUp and Takedown function.
// the setup function will be executed before starting the agent
// the takeDown function will be executed after all agent behaviour. before agent end of life
func NewAgentWithSetupAndTakeDown(name string, setUp setupFunc, takeDown takeDownFunc) {

	agent := Agent{
		Name:         name,
		Setup:        setUp,
		TakeDown:     takeDown,
		cancelChan:   make(chan struct{}, 1),
		dispatcher:   message.GetDispatcherInstance(),
		messagesChan: message.GetDispatcherInstance().Subscribe(name),
	}
	RegisterAgent(&agent)
}
