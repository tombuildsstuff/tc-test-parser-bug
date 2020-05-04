package main

import (
	"log"
	"testing"
)

func TestPassing(t *testing.T) {
	log.Printf("stdout - from the passing test")
	t.Log("This should always pass")
}

func TestSkipped(t *testing.T) {
	log.Printf("stdout - from the skipped test")
	t.Log("This should be skipped")
	t.Skip("Skipping because we should skip it")
}

func TestFailed(t *testing.T) {
	log.Printf("stdout - from the failed test")
	t.Log("This should always fail")
	t.Fail()
}