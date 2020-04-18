// Copyright 2020 Justice Nanhou. All rights reserved.
// license that can be found in the LICENSE file.

package chaos

import (
	"sync"
)

// AgentStore ...
type AgentStore struct {
	agents map[string]*Agent
	sync.RWMutex
}

var instance *AgentStore
var once sync.Once

// GetAgentStoreInstance ...
func GetAgentStoreInstance() *AgentStore {
	once.Do(func() {
		instance = &AgentStore{
			agents: make(map[string]*Agent, 10),
		}
	})
	return instance
}

// Add ...
func (s *AgentStore) Add(a *Agent) {
	s.Lock()
	defer s.Unlock()
	s.agents[a.Name] = a
}

// List ...
func (s *AgentStore) List() map[string]*Agent {
	return s.agents
}

// Count ...
func (s *AgentStore) Count() int {
	s.RLock()
	defer s.RUnlock()
	return len(s.agents)
}
