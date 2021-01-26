package container

import (
	"reflect"
	"testing"

	"github.com/mariomang/common-go/pkg/math/rand"
)

func Test_LinkedList(t *testing.T) {

	t.Run("InsertOfHead", func(tt *testing.T) {
		var originSlice = rand.RandomSliceInt(10, 0, 100)
		var linkedList = InitLinkedList()
		for _, v := range originSlice {
			linkedList.InsertOfHead(v)
		}

		var resultSlice = make([]int, 0)
		for _, v := range linkedList.ToSlice() {
			resultSlice = append(resultSlice, v.(int))
		}
		reverse(originSlice)
		if !reflect.DeepEqual(originSlice, resultSlice) {
			tt.Errorf("origin = %+v, result = %+v", originSlice, resultSlice)
		}
	})

	t.Run("InsertOfTail", func(tt *testing.T) {
		var originSlice = rand.RandomSliceInt(10, 0, 100)
		var linkedList = InitLinkedList()
		for _, v := range originSlice {
			linkedList.InsertOfTail(v)
		}

		var resultSlice = make([]int, 0)
		for _, v := range linkedList.ToSlice() {
			resultSlice = append(resultSlice, v.(int))
		}
		if !reflect.DeepEqual(originSlice, resultSlice) {
			tt.Errorf("origin = %+v, result = %+v", originSlice, resultSlice)
		}
	})

	t.Run("InsertOfIndex", func(tt *testing.T) {
		var originSlice = rand.RandomSliceInt(10, 0, 100)
		var linkedList = InitLinkedList()
		for _, v := range originSlice {
			linkedList.InsertOfIndex(v, 1)
		}

		var resultSlice = make([]int, 0)
		for _, v := range linkedList.ToSlice() {
			resultSlice = append(resultSlice, v.(int))
		}
		reverse(originSlice[1:])
		if !reflect.DeepEqual(originSlice, resultSlice) {
			tt.Errorf("origin = %+v, result = %+v", originSlice, resultSlice)
		}
	})

	t.Run("GetOfIndex", func(tt *testing.T) {
		var originSlice = rand.RandomSliceInt(10, 0, 100)
		var linkedList = InitLinkedList()
		for _, v := range originSlice {
			linkedList.InsertOfTail(v)
		}

		var resultSlice = make([]int, len(originSlice))
		for i := 0; i < len(originSlice); i++ {
			resultSlice[i] = linkedList.Get(i).(int)
		}
		if !reflect.DeepEqual(originSlice, resultSlice) {
			tt.Errorf("origin = %+v, result = %+v", originSlice, resultSlice)
		}
	})

	t.Run("RemoveOfIndex", func(tt *testing.T) {
		var originSlice = rand.RandomSliceInt(10, 0, 100)
		var linkedList = InitLinkedList()
		for _, v := range originSlice {
			linkedList.InsertOfTail(v)
		}

		var resultSlice = make([]int, len(originSlice))
		for i := 0; i < len(originSlice); i++ {
			resultSlice[i] = linkedList.Remove(0).(int)
		}
		if !reflect.DeepEqual(originSlice, resultSlice) {
			tt.Errorf("origin = %+v, result = %+v", originSlice, resultSlice)
		}
	})

	t.Run("RemoveOfIndex-middle", func(tt *testing.T) {
		var originSlice = rand.RandomSliceInt(10, 0, 100)
		var linkedList = InitLinkedList()
		for _, v := range originSlice {
			linkedList.InsertOfTail(v)
		}

		var r = linkedList.Remove(5).(int)
		if r != originSlice[5] {
			tt.Errorf("origin = %+v, result = %+v", originSlice, r)
		}

		var resultSlice = make([]int, len(originSlice) - 1)
		for i := 0; i < len(originSlice) - 1; i++ {
			resultSlice[i] = linkedList.Get(i).(int)
		}

		originSlice = append(originSlice[:5], originSlice[6:]...)

		if !reflect.DeepEqual(originSlice, resultSlice) {
			tt.Errorf("origin = %+v, result = %+v", originSlice, resultSlice)
		}
	})

	t.Run("RemoveOfIndex-end", func(tt *testing.T) {
		var originSlice = rand.RandomSliceInt(10, 0, 100)
		var linkedList = InitLinkedList()
		for _, v := range originSlice {
			linkedList.InsertOfTail(v)
		}

		var r = linkedList.Remove(9).(int)
		if r != originSlice[9] {
			tt.Errorf("origin = %+v, result = %+v", originSlice, r)
		}

		var resultSlice = make([]int, len(originSlice) - 1)
		for i := 0; i < len(originSlice) - 1; i++ {
			resultSlice[i] = linkedList.Get(i).(int)
		}


		if !reflect.DeepEqual(originSlice[:9], resultSlice) {
			tt.Errorf("origin = %+v, result = %+v", originSlice, resultSlice)
		}
	})
}
