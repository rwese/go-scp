package scp

import (
	"reflect"
	"testing"
	"time"

	"golang.org/x/crypto/ssh"
)

func TestNewClient(t *testing.T) {
	type args struct {
		host   string
		config *ssh.ClientConfig
	}
	tests := []struct {
		name string
		args args
		want Client
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClient(tt.args.host, tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewClientWithTimeout(t *testing.T) {
	type args struct {
		host    string
		config  *ssh.ClientConfig
		timeout time.Duration
	}
	tests := []struct {
		name string
		args args
		want Client
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClientWithTimeout(tt.args.host, tt.args.config, tt.args.timeout); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClientWithTimeout() = %v, want %v", got, tt.want)
			}
		})
	}
}
