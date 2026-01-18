// SPDX-FileCopyrightText: © 2025 W-A-T EU Operations Oü
// SPDX-License-Identifier: EUPL-1.2
// SPDX-FileContributor: Created by Jose Luis Tallon <jltallon@w-a-t.group>
// SPDX-FileContributor: Roberto Fernandez <rfernandez@quadrics.es>

package tags_test

import (
	"fmt"
	"testing"

	"github.com/WATgroup/assert"
	"github.com/WATgroup/tags"
)

var g_tagset = tags.FromStrings("one", "two", "three", "four", "five", "six")

func TestBasics(t *testing.T) {

	t.Run("emptyness0", func(t *testing.T) {
		ts0 := tags.NewTagset()
		assert.True(t, ts0.IsEmpty())
	})

	one := tags.New("one")
	ts1 := tags.NewTagset()
	ts1.AddString("one")

	t.Run("buildSet", func(t *testing.T) {

		// 	fmt.Println(ts1)

		ts2 := tags.NewTagset()
		ts2.Add(one)
		// 	fmt.Println(ts2)

		assert.True(t, tags.EqualSet(ts1, ts2))
	})

	t.Run("remove1", func(t *testing.T) {
		ts1.Remove("one")
		// 	fmt.Println(ts1)
		assert.True(t, ts1.IsEmpty())
	})
}

