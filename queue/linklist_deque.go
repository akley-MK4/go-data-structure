package queue

import (
	"container/list"
	"errors"
)

func NewLinkListDeque(capacity int) *LinkListDeque {
	return &LinkListDeque{
		list:     list.New(),
		capacity: capacity,
	}
}

type LinkListDeque struct {
	list     *list.List
	capacity int
}

func (t *LinkListDeque) GetLength() int {
	return t.list.Len()
}

func (t *LinkListDeque) IsEmpty() bool {
	return t.list.Len() == 0
}

func (t *LinkListDeque) IsFull() bool {
	if t.capacity < 0 {
		return false
	}

	return t.list.Len() >= t.capacity
}

func (t *LinkListDeque) GetAvailableCapacitySize() int {
	if t.capacity < 0 {
		return -1
	}

	return t.capacity - t.list.Len()
}

func (t *LinkListDeque) CheckAvailableCapacity(pushValueLen int) bool {
	availableCapSize := t.GetAvailableCapacitySize()
	if availableCapSize < 0 {
		return true
	}

	return availableCapSize >= pushValueLen
}

func (t *LinkListDeque) PushValueToBack(value any) error {
	if t.IsFull() {
		return errors.New("the queue capacity is already full")
	}

	t.list.PushBack(value)
	return nil
}

func (t *LinkListDeque) PushValuesToBack(values ...any) error {
	if !t.CheckAvailableCapacity(len(values)) {
		return errors.New("the capacity size of the queue is insufficient")
	}

	for _, value := range values {
		if err := t.PushValueToBack(value); err != nil {
			return err
		}
	}

	return nil
}

func (t *LinkListDeque) PushValuesToBackWithoutCheck(values ...any) (retPushedCount int) {
	for _, value := range values {
		if err := t.PushValueToBack(value); err != nil {
			return
		}
		retPushedCount += 1
	}

	return
}

func (t *LinkListDeque) PushValueToFront(value any) error {
	if t.IsFull() {
		return errors.New("the queue capacity is already full")
	}

	t.list.PushFront(value)
	return nil
}

func (t *LinkListDeque) PushValuesToFront(values ...any) error {
	if !t.CheckAvailableCapacity(len(values)) {
		return errors.New("the capacity size of the queue is insufficient")
	}

	for _, value := range values {
		if err := t.PushValueToFront(value); err != nil {
			return err
		}
	}

	return nil
}

func (t *LinkListDeque) PushValuesToFrontWithoutCheck(values ...any) (retPushedCount int) {
	for _, value := range values {
		if err := t.PushValueToFront(value); err != nil {
			return
		}
		retPushedCount += 1
	}

	return
}

func (t *LinkListDeque) PopValueFromFront() (any, bool) {
	valuesLen := t.list.Len()
	if valuesLen <= 0 {
		return nil, false
	}

	elem := t.list.Front()
	t.list.Remove(elem)
	return elem.Value, true
}

func (t *LinkListDeque) PopValuesFromFront(count int) (retValues []any) {
	valuesLen := t.list.Len()
	if valuesLen <= 0 {
		return
	}

	retValuesCap := count
	if valuesLen < count {
		retValuesCap = valuesLen
	}
	retValues = make([]any, 0, retValuesCap)

	for i := 0; i < count; i++ {
		val, valid := t.PopValueFromFront()
		if !valid {
			continue
		}
		retValues = append(retValues, val)
	}

	return
}

func (t *LinkListDeque) PopValuesFromFrontToListSpace(ptrListSpace *[]any) (retCount int, retErr error) {
	if ptrListSpace == nil {
		retErr = errors.New("the parameter listSpace is a nil value")
		return
	}

	listSpace := *ptrListSpace
	listSpaceLen := len(listSpace)
	if listSpaceLen > 0 {
		for i := 0; i < listSpaceLen; i++ {
			val, valid := t.PopValueFromFront()
			if !valid {
				return
			}

			listSpace[i] = val
			retCount += 1
		}
		return
	}

	listSpaceCap := cap(listSpace)
	if listSpaceCap <= 0 {
		retErr = errors.New("the capacity of the parameter listSpace is 0")
		return
	}

	for i := 0; i < listSpaceCap; i++ {
		val, valid := t.PopValueFromFront()
		if !valid {
			return
		}

		*ptrListSpace = append(*ptrListSpace, val)
		retCount += 1
	}

	return
}

func (t *LinkListDeque) PopValuesFromFrontWithFilterFunction(f func(value interface{}) bool) (retErr error) {
	if f == nil {
		return errors.New("the parameter f is a nil value")
	}

	for {
		value, valid := t.PopValueFromFront()
		if !valid {
			return
		}
		if !f(value) {
			return
		}
	}
	return
}

func (t *LinkListDeque) PopValueFromBack() (any, bool) {
	valuesLen := t.list.Len()
	if valuesLen <= 0 {
		return nil, false
	}

	elem := t.list.Back()
	t.list.Remove(elem)
	return elem.Value, true
}

func (t *LinkListDeque) PopValuesFromBack(count int) (retValues []any) {
	valuesLen := t.list.Len()
	if valuesLen <= 0 {
		return
	}

	retValuesCap := count
	if valuesLen < count {
		retValuesCap = valuesLen
	}
	retValues = make([]any, 0, retValuesCap)

	for i := 0; i < count; i++ {
		val, valid := t.PopValueFromBack()
		if !valid {
			continue
		}
		retValues = append(retValues, val)
	}

	return
}

func (t *LinkListDeque) PopValuesFromBackToListSpace(ptrListSpace *[]any) (retCount int, retErr error) {
	if ptrListSpace == nil {
		retErr = errors.New("the parameter listSpace is a nil value")
		return
	}

	listSpace := *ptrListSpace
	listSpaceLen := len(listSpace)
	if listSpaceLen > 0 {
		for i := 0; i < listSpaceLen; i++ {
			val, valid := t.PopValueFromBack()
			if !valid {
				return
			}

			listSpace[i] = val
			retCount += 1
		}
		return
	}

	listSpaceCap := cap(listSpace)
	if listSpaceCap <= 0 {
		retErr = errors.New("the capacity of the parameter listSpace is 0")
		return
	}

	for i := 0; i < listSpaceCap; i++ {
		val, valid := t.PopValueFromFront()
		if !valid {
			return
		}

		*ptrListSpace = append(*ptrListSpace, val)
		retCount += 1
	}

	return
}

func (t *LinkListDeque) PopValuesFromBackWithFilterFunction(f func(value interface{}) bool) (retErr error) {
	if f == nil {
		return errors.New("the parameter f is a nil value")
	}

	for {
		value, valid := t.PopValueFromBack()
		if !valid {
			return
		}
		if !f(value) {
			return
		}
	}
	return
}

func (t *LinkListDeque) ScanElementsFromFront(f func(elem *list.Element) bool) {
	for elem := t.list.Front(); elem != nil; elem = elem.Next() {
		if !f(elem) {
			return
		}
	}
}

func (t *LinkListDeque) ScanElementsFromBack(f func(elem *list.Element) bool) {
	for elem := t.list.Back(); elem != nil; elem = elem.Prev() {
		if !f(elem) {
			return
		}
	}
}

func (t *LinkListDeque) RemoveElements(elems []*list.Element) {
	for _, elem := range elems {
		t.list.Remove(elem)
	}
}
