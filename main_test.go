package main

import (
	"bytes"
	"io/ioutil"
	"testing"
)

var in string

func Test_ExecuteCommand(t *testing.T) {
	cmd := NewPinCmd()
	b := bytes.NewBufferString("")
	cmd.SetOutput(b)
	cmd.SetArgs([]string{"helloworld"})
	cmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	if string(out) != "4355696753\n" {
		t.Fatalf("expected \"%s\" got \"%s\"", "4355696753\n", string(out))
	}
}

func Test_ExecuteCommandVerbose(t *testing.T) {
	cmd := NewPinCmd()
	b := bytes.NewBufferString("")
	cmd.SetOutput(b)
	cmd.SetArgs([]string{"-v", "helloworld"})
	cmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	expectedOutput := "source: helloworld\n4355696753\n"
	if string(out) != expectedOutput {
		t.Fatalf("expected \"%s\" got \"%s\"", expectedOutput, string(out))
	}
}

func Test_ExecuteCommandIgnoreSpecial(t *testing.T) {
	cmd := NewPinCmd()
	b := bytes.NewBufferString("")
	cmd.SetOutput(b)
	cmd.SetArgs([]string{"-v", "-i", "hello-world"})
	cmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	expectedOutput := "source: hello-world\n43556-96753\n"
	if string(out) != expectedOutput {
		t.Fatalf("expected \"%s\" got \"%s\"", expectedOutput, string(out))
	}
}
