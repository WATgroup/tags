// SPDX-FileCopyrightText: © 2025 W-A-T EU Operations Oü
// SPDX-License-Identifier: EUPL-1.2 OR Proprietary
// SPDX-FileContributor: Created by Jose Luis Tallon <jltallon@w-a-t.group>

package tags

import "strings"

const (
	k_TAGSET_OUT_BEGIN = '['
	k_TAGSET_OUT_SEP   = ','
	k_TAGSET_OUT_END   = ']'
)

// Implement the "stringer" interface
// ..also needed for the flags.Value interface
func (ts Tagset) String() string {
	var sb strings.Builder
	sb.WriteRune(k_TAGSET_OUT_BEGIN) // more efficient than "["

	if len(ts) > 0 {
		sb.WriteString(string(ts[0]))
	}
	if len(ts) > 1 {
		for _, x := range ts[1:] {
			sb.WriteRune(k_TAGSET_OUT_SEP)
			sb.WriteString(string(x))
		}
	}
	sb.WriteRune(k_TAGSET_OUT_END)

	return sb.String()
}

func Parse(in string) (Tagset, error) {

	x := in
	if k_TAGSET_OUT_BEGIN == in[0] {
		// they used a delimited version... well (maybe it's our own output)
		if k_TAGSET_OUT_END != in[len(in)-1] {
			return nil, tagError("Invalid tagset")
		} else {
			x = in[1 : len(in)-1] // strip delimiters
		}
	}
	// XXX: Consider reimplementing, optimized, in terms of "Cut" (index+re-slice)
	parts := strings.Split(x, ",") // XXX: FIXME: k_TAGSET_OUT_SEP

	ret := make(Tagset, len(parts))
	for i, t := range parts {
		if v := tag(t); !v.Valid() {
			return nil, tagError("Invalid tag value '" + t + "' encountered")
		} else {
			ret[i] = v
		}
	}
	return ret, nil
}
