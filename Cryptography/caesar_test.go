package main

import "testing"

func Test_encrypt(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encrypt(tt.args.input)
		})
	}
}
