package set

import (
	"reflect"
	"testing"
)

func TestNewMapSet(t *testing.T) {
	type args struct {
		elems []interface{}
	}
	tests := []struct {
		name string
		args args
		want ISet
	}{
		{name: "1", args: args{elems: nil}, want: mapSet{}},
		{name: "2", args: args{elems: []interface{}{int64(1)}}, want: mapSet{int64(1): struct{}{}}},
		{name: "3", args: args{elems: []interface{}{"a"}}, want: mapSet{"a": struct{}{}}},
		{name: "4", args: args{elems: []interface{}{int64(1), "a"}}, want: mapSet{int64(1): struct{}{}, "a": struct{}{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMapSet(tt.args.elems...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMapSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mapSet_Empty(t *testing.T) {
	tests := []struct {
		name string
		s    mapSet
		want bool
	}{
		{name: "1", s: NewMapSet().(mapSet), want: true},
		{name: "2", s: NewMapSet(1).(mapSet), want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Empty(); got != tt.want {
				t.Errorf("Empty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mapSet_Singleton(t *testing.T) {
	tests := []struct {
		name string
		s    mapSet
		want bool
	}{
		{name: "1", s: NewMapSet().(mapSet), want: false},
		{name: "2", s: NewMapSet(1).(mapSet), want: true},
		{name: "3", s: NewMapSet("a").(mapSet), want: true},
		{name: "4", s: NewMapSet("a", 1).(mapSet), want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Singleton(); got != tt.want {
				t.Errorf("Singleton() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mapSet_Cardinality(t *testing.T) {
	tests := []struct {
		name string
		s    mapSet
		want int
	}{
		{name: "1", s: NewMapSet().(mapSet), want: 0},
		{name: "2", s: NewMapSet(1).(mapSet), want: 1},
		{name: "3", s: NewMapSet("a").(mapSet), want: 1},
		{name: "4", s: NewMapSet("a", 1).(mapSet), want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Cardinality(); got != tt.want {
				t.Errorf("Cardinality() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mapSet_ToSlice(t *testing.T) {
	tests := []struct {
		name string
		s    mapSet
		want []interface{}
	}{
		{name: "1", s: NewMapSet().(mapSet), want: []interface{}{}},
		{name: "2", s: NewMapSet(1).(mapSet), want: []interface{}{1}},
		{name: "3", s: NewMapSet("a").(mapSet), want: []interface{}{"a"}},
		{name: "4", s: NewMapSet("a", 1).(mapSet), want: []interface{}{"a", 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ToSlice(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mapSet_Adds(t *testing.T) {
	type args struct {
		elems []interface{}
	}
	tests := []struct {
		name string
		s    mapSet
		args args
		want bool
	}{
		{name: "1", s: NewMapSet().(mapSet), args: args{elems: []interface{}{1}}, want: true},
		{name: "2", s: NewMapSet(1).(mapSet), args: args{elems: []interface{}{1}}, want: false},
		{name: "3", s: NewMapSet("a").(mapSet), args: args{elems: []interface{}{1}}, want: true},
		{name: "4", s: NewMapSet("a", 1).(mapSet), args: args{elems: []interface{}{1}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Adds(tt.args.elems...); got != tt.want {
				t.Errorf("Adds() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mapSet_Removes(t *testing.T) {
	type args struct {
		elems []interface{}
	}
	tests := []struct {
		name string
		s    mapSet
		args args
	}{
		{name: "1", s: NewMapSet().(mapSet), args: args{elems: []interface{}{1}}},
		{name: "2", s: NewMapSet(1).(mapSet), args: args{elems: []interface{}{1}}},
		{name: "3", s: NewMapSet("a").(mapSet), args: args{elems: []interface{}{1}}},
		{name: "4", s: NewMapSet("a", 1).(mapSet), args: args{elems: []interface{}{1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Removes(tt.args.elems...)
			for _, elem := range tt.args.elems {
				if tt.s.Contains(elem) {
					t.Errorf("Removes() %v err", elem)
				}
			}
		})
	}
}

func Test_mapSet_Iterator(t *testing.T) {
	tests := []struct {
		name string
		s    mapSet
		want ISet
	}{
		{name: "1", s: NewMapSet().(mapSet), want: NewMapSet()},
		{name: "2", s: NewMapSet(1).(mapSet), want: NewMapSet(1)},
		{name: "3", s: NewMapSet("a").(mapSet), want: NewMapSet("a")},
		{name: "4", s: NewMapSet("a", 1).(mapSet), want: NewMapSet(1, "a")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIter := tt.s.Iterator()
			defer gotIter.Stop()
			var got = NewMapSet()
			for elem := range gotIter.C {
				got.Adds(elem)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Iterator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mapSet_IsSub(t *testing.T) {
	type args struct {
		other ISet
	}
	tests := []struct {
		name string
		s    mapSet
		args args
		want bool
	}{
		{name: "1", s: NewMapSet().(mapSet), args: args{other: NewMapSet()}, want: true},
		{name: "2", s: NewMapSet(1).(mapSet), args: args{other: NewMapSet()}, want: false},
		{name: "3", s: NewMapSet("a").(mapSet), args: args{other: NewMapSet("a", 1)}, want: true},
		{name: "4", s: NewMapSet("a", 1).(mapSet), args: args{other: NewMapSet(1)}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.IsSub(tt.args.other); got != tt.want {
				t.Errorf("IsSub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mapSet_Unions(t *testing.T) {
	type args struct {
		others []ISet
	}
	tests := []struct {
		name string
		s    mapSet
		args args
		want ISet
	}{
		{name: "1", s: NewMapSet().(mapSet), args: args{others: []ISet{NewMapSet()}}, want: NewMapSet()},
		{name: "2", s: NewMapSet().(mapSet), args: args{others: []ISet{NewMapSet(1)}}, want: NewMapSet(1)},
		{name: "3", s: NewMapSet(1).(mapSet), args: args{others: []ISet{NewMapSet(1)}}, want: NewMapSet(1)},
		{name: "4", s: NewMapSet(1, 2).(mapSet), args: args{others: []ISet{NewMapSet(1)}}, want: NewMapSet(1, 2)},
		{name: "5", s: NewMapSet(1, 2).(mapSet), args: args{others: []ISet{NewMapSet(3), NewMapSet(4)}}, want: NewMapSet(1, 2, 3, 4)},
		{name: "6", s: NewMapSet(1, 2).(mapSet), args: args{others: []ISet{NewMapSet(3), NewMapSet(1)}}, want: NewMapSet(1, 2, 3)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Unions(tt.args.others...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Unions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mapSet_Intersections(t *testing.T) {
	type args struct {
		others []ISet
	}
	tests := []struct {
		name string
		s    mapSet
		args args
		want ISet
	}{
		{name: "1", s: NewMapSet().(mapSet), args: args{others: []ISet{NewMapSet()}}, want: NewMapSet()},
		{name: "2", s: NewMapSet().(mapSet), args: args{others: []ISet{NewMapSet(1)}}, want: NewMapSet()},
		{name: "3", s: NewMapSet(1).(mapSet), args: args{others: []ISet{NewMapSet(1)}}, want: NewMapSet(1)},
		{name: "4", s: NewMapSet(1, 2).(mapSet), args: args{others: []ISet{NewMapSet(1)}}, want: NewMapSet(1)},
		{name: "5", s: NewMapSet(1, 2).(mapSet), args: args{others: []ISet{NewMapSet(3), NewMapSet(4)}}, want: NewMapSet()},
		{name: "6", s: NewMapSet(1, 2).(mapSet), args: args{others: []ISet{NewMapSet(1, 3), NewMapSet(1)}}, want: NewMapSet(1)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Intersections(tt.args.others...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Intersections() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mapSet_Complements(t *testing.T) {
	type args struct {
		others []ISet
	}
	tests := []struct {
		name string
		s    mapSet
		args args
		want ISet
	}{
		{name: "1", s: NewMapSet().(mapSet), args: args{others: []ISet{NewMapSet()}}, want: NewMapSet()},
		{name: "2", s: NewMapSet().(mapSet), args: args{others: []ISet{NewMapSet(1)}}, want: NewMapSet()},
		{name: "3", s: NewMapSet(1).(mapSet), args: args{others: []ISet{NewMapSet(1)}}, want: NewMapSet()},
		{name: "4", s: NewMapSet(1, 2).(mapSet), args: args{others: []ISet{NewMapSet(1)}}, want: NewMapSet(2)},
		{name: "5", s: NewMapSet(1, 2).(mapSet), args: args{others: []ISet{NewMapSet(3), NewMapSet(4)}}, want: NewMapSet(1, 2)},
		{name: "6", s: NewMapSet(1, 2).(mapSet), args: args{others: []ISet{NewMapSet(1, 3), NewMapSet(1)}}, want: NewMapSet(2)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Complements(tt.args.others...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Complements() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mapSet_Clear(t *testing.T) {
	tests := []struct {
		name string
		s    mapSet
	}{
		{name: "1", s: NewMapSet().(mapSet)},
		{name: "2", s: NewMapSet(1).(mapSet)},
		{name: "3", s: NewMapSet(1, "a").(mapSet)},
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

func Test_mapSet_Clone(t *testing.T) {
	tests := []struct {
		name string
		s    mapSet
		want ISet
	}{
		{name: "1", s: NewMapSet().(mapSet), want: NewMapSet()},
		{name: "2", s: NewMapSet(1).(mapSet), want: NewMapSet(1)},
		{name: "3", s: NewMapSet(1, "a").(mapSet), want: NewMapSet(1, "a")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Clone(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Clone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mapSet_Equal(t *testing.T) {
	type args struct {
		other ISet
	}
	tests := []struct {
		name string
		s    mapSet
		args args
		want bool
	}{
		{name: "1", s: NewMapSet().(mapSet), args: args{other: NewMapSet()}, want: true},
		{name: "2", s: NewMapSet(1).(mapSet), args: args{other: NewMapSet()}, want: false},
		{name: "3", s: NewMapSet("a").(mapSet), args: args{other: NewMapSet("a", 1)}, want: false},
		{name: "4", s: NewMapSet("a", 1).(mapSet), args: args{other: NewMapSet(1)}, want: false},
		{name: "5", s: NewMapSet("a", 1).(mapSet), args: args{other: NewMapSet(1, "a")}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Equal(tt.args.other); got != tt.want {
				t.Errorf("Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mapSet_Pop(t *testing.T) {
	tests := []struct {
		name string
		s    mapSet
		want interface{}
	}{
		{name: "1", s: NewMapSet().(mapSet), want: nil},
		{name: "2", s: NewMapSet(1).(mapSet), want: 1},
		{name: "3", s: NewMapSet(1, "a").(mapSet), want: 1},
		{name: "4", s: NewMapSet(1, "a", 0).(mapSet), want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Pop(); !reflect.DeepEqual(got, tt.want) && tt.s.Contains(got) {
				t.Errorf("Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mapSet_String(t *testing.T) {
	tests := []struct {
		name string
		s    mapSet
		want string
	}{
		{name: "1", s: NewMapSet().(mapSet), want: "{}"},
		{name: "2", s: NewMapSet(1).(mapSet), want: "{1}"},
		{name: "3", s: NewMapSet(1, "a").(mapSet), want: "{1,a}"},
		{name: "4", s: NewMapSet(1, "a", 0).(mapSet), want: "{1,a,0}"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
