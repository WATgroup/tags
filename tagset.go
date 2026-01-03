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

func (x *Tagset) AddString(tt string) error {
	t := tag(tt)
	if !t.Valid() {
		return tagError(`invalid tag ` + tt)
	}
	*x = append(*x, t)
	return nil
}

// Add 't' to the set; There can be duplicates
func (x *Tagset) Add(t tag) error {
	if !t.Valid() {
		return tagError(`invalid tag ` + t)
	}
	*x = append(*x, t)
	return nil
}

// Remove first occurence of "t" in the set
func (x *Tagset) Remove(t tag) {
	s := ([]tag)(*x)
	i := 0
	for ; i < len(s); i++ {
		if t == s[i] {
			copy(s[i:],s[i+1:])
			break
		}
	}
	*x = s[:len(s)-1] // reslice in-place; XXX: leaves extra capacity there
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

//////////////////////////////////////////////////////

func (x Tagset) Clone() (ret Tagset) {
	if 0 == len(x) {
		return nil
	}
	
	ret = make(Tagset,len(x))
	for i,t := range x {
		ret[i] = t
	}
	return // ret already contains result
}

// Remove extra space
func (x *Tagset) Clip() {
	s := ([]tag)(*x)
	*x = s[:len(s):len(s)]
}

// ** Equivalent to "uniq"
// Compact modifies the contents of the slice s and returns the modified slice,
// which may have a smaller length.
// Compact zeroes the elements between the new length and the original length.
// The result preserves the nilness of s.
func (x *Tagset) Compact() {
	if len(*x) < 2 {
		return
	}

	s := *x
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
			*x = s[:i]
		}
	}
	return
}

// Index returns the index of the first occurrence of v in s,
// or -1 if not present.
func (x Tagset) Index(t tag) int {
	for i := range x {
		if t == x[i] {
			return i
		}
	}
	return -1
}

func (x Tagset) Len() int {
	return len(x)
}

func (x Tagset) IsEmpty() bool {
	return 0 == len(x)
}
