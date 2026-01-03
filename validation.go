// SPDX-FileCopyrightText: © 2025 W-A-T EU Operations Oü
// SPDX-License-Identifier: EUPL-1.2 OR Proprietary
// SPDX-FileContributor: Created by Jose Luis Tallon <jltallon@w-a-t.group>

package tags

import (
	"unicode/utf8"
)

// Valid checks that a tag conforms to the spec for tags...
// ... that is, only ASCII alphanumeric, _/-/+
func (t tag) Valid() bool {

	for i, c := range t { // iterates as runes
		if c >= utf8.RuneSelf {
			return false
		}
		if _isAlnum(c) {
			continue
		}
		if i > 1 && _isValidSym(byte(c)) {
			continue
		}
		return false
	}
	return true
}

// XXX: Explore speeding up validation via table lookup
func _isAlnum(c rune) bool {
	return (c >= '0' && c <= '9') || (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z')
}

func _isValidSym(b byte) bool {
	return ('_' == b || '+' == b || '-' == b)
}
