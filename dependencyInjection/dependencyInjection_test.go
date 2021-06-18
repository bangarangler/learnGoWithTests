package dependencyInjection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Jon")

	got := buffer.String()
	want := "Hello, Jon"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
