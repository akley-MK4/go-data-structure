package queue

import (
	"errors"
	"sync"
)

type IDeque interface {
	GetLength() int
	IsEmpty() bool
	IsFull() bool
	GetAvailableCapacitySize() int
	CheckAvailableCapacity(pushValueLen int) bool

	PushValueToBack(value any) error
	PushValuesToBack(values ...any) error
	PushValuesToBackWithoutCheck(values ...any) (retPushedCount int)
	PushValueToFront(value any) error
	PushValuesToFront(values ...any) error
	PushValuesToFrontWithoutCheck(values ...any) (retPushedCount int)

	PopValueFromFront() (any, bool)
	PopValuesFromFront(count int) (retValues []any)
	PopValuesFromFrontToListSpace(ptrListSpace *[]any) (retCount int, retErr error)
	PopValueFromBack() (any, bool)
	PopValuesFromBack(count int) (retValues []any)
	PopValuesFromBackToListSpace(ptrListSpace *[]any) (retCount int, retErr error)
}

func NewSafetyDeque(newDequeFunc func() IDeque) (IDeque, error) {
	inst := newDequeFunc()
	if inst == nil {
		return nil, errors.New("the created deque instance is a nil value")
	}

	return &SafetyDeque{
		inst: inst,
	}, nil
}

type SafetyDeque struct {
	rwMutex sync.RWMutex
	inst    IDeque
}

func (t *SafetyDeque) GetQueueInstance() IDeque {
	return t.inst
}

func (t *SafetyDeque) GetLength() int {
	t.rwMutex.RLock()
	defer t.rwMutex.RUnlock()

	return t.inst.GetLength()
}

func (t *SafetyDeque) IsEmpty() bool {
	t.rwMutex.RLock()
	defer t.rwMutex.RUnlock()

	return t.inst.IsEmpty()
}

func (t *SafetyDeque) IsFull() bool {
	t.rwMutex.RLock()
	defer t.rwMutex.RUnlock()

	return t.inst.IsFull()
}

func (t *SafetyDeque) GetAvailableCapacitySize() int {
	t.rwMutex.RLock()
	defer t.rwMutex.RUnlock()

	return t.inst.GetAvailableCapacitySize()
}

func (t *SafetyDeque) CheckAvailableCapacity(pushValueLen int) bool {
	t.rwMutex.RLock()
	defer t.rwMutex.RUnlock()

	return t.inst.CheckAvailableCapacity(pushValueLen)
}

func (t *SafetyDeque) PushValueToBack(value any) error {
	t.rwMutex.Lock()
	defer t.rwMutex.Unlock()

	return t.inst.PushValueToBack(value)
}

func (t *SafetyDeque) PushValuesToBack(values ...any) error {
	t.rwMutex.Lock()
	defer t.rwMutex.Unlock()

	return t.inst.PushValuesToBack(values...)
}

func (t *SafetyDeque) PushValuesToBackWithoutCheck(values ...any) (retPushedCount int) {
	t.rwMutex.Lock()
	defer t.rwMutex.Unlock()

	retPushedCount = t.inst.PushValuesToBackWithoutCheck(values...)
	return
}

func (t *SafetyDeque) PushValueToFront(value any) error {
	t.rwMutex.Lock()
	defer t.rwMutex.Unlock()

	return t.inst.PushValueToFront(value)
}

func (t *SafetyDeque) PushValuesToFront(values ...any) error {
	t.rwMutex.Lock()
	defer t.rwMutex.Unlock()

	return t.inst.PushValuesToFront(values...)
}

func (t *SafetyDeque) PushValuesToFrontWithoutCheck(values ...any) (retPushedCount int) {
	t.rwMutex.Lock()
	defer t.rwMutex.Unlock()

	return t.inst.PushValuesToFrontWithoutCheck(values...)
}

func (t *SafetyDeque) PopValueFromFront() (any, bool) {
	t.rwMutex.Lock()
	defer t.rwMutex.Unlock()

	return t.inst.PopValueFromFront()
}

func (t *SafetyDeque) PopValuesFromFront(count int) (retValues []any) {
	t.rwMutex.Lock()
	defer t.rwMutex.Unlock()

	return t.inst.PopValuesFromFront(count)
}

func (t *SafetyDeque) PopValuesFromFrontToListSpace(ptrListSpace *[]any) (retCount int, retErr error) {
	t.rwMutex.Lock()
	defer t.rwMutex.Unlock()

	return t.inst.PopValuesFromFrontToListSpace(ptrListSpace)
}

func (t *SafetyDeque) PopValueFromBack() (any, bool) {
	t.rwMutex.Lock()
	defer t.rwMutex.Unlock()

	return t.inst.PopValueFromBack()
}

func (t *SafetyDeque) PopValuesFromBack(count int) (retValues []any) {
	t.rwMutex.Lock()
	defer t.rwMutex.Unlock()

	return t.inst.PopValuesFromBack(count)
}

func (t *SafetyDeque) PopValuesFromBackToListSpace(ptrListSpace *[]any) (retCount int, retErr error) {
	t.rwMutex.Lock()
	defer t.rwMutex.Unlock()

	return t.inst.PopValuesFromBackToListSpace(ptrListSpace)
}

func (t *SafetyDeque) SafeExecutionWriteFunction(f func()) {
	t.rwMutex.Lock()
	defer t.rwMutex.Unlock()

	f()
}

func (t *SafetyDeque) SafeExecutionReadFunction(f func()) {
	t.rwMutex.RLock()
	defer t.rwMutex.RUnlock()
	f()
}
