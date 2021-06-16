package service

import (
	"cryptoserver-clean-app/data"
	"reflect"
	"testing"

	"github.com/newrelic/go-agent/v3/newrelic"
)

func TestNewCryptoServerSvc(t *testing.T) {
	type args struct {
		config      *data.Config
	}
	tests := []struct {
		name string
		args args
		want Service
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCryptoServerSvc(tt.args.config, tt.args.newrelicApp, tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCryptoServerSvc() = %v, want %v", got, tt.want)
			}
		})
	}
}
