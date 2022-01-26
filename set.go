package set

type ISet interface {
	// Empty The empty set (or null set) is the unique set that has no members. It is denoted ∅.
	//Examples:
	//{}.Empty() return true
	//{1, 2}.Empty() return false
	Empty() bool
	// Singleton A singleton set is a set with exactly one element; such a set may also be called a unit set.
	//Any such set can be written as {x}, where x is the element. The set {x} and the element x mean different things;
	//Halmos draws the analogy that a box containing a hat is not the same as the hat.
	//Examples:
	//{1}.Singleton() return true
	//{1, 2}.Singleton() return false
	Singleton() bool
	// IsSub If every element of set A is also in B, then A is described as being a subset of B, or contained in B, written A ⊆ B.[34] B ⊇ A means B contains A, B includes A, or B is a superset of A; B ⊇ A is equivalent to A ⊆ B.[35][14] The relationship between sets established by ⊆ is called inclusion or containment. Two sets are equal if they contain each other: A ⊆ B and B ⊆ A is equivalent to A = B.[26]
	//If A is a subset of B, but A is not equal to B, then A is called a proper subset of B. This can be written A ⊊ B. Likewise, B ⊋ A means B is a proper superset of A, i.e. B contains A, and is not equal to A.
	//A third pair of operators ⊂ and ⊃ are used differently by different authors: some authors use A ⊂ B and B ⊃ A to mean A is any subset of B (and not necessarily a proper subset),[36][29] while others reserve A ⊂ B and B ⊃ A for cases where A is a proper subset of B.[34]
	//Examples:
	//The set of all humans is a proper subset of the set of all mammals.
	//{1, 3}.IsSub({1, 2, 3, 4}) return true
	//{1, 2, 3, 4}.IsSub({1, 2, 3, 4}) return true
	//The empty set is a subset of every set, and every set is a subset of itself.
	IsSub(ISet) bool
	// Cardinality The cardinality of a set S, denoted |S|, is the number of members of S.
	//For example, if B = {blue, white, red}, then |B| = 3.
	//Repeated members in roster notation are not counted, so |{blue, white, red, blue, white}| = 3, too.
	//The cardinality of the empty set is zero.
	//Examples:
	//{}.Cardinality() return 0
	//{1, 2}.Cardinality() return 2
	Cardinality() int
	// Unions Two sets can be "added" together. The union of A and B, denoted by A ∪ B, is the set of all things that are members of either A or B.
	//Examples:
	//{1, 2}.Unions({1, 2}) return {1, 2}.
	//{1, 2}.Unions({2, 3}) return {1, 2, 3}.
	//{1, 2, 3}.Unions({3}, {4, 5}) return {1, 2, 3, 4, 5}
	Unions(...ISet) ISet
	// Intersections A new set can also be constructed by determining which members two sets have "in common".
	//The intersection of A and B, denoted by A ∩ B, is the set of all things that are members of both A and B. If A ∩ B = ∅, then A and B are said to be disjoint.
	//The intersection of A and B, denoted A ∩ B.
	//Examples:
	//{1, 2}.Intersections({1, 2}) return {1, 2}.
	//{1, 2}.Intersections({2, 3}) return {2}.
	//{1, 2}.Intersections({3, 4}) return ∅.
	Intersections(...ISet) ISet
	// Complements Two sets can also be "subtracted". The relative complement of B in A (also called the set-theoretic difference of A and B),
	//denoted by A \ B (or A − B), is the set of all elements that are members of A, but not members of B.
	//It is valid to "subtract" members of a set that are not in the set, such as removing the element green from the set {1, 2, 3};
	//doing so will not affect the elements in the set.
	//In certain settings, all sets under discussion are considered to be subsets of a given universal set U.
	//In such cases, U \ A is called the absolute complement or simply complement of A, and is denoted by A′ or Ac.
	//A′ = U \ A
	//Examples:
	//{1, 2}.Complements({1, 2}) return ∅.
	//{1, 2, 3, 4}.Complements({1, 3}) return {2, 4}.
	Complements(...ISet) ISet
	// Adds many element to the set. Returns whether all the items was added.
	//Examples:
	//{1, 2}.Add(3,4)={1,2,3,4} return true
	//{1, 2}.Add(1,4)={1,2,4} return false
	//{1, 2}.Add(1,2)={1,2} return false
	Adds(...interface{}) bool
	// Clear removes all elements from the set, leaving the empty set.
	Clear()
	// Removes remove elements from the set. Returns whether all the items was Removed.
	//Examples:
	//{1, 2}.Removes(3,4)={1,2} return false
	//{1, 2}.Removes(1,4)={2} return false
	//{1, 2}.Removes(1,2)=∅ return true
	//{1, 2, 3}.Removes(1,3)={2} return true
	Removes(...interface{}) bool
	// Contains Returns whether the given items are all in the set.
	//Examples:
	//{1, 2}.Contains(1) return true
	//{1, 2}.Contains(1,2) return true
	//{1, 2}.Contains(1,2,3) return false
	//{1, 2}.Contains(3) return false
	Contains(...interface{}) bool
	// Clone Returns a clone of the set using the same implementation, duplicating all keys.
	Clone() ISet
	// Equal Determines if two sets are equal to each other.
	// If they have the same cardinality and contain the same elements, they are considered equal.
	// The order in which the elements were added is irrelevant.
	//Examples:
	//{1, 2}.Equal({3,4}) return false
	//{1, 2}.Equal({1,4}) return false
	//{1, 2}.Equal({1,2}) return true
	Equal(ISet) bool
	// Pop removes and returns an arbitrary item from the set.
	// return nil if set is full.
	//Examples:
	//{1}.Pop() return 1
	//{1, 2}.Pop() return 1 or 2
	//{}.Pop() return nil
	Pop() interface{}
	// ToSlice Returns the members of the set as a slice.
	//Examples:
	//{}.ToSlice() return nil
	//{1, 2}.ToSlice().Interface() return []interface{}{1,2}
	//{1, 2}.ToSlice().Int() return []int{}{1,2}
	ToSlice() ISlice
	// String Formatted output string
	// Examples:
	// NewMapSet(1,2,3).String()
	// output: {1,2,3}
	String() string
}

