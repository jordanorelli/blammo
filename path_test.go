package blammo

import (
	"testing"
)

func TestPath(t *testing.T) {
	p := NewPath("alice")
	if p.String() != "alice" {
		t.Error("bad root path generation")
	}

	p = p.Child("bob")
	if p.String() != "alice/bob" {
		t.Error("bad child path generation")
	}

	p = p.Child("carol")
	if p.String() != "alice/bob/carol" {
		t.Error("bad grandchild generation")
	}
}

func TestSafeNames(t *testing.T) {
	safeNames := []string{
		"one",
		"1",
		"1one",
		"niño",
		"garçon",
		"你好",
		"",
	}

	for _, n := range safeNames {
		if !IsSafeName(n) {
			t.Errorf("expected safe name is considered unsafe: %s", n)
		}
	}
}
