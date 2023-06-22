package http

import (
	"testing"
)

func TestToSnakeCase(t *testing.T) {
	tests := []struct {
		Name string
	}{{"Name"},{"age"},{"myName"},{"MyName"}}

	for _,val:=range tests{
		name := toSnakeCase(val.Name)
		t.Log(name)
	}
}