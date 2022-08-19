package sizeutil

import (
	"reflect"
)

func TrySetHeight[T any](m T, v int) (T, bool) {
	// 1) Does it have a (m T) SetHeight(int) method?
	ws, ok := any(m).(heightSetter)
	if ok {
		ws.SetHeight(v)
		return ws.(T), true
	}

	// 2) Does it have a (m *T) SetHeight(int) method?
	rv := reflect.ValueOf(m)
	switch rv.Kind() {
	case reflect.Struct:
		ptrV := reflect.New(rv.Type())
		ptrV.Elem().Set(rv)
		if ws, ok = ptrV.Interface().(heightSetter); ok {
			ws.SetHeight(v)
			return ptrV.Elem().Interface().(T), true
		}

		// 3) Does it have a Height field?
		wf := rv.FieldByName("Height")
		if wf.IsValid() && wf.CanSet() && wf.Type().Kind() == reflect.Int {
			wf.SetInt(int64(v))
			return wf.Interface().(T), true
		}
	}

	return m, false
}

func TrySetWidth[T any](m T, v int) (T, bool) {
	// 1) Does it have a (m T) SetWidth(int) method?
	ws, ok := any(m).(widthSetter)
	if ok {
		ws.SetWidth(v)
		return ws.(T), true
	}

	// 2) Does it have a (m *T) SetWidth(int) method?
	rv := reflect.ValueOf(m)
	switch rv.Kind() {
	case reflect.Struct:
		ptrV := reflect.New(rv.Type())
		ptrV.Elem().Set(rv)
		if ws, ok = ptrV.Interface().(widthSetter); ok {
			ws.SetWidth(v)
			return ptrV.Elem().Interface().(T), true
		}

		// 3) Does it have a Width field?
		wf := rv.FieldByName("Width")
		if wf.IsValid() && wf.CanSet() && wf.Type().Kind() == reflect.Int {
			wf.SetInt(int64(v))
			return wf.Interface().(T), true
		}
	}

	return m, false
}

type widthSetter interface {
	SetWidth(int)
}

type heightSetter interface {
	SetHeight(int)
}
