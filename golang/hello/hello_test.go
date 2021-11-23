package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("says hello to emili", func(t *testing.T) {
		got := Hello("Emili")
		want := "Hello, Emili"

		if got != want {
			t.Errorf("I wanted %q but I got %q", want, got)
		}
	})

	t.Run("says hello to Eldon", func(t *testing.T) {
		got := Hello("Eldon")
		want := "Hello, Eldon"

		if got != want {
			t.Errorf("I wanted %q but I got %q", want, got)
		}
	})
}
