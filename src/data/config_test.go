package data

import (
	"reflect"
	"testing"
)

func TestNewConfig(t *testing.T) {
	tests := []struct {
		name string
		want *Config
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_Init(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		cfg     *Config
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.cfg.Init(tt.args.filePath); (err != nil) != tt.wantErr {
				t.Errorf("Config.Init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_processFile(t *testing.T) {
	type args struct {
		filePath string
		cfg      *Config
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := processFile(tt.args.filePath, tt.args.cfg); (err != nil) != tt.wantErr {
				t.Errorf("processFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_processEnv(t *testing.T) {
	type args struct {
		cfg *Config
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := processEnv(tt.args.cfg); (err != nil) != tt.wantErr {
				t.Errorf("processEnv() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