func TestValidation(t *testing.T) {

	tt := []string{`group_W-A-T_2025+test`, `dell-sas`, `asus_u2nvme`, `cometnas+2025`}
	for _, x := range tt {
		ti := tags.New(x)
		assert.Valid(t, ti)
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

	ts2 := tags.FromStrings(`eight`, `f1ve`, `five`, `four`, `one`, `six`, `three`, `two`)
	// 	fmt.Println(ts2)

	assert.True(t, tags.EqualSet(ts, ts2))
}

func TestRemoval(t *testing.T) {

	ts := g_tagset.Clone()
	// 	fmt.Println(ts)

	t.Run("removal -simple", func(t *testing.T) {
		ts.Sort()
		// 	fmt.Println(ts)
		ts.Remove(tags.New("five"))
		// 	fmt.Println(ts)
		ts.Remove(tags.New("six"))
		// 	fmt.Println(ts)

		assert.True(t, tags.EqualSet(ts, tags.FromStrings("four", "one", "three", "two")))
	})

	t.Run("removal -not-there", func(t *testing.T) {
		// Remove one non-existing tag...
		// ...and check we didn't clobber anything valuable :S (bugfix coverage)
		ts.Remove(tags.New("not-there"))
		assert.Equal(t, ts.Len(), 4)
	})
}

func TestRemoval2(t *testing.T) {

	ts := g_tagset.Clone()
	//	fmt.Println(ts)

	t.Run("multi-removal", func(t *testing.T) {
		ts.Remove(tags.New("five"))
		//	fmt.Println(ts)

		ts.AddString("four")
		ts.AddString("nine")
		//	fmt.Println(ts)

		ts.Remove(tags.New("six"))
		//	fmt.Println(ts)
		ts.Remove(tags.New("five"))

		ts.Remove(tags.New("four"))

		ts.Sort()
		//	fmt.Println(ts)

		assert.True(t, tags.EqualSet(ts, tags.FromStrings("four", "nine", "one", "three", "two")))
	})
}

func TestContains(t *testing.T) {

	ts := g_tagset // g_tagset.Clone()

	assert.True(t, ts.Contains(tags.New("one")))
	assert.True(t, ts.ContainsString("three"))

	ts2 := tags.FromStrings("one", "two")
	assert.True(t, ts2.Contains(tags.New("one")))
	assert.False(t, ts2.Contains(tags.New("three")))

	assert.True(t, ts2.ContainsString("two"))
	assert.False(t, ts2.ContainsString("four"))
}

func TestOutput(t *testing.T) {
	ts := g_tagset //.Clone()
	fmt.Println(ts)
}

func TestParseRaw(t *testing.T) {

	tv, err := tags.Parse(`[one,two,three,four,five,six]`)
	if nil != err {
		t.Error(err)
	}

	assert.True(t, tags.EqualSet(tv, g_tagset))
}

func TestRoundTrip(t *testing.T) {

	out := g_tagset.String()

	in, err := tags.Parse(out)
	if nil != err {
		t.Error(err)
	}

	assert.True(t, tags.EqualSet(in, g_tagset))
	assert.Equal(t, out, in.String())
}

func TestAppend(t *testing.T) {

	ts1,_ := tags.Parse("[one,two,three]")
	ts2,_ := tags.Parse("[four,five,six]")

	t.Run("append-simple", func(t *testing.T) {
		ts := ts1.Clone()
		ts.Append(ts2)
		assert.True(t, tags.EqualSet(ts, g_tagset))
	})
	t.Run("append-sort", func(t *testing.T) {
		ts := ts1
		ts.Append(ts2)
		ts.Sort()
		ref := g_tagset.Clone()
		ref.Sort()
		assert.True(t, tags.EqualSet(ts, ref))
	})
}

func TestBinarySearch(t *testing.T) {

	ts := g_tagset.Clone()
	ts.Sort()
	// fmt.Println("ts sorted", ts)
	t.Run("found element", func(t *testing.T) {
		pos, found := ts.BinarySearch(tags.New("three"))
		assert.True(t, found)
		assert.Equal(t, 4, pos)
	})

	t.Run("not found in middle", func(t *testing.T) {
		// fmt.Println("ts before three2", ts)
		pos, found := ts.BinarySearch(tags.New("three2"))
		// fmt.Println("ts after three2", ts)
		assert.False(t, found)
		assert.Equal(t, 5, pos)
	})

	t.Run("before all elements", func(t *testing.T) {
		pos, found := ts.BinarySearch(tags.New("aaa"))
		assert.False(t, found)
		assert.Equal(t, 0, pos)
	})

	t.Run("after all elements", func(t *testing.T) {
		pos, found := ts.BinarySearch(tags.New("zzz"))
		assert.False(t, found)
		assert.Equal(t, ts.Len(), pos)
	})

	t.Run("empty tagset", func(t *testing.T) {
		empty := tags.NewTagset()
		pos, found := empty.BinarySearch(tags.New("one"))
		// fmt.Println("empty before empty tagset test", empty)
		assert.False(t, found)
		assert.Equal(t, 0, pos)
	})
}

func TestIndex(t *testing.T) {
	ts := g_tagset.Clone()

	t.Run("found element", func(t *testing.T) {
		index := ts.Index(tags.New("three"))
		// fmt.Println("ts____", ts)
		assert.Equal(t, 2, index)
	})
	t.Run("not found element", func(t *testing.T) {
		index := ts.Index(tags.New("three2"))
		assert.Equal(t, -1, index)
	})

	t.Run("empty tagset", func(t *testing.T) {
		empty := tags.NewTagset()
		index := empty.Index(tags.New("one"))
		assert.Equal(t, -1, index)
	})
}

func TestClip(t *testing.T) {
	ts := g_tagset.Clone()

	t.Run("clip capacity", func(t *testing.T) {
		ts2 := ts[0:3]
		beforeCapacity := cap(ts2)
		ts2.Clip()
		afterCapacity := cap(ts2)

		assert.Equal(t, 3, ts2.Len())
		assert.Equal(t, ts2.Len(), afterCapacity)
		assert.True(t, afterCapacity < beforeCapacity)
	})
}

func TestCompact(t *testing.T) {

	t.Run("consecutive duplicates", func(t *testing.T) {
		ts := tags.FromStrings("one", "one", "two", "three", "four")
		ts.Compact()
		//fmt.Println("ts in TestCompact after Compact()", ts)
		expected := tags.FromStrings("one", "two", "three", "four")
		assert.True(t, tags.EqualSet(ts, expected))
	})

	t.Run("no duplicates", func(t *testing.T) {
		ts := tags.FromStrings("one", "two", "three", "four")
		ts.Compact()
		expected := tags.FromStrings("one", "two", "three", "four")
		assert.True(t, tags.EqualSet(ts, expected))
	})

	t.Run("empty slice", func(t *testing.T) {
		ts := tags.NewTagset()
		ts.Compact()
		expected := tags.NewTagset()
		assert.True(t, tags.EqualSet(ts, expected))
	})

	t.Run("single element", func(t *testing.T) {
		ts := tags.FromStrings("one")
		ts.Compact()
		expected := tags.FromStrings("one")
		assert.True(t, tags.EqualSet(ts, expected))
	})

	t.Run("non consecutive duplicates", func(t *testing.T) {
		ts := tags.FromStrings("one", "two", "one", "three")
		ts.Compact()
		expected := tags.FromStrings("one", "two", "one", "three")
		assert.True(t, tags.EqualSet(ts, expected))
	})
}
func TestEqualSet(t *testing.T) {
	t.Run("equal tag sets", func(t *testing.T) {
		s1 := tags.FromStrings("a", "b", "c")
		s2 := tags.FromStrings("a", "b", "c")
		assert.True(t, tags.EqualSet(s1, s2))
	})

	t.Run("different tag sets", func(t *testing.T) {
		s1 := tags.FromStrings("a", "b")
		s2 := tags.FromStrings("a", "b", "c")
		assert.False(t, tags.EqualSet(s1, s2))
	})

	t.Run("different order", func(t *testing.T) {
		s1 := tags.FromStrings("a", "b", "c")
		s2 := tags.FromStrings("c", "b", "a")
		assert.False(t, tags.EqualSet(s1, s2))
	})

	t.Run("one empty tag set", func(t *testing.T) {
		s1 := tags.NewTagset()
		s2 := tags.FromStrings("a", "b")
		assert.False(t, tags.EqualSet(s1, s2))
	})

	t.Run("two empty tag sets", func(t *testing.T) {
		s1 := tags.NewTagset()
		s2 := tags.NewTagset()
		assert.True(t, tags.EqualSet(s1, s2))
	})
}

func TestClone(t *testing.T) {

	t.Run("non empty clone", func(t *testing.T) {
		original := tags.FromStrings("one", "two")
		clone := original.Clone()
		assert.True(t, tags.EqualSet(original, clone))

		clone.AddString("three")
		assert.False(t, tags.EqualSet(original, clone))
	})

	t.Run("empty clone return nil", func(t *testing.T) {
		empty := tags.NewTagset()
		clone := empty.Clone()
		assert.Equal(t, 0, clone.Len())
	})
}
