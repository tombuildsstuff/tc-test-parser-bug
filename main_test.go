package main

import "testing"

func TestPassing(t *testing.T) {
	t.Log("This should always pass")
}

func TestSkipped(t *testing.T) {
	t.Log("This should be skipped")
	t.Skip("Skipping because we should skip it")
}

func TestFailed(t *testing.T) {
	t.Log("This should always fail")
	t.Fail()
}