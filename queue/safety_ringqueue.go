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
	PushValuesWithCheck(values ...interface{}) (retPushedCount int)
	PopValue() (interface{}, bool)
	PopValues(count int) (retValues []interface{})
}

func NewSafetyRingDeque(newDequeFunc func() IRingQueue) (IRingQueue, error) {
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

func (t *SafetyRingQueue) PushValues(values ...interface{}) error {
	t.rwMutex.Lock()
	defer t.rwMutex.Unlock()
	return t.inst.PushValues(values...)
}

func (t *SafetyRingQueue) PushValuesWithCheck(values ...interface{}) (retPushedCount int) {
	t.rwMutex.Lock()
	defer t.rwMutex.Unlock()
	return t.inst.PushValuesWithCheck(values...)
}

func (t *SafetyRingQueue) PopValue() (interface{}, bool) {
	t.rwMutex.Lock()
	defer t.rwMutex.Unlock()
	return t.inst.PopValue()
}

func (t *SafetyRingQueue) PopValues(count int) (retValues []interface{}) {
	t.rwMutex.Lock()
	defer t.rwMutex.Unlock()
	return t.inst.PopValues(count)
}

func (t *SafetyRingQueue) SafeExecutionWriteFunction(f func()) {
	t.rwMutex.Lock()
	defer t.rwMutex.Unlock()
	f()
}

func (t *SafetyRingQueue) SafeExecutionReadFunction(f func()) {
	t.rwMutex.RLock()
	defer t.rwMutex.RUnlock()
	f()
}
