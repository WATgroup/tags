// SPDX-FileCopyrightText: © 2025 W-A-T EU Operations Oü
// SPDX-License-Identifier: EUPL-1.2 OR Proprietary
// SPDX-FileContributor: Created by Jose Luis Tallon <jltallon@w-a-t.group>
// Some functions based upon "slices", which is released under a BSD 3-Clause License

// Package tags implements a "Tag Collection"
package tags

const k_TAGCOLLSZ = 3

type Tagset []tag

func NewTagset() Tagset {
	return make(Tagset, 0, k_TAGCOLLSZ)
}

func FromStrings(tags ...string) (ret Tagset) {

	ret = make(Tagset, 0, len(tags))
	for _, t := range tags {
		ret = append(ret, tag(t))
	}
	return
}

func (ts *Tagset) AddString(tt string) error {
	t := tag(tt)
	if !t.Valid() {
		return tagError(`invalid tag ` + tt)
	}
	*ts = append(*ts, t)
	return nil
}

// Add 't' to the set; There can be duplicates
func (ts *Tagset) Add(t tag) error {
	if !t.Valid() {
		return tagError(`invalid tag ` + t)
	}
	*ts = append(*ts, t)
	return nil
}

func (ts *Tagset) Append(o Tagset) {
	for _, x := range o {
		*ts = append(*ts,x)
	}
}

// Remove occurences of "t" in the set
func (ts *Tagset) Remove(t tag) {
	s := ([]tag)(*ts)
	i := 0
	for ; i < len(s); i++ {
		if t == s[i] {
			copy(s[i:], s[i+1:])
			*ts = s[:len(s)-1] // reslice in-place; XXX: leaves extra capacity there
		}
	}
	return // might not have found the tag
}

func EqualSet(t1, t2 Tagset) bool {
	if len(t1) != len(t2) {
		return false
	}
	// 'tag' is comparable ...
	for i := range t1 {
		if t1[i] != t2[i] {
			return false
		}
	}
	return true
}

////////////////////////////////////////////////////////////////////////////////

func (ts Tagset) Len() int {
	return len(ts)
}

func (ts Tagset) IsEmpty() bool {
	return 0 == len(ts)
}

// Check if tagset contains a tag
// Linear search will be optimal, given that tagsets are tipically small
func (ts Tagset) Contains(t tag) bool {
	for i := range ts {
		if t == ts[i] {
			return true
		}
	}
	return false
}

func (ts Tagset) ContainsString(t string) bool {
	for _,x := range ts {
		if tag(t) == x {
			return true
		}
	}
	return false
}

// Index returns the index of the first occurrence of v in s,
// or -1 if not present.
func (ts Tagset) Index(t tag) int {
	for i,x := range ts {
		if t == x {
			return i
		}
	}
	return -1
}

////////////////////////////////////////////////////////////////////////////////

func (ts Tagset) Clone() (ret Tagset) {
	if 0 == len(ts) {
		return nil
	}
	ret = make(Tagset, len(ts))
// 	for i, t := range ts {
// 		ret[i] = t
// 	}
	copy(ret[:], ts[:])
	return // ret already contains result
}

// Remove extra space
func (ts *Tagset) Clip() {
	s := ([]tag)(*ts)
	*ts = s[:len(s):len(s)]
}

// ** Equivalent to "uniq"
// Compact modifies the contents of the slice s and returns the modified slice,
// which may have a smaller length.
// Compact zeroes the elements between the new length and the original length.
// The result preserves the nilness of s.
func (ts *Tagset) Compact() {
	if len(*ts) < 2 {
		return
	}

	s := *ts
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			s2 := s[i:]
			for k2 := 1; k2 < len(s2); k2++ {
				if s2[k2] != s2[k2-1] {
					s[i] = s2[k2]
					i++
				}
			}
			clear(s[i:]) // zero/nil out the obsolete elements, for GC
			*ts = s[:i]
		}
	}
	return
}
