package main_test

import (
	"fmt"
	"testing"
)

type Map map[string]string

var tests = []struct {
	existing Map
	input    Map
	expect   Map
}{
	{
		existing: Map{},
		input:    Map{},
		expect:   Map{},
	},
}

func Test(t *testing.T) {
	fmt.Println("TODO")
}
