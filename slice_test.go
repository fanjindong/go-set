package set

import (
	"reflect"
	"testing"
)

func TestSlice_Bool(t *testing.T) {
	tests := []struct {
		name    string
		s       Slice
		want    []bool
		wantErr bool
	}{
		{name: "1", s: Slice{}, want: []bool{}},
		{name: "2", s: Slice{true}, want: []bool{true}},
		{name: "3", s: Slice{false}, want: []bool{false}},
		{name: "4", s: Slice{true, false}, want: []bool{true, false}},
		{name: "5", s: Slice{1, true, false}, want: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Bool()
			if (err != nil) != tt.wantErr {
				t.Errorf("Bool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bool() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlice_Complex128(t *testing.T) {
	tests := []struct {
		name    string
		s       Slice
		want    []complex128
		wantErr bool
	}{
		{name: "1", s: Slice{}, want: []complex128{}},
		{name: "2", s: Slice{complex128(1)}, want: []complex128{1}},
		{name: "3", s: Slice{complex128(1), complex128(2)}, want: []complex128{1, 2}},
		{name: "4", s: Slice{1, true, false}, want: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Complex128()
			if (err != nil) != tt.wantErr {
				t.Errorf("Complex128() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Complex128() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlice_Complex64(t *testing.T) {
	tests := []struct {
		name    string
		s       Slice
		want    []complex64
		wantErr bool
	}{
		{name: "1", s: Slice{}, want: []complex64{}},
		{name: "2", s: Slice{complex64(1)}, want: []complex64{1}},
		{name: "3", s: Slice{complex64(1), complex64(2)}, want: []complex64{1, 2}},
		{name: "4", s: Slice{complex128(1)}, want: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Complex64()
			if (err != nil) != tt.wantErr {
				t.Errorf("Complex64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Complex64() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlice_Float32(t *testing.T) {
	tests := []struct {
		name    string
		s       Slice
		want    []float32
		wantErr bool
	}{
		{name: "1", s: Slice{}, want: []float32{}},
		{name: "2", s: Slice{float32(1)}, want: []float32{1}},
		{name: "3", s: Slice{float32(1), float32(2)}, want: []float32{1, 2}},
		{name: "4", s: Slice{1}, want: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Float32()
			if (err != nil) != tt.wantErr {
				t.Errorf("Float32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Float32() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlice_Float64(t *testing.T) {
	tests := []struct {
		name    string
		s       Slice
		want    []float64
		wantErr bool
	}{
		{name: "1", s: Slice{}, want: []float64{}},
		{name: "2", s: Slice{float64(1)}, want: []float64{1}},
		{name: "3", s: Slice{float64(1), float64(2)}, want: []float64{1, 2}},
		{name: "4", s: Slice{1}, want: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Float64()
			if (err != nil) != tt.wantErr {
				t.Errorf("Float64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Float64() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlice_Int(t *testing.T) {
	tests := []struct {
		name    string
		s       Slice
		want    []int
		wantErr bool
	}{
		{name: "1", s: Slice{}, want: []int{}},
		{name: "2", s: Slice{1}, want: []int{1}},
		{name: "3", s: Slice{1, 2}, want: []int{1, 2}},
		{name: "4", s: Slice{1, "a"}, want: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Int()
			if (err != nil) != tt.wantErr {
				t.Errorf("Int() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlice_Int16(t *testing.T) {
	tests := []struct {
		name    string
		s       Slice
		want    []int16
		wantErr bool
	}{
		{name: "1", s: Slice{}, want: []int16{}},
		{name: "2", s: Slice{int16(1)}, want: []int16{1}},
		{name: "3", s: Slice{int16(1), int16(2)}, want: []int16{1, 2}},
		{name: "4", s: Slice{1}, want: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Int16()
			if (err != nil) != tt.wantErr {
				t.Errorf("Int16() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int16() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlice_Int32(t *testing.T) {
	tests := []struct {
		name    string
		s       Slice
		want    []int32
		wantErr bool
	}{
		{name: "1", s: Slice{}, want: []int32{}},
		{name: "2", s: Slice{int32(1)}, want: []int32{1}},
		{name: "3", s: Slice{int32(1), int32(2)}, want: []int32{1, 2}},
		{name: "4", s: Slice{1}, want: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Int32()
			if (err != nil) != tt.wantErr {
				t.Errorf("Int32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int32() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlice_Int64(t *testing.T) {
	tests := []struct {
		name    string
		s       Slice
		want    []int64
		wantErr bool
	}{
		{name: "1", s: Slice{}, want: []int64{}},
		{name: "2", s: Slice{int64(1)}, want: []int64{1}},
		{name: "3", s: Slice{int64(1), int64(2)}, want: []int64{1, 2}},
		{name: "4", s: Slice{1}, want: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Int64()
			if (err != nil) != tt.wantErr {
				t.Errorf("Int64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int64() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlice_Int8(t *testing.T) {
	tests := []struct {
		name    string
		s       Slice
		want    []int8
		wantErr bool
	}{
		{name: "1", s: Slice{}, want: []int8{}},
		{name: "2", s: Slice{int8(1)}, want: []int8{1}},
		{name: "3", s: Slice{int8(1), int8(2)}, want: []int8{1, 2}},
		{name: "4", s: Slice{1}, want: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Int8()
			if (err != nil) != tt.wantErr {
				t.Errorf("Int8() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int8() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlice_Interface(t *testing.T) {
	tests := []struct {
		name string
		s    Slice
		want []interface{}
	}{
		{name: "1", s: Slice{}, want: []interface{}{}},
		{name: "2", s: Slice{1}, want: []interface{}{1}},
		{name: "3", s: Slice{1, 2}, want: []interface{}{1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Interface(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Interface() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlice_String(t *testing.T) {
	tests := []struct {
		name    string
		s       Slice
		want    []string
		wantErr bool
	}{
		{name: "1", s: Slice{}, want: []string{}},
		{name: "2", s: Slice{"a"}, want: []string{"a"}},
		{name: "3", s: Slice{"a", "1"}, want: []string{"a", "1"}},
		{name: "4", s: Slice{1}, want: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.String()
			if (err != nil) != tt.wantErr {
				t.Errorf("String() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("String() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlice_Uint(t *testing.T) {
	tests := []struct {
		name    string
		s       Slice
		want    []uint
		wantErr bool
	}{
		{name: "1", s: Slice{}, want: []uint{}},
		{name: "2", s: Slice{uint(1)}, want: []uint{1}},
		{name: "3", s: Slice{uint(1), uint(2)}, want: []uint{1, 2}},
		{name: "4", s: Slice{1}, want: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Uint()
			if (err != nil) != tt.wantErr {
				t.Errorf("Uint() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlice_Uint16(t *testing.T) {
	tests := []struct {
		name    string
		s       Slice
		want    []uint16
		wantErr bool
	}{
		{name: "1", s: Slice{}, want: []uint16{}},
		{name: "2", s: Slice{uint16(1)}, want: []uint16{1}},
		{name: "3", s: Slice{uint16(1), uint16(2)}, want: []uint16{1, 2}},
		{name: "4", s: Slice{1}, want: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Uint16()
			if (err != nil) != tt.wantErr {
				t.Errorf("Uint16() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint16() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlice_Uint32(t *testing.T) {
	tests := []struct {
		name    string
		s       Slice
		want    []uint32
		wantErr bool
	}{
		{name: "1", s: Slice{}, want: []uint32{}},
		{name: "2", s: Slice{uint32(1)}, want: []uint32{1}},
		{name: "3", s: Slice{uint32(1), uint32(2)}, want: []uint32{1, 2}},
		{name: "4", s: Slice{1}, want: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Uint32()
			if (err != nil) != tt.wantErr {
				t.Errorf("Uint32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint32() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlice_Uint64(t *testing.T) {
	tests := []struct {
		name    string
		s       Slice
		want    []uint64
		wantErr bool
	}{
		{name: "1", s: Slice{}, want: []uint64{}},
		{name: "2", s: Slice{uint64(1)}, want: []uint64{1}},
		{name: "3", s: Slice{uint64(1), uint64(2)}, want: []uint64{1, 2}},
		{name: "4", s: Slice{1}, want: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Uint64()
			if (err != nil) != tt.wantErr {
				t.Errorf("Uint64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint64() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlice_Uint8(t *testing.T) {
	tests := []struct {
		name    string
		s       Slice
		want    []uint8
		wantErr bool
	}{
		{name: "1", s: Slice{}, want: []uint8{}},
		{name: "2", s: Slice{uint8(1)}, want: []uint8{1}},
		{name: "3", s: Slice{uint8(1), uint8(2)}, want: []uint8{1, 2}},
		{name: "4", s: Slice{1}, want: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Uint8()
			if (err != nil) != tt.wantErr {
				t.Errorf("Uint8() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint8() got = %v, want %v", got, tt.want)
			}
		})
	}
}
