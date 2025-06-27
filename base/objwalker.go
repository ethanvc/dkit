package base

import (
	"reflect"
)

type ObjWalker struct {
}

func NewObjWalker() *ObjWalker {
	return &ObjWalker{}
}

type VisitResult int

const (
	VisitResultContinue VisitResult = iota
	VisitResultStop
	VisitResultSkipCurrentValue
)

type VisitFunc func(obj, key, val reflect.Value) (VisitResult, reflect.Value)

func (w *ObjWalker) Walk(obj any, fn VisitFunc) {
	w.walk(reflect.ValueOf(obj), fn)
}

func (w *ObjWalker) walk(obj reflect.Value, fn VisitFunc) bool {
	for {
		kind := obj.Kind()
		switch kind {
		case reflect.Map:
			return w.walkMap(obj, fn)
		case reflect.Array:
		case reflect.Slice:
			return w.walkSlice(obj, fn)
		case reflect.Interface, reflect.Pointer:
			obj = obj.Elem()
		default:
			return true
		}
	}
}

func (w *ObjWalker) walkMap(obj reflect.Value, fn VisitFunc) bool {
	iter := obj.MapRange()
	for iter.Next() {
		key := iter.Key()
		val := iter.Value()
		result, newVal := fn(obj, key, val)
		if val != newVal {
			obj.SetMapIndex(key, newVal)
		}
		switch result {
		case VisitResultStop:
			return false
		case VisitResultSkipCurrentValue:
			continue
		case VisitResultContinue:
			if !w.walk(newVal, fn) {
				return false
			}
		}
	}
	return true
}

func (w *ObjWalker) walkSlice(obj reflect.Value, fn VisitFunc) bool {
	length := obj.Len()
	for i := 0; i < length; i++ {
		val := obj.Index(i)
		result, newVal := fn(obj, reflect.ValueOf(i), val)
		if val != newVal {
			if newVal.Type().AssignableTo(val.Type()) {
				val.Set(newVal)
			} else {
				continue
			}
		}
		switch result {
		case VisitResultStop:
			return false
		case VisitResultContinue:
			if !w.walk(newVal, fn) {
				return false
			}
		case VisitResultSkipCurrentValue:
			continue
		}
	}
	return true
}
