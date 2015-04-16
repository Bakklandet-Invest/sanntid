// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// MODIFIED FOR EDUCATIONAL USE
// The Project Intellivator 2000 Authors
// ADDED SORTING

// Package list implements a doubly linked list.
//
// To iterate over a list (where l is a *LinkedList):
//	for e := l.Front(); e != nil; e = e.Next() {
//		// do something with e.Value
//	}
//
package util

// Element is an element of a linked list.
type Element struct {
	// Next and previous pointers in the doubly-linked list of elements.
	// To simplify the implementation, internally a list l is implemented
	// as a ring, such that &l.root is both the next element of the last
	// list element (l.Back()) and the previous element of the first list
	// element (l.Front()).
	next, prev *Element

	// The list to which this element belongs.
	list *LinkedList

	// The value stored with this element.
	Value interface{}
}

// Next returns the next list element or nil.
func (e *Element) Next() *Element {
	if p := e.next; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

// Prev returns the previous list element or nil.
func (e *Element) Prev() *Element {
	if p := e.prev; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

// List represents a doubly linked list.
// The zero value for List is an empty list ready to use.
type LinkedList struct {
	root Element // sentinel list element, only &root, root.prev, and root.next are used
	len  int     // current list length excluding (this) sentinel element
}

// Init initializes or clears list l.
func (l *LinkedList) Init() *LinkedList {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.len = 0
	return l
}

// New returns an initialized list.
func New() *LinkedList { return new(LinkedList).Init() }

// Len returns the number of elements of list l.
// The complexity is O(1).
func (l *LinkedList) Len() int { return l.len }

// Front returns the first element of list l or nil
func (l *LinkedList) Front() *Element {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}

// Back returns the last element of list l or nil.
func (l *LinkedList) Back() *Element {
	if l.len == 0 {
		return nil
	}
	return l.root.prev
}

// lazyInit lazily initializes a zero List value.
func (l *LinkedList) lazyInit() {
	if l.root.next == nil {
		l.Init()
	}
}

// insert inserts e after at, increments l.len, and returns e.
func (l *LinkedList) insert(e, at *Element) *Element {
	n := at.next
	at.next = e
	e.prev = at
	e.next = n
	n.prev = e
	e.list = l
	l.len++
	return e
}

// insertValue is a convenience wrapper for insert(&Element{Value: v}, at).
func (l *LinkedList) insertValue(v interface{}, at *Element) *Element {
	return l.insert(&Element{Value: v}, at)
}

// remove removes e from its list, decrements l.len, and returns e.
func (l *LinkedList) remove(e *Element) *Element {
	e.prev.next = e.next
	e.next.prev = e.prev
	e.next = nil // avoid memory leaks
	e.prev = nil // avoid memory leaks
	e.list = nil
	l.len--
	return e
}

// Remove removes e from l if e is an element of list l.
// It returns the element value e.Value.
func (l *LinkedList) Remove(e *Element) interface{} {
	if e.list == l {
		// if e.list == l, l must have been initialized when e was inserted
		// in l or l == nil (e is a zero Element) and l.remove will crash
		l.remove(e)
	}
	return e.Value
}

// PushFront inserts a new element e with value v at the front of list l and returns e.
func (l *LinkedList) PushFront(v interface{}) *Element {
	l.lazyInit()
	return l.insertValue(v, &l.root)
}

// PushBack inserts a new element e with value v at the back of list l and returns e.
func (l *LinkedList) PushBack(v interface{}) *Element {
	l.lazyInit()
	return l.insertValue(v, l.root.prev)
}

// InsertBefore inserts a new element e with value v immediately before mark and returns e.
// If mark is not an element of l, the list is not modified.
func (l *LinkedList) InsertBefore(v interface{}, mark *Element) *Element {
	if mark.list != l {
		return nil
	}
	// see comment in List.Remove about initialization of l
	return l.insertValue(v, mark.prev)
}

// InsertAfter inserts a new element e with value v immediately after mark and returns e.
// If mark is not an element of l, the list is not modified.
func (l *LinkedList) InsertAfter(v interface{}, mark *Element) *Element {
	if mark.list != l {
		return nil
	}
	// see comment in List.Remove about initialization of l
	return l.insertValue(v, mark)
}

// MoveToFront moves element e to the front of list l.
// If e is not an element of l, the list is not modified.
func (l *LinkedList) MoveToFront(e *Element) {
	if e.list != l || l.root.next == e {
		return
	}
	// see comment in List.Remove about initialization of l
	l.insert(l.remove(e), &l.root)
}

// MoveToBack moves element e to the back of list l.
// If e is not an element of l, the list is not modified.
func (l *LinkedList) MoveToBack(e *Element) {
	if e.list != l || l.root.prev == e {
		return
	}
	// see comment in List.Remove about initialization of l
	l.insert(l.remove(e), l.root.prev)
}

// MoveBefore moves element e to its new position before mark.
// If e is not an element of l, or e == mark, the list is not modified.
func (l *LinkedList) MoveBefore(e, mark *Element) {
	if e.list != l || e == mark {
		return
	}
	l.insert(l.remove(e), mark.prev)
}

// MoveAfter moves element e to its new position after mark.
// If e is not an element of l, or e == mark, the list is not modified.
func (l *LinkedList) MoveAfter(e, mark *Element) {
	if e.list != l || e == mark {
		return
	}
	l.insert(l.remove(e), mark)
}

// PushBackList inserts a copy of an other list at the back of list l.
// The lists l and other may be the same.
func (l *LinkedList) PushBackList(other *LinkedList) {
	l.lazyInit()
	for i, e := other.Len(), other.Front(); i > 0; i, e = i-1, e.Next() {
		l.insertValue(e.Value, l.root.prev)
	}
}

// PushFrontList inserts a copy of an other list at the front of list l.
// The lists l and other may be the same.
func (l *LinkedList) PushFrontList(other *LinkedList) {
	l.lazyInit()
	for i, e := other.Len(), other.Back(); i > 0; i, e = i-1, e.Prev() {
		l.insertValue(e.Value, &l.root)
	}
}

// ADDED BY US (OBAMA & CO)

func (l *LinkedList) Contains(val int) bool {
	for e := l.Front(); e != nil; e = e.Next() {
		if e.Value == val {
			return true
		}
	}
	return false
}

func (l *LinkedList) Sort(direction int) {
	for e := l.Front(); e != nil; e = e.Next() {
		for f := l.Back(); f != e; f = f.Prev() {
			if direction < 0 {
				// max sort
				if e.Value.(int) < f.Value.(int) {
					l.Swap(e,f)
					break
				}
			} else {
				// min sort
				if e.Value.(int) > f.Value.(int) {
					l.Swap(e,f)
					break
				}
			}
		}
	}
}
func (l *LinkedList) Swap(e,f *Element) {
	if l.len < 2 { return }
	
	if e.Prev() != nil{
		temp := e.Prev()
		l.MoveAfter(e,f)
		l.MoveAfter(f,temp)
	} else if e.Next() != nil {
		temp := e.Next()
		l.MoveAfter(e,f)
		l.MoveBefore(f,temp)
	}
}

/*unødvendig, bruka mæpp
func (l *LinkedList) FindElevID(id int) *Element{
    //returns a pointer to the element with elevatorID == id, or nil if not found
    //assuming the list consists of elevators
    
    for e := l.Front(); e != nil; e = e.Next() {
        elevator := e.Value.(*Elevator)
        if elevator.GetID() == id {
            return e
        }
    }
    
    return nil
    
}
*/
func (l *LinkedList) Find(id int) *Element {
    //return a pointer to the element
   for e := l.Front(); e != nil; e = e.Next() {
        
        if e.Value.(int) == id {
            return e
        }
    }
    
    return nil
}

func (l *LinkedList) GetFirstInt() int{
    return l.Front().Value.(int)
}

func (l *LinkedList) Get(index int) int {
	// returns the value stored in element i in the list (assumes int)
	if index >= l.Len() {return -1}
	
	e := l.Front()
	
	for index > 0 {
		e = e.Next()
		index--
	}
	
	return e.Value.(int)	
}

func GetValue(e *Element) int {
	if e == nil { return -1 }
	return e.Value.(int)
}