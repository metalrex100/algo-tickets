package main

import (
	"fmt"
	"github.com/metalrex100/algo-tester"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"runtime"
	"testing"
)

func TestStrLenTask(t *testing.T) {
	err := algo_tester.RunTests(t, getStrLenTask(), fmt.Sprintf("%s/test-data/0.String", getPathToCurrentDir()))

	assert.NoError(t, err)
}

func TestLuckyTicketsTask(t *testing.T) {
	err := algo_tester.RunTests(t, getLuckyTicketsTask(), fmt.Sprintf("%s/test-data/1.Tickets", getPathToCurrentDir()))

	assert.NoError(t, err)
}

func getPathToCurrentDir() string {
	_, b, _, _ := runtime.Caller(0)

	return filepath.Dir(b)
}
