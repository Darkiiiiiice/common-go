package container

import (
	"reflect"
	"testing"
)

func Test_ArrayList(t *testing.T) {

	t.Run("ArrayList Insert Of Head", func(t *testing.T) {
		var list = InitArrayList(0)
		list.Insert(0, 1)
		list.Insert(0, 2)
		list.Insert(0, 3)
		list.Insert(0, 4)
		list.Insert(0, 5)
		list.Insert(0, 6)
		list.Insert(0, 7)
		list.Insert(0, 8)
		list.Insert(0, 9)

		var slice = []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
		if !array_equals_with_slice(list, slice) {
			t.Errorf("array_list = %+v, result = %+v", list, slice)
		}
	})

	t.Run("ArrayList Insert Of Index", func(t *testing.T) {
		var list = InitArrayList(0)
		list.Insert(0, 1)
		list.Insert(1, 2)
		list.Insert(1, 3)
		list.Insert(1, 4)
		list.Insert(1, 5)
		list.Insert(1, 6)
		list.Insert(2, 7)
		list.Insert(1, 8)
		list.Insert(1, 9)

		var slice = []int{1, 9, 8, 6, 7, 5, 4, 3, 2}
		if !array_equals_with_slice(list, slice) {
			t.Errorf("array_list = %+v, result = %+v", list, slice)
		}
	})

	t.Run("ArrayList Insert Of Tail", func(t *testing.T) {
		var list = InitArrayList(0)
		list.Insert(0, 1)
		list.Insert(1, 2)
		list.Insert(2, 3)
		list.Insert(3, 4)
		list.Insert(4, 5)
		list.Insert(5, 6)
		list.Insert(6, 7)
		list.Insert(7, 8)
		list.Insert(10, 9)

		var slice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		if !array_equals_with_slice(list, slice) {
			t.Errorf("array_list = %+v, result = %+v", list, slice)
		}
	})

	t.Run("ArrayList Clear", func(t *testing.T) {
		var list = InitArrayList(0)
		list.Insert(0, 1)
		list.Insert(1, 2)
		list.Insert(2, 3)
		list.Insert(3, 4)
		list.Insert(4, 5)
		list.Insert(5, 6)
		list.Insert(6, 7)
		list.Insert(7, 8)
		list.Insert(10, 9)

		list.Clear()
		if list.Length() != 0 || list.capacity != 0 {
			t.Errorf("array_list = %+v, result = %+v", list, 0)
		}
	})

	t.Run("ArrayList Get", func(t *testing.T) {
		var list = InitArrayList(0)
		list.Insert(0, 1)
		list.Insert(1, 2)
		list.Insert(2, 3)
		list.Insert(3, 4)
		list.Insert(4, 5)
		list.Insert(5, 6)
		list.Insert(6, 7)
		list.Insert(7, 8)
		list.Insert(10, 9)

		var v = list.Get(6)
		if v != 7 {
			t.Errorf("array_list = %+v, result = %+v", list, v)
		}

		v = list.Get(22)
		if v != nil {
			t.Errorf("array_list = %+v, result = %+v", list, v)
		}

		v = list.Get(0)
		if v != 1 {
			t.Errorf("array_list = %+v, result = %+v", list, v)
		}

	})

	t.Run("ArrayList Insert Of Delete", func(t *testing.T) {
		var list = InitArrayList(0)
		list.Insert(0, 1)
		list.Insert(1, 2)
		list.Insert(2, 3)
		list.Insert(3, 4)
		list.Insert(4, 5)
		list.Insert(5, 6)
		list.Insert(6, 7)
		list.Insert(7, 8)
		list.Insert(10, 9)

		list.Delete(0)
		var slice = []int{ 2, 3, 4, 5, 6, 7, 8, 9}
		if !array_equals_with_slice(list, slice) {
			t.Errorf("array_list = %+v, result = %+v", list, slice)
		}

		list .Delete(-1)
		slice = []int{2, 3, 4, 5, 6, 7, 8, 9}
		if !array_equals_with_slice(list, slice) {
			t.Errorf("array_list = %+v, result = %+v", list, slice)
		}

		list .Delete(8)
		slice = []int{2, 3, 4, 5, 6, 7, 8, 9}
		if !array_equals_with_slice(list, slice) {
			t.Errorf("array_list = %+v, result = %+v", list, slice)
		}

		list .Delete(7)
		slice = []int{2, 3, 4, 5, 6, 7, 8}
		if !array_equals_with_slice(list, slice) {
			t.Errorf("array_list = %+v, result = %+v", list, slice)
		}

		list .Delete(3)
		slice = []int{2, 3, 4, 6, 7, 8}
		if !array_equals_with_slice(list, slice) {
			t.Errorf("array_list = %+v, result = %+v", list, slice)
		}
	})
}

func array_equals_with_slice(arrayList *ArrayList, slice []int) bool {

	if arrayList.length != len(slice) {
		return false
	}

	for i, v := range slice {
		if !reflect.DeepEqual(arrayList.nodes[i].value, v) {
			return false
		}
	}
	return true
}
