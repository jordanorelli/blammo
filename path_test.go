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

	p = p.Child(" dave ")
	if p.String() != "alice/bob/carol/dave" {
		t.Error("bad sanitation transformation")
	}
}

func TestSafeNames(t *testing.T) {
	safeNames := []string{
		"one",
		"1",
		"1one",
		"niño",
		"garçon",
		"alice-bob",
		"alice_bob",
		"alice:bob",
		"你好",
		// this string contains a unicode zero-width non-joiner character. Not
		// sure how I feel about this being considered safe. On the one hand
		// it's necessary for some languages, on the other hand it has the
		// propensity to create confusing homoglyph situations.
		string([]rune{'o', 0x8204, 'n', 'e'}),
		"",
	}

	for _, n := range safeNames {
		if !IsSafeName(n) {
			t.Errorf("expected safe name is considered unsafe: %s", n)
		}
	}

	unsafeNames := []string{
		" one",
		"one ",
		"alice/bob",
		"alice bob",
	}
	for _, n := range unsafeNames {
		if IsSafeName(n) {
			t.Errorf("expected unsafe name is considered safe: %s", n)
		}
	}
}
