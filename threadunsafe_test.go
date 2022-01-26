package set

import (
	"reflect"
	"testing"
)

func TestNewThreadUnsafeSet(t *testing.T) {
	type args struct {
		elems []interface{}
	}
	tests := []struct {
		name string
		args args
		want ISet
	}{
		{name: "1", args: args{elems: nil}, want: &threadUnsafeSet{}},
		{name: "2", args: args{elems: []interface{}{int64(1)}}, want: &threadUnsafeSet{int64(1): struct{}{}}},
		{name: "3", args: args{elems: []interface{}{"a"}}, want: &threadUnsafeSet{"a": struct{}{}}},
		{name: "4", args: args{elems: []interface{}{int64(1), "a"}}, want: &threadUnsafeSet{int64(1): struct{}{}, "a": struct{}{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewThreadUnsafeSet(tt.args.elems...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewThreadUnsafeSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_threadUnsafeSet_Empty(t *testing.T) {
	tests := []struct {
		name string
		s    ISet
		want bool
	}{
		{name: "1", s: NewThreadUnsafeSet(), want: true},
		{name: "1", s: NewSet(), want: true},
		{name: "2", s: NewThreadUnsafeSet(1), want: false},
		{name: "2", s: NewSet(1), want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Empty(); got != tt.want {
				t.Errorf("Empty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_threadUnsafeSet_Singleton(t *testing.T) {
	tests := []struct {
		name string
		s    ISet
		want bool
	}{
		{name: "1", s: NewThreadUnsafeSet(), want: false},
		{name: "2", s: NewThreadUnsafeSet(1), want: true},
		{name: "3", s: NewThreadUnsafeSet("a"), want: true},
		{name: "4", s: NewThreadUnsafeSet("a", 1), want: false},
		{name: "1", s: NewSet(), want: false},
		{name: "2", s: NewSet(1), want: true},
		{name: "3", s: NewSet("a"), want: true},
		{name: "4", s: NewSet("a", 1), want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Singleton(); got != tt.want {
				t.Errorf("Singleton() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_threadUnsafeSet_Cardinality(t *testing.T) {
	tests := []struct {
		name string
		s    ISet
		want int
	}{
		{name: "1", s: NewThreadUnsafeSet(), want: 0},
		{name: "2", s: NewThreadUnsafeSet(1), want: 1},
		{name: "3", s: NewThreadUnsafeSet("a"), want: 1},
		{name: "4", s: NewThreadUnsafeSet("a", 1), want: 2},
		{name: "1", s: NewSet(), want: 0},
		{name: "2", s: NewSet(1), want: 1},
		{name: "3", s: NewSet("a"), want: 1},
		{name: "4", s: NewSet("a", 1), want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Cardinality(); got != tt.want {
				t.Errorf("Cardinality() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_threadUnsafeSet_ToSlice(t *testing.T) {
	tests := []struct {
		name string
		s    ISet
		want []interface{}
	}{
		{name: "1", s: NewThreadUnsafeSet(), want: []interface{}{}},
		{name: "2", s: NewThreadUnsafeSet(1), want: []interface{}{1}},
		{name: "3", s: NewThreadUnsafeSet("a"), want: []interface{}{"a"}},
		{name: "4", s: NewThreadUnsafeSet("a", 1), want: []interface{}{"a", 1}},
		{name: "1", s: NewSet(), want: []interface{}{}},
		{name: "2", s: NewSet(1), want: []interface{}{1}},
		{name: "3", s: NewSet("a"), want: []interface{}{"a"}},
		{name: "4", s: NewSet("a", 1), want: []interface{}{"a", 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ToSlice().Interface(); !reflect.DeepEqual(slice2map(got), slice2map(tt.want)) {
				t.Errorf("ToSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_threadUnsafeSet_Adds(t *testing.T) {
	type args struct {
		elems []interface{}
	}
	tests := []struct {
		name string
		s    ISet
		args args
		want bool
	}{
		{name: "1", s: NewThreadUnsafeSet(), args: args{elems: []interface{}{1}}, want: true},
		{name: "2", s: NewThreadUnsafeSet(1), args: args{elems: []interface{}{1}}, want: false},
		{name: "3", s: NewThreadUnsafeSet("a"), args: args{elems: []interface{}{1}}, want: true},
		{name: "4", s: NewThreadUnsafeSet("a", 1), args: args{elems: []interface{}{1}}, want: false},
		{name: "1", s: NewSet(), args: args{elems: []interface{}{1}}, want: true},
		{name: "2", s: NewSet(1), args: args{elems: []interface{}{1}}, want: false},
		{name: "3", s: NewSet("a"), args: args{elems: []interface{}{1}}, want: true},
		{name: "4", s: NewSet("a", 1), args: args{elems: []interface{}{1}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Adds(tt.args.elems...); got != tt.want {
				t.Errorf("Adds() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_threadUnsafeSet_Removes(t *testing.T) {
	type args struct {
		elems []interface{}
	}
	tests := []struct {
		name string
		s    ISet
		args args
		want bool
	}{
		{name: "1", s: NewThreadUnsafeSet(), args: args{elems: []interface{}{1}}, want: false},
		{name: "2", s: NewThreadUnsafeSet(1), args: args{elems: []interface{}{1}}, want: true},
		{name: "3", s: NewThreadUnsafeSet("a"), args: args{elems: []interface{}{1}}, want: false},
		{name: "4", s: NewThreadUnsafeSet("a", 1), args: args{elems: []interface{}{1}}, want: true},
		{name: "1", s: NewSet(), args: args{elems: []interface{}{1}}, want: false},
		{name: "2", s: NewSet(1), args: args{elems: []interface{}{1}}, want: true},
		{name: "3", s: NewSet("a"), args: args{elems: []interface{}{1}}, want: false},
		{name: "4", s: NewSet("a", 1), args: args{elems: []interface{}{1}}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.Removes(tt.args.elems...)
			if got != tt.want {
				t.Errorf("Removes() %v, want %v", got, tt.want)
			}
			for _, elem := range tt.args.elems {
				if tt.s.Contains(elem) {
					t.Errorf("Removes() %v err", elem)
				}
			}
		})
	}
}

func Test_threadUnsafeSet_IsSub(t *testing.T) {
	type args struct {
		other ISet
	}
	tests := []struct {
		name string
		s    ISet
		args args
		want bool
	}{
		{name: "1", s: NewThreadUnsafeSet(), args: args{other: NewThreadUnsafeSet()}, want: true},
		{name: "2", s: NewThreadUnsafeSet(1), args: args{other: NewThreadUnsafeSet()}, want: false},
		{name: "3", s: NewThreadUnsafeSet("a"), args: args{other: NewThreadUnsafeSet("a", 1)}, want: true},
		{name: "4", s: NewThreadUnsafeSet("a", 1), args: args{other: NewThreadUnsafeSet(1)}, want: false},
		{name: "1", s: NewSet(), args: args{other: NewSet()}, want: true},
		{name: "2", s: NewSet(1), args: args{other: NewSet()}, want: false},
		{name: "3", s: NewSet("a"), args: args{other: NewSet("a", 1)}, want: true},
		{name: "4", s: NewSet("a", 1), args: args{other: NewSet(1)}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.IsSub(tt.args.other); got != tt.want {
				t.Errorf("IsSub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_threadUnsafeSet_Unions(t *testing.T) {
	type args struct {
		others []ISet
	}
	tests := []struct {
		name string
		s    ISet
		args args
		want ISet
	}{
		{name: "1", s: NewThreadUnsafeSet(), args: args{others: []ISet{NewThreadUnsafeSet()}}, want: NewThreadUnsafeSet()},
		{name: "2", s: NewThreadUnsafeSet(), args: args{others: []ISet{NewThreadUnsafeSet(1)}}, want: NewThreadUnsafeSet(1)},
		{name: "3", s: NewThreadUnsafeSet(1), args: args{others: []ISet{NewThreadUnsafeSet(1)}}, want: NewThreadUnsafeSet(1)},
		{name: "4", s: NewThreadUnsafeSet(1, 2), args: args{others: []ISet{NewThreadUnsafeSet(1)}}, want: NewThreadUnsafeSet(1, 2)},
		{name: "5", s: NewThreadUnsafeSet(1, 2), args: args{others: []ISet{NewThreadUnsafeSet(3), NewThreadUnsafeSet(4)}}, want: NewThreadUnsafeSet(1, 2, 3, 4)},
		{name: "6", s: NewThreadUnsafeSet(1, 2), args: args{others: []ISet{NewThreadUnsafeSet(3), NewThreadUnsafeSet(1)}}, want: NewThreadUnsafeSet(1, 2, 3)},
		{name: "1", s: NewSet(), args: args{others: []ISet{NewSet()}}, want: NewSet()},
		{name: "2", s: NewSet(), args: args{others: []ISet{NewSet(1)}}, want: NewSet(1)},
		{name: "3", s: NewSet(1), args: args{others: []ISet{NewSet(1)}}, want: NewSet(1)},
		{name: "4", s: NewSet(1, 2), args: args{others: []ISet{NewSet(1)}}, want: NewSet(1, 2)},
		{name: "5", s: NewSet(1, 2), args: args{others: []ISet{NewSet(3), NewSet(4)}}, want: NewSet(1, 2, 3, 4)},
		{name: "6", s: NewSet(1, 2), args: args{others: []ISet{NewSet(3), NewSet(1)}}, want: NewSet(1, 2, 3)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Unions(tt.args.others...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Unions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_threadUnsafeSet_Intersections(t *testing.T) {
	type args struct {
		others []ISet
	}
	tests := []struct {
		name string
		s    ISet
		args args
		want ISet
	}{
		{name: "1", s: NewThreadUnsafeSet(), args: args{others: []ISet{NewThreadUnsafeSet()}}, want: NewThreadUnsafeSet()},
		{name: "2", s: NewThreadUnsafeSet(), args: args{others: []ISet{NewThreadUnsafeSet(1)}}, want: NewThreadUnsafeSet()},
		{name: "3", s: NewThreadUnsafeSet(1), args: args{others: []ISet{NewThreadUnsafeSet(1)}}, want: NewThreadUnsafeSet(1)},
		{name: "4", s: NewThreadUnsafeSet(1, 2), args: args{others: []ISet{NewThreadUnsafeSet(1)}}, want: NewThreadUnsafeSet(1)},
		{name: "5", s: NewThreadUnsafeSet(1, 2), args: args{others: []ISet{NewThreadUnsafeSet(3), NewThreadUnsafeSet(4)}}, want: NewThreadUnsafeSet()},
		{name: "6", s: NewThreadUnsafeSet(1, 2), args: args{others: []ISet{NewThreadUnsafeSet(1, 3), NewThreadUnsafeSet(1)}}, want: NewThreadUnsafeSet(1)},
		{name: "1", s: NewSet(), args: args{others: []ISet{NewSet()}}, want: NewSet()},
		{name: "2", s: NewSet(), args: args{others: []ISet{NewSet(1)}}, want: NewSet()},
		{name: "3", s: NewSet(1), args: args{others: []ISet{NewSet(1)}}, want: NewSet(1)},
		{name: "4", s: NewSet(1, 2), args: args{others: []ISet{NewSet(1)}}, want: NewSet(1)},
		{name: "5", s: NewSet(1, 2), args: args{others: []ISet{NewSet(3), NewSet(4)}}, want: NewSet()},
		{name: "6", s: NewSet(1, 2), args: args{others: []ISet{NewSet(1, 3), NewSet(1)}}, want: NewSet(1)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Intersections(tt.args.others...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Intersections() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_threadUnsafeSet_Complements(t *testing.T) {
	type args struct {
		others []ISet
	}
	tests := []struct {
		name string
		s    ISet
		args args
		want ISet
	}{
		{name: "1", s: NewThreadUnsafeSet(), args: args{others: []ISet{NewThreadUnsafeSet()}}, want: NewThreadUnsafeSet()},
		{name: "2", s: NewThreadUnsafeSet(), args: args{others: []ISet{NewThreadUnsafeSet(1)}}, want: NewThreadUnsafeSet()},
		{name: "3", s: NewThreadUnsafeSet(1), args: args{others: []ISet{NewThreadUnsafeSet(1)}}, want: NewThreadUnsafeSet()},
		{name: "4", s: NewThreadUnsafeSet(1, 2), args: args{others: []ISet{NewThreadUnsafeSet(1)}}, want: NewThreadUnsafeSet(2)},
		{name: "5", s: NewThreadUnsafeSet(1, 2), args: args{others: []ISet{NewThreadUnsafeSet(3), NewThreadUnsafeSet(4)}}, want: NewThreadUnsafeSet(1, 2)},
		{name: "6", s: NewThreadUnsafeSet(1, 2), args: args{others: []ISet{NewThreadUnsafeSet(1, 3), NewThreadUnsafeSet(1)}}, want: NewThreadUnsafeSet(2)},
		{name: "1", s: NewSet(), args: args{others: []ISet{NewSet()}}, want: NewSet()},
		{name: "2", s: NewSet(), args: args{others: []ISet{NewSet(1)}}, want: NewSet()},
		{name: "3", s: NewSet(1), args: args{others: []ISet{NewSet(1)}}, want: NewSet()},
		{name: "4", s: NewSet(1, 2), args: args{others: []ISet{NewSet(1)}}, want: NewSet(2)},
		{name: "5", s: NewSet(1, 2), args: args{others: []ISet{NewSet(3), NewSet(4)}}, want: NewSet(1, 2)},
		{name: "6", s: NewSet(1, 2), args: args{others: []ISet{NewSet(1, 3), NewSet(1)}}, want: NewSet(2)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Complements(tt.args.others...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Complements() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_threadUnsafeSet_Clear(t *testing.T) {
	tests := []struct {
		name string
		s    ISet
	}{
		{name: "1", s: NewThreadUnsafeSet()},
		{name: "2", s: NewThreadUnsafeSet(1)},
		{name: "3", s: NewThreadUnsafeSet(1, "a")},
		{name: "1", s: NewSet()},
		{name: "2", s: NewSet(1)},
		{name: "3", s: NewSet(1, "a")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Clear()
			if tt.s.Cardinality() != 0 {
				t.Errorf("Clear() %v", tt.s)
			}
		})
	}
}

func Test_threadUnsafeSet_Clone(t *testing.T) {
	tests := []struct {
		name string
		s    ISet
		want ISet
	}{
		{name: "1", s: NewThreadUnsafeSet(), want: NewThreadUnsafeSet()},
		{name: "2", s: NewThreadUnsafeSet(1), want: NewThreadUnsafeSet(1)},
		{name: "3", s: NewThreadUnsafeSet(1, "a"), want: NewThreadUnsafeSet(1, "a")},
		{name: "1", s: NewSet(), want: NewSet()},
		{name: "2", s: NewSet(1), want: NewSet(1)},
		{name: "3", s: NewSet(1, "a"), want: NewSet(1, "a")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Clone(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Clone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_threadUnsafeSet_Equal(t *testing.T) {
	type args struct {
		other ISet
	}
	tests := []struct {
		name string
		s    ISet
		args args
		want bool
	}{
		{name: "1", s: NewThreadUnsafeSet(), args: args{other: NewThreadUnsafeSet()}, want: true},
		{name: "2", s: NewThreadUnsafeSet(1), args: args{other: NewThreadUnsafeSet()}, want: false},
		{name: "3", s: NewThreadUnsafeSet("a"), args: args{other: NewThreadUnsafeSet("a", 1)}, want: false},
		{name: "4", s: NewThreadUnsafeSet("a", 1), args: args{other: NewThreadUnsafeSet(1)}, want: false},
		{name: "5", s: NewThreadUnsafeSet("a", 1), args: args{other: NewThreadUnsafeSet(1, "a")}, want: true},
		{name: "1", s: NewSet(), args: args{other: NewSet()}, want: true},
		{name: "2", s: NewSet(1), args: args{other: NewSet()}, want: false},
		{name: "3", s: NewSet("a"), args: args{other: NewSet("a", 1)}, want: false},
		{name: "4", s: NewSet("a", 1), args: args{other: NewSet(1)}, want: false},
		{name: "5", s: NewSet("a", 1), args: args{other: NewSet(1, "a")}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Equal(tt.args.other); got != tt.want {
				t.Errorf("Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_threadUnsafeSet_Pop(t *testing.T) {
	tests := []struct {
		name string
		s    ISet
		want interface{}
	}{
		{name: "1", s: NewThreadUnsafeSet(), want: nil},
		{name: "2", s: NewThreadUnsafeSet(1), want: 1},
		{name: "3", s: NewThreadUnsafeSet(1, "a"), want: 1},
		{name: "4", s: NewThreadUnsafeSet(1, "a", 0), want: 1},
		{name: "1", s: NewSet(), want: nil},
		{name: "2", s: NewSet(1), want: 1},
		{name: "3", s: NewSet(1, "a"), want: 1},
		{name: "4", s: NewSet(1, "a", 0), want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Pop(); !reflect.DeepEqual(got, tt.want) && tt.s.Contains(got) {
				t.Errorf("Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_threadUnsafeSet_String(t *testing.T) {
	tests := []struct {
		name string
		s    ISet
		want ISet
	}{
		{name: "1", s: NewThreadUnsafeSet(), want: NewThreadUnsafeSet("{}")},
		{name: "2", s: NewThreadUnsafeSet(1), want: NewThreadUnsafeSet("{1}")},
		{name: "3", s: NewThreadUnsafeSet(1, "a"), want: NewThreadUnsafeSet("{1,a}", "{a,1}")},
		{name: "4", s: NewThreadUnsafeSet(1, "a", 0), want: NewThreadUnsafeSet("{1,a,0}", "{1,0,a}", "{0,a,1}", "{0,1,a}", "{a,0,1}", "{a,1,0}")},
		{name: "1", s: NewSet(), want: NewSet("{}")},
		{name: "2", s: NewSet(1), want: NewSet("{1}")},
		{name: "3", s: NewSet(1, "a"), want: NewSet("{1,a}", "{a,1}")},
		{name: "4", s: NewSet(1, "a", 0), want: NewSet("{1,a,0}", "{1,0,a}", "{0,a,1}", "{0,1,a}", "{a,0,1}", "{a,1,0}")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.String(); !tt.want.Contains(got) {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func slice2map(s []interface{}) map[interface{}]struct{} {
	m := make(map[interface{}]struct{}, len(s))
	for _, item := range s {
		m[item] = struct{}{}
	}
	return m
}
