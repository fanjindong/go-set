package set

import "sync"

func NewSet(elems ...interface{}) ISet {
	s := &threadSafeSet{m: NewThreadUnsafeSet(elems...).(*threadUnsafeSet)}
	return s
}

type threadSafeSet struct {
	rwm sync.RWMutex
	m   *threadUnsafeSet
}

func (s *threadSafeSet) Empty() bool {
	return s.Cardinality() == 0
}

func (s *threadSafeSet) Singleton() bool {
	return s.Cardinality() == 1
}

func (s *threadSafeSet) Cardinality() int {
	s.rwm.RLock()
	defer s.rwm.RUnlock()
	return s.m.Cardinality()
}

func (s *threadSafeSet) ToSlice() ISlice {
	s.rwm.RLock()
	defer s.rwm.RUnlock()
	return s.m.ToSlice()
}

func (s *threadSafeSet) Adds(elems ...interface{}) bool {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	return s.m.Adds(elems...)
}

func (s *threadSafeSet) Removes(elems ...interface{}) bool {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	return s.m.Removes(elems...)
}

func (s *threadSafeSet) IsSub(other ISet) bool {
	if s.Cardinality() > other.Cardinality() {
		return false
	}
	return other.Contains(s.ToSlice().Interface()...)
}

func (s *threadSafeSet) Unions(others ...ISet) ISet {
	result := s.Clone()
	for _, other := range others {
		result.Adds(other.ToSlice().Interface()...)
	}
	return result
}

func (s *threadSafeSet) Intersections(others ...ISet) ISet {
	result := NewSet()
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

func (s *threadSafeSet) Complements(others ...ISet) ISet {
	result := s.Clone()
	for _, other := range others {
		result.Removes(other.ToSlice().Interface()...)
	}
	return result
}

func (s *threadSafeSet) Clear() {
	s.Removes(s.ToSlice().Interface()...)
}

func (s *threadSafeSet) Contains(elems ...interface{}) bool {
	s.rwm.RLock()
	defer s.rwm.RUnlock()
	return s.m.Contains(elems...)
}

func (s *threadSafeSet) Clone() ISet {
	return NewSet(s.ToSlice().Interface()...)
}

func (s *threadSafeSet) Equal(other ISet) bool {
	if other.Cardinality() != s.Cardinality() {
		return false
	}
	return s.Contains(other.ToSlice().Interface()...)
}

func (s *threadSafeSet) Pop() interface{} {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	return s.m.Pop()
}

func (s *threadSafeSet) String() string {
	s.rwm.RLock()
	defer s.rwm.RUnlock()
	return s.m.String()
}
