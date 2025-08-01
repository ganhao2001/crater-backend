/*
Copyright 2022 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This file was copied from client-go/tools/cache/heap.go and modified
// for our non thread-safe heap

package queue

import (
	"testing"
)

func testHeapObjectKeyFunc(obj any) string {
	return obj.(testHeapObject).name
}

type testHeapObject struct {
	name string
	val  any
}

func mkHeapObj(name string, val any) testHeapObject {
	return testHeapObject{name: name, val: val}
}

func compareInts(val1, val2 any) bool {
	first := val1.(testHeapObject).val.(int)
	second := val2.(testHeapObject).val.(int)
	return first < second
}

// TestHeapBasic tests Heap invariant
func TestHeapBasic(t *testing.T) {
	h := New(testHeapObjectKeyFunc, compareInts)
	const amount = 500
	var i int

	for i = amount; i > 0; i-- {
		h.PushOrUpdate(mkHeapObj(string([]rune{'a', rune(i)}), i))
	}

	// Make sure that the numbers are popped in ascending order.
	prevNum := 0
	for i := 0; i < amount; i++ {
		obj := h.Pop()
		num := obj.(testHeapObject).val.(int)
		// All the items must be sorted.
		if prevNum > num {
			t.Errorf("got %v out of order, last was %v", obj, prevNum)
		}
		prevNum = num
	}
}

// Tests Heap.PushOrUpdate and ensures that heap invariant is preserved after adding items.
func TestHeap_Add(t *testing.T) {
	h := New(testHeapObjectKeyFunc, compareInts)
	h.PushOrUpdate(mkHeapObj("foo", 10))
	h.PushOrUpdate(mkHeapObj("bar", 1))
	h.PushOrUpdate(mkHeapObj("baz", 11))
	h.PushOrUpdate(mkHeapObj("zab", 30))
	h.PushOrUpdate(mkHeapObj("foo", 13)) // This updates "foo".

	item := h.Pop()
	if e, a := 1, item.(testHeapObject).val; a != e {
		t.Fatalf("expected %d, got %d", e, a)
	}
	item = h.Pop()
	if e, a := 11, item.(testHeapObject).val; a != e {
		t.Fatalf("expected %d, got %d", e, a)
	}
	h.Delete("baz")                      // Nothing is deleted.
	h.PushOrUpdate(mkHeapObj("foo", 14)) // foo is updated.
	item = h.Pop()
	if e, a := 14, item.(testHeapObject).val; a != e {
		t.Fatalf("expected %d, got %d", e, a)
	}
	item = h.Pop()
	if e, a := 30, item.(testHeapObject).val; a != e {
		t.Fatalf("expected %d, got %d", e, a)
	}
}

// TestHeap_PushIfNotPresent tests Heap.PushIfNotPresent and ensures that heap
// invariant is preserved after adding items.
func TestHeap_PushIfNotPresent(t *testing.T) {
	h := New(testHeapObjectKeyFunc, compareInts)
	_ = h.PushIfNotPresent(mkHeapObj("foo", 10))
	_ = h.PushIfNotPresent(mkHeapObj("bar", 1))
	_ = h.PushIfNotPresent(mkHeapObj("baz", 11))
	_ = h.PushIfNotPresent(mkHeapObj("zab", 30))
	_ = h.PushIfNotPresent(mkHeapObj("foo", 13)) // This is not added.

	if length := len(h.items); length != 4 {
		t.Errorf("unexpected number of items: %d", length)
	}
	if val := h.items["foo"].obj.(testHeapObject).val; val != 10 {
		t.Errorf("unexpected value: %d", val)
	}
	item := h.Pop()
	if e, a := 1, item.(testHeapObject).val; a != e {
		t.Fatalf("expected %d, got %d", e, a)
	}
	item = h.Pop()
	if e, a := 10, item.(testHeapObject).val; a != e {
		t.Fatalf("expected %d, got %d", e, a)
	}
	// bar is already popped. Let's add another one.
	_ = h.PushIfNotPresent(mkHeapObj("bar", 14))
	item = h.Pop()
	if e, a := 11, item.(testHeapObject).val; a != e {
		t.Fatalf("expected %d, got %d", e, a)
	}
	item = h.Pop()
	if e, a := 14, item.(testHeapObject).val; a != e {
		t.Fatalf("expected %d, got %d", e, a)
	}
}

// TestHeap_PushOrUpdate tests Heap.PushOrUpdate and ensures that heap
// invariant is preserved after adding/updating items.
func TestHeap_PushOrUpdate(t *testing.T) {
	h := New(testHeapObjectKeyFunc, compareInts)
	h.PushOrUpdate(mkHeapObj("foo", 100))
	h.PushOrUpdate(mkHeapObj("baz", 20))
	h.PushOrUpdate(mkHeapObj("foo", 1)) // This behaviors as update.
	h.PushOrUpdate(mkHeapObj("zab", 8)) // This behaviors as add.

	if length := len(h.items); length != 3 {
		t.Errorf("unexpected number of items: %d", length)
	}
	item := h.Pop()
	if e, a := 1, item.(testHeapObject).val; a != e {
		t.Fatalf("expected %d, got %d", e, a)
	}
	item = h.Pop()
	if e, a := 8, item.(testHeapObject).val; a != e {
		t.Fatalf("expected %d, got %d", e, a)
	}
}

// TestHeap_Delete tests Heap.Delete and ensures that heap invariant is
// preserved after deleting items.
func TestHeap_Delete(t *testing.T) {
	h := New(testHeapObjectKeyFunc, compareInts)
	h.PushOrUpdate(mkHeapObj("foo", 10))
	h.PushOrUpdate(mkHeapObj("bar", 1))
	h.PushOrUpdate(mkHeapObj("bal", 31))
	h.PushOrUpdate(mkHeapObj("baz", 11))

	// Delete head. Delete should work with "key" and doesn't care about the value.
	h.Delete("bar")
	item := h.Pop()
	if e, a := 10, item.(testHeapObject).val; a != e {
		t.Fatalf("expected %d, got %d", e, a)
	}
	h.PushOrUpdate(mkHeapObj("zab", 30))
	h.PushOrUpdate(mkHeapObj("faz", 30))
	length := h.Len()
	h.Delete("non-existent")
	if length != h.Len() {
		t.Fatalf("Didn't expect any item removal")
	}
	// Delete tail.
	h.Delete("bal")
	// Delete one of the items with value 30.
	h.Delete("zab")
	item = h.Pop()
	if e, a := 11, item.(testHeapObject).val; a != e {
		t.Fatalf("expected %d, got %d", e, a)
	}
	item = h.Pop()
	if e, a := 30, item.(testHeapObject).val; a != e {
		t.Fatalf("expected %d, got %d", e, a)
	}
	if h.Len() != 0 {
		t.Fatalf("expected an empty heap.")
	}
}

// TestHeap_Update tests Heap.PushOrUpdate and ensures that heap invariant is
// preserved after adding items.
func TestHeap_Update(t *testing.T) {
	h := New(testHeapObjectKeyFunc, compareInts)
	h.PushOrUpdate(mkHeapObj("foo", 10))
	h.PushOrUpdate(mkHeapObj("bar", 1))
	h.PushOrUpdate(mkHeapObj("bal", 31))
	h.PushOrUpdate(mkHeapObj("baz", 11))

	// Update an item to a value that should push it to the head.
	h.PushOrUpdate(mkHeapObj("baz", 0))
	if h.keys[0] != "baz" || h.items["baz"].index != 0 {
		t.Fatalf("expected baz to be at the head")
	}
	item := h.Pop()
	if e, a := 0, item.(testHeapObject).val; a != e {
		t.Fatalf("expected %d, got %d", e, a)
	}
	// Update bar to push it farther back in the queue.
	h.PushOrUpdate(mkHeapObj("bar", 100))
	if h.keys[0] != "foo" || h.items["foo"].index != 0 {
		t.Fatalf("expected foo to be at the head")
	}
}

// TestHeap_Get tests Heap.Get.
func TestHeap_Get(t *testing.T) {
	h := New(testHeapObjectKeyFunc, compareInts)
	h.PushOrUpdate(mkHeapObj("foo", 10))
	h.PushOrUpdate(mkHeapObj("bar", 1))
	h.PushOrUpdate(mkHeapObj("bal", 31))
	h.PushOrUpdate(mkHeapObj("baz", 11))

	// Get works with the key.
	obj := h.Get(mkHeapObj("baz", 0))
	if obj == nil || obj.(testHeapObject).val != 11 {
		t.Fatalf("unexpected error in getting element")
	}
	// Get non-existing object.
	if obj = h.Get(mkHeapObj("non-existing", 0)); obj != nil {
		t.Fatalf("didn't expect to get any object")
	}
}

// TestHeap_GetByKey tests Heap.GetByKey and is very similar to TestHeap_Get.
func TestHeap_GetByKey(t *testing.T) {
	h := New(testHeapObjectKeyFunc, compareInts)
	h.PushOrUpdate(mkHeapObj("foo", 10))
	h.PushOrUpdate(mkHeapObj("bar", 1))
	h.PushOrUpdate(mkHeapObj("bal", 31))
	h.PushOrUpdate(mkHeapObj("baz", 11))

	obj := h.GetByKey("baz")
	if obj == nil || obj.(testHeapObject).val != 11 {
		t.Fatalf("unexpected error in getting element")
	}
	// Get non-existing object.
	if obj = h.GetByKey("non-existing"); obj != nil {
		t.Fatalf("didn't expect to get any object")
	}
}

// TestQueue_List tests Queue.List function.
func TestQueue_List(t *testing.T) {
	h := New(testHeapObjectKeyFunc, compareInts)
	list := h.List()
	if len(list) != 0 {
		t.Errorf("expected an empty list")
	}
	items := []testHeapObject{
		{"foo", 10},
		{"bar", 1},
		{"bal", 30},
		{"baz", 11},
		{"faz", 30},
	}
	for _, i := range items {
		h.PushOrUpdate(i)
	}
	list = h.List()
	if len(list) != len(items) {
		t.Errorf("expected %d items, got %d", len(items), len(list))
	}
	for _, obj := range list {
		heapObj := obj.(testHeapObject)
		//nolint:gocritic // TODO: what is this?
		// v, ok := items[heapObj.name]
		t.Logf("%v:%v\n", heapObj.name, heapObj.val)
		//nolint:gocritic // TODO: what is this?
		// if !ok || v != heapObj.val {
		// 	t.Errorf("unexpected item in the list: %v", heapObj)
		// }
	}
	h.Reorder()
	t.Logf("---------\n")
	for h.Len() > 0 {
		heapObj := h.Pop()
		t.Logf("%v:%v\n", heapObj.(testHeapObject).name, heapObj.(testHeapObject).val)
	}
}
