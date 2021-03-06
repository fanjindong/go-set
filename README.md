# go-set
[![CI](https://github.com/fanjindong/go-set/actions/workflows/main.yml/badge.svg?branch=master)](https://github.com/fanjindong/go-set/actions/workflows/main.yml)
[![GoDoc](https://godoc.org/github.com/fanjindong/go-set?status.svg)](https://pkg.go.dev/github.com/fanjindong/go-set)

Implementation of set data structure in golang language

![image](./images/LoonyDook_ZH-CN1879420705_1920x1080.jpg)

## Install

`go get -u github.com/fanjindong/go-set`

## Fast Start

```go
import "github.com/fanjindong/go-set"

func main() {
    s := set.NewSet()
    s.Adds(1)
    s.Adds(2,3)
    
    fmt.Println(s.String()) // {1,2,3}
    fmt.Println(s.Contains(1)) // true
    fmt.Println(s.Cardinality()) // 3
    fmt.Println(s.ToSlice().Int()) // []int{1,2,3}, nil
}

```

## Complete Tutorial
ISet contains threadSafeSet: `NewSet()`, threadUnsafeSet: `NewThreadUnsafeSet()`.

List of interface methods
* [Cardinality() int](#cardinality-int)
* [Adds(\.\.\.interface\{\}) bool](#addsinterface-bool)
* [Clear()](#clear)
* [Removes(\.\.\.interface\{\}) bool](#removesinterface-bool)
* [Contains(\.\.\.interface\{\}) bool](#containsinterface-bool)
* [Empty() bool](#empty-bool)
* [Singleton() bool](#singleton-bool)
* [IsSub(ISet) bool](#issubiset-bool)
* [Unions(\.\.\.ISet) ISet](#unionsiset-iset)
* [Intersections(\.\.\.ISet) ISet](#intersectionsiset-iset)
* [Complements(\.\.\.ISet) ISet](#complementsiset-iset)
* [Clone() ISet](#clone-iset)
* [Equal(ISet) bool](#equaliset-bool)
* [Pop() interface\{\}](#pop-interface)
* [ToSlice() Slice](#toslice-slice)
* [String() string](#string-string)

Slice

### Cardinality() int

The cardinality of a set S, denoted |S|, is the number of members of S.
For example, if B = {blue, white, red}, then |B| = 3.
Repeated members in roster notation are not counted, so |{blue, white, red, blue, white}| = 3, too.
The cardinality of the empty set is zero.

Examples:
```go
NewSet().Cardinality() // 0
NewSet(1,2).Cardinality() // 2
```

### Adds(...interface{}) bool

Adds many elements to the set. Returns whether all the items was added.

Examples:
```go
NewSet(1,2).Add(3,4) // true
NewSet(1,2).Add(1,4) // false
NewSet(1,2).Add(1,2) // false
```

### Clear()

removes all elements from the set, leaving the empty set.

### Removes(...interface{}) bool

remove elements from the set.

Examples:
```go
NewSet(1,2).Removes(3,4) // false
NewSet(1,2).Removes(1,4) // false
NewSet(1,2).Removes(1,2) // true
```

### Contains(...interface{}) bool

Returns whether the given items are all in the set.

Examples:
```go
NewSet(1,2).Contains(1) // true
NewSet(1,2).Contains(1,2) // true
NewSet(1,2).Contains(1,2,3) // false
NewSet(1,2).Contains(3) // false
```

### Empty() bool

The empty set (or null set) is the unique set that has no members. It is denoted ???({}).

Examples:
```go
NewSet().Empty() // true
NewSet(1,2).Empty() // false
```

### Singleton() bool

A singleton set is a set with exactly one element; such a set may also be called a unit set.
Any such set can be written as {x}, where x is the element. The set {x} and the element x mean different things;
Halmos draws the analogy that a box containing a hat is not the same as the hat.

Examples:
```go
NewSet(1).Singleton() //true
NewSet(1,2).Singleton() //false
```

### IsSub(ISet) bool

If every element of set A is also in B, then A is described as being a subset of B, or contained in B, written A ??? B. B ??? A means B contains A, B includes A, or B is a superset of A; B ??? A is equivalent to A ??? B. The relationship between sets established by ??? is called inclusion or containment. Two sets are equal if they contain each other: A ??? B and B ??? A is equivalent to A = B.
If A is a subset of B, but A is not equal to B, then A is called a proper subset of B. This can be written A ??? B. Likewise, B ??? A means B is a proper superset of A, i.e. B contains A, and is not equal to A.
A third pair of operators ??? and ??? are used differently by different authors: some authors use A ??? B and B ??? A to mean A is any subset of B (and not necessarily a proper subset), while others reserve A ??? B and B ??? A for cases where A is a proper subset of B.

Examples:
```go
// The set of all humans is a proper subset of the set of all mammals.
NewSet(1,3).IsSub(NewSet(1,2,3,4)) // true
NewSet(1,2,3,4).IsSub(NewSet(1,2,3,4)) // true
```
The empty set is a subset of every set, and every set is a subset of itself.

### Unions(...ISet) ISet

Two sets can be "added" together. The union of A and B, denoted by A ??? B, is the set of all things that are members of either A or B.

Examples:
```go
NewSet(1,2).Unions(NewSet(1,2)) // NewSet(1,2)
NewSet(1,2).Unions(NewSet(3,2)) // NewSet(1,2,3)
NewSet(1,2,3).Unions(NewSet(3), NewSet(4,5)) // NewSet(1,2,3,4,5)
```

### Intersections(...ISet) ISet

A new set can also be constructed by determining which members two sets have "in common".
The intersection of A and B, denoted by A ??? B, is the set of all things that are members of both A and B. If A ??? B = ???, then A and B are said to be disjoint.
The intersection of A and B, denoted A ??? B.

Examples:
```go
NewSet(1,2).Intersections(NewSet(1,2)) // NewSet(1,2)
NewSet(1,2).Intersections(NewSet(3,2)) // NewSet(2)
NewSet(1,2).Intersections(NewSet(3,4)) // NewSet()
```

### Complements(...ISet) ISet

Two sets can also be "subtracted". The relative complement of B in A (also called the set-theoretic difference of A and B),
denoted by A \ B (or A ??? B), is the set of all elements that are members of A, but not members of B.
It is valid to "subtract" members of a set that are not in the set, such as removing the element green from the set {1, 2, 3};
doing so will not affect the elements in the set.
In certain settings, all sets under discussion are considered to be subsets of a given universal set U.
In such cases, U \ A is called the absolute complement or simply complement of A, and is denoted by A??? or Ac.
A??? = U \ A.

Examples:
```go
NewSet(1,2).Complements(NewSet(1,2)) // NewSet().
NewSet(1,2,3,4).Complements(NewSet(1,3)) // NewSet(2,4).
```

### Clone() ISet

Returns a clone of the set using the same implementation, duplicating all keys.

### Equal(ISet) bool

Determines if two sets are equal to each other.
If they have the same cardinality and contain the same elements, they are considered equal.
The order in which the elements were added is irrelevant.

Examples:
```go
NewSet(1,2).Equal(NewSet(3,4)) // false
NewSet(1,2).Equal(NewSet(1,4)) // false
NewSet(1,2).Equal(NewSet(1,2)) // true
```

### Pop() interface{}

removes and returns an arbitrary item from the set.
return nil if set is full.

Examples:
```go
NewSet(1).Pop() // 1
NewSet(1,2).Pop() // 1 or 2
NewSet().Pop() // nil
```

### ToSlice() Slice

Returns the members of the set as a slice.

Examples:
```go
NewSet().ToSlice() // nil
NewSet(1,2).ToSlice().Interface() // []interface{}{1,2}
NewSet(1,2).ToSlice().Int() // []int{1,2}, nil
NewSet(1,2).ToSlice().Int64() // []int64{1,2}, nil
```

### String() string

Formatted output string.

Examples:
```go
NewSet(1,2,3).String() // {1,2,3}
```

