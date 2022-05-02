package set

// Set definition.
type Set[V comparable] map[V]struct{}

// Add a value to the set.
func (set Set[V]) Add(value V) {
	set[value] = struct{}{}
}

// Remove a value from the set.
func (set Set[V]) Remove(value V) {
	delete(set, value)
}

// Contains returns true if the set contains the value.
func (set Set[V]) Contains(value V) bool {
	_, ok := set[value]
	return ok
}

// Merge another set into this.
func (set Set[V]) Merge(other Set[V]) {
	for k := range other {
		set.Add(k)
	}
}

// Compare two sets.
func (set Set[V]) Compare(other Set[V]) bool {
	if len(set) != len(other) {
		return false
	}

	t := true
	for k := range set {
		t = t && other.Contains(k)
	}

	return t
}
