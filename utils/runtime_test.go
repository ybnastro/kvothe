package utils

import (
	"testing"
)

func TestUtils_GetFuncName(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "normal case",
			want: "func1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFuncName(); got != tt.want {
				t.Errorf("GetFuncName() = %v, want %v", got, tt.want)
			}
		})
	}
}
