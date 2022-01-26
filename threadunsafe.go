package set

import (
	"fmt"
	"strings"
)

func NewThreadUnsafeSet(elems ...interface{}) ISet {
	s := make(threadUnsafeSet, len(elems))
	s.Adds(elems...)
	return &s
}

type threadUnsafeSet map[interface{}]struct{}

func (s *threadUnsafeSet) Empty() bool {
	return s.Cardinality() == 0
}

func (s *threadUnsafeSet) Singleton() bool {
	return s.Cardinality() == 1
}

func (s *threadUnsafeSet) Cardinality() int {
	return len(*s)
}

func (s *threadUnsafeSet) ToSlice() ISlice {
	result := make(Slice, 0, len(*s))
	for elem := range *s {
		result = append(result, elem)
	}
	return result
}

func (s *threadUnsafeSet) Adds(elems ...interface{}) bool {
	var exist bool
	for i := 0; i < len(elems); i++ {
		if _, ok := (*s)[elems[i]]; ok {
			exist = true
		} else {
			(*s)[elems[i]] = struct{}{}
		}
	}
	return !exist
}

func (s *threadUnsafeSet) Removes(elems ...interface{}) bool {
	var notExist bool
	for i := 0; i < len(elems); i++ {
		if _, ok := (*s)[elems[i]]; ok {
			delete(*s, elems[i])
		} else {
			notExist = true
		}
	}
	return !notExist
}

func (s *threadUnsafeSet) IsSub(other ISet) bool {
	if s.Cardinality() > other.Cardinality() {
		return false
	}
	return other.Contains(s.ToSlice().Interface()...)
}

func (s *threadUnsafeSet) Unions(others ...ISet) ISet {
	result := s.Clone()
	for _, other := range others {
		result.Adds(other.ToSlice().Interface()...)
	}
	return result
}

func (s *threadUnsafeSet) Intersections(others ...ISet) ISet {
	result := NewThreadUnsafeSet()
	var baseSet ISet = s
	var diffSets []ISet
	for _, other := range others {
		if other.Cardinality() < baseSet.Cardinality() {
			diffSets = append(diffSets, baseSet)
			baseSet = other
		} else {
			diffSets = append(diffSets, other)
		}
	}
Loop:
	for _, elem := range baseSet.ToSlice().Interface() {
		for _, diffSet := range diffSets {
			if !diffSet.Contains(elem) {
				continue Loop
			}
		}
		result.Adds(elem)
	}
	return result
}

func (s *threadUnsafeSet) Complements(others ...ISet) ISet {
	result := s.Clone()
	for _, other := range others {
		result.Removes(other.ToSlice().Interface()...)
	}
	return result
}

func (s *threadUnsafeSet) Clear() {
	*s = make(threadUnsafeSet)
}

func (s *threadUnsafeSet) Contains(elems ...interface{}) bool {
	for i := 0; i < len(elems); i++ {
		if _, ok := (*s)[elems[i]]; !ok {
			return false
		}
	}
	return true
}

func (s *threadUnsafeSet) Clone() ISet {
	return NewThreadUnsafeSet(s.ToSlice().Interface()...)
}

func (s *threadUnsafeSet) Equal(other ISet) bool {
	if other.Cardinality() != s.Cardinality() {
		return false
	}
	return s.Contains(other.ToSlice().Interface()...)
}

func (s *threadUnsafeSet) Pop() interface{} {
	for elem := range *s {
		delete(*s, elem)
		return elem
	}
	return nil
}

func (s *threadUnsafeSet) String() string {
	elems := make([]string, 0, s.Cardinality())
	for elem := range *s {
		elems = append(elems, fmt.Sprintf("%v", elem))
	}
	return "{" + strings.Join(elems, ",") + "}"
}
