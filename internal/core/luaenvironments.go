package core

import (
	"path"
	"strings"

	"github.com/apyrr/tlua/internal/tspath"
)

// LuaEnvironment describes globals a Lua host installs before it runs a group
// of chunks. Every path Root matches (a directory or a single file) is one
// group. In the group's files each name in Globals is an open table seeded with
// the named global type, and writes there add members to that group's table
// only. With Registry set (and exactly one global), each group also becomes a
// member of the named global interface, keyed by the group's last path segment
// and typed as the group's table, so a host's lookup-by-name functions can
// return the class a name denotes.
type LuaEnvironment struct {
	Root     string            `json:"root"`
	Globals  map[string]string `json:"globals"`
	Registry string            `json:"registry,omitzero"`
}

// LuaEnvironmentGroup is the group a file belongs to: the luaEnvironments
// entry that matched it, and the path that entry's root matched. Every file
// under one matched path shares one set of host globals.
type LuaEnvironmentGroup struct {
	// Entry indexes CompilerOptions.LuaEnvironments.
	Entry int
	// Path is the matched root path, relative to the Lua search root, with the
	// extension removed so a single-file group and a folder of the same name
	// agree. Its last segment names the group.
	Path string
}

// Name is the group's last path segment: the class name a host registers the
// group under.
func (group LuaEnvironmentGroup) Name() string {
	return group.Path[strings.LastIndex(group.Path, "/")+1:]
}

// MatchLuaEnvironmentGroup returns the group of a file, given its path relative
// to the Lua search root. For each entry the shortest path its root matches is
// the group's path; between entries the deepest group wins, so a nested scope
// (stools inside weapons) overrides the one around it, and a later entry wins a
// tie. Roots are normalized by option parsing.
func MatchLuaEnvironmentGroup(environments []*LuaEnvironment, relativePath string, useCaseSensitiveFileNames bool) (LuaEnvironmentGroup, bool) {
	segments := strings.Split(relativePath, "/")
	compared := segments
	if !useCaseSensitiveFileNames {
		compared = strings.Split(strings.ToLower(relativePath), "/")
	}
	best, bestDepth := -1, 0
	for index, environment := range environments {
		root := environment.Root
		if !useCaseSensitiveFileNames {
			root = strings.ToLower(root)
		}
		pattern := strings.Split(root, "/")
		for depth := 1; depth <= len(compared); depth++ {
			if matchLuaEnvironmentSegments(pattern, compared[:depth]) {
				if depth >= bestDepth {
					best, bestDepth = index, depth
				}
				break
			}
		}
	}
	if best < 0 {
		return LuaEnvironmentGroup{}, false
	}
	groupPath := strings.Join(segments[:bestDepth], "/")
	if bestDepth == len(segments) {
		groupPath = tspath.RemoveFileExtension(groupPath)
	}
	return LuaEnvironmentGroup{Entry: best, Path: groupPath}, true
}

// matchLuaEnvironmentSegments matches glob segments against path segments:
// `**` spans any number of segments, and other segments match one path segment
// with path.Match wildcards.
func matchLuaEnvironmentSegments(pattern []string, segments []string) bool {
	if len(pattern) == 0 {
		return len(segments) == 0
	}
	if pattern[0] == "**" {
		for skip := 0; skip <= len(segments); skip++ {
			if matchLuaEnvironmentSegments(pattern[1:], segments[skip:]) {
				return true
			}
		}
		return false
	}
	if len(segments) == 0 {
		return false
	}
	// Roots are validated when the option is parsed, so the pattern is well formed.
	if matched, _ := path.Match(pattern[0], segments[0]); !matched { //nolint:forbidigo // glob segment matching, not path handling
		return false
	}
	return matchLuaEnvironmentSegments(pattern[1:], segments[1:])
}
