package util

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestCustomLogFormatter_GetLogger(t *testing.T) {
	type args struct {
		app      string
		logLevel string
	}
	tests := []struct {
		name       string
		f          *CustomLogFormatter
		args       args
		want       *logrus.Entry
		wantWriter string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := &bytes.Buffer{}
			if got := tt.f.GetLogger(writer, tt.args.app, tt.args.logLevel); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CustomLogFormatter.GetLogger() = %v, want %v", got, tt.want)
			}
			if gotWriter := writer.String(); gotWriter != tt.wantWriter {
				t.Errorf("CustomLogFormatter.GetLogger() = %v, want %v", gotWriter, tt.wantWriter)
			}
		})
	}
}

func TestNewCustomLogFormatter(t *testing.T) {
	tests := []struct {
		name string
		want *CustomLogFormatter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCustomLogFormatter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCustomLogFormatter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCustomLogFormatter_Format(t *testing.T) {
	type args struct {
		entry *logrus.Entry
	}
	tests := []struct {
		name    string
		f       *CustomLogFormatter
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.f.Format(tt.args.entry)
			if (err != nil) != tt.wantErr {
				t.Errorf("CustomLogFormatter.Format() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CustomLogFormatter.Format() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCustomLogFormatter_getCaller(t *testing.T) {
	type args struct {
		skip int
	}
	tests := []struct {
		name         string
		f            *CustomLogFormatter
		args         args
		wantLine     int
		wantFileName string
		wantFuncName string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLine, gotFileName, gotFuncName := tt.f.getCaller(tt.args.skip)
			if gotLine != tt.wantLine {
				t.Errorf("CustomLogFormatter.getCaller() gotLine = %v, want %v", gotLine, tt.wantLine)
			}
			if gotFileName != tt.wantFileName {
				t.Errorf("CustomLogFormatter.getCaller() gotFileName = %v, want %v", gotFileName, tt.wantFileName)
			}
			if gotFuncName != tt.wantFuncName {
				t.Errorf("CustomLogFormatter.getCaller() gotFuncName = %v, want %v", gotFuncName, tt.wantFuncName)
			}
		})
	}
}

func TestLogger(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
		want *logrus.Entry
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Logger(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Logger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithLogger(t *testing.T) {
	type args struct {
		ctx   context.Context
		entry *logrus.Entry
	}
	tests := []struct {
		name string
		args args
		want context.Context
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithLogger(tt.args.ctx, tt.args.entry); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithLogger() = %v, want %v", got, tt.want)
			}
		})
	}
}
