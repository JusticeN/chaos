// Copyright 2020 Justice Nanhou. All rights reserved.
// license that can be found in the LICENSE file.

package globals

import (
	"sync"
)

// Data map representing kv stored for each Agent
type Data map[string]string

// GlobalStore responsible ofsaving agents global state like variables in form of key values pairs
type GlobalStore struct {
	store sync.Map
}

var store *GlobalStore
var once sync.Once

// GetInstance return the current using store instance
func GetInstance() *GlobalStore {
	once.Do(func() {
		store = &GlobalStore{}
	})
	return store
}

// Get return a value for a given Agent key
func (gs *GlobalStore) Get(agentID string, key string) string {
	if data, ok := gs.store.Load(agentID); ok {
		val := data.(Data)
		value := val[key]
		return value
	}
	return ""
}

// Set add a new key value pair in the stare
func (gs *GlobalStore) Set(agentID string, key string, value string) {
	// data := value.(string)
	var val Data
	if data, ok := gs.store.Load(agentID); ok {
		val = data.(Data)
	} else {
		val = make(Data, 0)
	}
	val[key] = value
	gs.store.Store(agentID, val)
}