//
//func NewMapSet(elems ...interface{}) ISet {
//	s := make(mapSet, len(elems))
//	s.Adds(elems...)
//	return s
//}
//
//type mapSet map[interface{}]struct{}
//
//func (s mapSet) Empty() bool {
//	return s.Cardinality() == 0
//}
//
//func (s mapSet) Singleton() bool {
//	return s.Cardinality() == 1
//}
//
//func (s mapSet) Cardinality() int {
//	return len(s)
//}
//
//func (s mapSet) ToSlice() ISlice {
//	result := make(Slice, 0, len(s))
//	for elem := range s {
//		result = append(result, elem)
//	}
//	return result
//}
//
//func (s mapSet) Adds(elems ...interface{}) bool {
//	var exist bool
//	for i := 0; i < len(elems); i++ {
//		if _, ok := s[elems[i]]; ok {
//			exist = true
//		} else {
//			s[elems[i]] = struct{}{}
//		}
//	}
//	return !exist
//}
//
//func (s mapSet) Removes(elems ...interface{}) {
//	for i := 0; i < len(elems); i++ {
//		if _, ok := s[elems[i]]; ok {
//			delete(s, elems[i])
//		}
//	}
//}
//
//func (s mapSet) IsSub(other ISet) bool {
//	if s.Cardinality() > other.Cardinality() {
//		return false
//	}
//	return other.Contains(s.ToSlice().Interface()...)
//}
//
//func (s mapSet) Unions(others ...ISet) ISet {
//	result := s.Clone()
//	for _, other := range others {
//		result.Adds(other.ToSlice().Interface()...)
//	}
//	return result
//}
//
//func (s mapSet) Intersections(others ...ISet) ISet {
//	result := NewMapSet()
//	var baseSet ISet = s
//	var diffSets []ISet
//	for _, other := range others {
//		if other.Cardinality() < baseSet.Cardinality() {
//			diffSets = append(diffSets, baseSet)
//			baseSet = other
//		} else {
//			diffSets = append(diffSets, other)
//		}
//	}
//Loop:
//	for _, elem := range baseSet.ToSlice().Interface() {
//		for _, diffSet := range diffSets {
//			if !diffSet.Contains(elem) {
//				continue Loop
//			}
//		}
//		result.Adds(elem)
//	}
//	return result
//}
//
//func (s mapSet) Complements(others ...ISet) ISet {
//	result := s.Clone()
//	for _, other := range others {
//		result.Removes(other.ToSlice().Interface()...)
//	}
//	return result
//}
//
//func (s mapSet) Clear() {
//	s.Removes(s.ToSlice().Interface()...)
//}
//
//func (s mapSet) Contains(elems ...interface{}) bool {
//	for i := 0; i < len(elems); i++ {
//		if _, ok := s[elems[i]]; !ok {
//			return false
//		}
//	}
//	return true
//}
//
//func (s mapSet) Clone() ISet {
//	return NewMapSet(s.ToSlice().Interface()...)
//}
//
//func (s mapSet) Equal(other ISet) bool {
//	if other.Cardinality() != s.Cardinality() {
//		return false
//	}
//	return s.Contains(other.ToSlice().Interface()...)
//}
//
//func (s mapSet) Pop() interface{} {
//	for elem := range s {
//		delete(s, elem)
//		return elem
//	}
//	return nil
//}
//
//func (s mapSet) String() string {
//	elems := make([]string, 0, s.Cardinality())
//	for elem, _ := range s {
//		elems = append(elems, fmt.Sprintf("%v", elem))
//	}
//	return "{" + strings.Join(elems, ",") + "}"
//}
//
//func NewConcurrentMapSet(elems ...interface{}) ISet {
//	s := &cMapSet{m: NewMapSet(elems...)}
//	return s
//}
//
//type cMapSet struct {
//	rwm sync.RWMutex
//	m   ISet
//}
//
//func (s *cMapSet) Empty() bool {
//	return s.Cardinality() == 0
//}
//
//func (s *cMapSet) Singleton() bool {
//	return s.Cardinality() == 1
//}
//
//func (s *cMapSet) Cardinality() int {
//	s.rwm.RLock()
//	defer s.rwm.RUnlock()
//	return s.m.Cardinality()
//}
//
//func (s *cMapSet) ToSlice() ISlice {
//	s.rwm.RLock()
//	defer s.rwm.RUnlock()
//	return s.m.ToSlice()
//}
//
//func (s *cMapSet) Adds(elems ...interface{}) bool {
//	s.rwm.Lock()
//	defer s.rwm.Unlock()
//	return s.m.Adds(elems...)
//}
//
//func (s *cMapSet) Removes(elems ...interface{}) {
//	s.rwm.Lock()
//	defer s.rwm.Unlock()
//	s.m.Removes(elems...)
//}
//
//func (s *cMapSet) IsSub(other ISet) bool {
//	if s.Cardinality() > other.Cardinality() {
//		return false
//	}
//	return other.Contains(s.ToSlice().Interface()...)
//}
//
//func (s *cMapSet) Unions(others ...ISet) ISet {
//	result := s.Clone()
//	for _, other := range others {
//		result.Adds(other.ToSlice().Interface()...)
//	}
//	return result
//}
//
//func (s *cMapSet) Intersections(others ...ISet) ISet {
//	result := NewConcurrentMapSet()
//	var baseSet ISet = s
//	var diffSets []ISet
//	for _, other := range others {
//		if other.Cardinality() < baseSet.Cardinality() {
//			diffSets = append(diffSets, baseSet)
//			baseSet = other
//		} else {
//			diffSets = append(diffSets, other)
//		}
//	}
//Loop:
//	for _, elem := range baseSet.ToSlice().Interface() {
//		for _, diffSet := range diffSets {
//			if !diffSet.Contains(elem) {
//				continue Loop
//			}
//		}
//		result.Adds(elem)
//	}
//	return result
//}
//
//func (s *cMapSet) Complements(others ...ISet) ISet {
//	result := s.Clone()
//	for _, other := range others {
//		result.Removes(other.ToSlice().Interface()...)
//	}
//	return result
//}
//
//func (s *cMapSet) Clear() {
//	s.Removes(s.ToSlice().Interface()...)
//}
//
//func (s *cMapSet) Contains(elems ...interface{}) bool {
//	s.rwm.RLock()
//	defer s.rwm.RUnlock()
//	return s.m.Contains(elems...)
//}
//
//func (s *cMapSet) Clone() ISet {
//	return NewConcurrentMapSet(s.ToSlice().Interface()...)
//}
//
//func (s *cMapSet) Equal(other ISet) bool {
//	if other.Cardinality() != s.Cardinality() {
//		return false
//	}
//	return s.Contains(other.ToSlice().Interface()...)
//}
//
//func (s *cMapSet) Pop() interface{} {
//	s.rwm.Lock()
//	defer s.rwm.Unlock()
//	return s.m.Pop()
//}
//
//func (s *cMapSet) String() string {
//	s.rwm.RLock()
//	defer s.rwm.RUnlock()
//	return s.m.String()
//}
