package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMain(m *testing.M) {
	// before unit test
	fmt.Print("before unit test running")

	m.Run()

	fmt.Print("after unit test running")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Fair")
	if result != "Hello Fair" {
		t.Error("Result keliru bosku, tulung sampeyan cehck maneh!")
	}
}

func TestHelloWorld2(t *testing.T) {
	result := HelloWorld("Bambang")
	assert.Equal(t, "Hello Bambang", result, "Result must be 'Hello Bambang'")
}
