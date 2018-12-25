package container

import (
	"reflect"
)

// Contains 用于检测在 source 中是否包含 target
// source 可以为 map, slice 或者 array
func Contains(source, target interface{}) bool {
	if source == nil {
		return false
	}

	var sourceValue = reflect.ValueOf(source)
	if sourceValue.IsNil() {
		return false
	}

	switch sourceValue.Kind() {
	case reflect.Array, reflect.Slice:
		var l = sourceValue.Len()
		for i := 0; i < l; i++ {
			if sourceValue.Index(i).Interface() == target {
				return true
			}
		}
	case reflect.Map:
		var targetValue = reflect.ValueOf(target)
		if targetValue.IsValid() {
			if sourceValue.MapIndex(targetValue).IsValid() {
				return true
			}
		}
	}
	return false
}

// Remove 用于移除 source 中的 target
// source 可以为 map 或者 slice
func Remove(source, target interface{}) bool {
	return removeAll(source, target, false)
}

func RemoveAll(source, target interface{}) bool {
	return removeAll(source, target, true)
}

func removeAll(source, target interface{}, removeAll bool) bool {
	if source == nil || target == nil {
		return false
	}

	var sourceValue = reflect.ValueOf(source)
	if sourceValue.IsNil() {
		return false
	}

	switch sourceValue.Kind() {
	case reflect.Ptr:
		var sliceValue = sourceValue.Elem()
		if sliceValue.Len() == 0 {
			return false
		}
		var indexList = indexs(sliceValue.Interface(), target, removeAll)
		if len(indexList) > 0 {
			var newSlice = reflect.MakeSlice(sliceValue.Type(), 0, 0)
			var beginIndex = 0
			var endIndex = 0
			for index, _ := range indexList {

				if index == 0 {
					beginIndex = 0
				} else {
					beginIndex = indexList[index-1] + 1
				}

				endIndex = indexList[index]
				newSlice = reflect.AppendSlice(newSlice, sliceValue.Slice(beginIndex, endIndex))
			}
			newSlice = reflect.AppendSlice(newSlice, sliceValue.Slice(endIndex+1, sliceValue.Len()))

			sliceValue.Set(newSlice)
			return true
		}
	case reflect.Map:
		sourceValue.SetMapIndex(reflect.ValueOf(target), reflect.Value{})
		return true
	}
	return false
}

// Index 用于获取 target 在 source 的索引位置
// source 可以为 slice 或者 array
func Index(source, target interface{}) int {
	var indexList = indexs(source, target, false)
	if len(indexList) > 0 {
		return indexList[0]
	}
	return -1
}

func Indexs(source, target interface{}) []int {
	return indexs(source, target, true)
}

func indexs(source, target interface{}, findAll bool) []int {
	if source == nil {
		return nil
	}

	var sourceValue = reflect.ValueOf(source)
	if sourceValue.IsNil() {
		return nil
	}

	var indexList []int
	switch sourceValue.Kind() {
	case reflect.Array, reflect.Slice:
		var l = sourceValue.Len()

		for i := 0; i < l; i++ {
			if sourceValue.Index(i).Interface() == target {
				indexList = append(indexList, i)
				if !findAll {
					break
				}
			}
		}
	}
	return indexList
}
