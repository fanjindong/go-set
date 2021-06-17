package set

import (
	"fmt"
)

type ISlice interface {
	Interface() []interface{}
	Int() ([]int, error)
	Int8() ([]int8, error)
	Int16() ([]int16, error)
	Int32() ([]int32, error)
	Int64() ([]int64, error)
	Uint() ([]uint, error)
	Uint8() ([]uint8, error)
	Uint16() ([]uint16, error)
	Uint32() ([]uint32, error)
	Uint64() ([]uint64, error)
	Float32() ([]float32, error)
	Float64() ([]float64, error)
	Complex64() ([]complex64, error)
	Complex128() ([]complex128, error)
	String() ([]string, error)
	Bool() ([]bool, error)
}

type Slice []interface{}

func (s Slice) Interface() []interface{} {
	return s
}
func (s Slice) Int() ([]int, error) {
	result := make([]int, 0, len(s))
	for i := 0; i < len(s); i++ {
		v, ok := s[i].(int)
		if !ok {
			return nil, fmt.Errorf("go-set: Slice Int() err, value: %+v", s[i])
		}
		result = append(result, v)
	}
	return result, nil
}
func (s Slice) Int8() ([]int8, error) {
	result := make([]int8, 0, len(s))
	for i := 0; i < len(s); i++ {
		v, ok := s[i].(int8)
		if !ok {
			return nil, fmt.Errorf("go-set: Slice Int8() err, value: %+v", s[i])
		}
		result = append(result, v)
	}
	return result, nil
}
func (s Slice) Int16() ([]int16, error) {
	result := make([]int16, 0, len(s))
	for i := 0; i < len(s); i++ {
		v, ok := s[i].(int16)
		if !ok {
			return nil, fmt.Errorf("go-set: Slice Int16() err, value: %+v", s[i])
		}
		result = append(result, v)
	}
	return result, nil
}
func (s Slice) Int32() ([]int32, error) {
	result := make([]int32, 0, len(s))
	for i := 0; i < len(s); i++ {
		v, ok := s[i].(int32)
		if !ok {
			return nil, fmt.Errorf("go-set: Slice Int32() err, value: %+v", s[i])
		}
		result = append(result, v)
	}
	return result, nil
}
func (s Slice) Int64() ([]int64, error) {
	result := make([]int64, 0, len(s))
	for i := 0; i < len(s); i++ {
		v, ok := s[i].(int64)
		if !ok {
			return nil, fmt.Errorf("go-set: Slice Int64() err, value: %+v", s[i])
		}
		result = append(result, v)
	}
	return result, nil
}

func (s Slice) Uint() ([]uint, error) {
	result := make([]uint, 0, len(s))
	for i := 0; i < len(s); i++ {
		v, ok := s[i].(uint)
		if !ok {
			return nil, fmt.Errorf("go-set: Slice Int() err, value: %+v", s[i])
		}
		result = append(result, v)
	}
	return result, nil
}
func (s Slice) Uint8() ([]uint8, error) {
	result := make([]uint8, 0, len(s))
	for i := 0; i < len(s); i++ {
		v, ok := s[i].(uint8)
		if !ok {
			return nil, fmt.Errorf("go-set: Slice Uint8() err, value: %+v", s[i])
		}
		result = append(result, v)
	}
	return result, nil
}
func (s Slice) Uint16() ([]uint16, error) {
	result := make([]uint16, 0, len(s))
	for i := 0; i < len(s); i++ {
		v, ok := s[i].(uint16)
		if !ok {
			return nil, fmt.Errorf("go-set: Slice Uint16() err, value: %+v", s[i])
		}
		result = append(result, v)
	}
	return result, nil
}
func (s Slice) Uint32() ([]uint32, error) {
	result := make([]uint32, 0, len(s))
	for i := 0; i < len(s); i++ {
		v, ok := s[i].(uint32)
		if !ok {
			return nil, fmt.Errorf("go-set: Slice Uint32() err, value: %+v", s[i])
		}
		result = append(result, v)
	}
	return result, nil
}
func (s Slice) Uint64() ([]uint64, error) {
	result := make([]uint64, 0, len(s))
	for i := 0; i < len(s); i++ {
		v, ok := s[i].(uint64)
		if !ok {
			return nil, fmt.Errorf("go-set: Slice Uint64() err, value: %+v", s[i])
		}
		result = append(result, v)
	}
	return result, nil
}
func (s Slice) Float32() ([]float32, error) {
	result := make([]float32, 0, len(s))
	for i := 0; i < len(s); i++ {
		v, ok := s[i].(float32)
		if !ok {
			return nil, fmt.Errorf("go-set: Slice Float32() err, value: %+v", s[i])
		}
		result = append(result, v)
	}
	return result, nil
}
func (s Slice) Float64() ([]float64, error) {
	result := make([]float64, 0, len(s))
	for i := 0; i < len(s); i++ {
		v, ok := s[i].(float64)
		if !ok {
			return nil, fmt.Errorf("go-set: Slice Float64() err, value: %+v", s[i])
		}
		result = append(result, v)
	}
	return result, nil
}
func (s Slice) Complex64() ([]complex64, error) {
	result := make([]complex64, 0, len(s))
	for i := 0; i < len(s); i++ {
		v, ok := s[i].(complex64)
		if !ok {
			return nil, fmt.Errorf("go-set: Slice Complex64() err, value: %+v", s[i])
		}
		result = append(result, v)
	}
	return result, nil
}
func (s Slice) Complex128() ([]complex128, error) {
	result := make([]complex128, 0, len(s))
	for i := 0; i < len(s); i++ {
		v, ok := s[i].(complex128)
		if !ok {
			return nil, fmt.Errorf("go-set: Slice Complex128() err, value: %+v", s[i])
		}
		result = append(result, v)
	}
	return result, nil
}
func (s Slice) String() ([]string, error) {
	result := make([]string, 0, len(s))
	for i := 0; i < len(s); i++ {
		v, ok := s[i].(string)
		if !ok {
			return nil, fmt.Errorf("go-set: Slice String() err, value: %+v", s[i])
		}
		result = append(result, v)
	}
	return result, nil
}
func (s Slice) Bool() ([]bool, error) {
	result := make([]bool, 0, len(s))
	for i := 0; i < len(s); i++ {
		v, ok := s[i].(bool)
		if !ok {
			return nil, fmt.Errorf("go-set: Slice Bool() err, value: %+v", s[i])
		}
		result = append(result, v)
	}
	return result, nil
}
