package queue

import (
	"errors"
	"sync"
)

type RingQueue struct {
	lock     sync.Mutex
	capacity int
	values   []interface{}
	front    int
	back     int
}

func NewRingQueue(capacity int) *RingQueue {
	return &RingQueue{
		capacity: capacity,
		values:   make([]interface{}, capacity),
		front:    0,
		back:     0,
	}
}

func (t *RingQueue) IsEmpty() bool {
	return t.front == t.back
}

func (t *RingQueue) IsFull() bool {
	return t.front == ((t.back + 1) % t.capacity)
	//return t.front == (t.back % t.capacity)
}

func (t *RingQueue) GetLength() int {
	return (t.back - t.front + t.capacity) % t.capacity
}

func (t *RingQueue) GetAvailableCapacitySize() int {
	return t.capacity - t.GetLength()
}

func (t *RingQueue) PushValue(value interface{}) error {
	if t.IsFull() {
		return errors.New("the queue capacity is already full")
	}

	t.values[t.back] = value
	t.back = (t.back + 1) % t.capacity
	return nil
}

func (t *RingQueue) PushValues(values ...interface{}) error {
	valuesLen := len(values)
	if valuesLen <= 0 {
		return nil
	}

	if valuesLen > t.GetAvailableCapacitySize() {
		return errors.New("the capacity size of the queue is insufficient")
	}

	for _, value := range values {
		if err := t.PushValue(value); err != nil {
			return err
		}
	}

	return nil
}

func (t *RingQueue) PushValuesWithoutCheck(values ...interface{}) (retPushedCount int) {
	for _, value := range values {
		if err := t.PushValue(value); err != nil {
			return
		}
		retPushedCount += 1
	}

	return
}

func (t *RingQueue) PopValue() (interface{}, bool) {
	if t.IsEmpty() {
		return nil, false
	}

	retValue := t.values[t.front]
	t.values[t.front] = nil
	t.front = (t.front + 1) % t.capacity
	return retValue, true
}

func (t *RingQueue) PopValues(count int) (retValues []interface{}) {
	for i := 0; i < count; i++ {
		value, valid := t.PopValue()
		if !valid {
			return
		}
		retValues = append(retValues, value)
	}

	return
}

func (t *RingQueue) PopValuesToListSpace(ptrListSpace *[]any) (retCount int, retErr error) {
	if ptrListSpace == nil {
		retErr = errors.New("the parameter listSpace is a nil value")
		return
	}

	listSpace := *ptrListSpace
	listSpaceLen := len(listSpace)
	if listSpaceLen > 0 {
		for i := 0; i < listSpaceLen; i++ {
			val, valid := t.PopValue()
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
		val, valid := t.PopValue()
		if !valid {
			return
		}

		*ptrListSpace = append(*ptrListSpace, val)
		retCount += 1
	}

	return
}

func (t *RingQueue) PopValuesWithFilterFunction(f func(value interface{}) bool) (retErr error) {
	if f == nil {
		return errors.New("the parameter f is a nil value")
	}

	for {
		value, valid := t.PopValue()
		if !valid {
			return
		}
		if !f(value) {
			return
		}
	}

	return
}

func (t *RingQueue) ScanElements(f func(value interface{}) bool) error {
	if f == nil {
		return errors.New("the parameter f is a nil value")
	}

	for _, v := range t.values {
		f(v)
	}

	return nil
}
