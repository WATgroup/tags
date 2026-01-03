// SPDX-FileCopyrightText: © 2025 W-A-T EU Operations Oü
// SPDX-License-Identifier: EUPL-1.2 OR Proprietary
// SPDX-FileContributor: Created by Jose Luis Tallon <jltallon@w-a-t.group>

package tags


type tag string // private type, to force via constructor

func New(t string) tag {
	// XXX: TODO: validate
	return tag(t)
}



////////////////////////////////////////////////////////////////////////////////
type tagError string

func (e tagError) Error() string {
	return string(e)
}
