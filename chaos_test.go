// Copyright 2020 Justice Nanhou. All rights reserved.

// license that can be found in the LICENSE file.

package chaos

import (
	"testing"
)

func TestRegisterAgent(t *testing.T) {
	type args struct {
		a *Agent
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"adding agent", args{&Agent{Name: "agent1", Setup: nil}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RegisterAgent(tt.args.a)
			if got := GetAgentStoreInstance().Count(); got != tt.want {
				t.Errorf("RegisterAgent(..) saved = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_genID(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"call genID", genID()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := genID(); got == tt.want {
				t.Errorf("genID() = %v, want %v", got, tt.want)
				lenGot := len(got)
				lenWant := len(tt.want)
				if lenGot != lenWant {
					t.Errorf("len of = %v, and %v missmatch", got, tt.want)
				}
			}
		})
	}
}
