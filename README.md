# Chaos
chaos is a multiagent application Inspired by the jade library. https://en.wikipedia.org/wiki/Java_Agent_Development_Framework 
i try to build a POC to use the concurency model of Go and make the developement easier.
##### How to create an agent
a simple way to create an Agent without any behaviour
```golang
// main
func main() {
     neo := &chaos.Agent{
		Name:  "Neo",
		Setup: func(agent *chaos.Agent) {},
	}
	chaos.RegisterAgent(neo)
	chaos.Start()
}

```

##### Create an agent with Behaviour
for each agent a behaviour can be added in the setup method
there are different type of beahviours
###### one shot behaviour
one shot behaviour have these signature ` func(agent *chaos.Agent) error`
and are added like this `agent.AddOneShotBehaviour(...)`

```go
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

```
###### Cylic behaviour
TODO: in progress
###### Waker Behabiour
TODO: in progress
###### Ticker Behaviour
TODO: in progress

#### Agent can send Message
TODO: in progress
#### Agents have global state
TODO: in progress
