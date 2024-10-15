package revisit

import "testing"

func TestRevisit(t *testing.T) {
	want := "hello world"
	if got := Revisit("world"); got != want {
		t.Errorf("revisit() = %q, want %q", got, want)
	}
}
