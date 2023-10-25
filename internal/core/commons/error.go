package commons

import "strings"

const duplicatedEntryMsgSubstring = "duplicate entry"

func DuplicatedEntryError(err error) bool {
	return strings.Contains(strings.ToLower(err.Error()), duplicatedEntryMsgSubstring)
}
