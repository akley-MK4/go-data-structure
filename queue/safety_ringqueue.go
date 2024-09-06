package queue

import (
	"errors"
	"sync"
)

type IRingQueue interface {
	GetLength() int
	IsEmpty() bool
	IsFull() bool
	GetAvailableCapacitySize() int
	PushValue(value interface{}) error
	PushValues(values ...interface{}) error
	PopValue() (interface{}, bool)
	PopValues(count int) (retValues []interface{})
	PopValuesToListSpace(ptrListSpace *[]any) (retCount int, retErr error)
	PopValuesWithFilterFunction(f func(value interface{}) bool) (retErr error)
}

func NewSafetyRingDeque(newDequeFunc func() IRingQueue) (*SafetyRingQueue, error) {
	inst := newDequeFunc()
	if inst == nil {
		return nil, errors.New("the created ring queue instance is a nil value")
	}

	return &SafetyRingQueue{
		inst: inst,
	}, nil
}

type SafetyRingQueue struct {
	rwMutex sync.RWMutex
	inst    IRingQueue
}

func (t *SafetyRingQueue) GetQueueInstance() IRingQueue {
	return t.inst
}

func (t *SafetyRingQueue) GetLength() int {
	t.rwMutex.RLock()
	defer t.rwMutex.RUnlock()
	return t.inst.GetLength()
}

func (t *SafetyRingQueue) IsEmpty() bool {
	t.rwMutex.RLock()
	defer t.rwMutex.RUnlock()
	return t.inst.IsEmpty()
}

func (t *SafetyRingQueue) IsFull() bool {
	t.rwMutex.RLock()
	defer t.rwMutex.RUnlock()
	return t.inst.IsFull()
}

func (t *SafetyRingQueue) GetAvailableCapacitySize() int {
	t.rwMutex.RLock()
	defer t.rwMutex.RUnlock()
	return t.inst.GetAvailableCapacitySize()
}

func (t *SafetyRingQueue) PushValue(value interface{}) error {
	t.rwMutex.Lock()
	defer t.rwMutex.Unlock()
	return t.inst.PushValue(value)
}

func (t *SafetyRingQueue) PushValueAndRetLength(value interface{}) (retLen int, retErr error) {
	t.rwMutex.Lock()
	defer t.rwMutex.Unlock()

	if retErr = t.inst.PushValue(value); retErr != nil {
		return
	}
	retLen = t.inst.GetLength()
	return
}

func (t *SafetyRingQueue) PushValues(values ...interface{}) error {
	t.rwMutex.Lock()
	defer t.rwMutex.Unlock()
	return t.inst.PushValues(values...)
}

func (t *SafetyRingQueue) PushValuesAndRetLength(values ...interface{}) (retLen int, retErr error) {
	t.rwMutex.Lock()
	defer t.rwMutex.Unlock()

	if retErr = t.inst.PushValues(values...); retErr != nil {
		return
	}
	retLen = t.inst.GetLength()
	return
}

func (t *SafetyRingQueue) PopValue() (interface{}, bool) {
	t.rwMutex.Lock()
	defer t.rwMutex.Unlock()
	return t.inst.PopValue()
}

func (t *SafetyRingQueue) PopValueAndRetLength() (retVal interface{}, retOk bool, retLen int) {
	t.rwMutex.Lock()
	defer t.rwMutex.Unlock()

	retVal, retOk = t.inst.PopValue()
	retLen = t.inst.GetLength()
	return
}

func (t *SafetyRingQueue) PopValues(count int) (retValues []interface{}) {
	t.rwMutex.Lock()
	defer t.rwMutex.Unlock()
	return t.inst.PopValues(count)
}

func (t *SafetyRingQueue) PopValuesAndRetLength(count int) (retValues []interface{}, retLen int) {
	t.rwMutex.Lock()
	defer t.rwMutex.Unlock()

	retValues = t.inst.PopValues(count)
	retLen = t.inst.GetLength()
	return
}

func (t *SafetyRingQueue) PopValuesToListSpace(ptrListSpace *[]any) (retCount int, retErr error) {
	t.rwMutex.Lock()
	defer t.rwMutex.Unlock()
	return t.inst.PopValuesToListSpace(ptrListSpace)
}

func (t *SafetyRingQueue) PopValuesWithFilterFunction(f func(value interface{}) bool) (retErr error) {
	t.rwMutex.Lock()
	defer t.rwMutex.Unlock()
	return t.inst.PopValuesWithFilterFunction(f)
}

func (t *SafetyRingQueue) ExecuteWriteMethod(f func()) {
	t.rwMutex.Lock()
	defer t.rwMutex.Unlock()
	f()
}

func (t *SafetyRingQueue) ExecuteReadMethod(f func()) {
	t.rwMutex.RLock()
	defer t.rwMutex.RUnlock()
	f()
}
