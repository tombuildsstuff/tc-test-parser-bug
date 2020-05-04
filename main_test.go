package main

import (
	"log"
	"os"
	"testing"
)

func TestPassing(t *testing.T) {
	os.Stderr.WriteString("stderr - from the passing test\n")
	log.Printf("stdout - from the passing test")
	t.Log("This should always pass")
}

func TestSkipped(t *testing.T) {
	os.Stderr.WriteString("stderr - from the skipped test\n")
	log.Printf("stdout - from the skipped test")
	t.Log("This should be skipped")
	t.Skip("Skipping because we should skip it")
}

func TestFailed(t *testing.T) {
	os.Stderr.WriteString("stderr - from the failed test\n")
	log.Printf("stdout - from the failed test")
	t.Log("This should always fail")
	t.Fail()
}