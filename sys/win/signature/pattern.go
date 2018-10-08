package signature

import (
	"encoding/hex"
	"fmt"
	"strings"
)

const (
	WildcardChar = "?"
)

type PatternEntry struct {
	wildcard bool
	value    byte
}

type Pattern struct {
	entries []PatternEntry
}

func FromString(s string) (*Pattern, error) {
	pattern := &Pattern{
		entries: make([]PatternEntry, 0),
	}

	groups := strings.Fields(s)

	for _, group := range groups {
		if strings.ContainsAny(group, WildcardChar) {
			pattern.entries = append(pattern.entries, PatternEntry{
				value:    byte(0),
				wildcard: true,
			})
		} else {
			bytes, err := hex.DecodeString(group)
			if err != nil {
				return nil, fmt.Errorf("fail to decode hex from signature group [ %s ], becuase %s", group, err.Error())
			}

			if len(bytes) < 1 {
				return nil, fmt.Errorf("fail to decode hex from signature group [ %s ], because decode bytes empty", group)
			}
			pattern.entries = append(pattern.entries, PatternEntry{
				value:    bytes[0],
				wildcard: false,
			})
		}
	}

	return pattern, nil
}

func (pattern *Pattern) String() string {
	var values []string
	for _, e := range pattern.entries {
		if e.wildcard {
			values = append(values, strings.Repeat(WildcardChar, 2))
		} else {
			values = append(values, fmt.Sprintf("%02X", e.value))
		}
	}
	return strings.Join(values, " ")
}

func (pattern *Pattern) MatchAt(index int, b byte) bool {
	if index >= len(pattern.entries) {
		return false
	}
	entry := pattern.entries[index]
	if entry.wildcard {
		return true
	}
	return entry.value == b
}
