package lists

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

type List struct {
	Values    []string
	Separator string
	Index     int
}

var OutOfBounds = errors.New("index out of bounds")

// ParseArray returns a new List.
func ParseArray(values []string, separator string) (list List) {
	return List{values, separator, -1}
}

// ParseString parses string to a new List.
func ParseString(values string, separator string) (list List) {
	list.Values = strings.Split(values, separator)
	list.Separator = separator
	list.Index = -1
	return
}

// ToString returns the list as string.
func (list *List) ToString() (tostr string) {

	for i, l := range list.Values {

		if i > 0 {
			tostr += list.Separator
		}

		tostr += l
	}

	return
}

// ToArray returns the list as array.
func (list *List) ToArray() []string {
	return list.Values
}

// Len returns the number of elements in the list.
func (list *List) Len() int {
	return len(list.Values)
}

// First returns the element at the first position.
func (list *List) First() string {
	list.Index = 0
	return list.Values[list.Index]
}

// Last returns the element at the last position.
func (list *List) Last() string {
	list.Index = list.Len() - 1
	return list.Values[list.Index]
}

// Find returns the index of the first occurrence of a value within a list or returns -1 if no value is found.
// The search is case-sensitive.
func (list *List) Find(value string) int {

	for i, v := range list.Values {
		if v == value {
			return i
		}
	}

	return -1
}

// Exist returns true or false when a value is or not found.
func (list *List) Exist(value string) bool {
	return list.Find(value) > -1
}

// Exist checks if there are more elements to move forward
func (list *List) c() bool {
	i := list.Len() - list.Index
	return i > 1
}

// HasNext checks if there are more elements to move forward
func (list *List) HasNext() bool {
	i := list.Len() - list.Index
	return i > 1
}

// Next returns the next element.
func (list *List) Next() (string, error) {

	// check index position
	if list.IsOutOfBound(list.Index) {
		return "", OutOfBounds
	}

	// go to the next element
	list.Index++

	// get value
	value, _ := list.GetAt(list.Index)

	// return it
	return value, nil
}

// MoveTo moves the index to a given position.
func (list *List) MoveTo(index int) error {

	// check list size
	if list.IsOutOfBound(index) {
		return OutOfBounds
	}

	// move index
	list.Index = index

	return nil
}

// Rewind moves the index before the first element.
func (list *List) Rewind() {
	list.Index = -1
}

// IsEmpty returns true if list has no element.
func (list *List) IsEmpty() bool {
	return list.Len() == 0
}

// IsOutOfBound returns true if index is out of bound.
func (list *List) IsOutOfBound(index int) bool {
	return index >= list.Len()
}

// Get returns the element at the current position.
func (list *List) Get() string {
	return list.Values[list.Index]
}

// GetAt returns the element at the specified position.
func (list *List) GetAt(position int) (string, error) {

	// check list size
	if list.IsOutOfBound(position) {
		return "", OutOfBounds
	}

	// return element
	return list.Values[position], nil
}

// Delete removes the element at the current position.
func (list *List) Delete() error {
	return list.DeleteAt(list.Index)
}

// DeleteAt returns list with element deleted at the specified position.
func (list *List) DeleteAt(position int) error {

	// check list size
	if list.IsOutOfBound(position) {
		return OutOfBounds
	}

	// remove element
	list.Values = append(list.Values[:position], list.Values[position+1:]...)

	// ok!
	return nil
}

// DeleteFirst removes the first occurrence of the given element.
func (list *List) DeleteFirst(element string) {
	if i := list.Find(element); i > -1 {
		list.DeleteAt(i)
	}
}

// DeleteLast removes the last occurrence of the given element.
func (list *List) DeleteLast(element string) {
	for i := list.Len() - 1; i >= 0; i-- {
		if value, _ := list.GetAt(i); value == element {
			list.DeleteAt(i)
			return
		}
	}
}

// DeleteAll removes matching elements.
func (list *List) DeleteAll(elements ...string) {
	for _, e := range elements {
		if i := list.Find(e); i > -1 {
			list.DeleteAt(i)  // delete element
			list.DeleteAll(e) // call it again
		}
	}
}

