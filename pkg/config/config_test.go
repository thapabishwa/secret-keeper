package config

import (
	"reflect"
	"testing"
)

func TestNewConfig(t *testing.T) {
	tests := []struct {
		name string
		want *Config
	}{
		{
			name: "TestNewConfig",
			want: &Config{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
