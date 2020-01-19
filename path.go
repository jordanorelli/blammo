package blammo

import (
	"fmt"
	"strings"
	"unicode"
)

// Path provides the basis of a hierarchical pathing system. Each path is a
// singly linked list of strings. This is most often used to form the prefix
// for a log line in a log file.
//
// Path is immutable by design. This allows us to make assumptions about a
// given path. We know, for example, that because at each step the name cannot
// be changed, the name of a given path's parent is never changed.
type Path struct {
	name   string
	parent *Path
}

// Name gives the tail member of a path. This is analagous to the notion of a
// file path's basename, or the last element of a linked list.
func (p *Path) Name() string { return p.name }

// Parent gives the parent path of this path. The expected use-case for
// retrieving a path's parent is to facilitate the construction of custom
// blammo.Writer implementations.
func (p *Path) Parent() *Path { return p.parent }

// String represents a canonical stringification of a path, represented as a
// slash-delimited absolute path similar to a filesystem path.
func (p *Path) String() string {
	if p.parent == nil {
		return p.name
	}
	return strings.TrimLeft(fmt.Sprintf("%s/%s", p.parent.String(), p.name), "/")
}

// Child creates a child path having parent p. This is the recommended way of
// constructing a hierarchical path.
func (p *Path) Child(name string) *Path {
	if name == "" {
		name = "-"
	}
	return &Path{
		name:   MakeSafeName(name),
		parent: p,
	}
}

// NewPath creates a new path with a given root name. The supplied name string
// is cleaned of unsafe characters.
func NewPath(name string) *Path {
	return &Path{name: MakeSafeName(name)}
}

// IsSafeName determines whether a provided path name is considered to be a
// "safe" name by the standards of blammo. A safe name in the context of blammo
// is a name that consists of only unicode letters and numbers, plus the hyphen
// (-), underscore (_), and colon (:) characters. Note that the character
// classes being tested against are unicode letters and numbers, not ascii
// letters and numbers; letters with accents and letters that do not appear in
// English are permitted.
//
// The goal of the safe name checker is to ensure that logs written by
// blammo can be written in any (human) language while maintaining a few rules
// to ensure the logs can be reasonably straightforward to parse and search.
func IsSafeName(name string) bool {
	runes := []rune(name)
	for _, r := range runes {
		if r == '-' || r == '_' || r == ':' {
			continue
		}
		if !unicode.In(r, unicode.Letter, unicode.Number) {
			return false
		}
	}
	return true
}

// MakeSafeName takes a string and transforms it, if necessary, into a string
// that is considered to be a safe name. The transformation strips all leading
// and trailing whitespace, converts intermediate spacing characters into
// underscores, and converts other unsafe characters into hyphens.
func MakeSafeName(name string) string {
	if IsSafeName(name) {
		return name
	}
	name = strings.TrimSpace(name)
	runes := []rune(name)
	out := make([]rune, 0, len(runes))
	for _, r := range runes {
		if r == '-' || r == '_' || r == ':' {
			continue
		}
		if unicode.In(r, unicode.Letter, unicode.Number) {
			out = append(out, r)
			continue
		}
		if unicode.IsSpace(r) {
			out = append(out, '_')
			continue
		}
		out = append(out, '-')
	}
	return string(out)
}
