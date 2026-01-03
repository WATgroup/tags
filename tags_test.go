// SPDX-FileCopyrightText: © 2025 W-A-T EU Operations Oü
// SPDX-License-Identifier: EUPL-1.2
// SPDX-FileContributor: Created by Jose Luis Tallon <jltallon@w-a-t.group>

package tags_test

import (
	"testing"
 	"github.com/WATgroup/assert"
	"github.com/WATgroup/tags"
//  	"fmt"
)

var g_tagset = tags.FromStrings("one","two","three","four","five","six")


func TestBasics(t *testing.T) {
	
	t.Run("emptyness0", func(t *testing.T){
		ts0 := tags.NewTagset()
		assert.True(t, ts0.IsEmpty())
	})
	
	
	one := tags.New("one")
	ts1 := tags.NewTagset()
	ts1.AddString("one")
		
	t.Run("buildSet", func(t *testing.T) {
		
// 		fmt.Println(ts1)
	
		ts2 := tags.NewTagset()
		ts2.Add(one)
// 		fmt.Println(ts2)
	
		assert.True(t, tags.EqualSet(ts1,ts2))
	})
	
	t.Run("remove1", func(t *testing.T) {
		ts1.Remove("one")
// 		fmt.Println(ts1)
		assert.True(t, ts1.IsEmpty())
	})
}

func TestValidation(t *testing.T) {

	tt := []string{`group_W-A-T_2025+test`, `dell-sas`, `asus_u2nvme`, `cometnas+2025`}
	for _,x := range tt {
		ti := tags.New(x)
		assert.Valid(t,ti)
	}
	
	tf1 := tags.New("-invalidTag")
	assert.False(t, tf1.Valid())
	
	tf2 := tags.New("com.invalid.tag")
	assert.False(t, tf2.Valid())
	
}

func TestSorting(t *testing.T) {
	
	ts := g_tagset.Clone()
	ts.Sort()
// 	fmt.Println(ts)

	ts.Add(tags.New("eight"))
// 	fmt.Println(ts)
	
	assert.False(t, ts.IsSorted())
	
	ts.Add(tags.New("f1ve"))
	
	ts.Sort()
// 	fmt.Println(ts)
	
	assert.True(t, ts.IsSorted())
	
	ts2 := tags.FromStrings(`eight`,`f1ve`,`five`,`four`,`one`,`six`,`three`,`two`)
// 	fmt.Println(ts2)
	
	assert.True(t, tags.EqualSet(ts,ts2))
}

func TestRemovals(t *testing.T) {
	
	ts := g_tagset.Clone()
// 	fmt.Println(ts)
	
	ts.Sort()
// 	fmt.Println(ts)
	
	ts.Remove(tags.New("five"))
// 	fmt.Println(ts)
	ts.Remove(tags.New("six"))
// 	fmt.Println(ts)
	
	assert.True(t, tags.EqualSet(ts, tags.FromStrings("four","one","three","two")))
}
