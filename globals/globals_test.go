// Copyright 2020 Justice Nanhou. All rights reserved.
// license that can be found in the LICENSE file.

package globals

import (
	"reflect"
	"testing"
)

func TestGetInstance(t *testing.T) {
	tests := []struct {
		name string
		want *GlobalStore
	}{
		{"singleton instance", GetInstance()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetInstance(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetInstance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGlobalStore_Get(t *testing.T) {
	type args struct {
		agentID string
		key     string
	}
	store := GetInstance()
	store.store.Store("agent1", Data{"key1": "value1"})
	store.store.Store("agent2", Data{"key2": "value2"})
	tests := []struct {
		name string
		gs   *GlobalStore
		args args
		want string
	}{
		// TODO: Add test cases.
		{"get first item", GetInstance(), args{"agent1", "key1"}, "value1"},
		{"get second item", GetInstance(), args{"agent2", "key2"}, "value2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.gs.Get(tt.args.agentID, tt.args.key); got != tt.want {
				t.Errorf("GlobalStore.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGlobalStore_Set(t *testing.T) {
	type args struct {
		agentID string
		key     string
		value   string
	}
	store := GetInstance()
	tests := []struct {
		name string
		gs   *GlobalStore
		args args
	}{
		{"setValue1", store, args{"agent1", "key1", "value1"}},
		{"setValue2", store, args{"agent2", "key2", "value2"}},
		{"setValue3", store, args{"agent3", "key3", "value3"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.gs.Set(tt.args.agentID, tt.args.key, tt.args.value)
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.gs.Get(tt.args.agentID, tt.args.key); got != tt.args.value {
				t.Errorf("GlobalStore.Get() = %v, want %v", got, tt.args.value)
			}
		})

	}
}
