package main

import "testing"

func TestSimpleFactory(t *testing.T) {
	got := SimpleFactory("example.com:4444")

	if got.Name != "Hello" {
		t.Errorf("Name = %q, want %q", got.Name, "Hello")
	}
	if got.Description != "World" {
		t.Errorf("Description = %q, want %q", got.Description, "World1")
	}
	if got.Url != "example.com:4444" {
		t.Errorf("Url = %q, want %q", got.Url, "example.com:4444")
	}
}
