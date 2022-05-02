package set_test

import (
	"testing"

	"github.com/grimdork/set"
)

func TestIntSet(t *testing.T) {
	s1 := set.Set[int]{}
	s1.Add(42)
	s1.Add(42)
	s1.Add(43)
	s1.Add(44)
	t.Logf("Set 1: %v", s1)
	if len(s1) != 3 {
		t.Errorf("Set 1 should have 3 elements, but has %d", len(s1))
		t.Fail()
	}

	if !s1.Contains(42) {
		t.Errorf("Set 1 does not contain 42")
		t.Fail()
	}

	t.Logf("Does set 1 contain 42? %t", s1.Contains(42))

	s2 := set.Set[int]{}
	s2.Add(42)
	s2.Add(52)
	s2.Add(62)
	s2.Add(72)
	t.Logf("Set 2: %v", s2)

	t.Logf("Set 1 and set 2 equality: %t", s1.Compare(s2))
	t.Log("Merging set 2 into set 1")
	s1.Merge(s2)
	t.Logf("Set 1: %v", s1)

	t.Log("Merging set 1 into set 2")
	s2.Merge(s1)
	t.Logf("Set 2: %v", s2)
	if !s1.Compare(s2) {
		t.Error("Set 1 and set 2 are not equal, but they should be.")
		t.Fail()
	}

	t.Logf("Set 1 and set 2 equality: %t", s1.Compare(s2))
	s1.Remove(42)
	t.Log("Removed 42 from set 1")
	if s1.Contains(42) {
		t.Errorf("Set 1 should not contain 42, but does.")
		t.Fail()
	}
}

func TestStringSet(t *testing.T) {
	s1 := set.Set[string]{}
	s1.Add("Moo")
	s1.Add("Moo")
	s1.Add("Nyar")
	s1.Add("Mbleh")
	t.Logf("Set 1: %v", s1)
	if !s1.Contains("Moo") {
		t.Errorf("Set 1 does not contain Moo")
		t.Fail()
	}

	s2 := set.Set[string]{}
	s2.Add("Moo")
	s2.Add("Hey")
	s2.Add("Ho")
	t.Logf("Set 2: %v", s2)

	if s1.Compare(s2) {
		t.Error("Set 1 and set 2 are equal, but they shouldn't be.")
		t.Fail()
	}
	t.Logf("Set 1 and set 2 equality: %t", s1.Compare(s2))
	t.Log("Merging set 2 into set 1")
	s1.Merge(s2)
	t.Logf("Set 1: %v", s1)
	if s1.Compare(s2) {
		t.Error("Set 1 and set 2 are equal, but they shouldn't be.")
		t.Fail()
	}
}

func TestBooleanSet(t *testing.T) {
	s1 := set.Set[bool]{}
	s1.Add(true)
	s1.Add(true)
	s1.Add(false)
	s1.Add(false)
	t.Logf("Set 1: %v", s1)

	s2 := set.Set[bool]{}
	s2.Add(true)
	s2.Add(false)
	t.Logf("Set 2: %v", s2)

	if !s1.Compare(s2) {
		t.Error("Set 1 and set 2 are not equal, but they should be.")
		t.Fail()
	}
}
