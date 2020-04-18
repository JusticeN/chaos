// Copyright 2020 Justice Nanhou. All rights reserved.
// license that can be found in the LICENSE file.

package message

import "sync"

// Dispatcher ...
type Dispatcher struct {
	mu     sync.RWMutex
	subs   map[string][]chan Message
	closed bool
}

var dispatcherInstance *Dispatcher
var once sync.Once

// GetDispatcherInstance ...
func GetDispatcherInstance() *Dispatcher {
	once.Do(func() {
		dispatcherInstance = &Dispatcher{}
		dispatcherInstance.subs = make(map[string][]chan Message)
	})
	return dispatcherInstance
}

// Subscribe ...
func (ps *Dispatcher) Subscribe(agentID string) chan Message {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	ch := make(chan Message)
	ps.subs[agentID] = append(ps.subs[agentID], ch)
	return ch
}

// Publish ...
func (ps *Dispatcher) Publish(msg Message) {
	ps.mu.RLock()
	defer ps.mu.RUnlock()

	if ps.closed {
		return
	}
	for _, ch := range ps.subs[msg.Receiver] {
		// go func(ch chan Message) {
		// 	ch <- msg
		// }(ch)
		ch <- msg
	}
}

// Close ...
func (ps *Dispatcher) Close() {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	if !ps.closed {
		ps.closed = true
		for _, subs := range ps.subs {
			for _, ch := range subs {
				close(ch)
			}
		}
	}
}