// Filter returns a new list with elements that match the given elements.
func (list *List) Filter(elements ...string) {

	var filtered []string

	for _, value := range list.Values {
		for _, e := range elements {
			if value == e {
				filtered = append(filtered, value)
			}
		}
	}

	list.Values = filtered
}

// Range returns the subset of a list.
func (list *List) Range(start, end int) List {
	return ParseArray(list.Values[start:end], list.Separator)
}

// From returns the subset of a list from an index until the end.
func (list *List) From(start int) List {
	return ParseArray(list.Values[start:], list.Separator)
}

// Until returns the subset of a list from its beginning until the given position.
func (list *List) Until(end int) List {
	return ParseArray(list.Values[:end], list.Separator)
}

// Set sets a new value assigned to its element based on current index.
func (list *List) Set(value string) {
	list.Values[list.Index] = value
}

// SetLast sets a new value assigned to the last element.
func (list *List) SetLast(value string) {
	list.Values[list.Len()-1] = value
}

// SetAt sets a new value assigned to its element based on a given position.
func (list *List) SetAt(position int, value string) error {

	// check list size
	if list.IsOutOfBound(position) {
		return OutOfBounds
	}

	// set new value
	list.Values[position] = value

	return nil
}

// Append adds values to the end of the list.
func (list *List) Append(values ...string) {
	for _, value := range values {
		list.Values = append(list.Values, value)
	}
}

// Insert adds a new element to the list at a specified position.
func (list *List) Insert(position int, value string) {
	list.Values = append(list.Values[:position], append([]string{value}, list.Values[position:]...)...)
}

// Swap swaps the position of two elements.
func (list *List) Swap(x, y int) {
	list.Values[x], list.Values[y] = list.Values[y], list.Values[x]
}

// Shift moves an element from a position to another.
func (list *List) Shift(from, to int) {
	list.Values = append(list.Values[:from], append(list.Values[from+1:to], append(list.Values[from:from+1], list.Values[to:]...)...)...)
}

// Split splits a element in 'n' new elements
func (list *List) Split(position int, interval ...int) {

	from := 0
	elements := make([]string, len(interval))

	for i := 0; i < len(interval); i++ {
		elements[i] = list.Values[position][from:interval[i]]
		from = interval[i]
	}

	list.Values = append(list.Values[:position], append(elements, list.Values[position+1:]...)...)
}

// Upper sets all elements as uppercase.
func (list *List) Upper() {
	for i, value := range list.Values {
		list.Values[i] = strings.ToUpper(value)
	}
}

// Lower sets all elements as lowercase.
func (list *List) Lower() {
	for i, value := range list.Values {
		list.Values[i] = strings.ToLower(value)
	}
}

// Sort sorts elements. (A to Z)
func (list *List) Sort() {
	var sorter sort.StringSlice = list.Values
	sorter.Sort()
	list.Values = sorter[:]
}

// Reverse reverses elements. (Z to A)
func (list *List) Reverse() {
	var sorter sort.StringSlice = list.Values
	sort.Sort(sort.Reverse(sorter))
	list.Values = sorter[:]
}

// Count counts the number of occurrences
func (list *List) Count(element string) (count int) {
	for _, value := range list.Values {
		if value == element {
			count++
		}
	}
	return
}

// Dedup deletes duplicated elements in the list.
func (list *List) Dedup() {
	for i := list.Len() - 1; i >= 0; i-- {
		value, _ := list.GetAt(i)
		if count := list.Count(value); count > 1 {
			list.DeleteLast(value)
		}
	}
}

// Quote wraps elements with a given rune.
func (list *List) Quote(r rune) {
	for i := range list.Values {
		list.Values[i] = fmt.Sprintf("%s%s%s", string(r), list.Values[i], string(r))
	}
}

// AppendNew appends a new value to the list only if the new element does not exist.
func (list *List) AppendNew(values ...string) {
	for _, value := range values {
		if !list.Exist(value) {
			list.Values = append(list.Values, value)
		}
	}
}

// Replacer replaces elements with a given map.
func (list *List) Replacer(fromTo map[string]string) {
	for i, value := range list.Values {
		if v, ok := fromTo[value]; ok {
			list.Values[i] = v
		}
	}
}
