# Chaos

##### create an agent
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

##### create an agent with Behaviour
for each agent a behaviour can be added in the setup method
there are different type of beahviours
## one shot behaviour
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
## Cylic behaviour
## waker Behabiour
## Ticker Behaviour


#### Agent can send Message

#### Agents have global state