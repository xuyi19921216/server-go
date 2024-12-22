package reflection

import (
	"errors"
	"reflect"
)

// Max 使用反射比较两个值（目前支持基本数字类型），返回较大的值以及可能的错误
func Max(a, b interface{}) (interface{}, error) {
	va := reflect.ValueOf(a)
	vb := reflect.ValueOf(b)

	// 检查类型是否一致且是支持的数字类型
	if va.Type() != vb.Type() {
		return nil, errors.New("a and b are not of equal type")
	}
	switch va.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if va.Int() > vb.Int() {
			return a, nil
		}
		return b, nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if va.Uint() > vb.Uint() {
			return a, nil
		}
		return b, nil
	case reflect.Float32, reflect.Float64:
		if va.Float() > vb.Float() {
			return a, nil
		}
		return b, nil
	default:
		return nil, errors.New("unsupported kind")
	}
}
