package lists

import (
	"errors"
	"sort"
	"strings"
)

type List struct {
	Values    []string
	Separator string
	Index     int
}

var OutOfBounds = errors.New("index out of bounds")

/**
 ** Returns a new List.
 **/
func ParseArray(values []string, separator string) (list List) {
	return List{values, separator, -1}
}

/**
 ** Parse string to a new List.
 **/
func ParseString(values string, separator string) (list List) {
	list.Values = strings.Split(values, separator)
	list.Separator = separator
	list.Index = -1
	return
}

/**
 ** Returns the list as string.
 **/
func (list *List) ToString() (tostr string) {

	for i, l := range list.Values {

		if i > 0 {
			tostr += list.Separator
		}

		tostr += l
	}

	return
}

/**
 ** Returns the list as array.
 **/
func (list *List) ToArray() []string {
	return list.Values
}

/**
 ** Returns the number of elements in the list.
 **/
func (list *List) Len() int {
	return len(list.Values)
}

/**
 ** Returns the element at the first position.
 **/
func (list *List) First() string {
	list.Index = 0
	return list.Values[list.Index]
}

/**
 ** Returns the element at the last position.
 **/
func (list *List) Last() string {
	list.Index = list.Len() - 1
	return list.Values[list.Index]
}

/**
 ** Returns the index of the first occurrence of a value within a list.
 ** Returns -1 if no value is found. The search is case-sensitive.
 **/
func (list *List) Find(value string) int {

	for i, v := range list.Values {
		if v == value {
			return i
		}
	}

	return -1
}

/**
 ** Returns true or false when a value is or not found.
 **/
func (list *List) Exist(value string) bool {
	return list.Find(value) > -1
}

/**
 ** Check if there more elements to move forward
 **/
func (list *List) HasNext() bool {
	i := list.Len() - list.Index
	return i > 1
}

/**
 ** Returns next element.
 **/
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

/**
 ** Move index to given position
 **/
func (list *List) MoveTo(index int) error {

	// check list size
	if list.IsOutOfBound(index) {
		return OutOfBounds
	}

	// move index
	list.Index = index

	return nil
}

/**
 ** Move index to before first
 **/
func (list *List) Rewind() {
	list.Index = -1
}

/**
 ** Returns true if list has no element.
 **/
func (list *List) IsEmpty() bool {
	return list.Len() == 0
}

/**
 ** Returns true if index is out of bound.
 **/
func (list *List) IsOutOfBound(index int) bool {
	return index >= list.Len()
}

/**
 ** Returns the element based on index.
 **/
func (list *List) Get() string {
	return list.Values[list.Index]
}

/**
 ** Returns the element at a given position.
 **/
func (list *List) GetAt(position int) (string, error) {

	// check list size
	if list.IsOutOfBound(position) {
		return "", OutOfBounds
	}

	// return element
	return list.Values[position], nil
}

/**
 ** Delete the element based on index.
 **/
func (list *List) Delete() error {
	return list.DeleteAt(list.Index)
}

/**
 ** Returns list with element deleted at the specified position.
 **/
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

/**
 ** Delete first occurrence of the given element.
 **/
func (list *List) DeleteFirst(element string) {
	if i := list.Find(element); i > -1 {
		list.DeleteAt(i)
	}
}

/**
** Delete last occurrence of the given element.
 **/
func (list *List) DeleteLast(element string) {
	for i := list.Len() - 1; i >= 0; i-- {
		if value, _ := list.GetAt(i); value == element {
			list.DeleteAt(i)
			return
		}
	}
}

/**
 ** Delete elements in the list.
 **/
func (list *List) DeleteAll(elements ...string) {
	for _, e := range elements {
		if i := list.Find(e); i > -1 {
			list.DeleteAt(i)  // delete element
			list.DeleteAll(e) // call it again
		}
	}
}

/**
 ** Filter list.
 **/
func (list *List) Filter(elements ...string) {

	filtered := []string{}

	for _, value := range list.Values {
		for _, e := range elements {
			if value == e {
				filtered = append(filtered, value)
			}
		}
	}

	list.Values = filtered
}

/**
 ** Returns the subset of a list.
 **/
func (list *List) Range(start, end int) List {
	return ParseArray(list.Values[start:end], list.Separator)
}

/**
 ** Returns the subset of a list from an index until the end.
 **/
func (list *List) From(start int) List {
	return ParseArray(list.Values[start:], list.Separator)
}

/**
 ** Returns the subset of a list from its beginning until the given position.
 **/
func (list *List) Until(end int) List {
	return ParseArray(list.Values[:end], list.Separator)
}

/**
 ** Set a new value assigned to its element based on index.
 **/
func (list *List) Set(value string) {
	list.Values[list.Index] = value
}

/**
 ** Set a new value assigned to last element.
 **/
func (list *List) SetLast(value string) {
	list.Values[list.Len()-1] = value
}

/**
 ** Set a new value assigned to its element at specified position.
 **/
func (list *List) SetAt(position int, value string) error {

	// check list size
	if list.IsOutOfBound(position) {
		return OutOfBounds
	}

	// set new value
	list.Values[position] = value

	return nil
}

/**
 ** Append a new value to the list.
 **/
func (list *List) Append(values ...string) {
	for _, value := range values {
		list.Values = append(list.Values, value)
	}
}

/**
 ** Insert a new element to the list at a specified position.
 **/
func (list *List) Insert(position int, value string) {
	list.Values = append(list.Values[:position], append([]string{value}, list.Values[position:]...)...)
}

/**
 ** Swap elements in the list
 **/
func (list *List) Swap(x, y int) {
	list.Values[x], list.Values[y] = list.Values[y], list.Values[x]
}

/**
 ** Shift a element in the list
 **/
func (list *List) Shift(from, to int) {
	list.Values = append(list.Values[:from], append(list.Values[from+1:to], append(list.Values[from:from+1], list.Values[to:]...)...)...)
}

/**
 ** Split a element in 'n' new elements
 **/
func (list *List) Split(position int, interval ...int) {

	var (
		elements = make([]string, len(interval))
		from     = 0
	)

	for i := 0; i < len(interval); i++ {
		elements[i] = list.Values[position][from:interval[i]]
		from = interval[i]
	}

	list.Values = append(list.Values[:position], append(elements, list.Values[position+1:]...)...)
}

/**
 ** Set all elements as uppercase
 **/
func (list *List) Upper() {
	for i, value := range list.Values {
		list.Values[i] = strings.ToUpper(value)
	}
}

/**
 ** Set all elements as lowercase
 **/
func (list *List) Lower() {
	for i, value := range list.Values {
		list.Values[i] = strings.ToLower(value)
	}
}

/**
 ** Sort elements. (A to Z)
 **/
func (list *List) Sort() {
	var sorter sort.StringSlice = list.Values
	sorter.Sort()
	list.Values = sorter[:]
}

/**
 ** Reverse elements. (Z to A)
 **/
func (list *List) Reverse() {
	var sorter sort.StringSlice = list.Values
	sort.Sort(sort.Reverse(sorter))
	list.Values = sorter[:]
}

/**
 ** Counts the number of occurrences
 **/
func (list *List) Count(element string) (count int) {
	for _, value := range list.Values {
		if value == element {
			count++
		}
	}
	return
}

/**
 ** Delete duplicate elements in the list.
 **/
func (list *List) Dedup() {
	for i:=list.Len()-1; i>=0; i-- {
		value, _ := list.GetAt(i)
		if count := list.Count(value); count > 1 {
			list.DeleteLast(value)
		}
	}
}
