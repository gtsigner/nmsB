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

func entriesFromString(s string) ([]PatternEntry, error) {
	// create the empty entries array
	entries := make([]PatternEntry, 0)
	// split the string on whitespaces
	groups := strings.Fields(s)
	// inspect each group
	for _, group := range groups {
		// check if the group is a wildcard
		if strings.ContainsAny(group, WildcardChar) {
			// append a wildcard entry
			entries = append(entries, PatternEntry{
				value:    byte(0),
				wildcard: true,
			})
		} else {
			// decode the group as hex
			bytes, err := hex.DecodeString(group)
			if err != nil {
				return nil, fmt.Errorf("fail to decode hex from signature group [ %s ], becuase %s", group, err.Error())
			}
			// check if the bytes not empty
			if len(bytes) < 1 {
				return nil, fmt.Errorf("fail to decode hex from signature group [ %s ], because decode bytes empty", group)
			}
			// append the byte entry
			entries = append(entries, PatternEntry{
				value:    bytes[0],
				wildcard: false,
			})
		}
	}

	return entries, nil
}

func FromString(s string) (*Pattern, error) {
	// extract the entries
	entries, err := entriesFromString(s)
	if err != nil {
		return nil, err
	}

	// create the pattern with the parst entries
	pattern := &Pattern{
		entries: entries,
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

func (pattern *Pattern) Length() int {
	return len(pattern.entries)
}

func (pattern *Pattern) MarshalJSON() ([]byte, error) {
	sigString := pattern.String()
	return []byte(sigString), nil
}

func (pattern *Pattern) UnmarshalJSON(bytes []byte) error {
	// make the bytes to string
	sigString := string(bytes)
	// make the string to entries
	entries, err := entriesFromString(sigString)
	if err != nil {
		return err
	}
	pattern.entries = entries

	return nil
}
