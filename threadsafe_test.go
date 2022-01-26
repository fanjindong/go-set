package set

import (
	"math/rand"
	"os"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
)

var elems []int

func TestMain(m *testing.M) {
	runtime.GOMAXPROCS(2)
	elems = rand.Perm(1000)
	os.Exit(m.Run())
}

func Test_threadSafeSet_Adds(t *testing.T) {
	s := NewSet()
	wg := sync.WaitGroup{}
	for i := range elems {
		wg.Add(1)
		go func(i int) {
			s.Adds(elems[i])
			wg.Done()
		}(i)
	}
	wg.Wait()
	if s.Cardinality() != len(elems) {
		t.Errorf("Adds.Cardinality() = %v, want %v", s.Cardinality(), len(elems))
	}
	for i := range elems {
		if !s.Contains(elems[i]) {
			t.Errorf("Adds.Contains() = %v, want %v", nil, elems[i])
		}
	}
}

func Test_threadSafeSet_Cardinality(t *testing.T) {
	s := NewSet()
	wg := sync.WaitGroup{}
	for i := range elems {
		wg.Add(2)
		go func(i int) {
			s.Adds(elems[i])
			wg.Done()
		}(i)
		go func() {
			s.Cardinality()
			wg.Done()
		}()
	}
	wg.Wait()
	if s.Cardinality() != len(elems) {
		t.Errorf("Cardinality() = %v, want %v", s.Cardinality(), len(elems))
	}
}

func Test_threadSafeSet_Clear(t *testing.T) {
	s := NewSet()
	wg := sync.WaitGroup{}
	for i := range elems {
		wg.Add(2)
		go func(i int) {
			s.Adds(elems[i])
			wg.Done()
		}(i)
		go func() {
			s.Clear()
			wg.Done()
		}()
	}
	wg.Wait()
	s.Clear()
	if s.Cardinality() != 0 {
		t.Errorf("Clear() = %v, want %v", s.Cardinality(), 0)
	}
}

func Test_threadSafeSet_Clone(t *testing.T) {
	s := NewSet()
	wg := sync.WaitGroup{}
	for i := range elems {
		wg.Add(2)
		go func(i int) {
			s.Adds(elems[i])
			wg.Done()
		}(i)
		go func() {
			s.Clone()
			wg.Done()
		}()
	}
	wg.Wait()
	sCopy := s.Clone()
	if !sCopy.Equal(s) {
		t.Errorf("Clone() = %v, want %v", sCopy.String(), s.String())
	}
}

func Test_threadSafeSet_Contains(t *testing.T) {
	s := NewSet()
	wg := sync.WaitGroup{}
	for i := range elems {
		wg.Add(2)
		go func(i int) {
			s.Adds(elems[i])
			wg.Done()
		}(i)
		go func(i int) {
			s.Contains(elems[i])
			wg.Done()
		}(i)
	}
	wg.Wait()
	wg = sync.WaitGroup{}
	for i := range elems {
		wg.Add(1)
		go func(i int) {
			if !s.Contains(elems[i]) {
				t.Errorf("Contains() = %v, want %v", nil, elems[i])
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func Test_threadSafeSet_Pop(t *testing.T) {
	s := NewSet()
	wg := sync.WaitGroup{}
	var hasPop int64
	for i := range elems {
		wg.Add(2)
		go func(i int) {
			s.Adds(elems[i])
			wg.Done()
		}(i)
		go func() {
			if v := s.Pop(); v != nil {
				atomic.AddInt64(&hasPop, 1)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	if s.Cardinality()+int(hasPop) != len(elems) {
		t.Errorf("Cardinality() = %v, want %v", s.Cardinality()+int(hasPop), len(elems))
	}
}

func Test_threadSafeSet_Removes(t *testing.T) {
	s := NewSet()
	wg := sync.WaitGroup{}
	var hasRemoved int64
	for i := range elems {
		wg.Add(2)
		go func(i int) {
			s.Adds(elems[i])
			wg.Done()
		}(i)
		go func(i int) {
			if v := s.Removes(elems[i]); v {
				atomic.AddInt64(&hasRemoved, 1)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	if s.Cardinality()+int(hasRemoved) != len(elems) {
		t.Errorf("Cardinality() = %v, want %v", s.Cardinality()+int(hasRemoved), len(elems))
	}
}

func Test_threadSafeSet_ToSlice(t *testing.T) {
	s := NewSet()
	wg := sync.WaitGroup{}
	for i := range elems {
		wg.Add(2)
		go func(i int) {
			s.Adds(elems[i])
			wg.Done()
		}(i)
		go func() {
			s.ToSlice()
			wg.Done()
		}()
	}
	wg.Wait()
	if len(s.ToSlice().Interface()) != len(elems) {
		t.Errorf("ToSlice() = %v, want %v", len(s.ToSlice().Interface()), len(elems))
	}
}
