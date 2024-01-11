package test

import (
	"github.com/akley-MK4/go-data-structure/queue"
	"testing"
)

//const (
//	ringQueueCapacitySize  = 50
//	ringQueueElemValuesLen = 30
//)

func TestSafetyRingQueuePushValues_1(t *testing.T) {
	ringQueueCapacitySize := 50
	ringQueueElemValuesLen := 30

	safetyQueue, newQueueErr := queue.NewSafetyRingDeque(func() queue.IRingQueue {
		return queue.NewRingQueue(ringQueueCapacitySize)
	})
	if newQueueErr != nil {
		t.Errorf("Failed to create a safety ring queue, %v", newQueueErr)
		return
	}

	var elemValues []interface{}
	for i := 0; i < ringQueueElemValuesLen; i++ {
		elemValues = append(elemValues, i)
	}

	if err := safetyQueue.PushValues(elemValues...); err != nil {
		t.Errorf("Failed to push the value to ring queue, %v", err)
		return
	}

	maxTestNum := 3
	for i := 0; i < maxTestNum; i++ {
		oncePopCount := dequeElemValuesLen / maxTestNum
		poppedValues := safetyQueue.PopValues(oncePopCount)
		if len(poppedValues) != oncePopCount {
			t.Error("The pop-up value does not match the result")
			return
		}

		for idx, v := range poppedValues {
			if v != elemValues[i*oncePopCount+idx] {
				t.Error("The pop-up value does not match the result")
				return
			}
		}

		if safetyQueue.GetLength() != (dequeElemValuesLen - oncePopCount*(i+1)) {
			t.Error("Wrong number of remaining elements")
			return
		}
	}

	if safetyQueue.GetLength() > 0 {
		t.Error("Not all queue elements popped up")
		return
	}

	queueInst := safetyQueue.GetQueueInstance().(*queue.RingQueue)
	safetyQueue.ExecutionInstanceReadMethod(func() {
		queueInst.ScanElements(
			func(value interface{}) bool {
				if value != nil {
					t.Error("There are still non empty elements in the queue")
					return false
				}
				return true
			})
	})
}

func TestSafetyRingQueuePushValuesWithException_2(t *testing.T) {
	ringQueueCapacitySize := 50

	safetyQueue, newQueueErr := queue.NewSafetyRingDeque(func() queue.IRingQueue {
		return queue.NewRingQueue(ringQueueCapacitySize)
	})
	if newQueueErr != nil {
		t.Errorf("Failed to create a safety ring queue, %v", newQueueErr)
		return
	}

	var elemValues []interface{}
	for i := 0; i < ringQueueCapacitySize+1; i++ {
		elemValues = append(elemValues, i)
	}

	if err := safetyQueue.PushValues(elemValues...); err == nil {
		t.Errorf("Exceeding capacity size without throwing an error, %v", err)
		return
	}

	// The last element of the circular queue cannot be used
	for i := 0; i < ringQueueCapacitySize-1; i++ {
		if err := safetyQueue.PushValue(i); err != nil {
			t.Errorf("Failed to push the value to ring queue, %v", err)
			return
		}
	}
	if err := safetyQueue.PushValue(0); err == nil {
		t.Errorf("Exceeding capacity size without throwing an error, %v", err)
		return
	}

}

func TestSafetyRingQueuePushValues_3(t *testing.T) {
	ringQueueCapacitySize := 50
	ringQueueElemValuesLen := 30

	safetyQueue, newQueueErr := queue.NewSafetyRingDeque(func() queue.IRingQueue {
		return queue.NewRingQueue(ringQueueCapacitySize)
	})
	if newQueueErr != nil {
		t.Errorf("Failed to create a safety ring queue, %v", newQueueErr)
		return
	}

	var elemValues []interface{}
	for i := 0; i < (ringQueueElemValuesLen - 1); i++ {
		elemValues = append(elemValues, i)
	}

	for _, v := range elemValues {
		if err := safetyQueue.PushValue(v); err != nil {
			t.Errorf("Failed to push the value to ring queue, %v", err)
			return
		}
	}

	popListSpace := make([]interface{}, 0, ringQueueElemValuesLen-1)
	poppedCount, popErr := safetyQueue.PopValuesToListSpace(&popListSpace)
	if popErr != nil {
		t.Errorf("Failed to pop values to list space, %v", popErr)
		return
	}
	popListSpace = popListSpace[:poppedCount]

	for idx, v := range popListSpace {
		if v != elemValues[idx] {
			t.Error("The pop-up value does not match the result")
			return
		}
	}

	if err := safetyQueue.PushValues(elemValues...); err != nil {
		t.Errorf("Failed to push the value to ring queue, %v", err)
		return
	}

	popListSpace = make([]interface{}, ringQueueElemValuesLen-1)
	poppedCount, popErr = safetyQueue.PopValuesToListSpace(&popListSpace)
	if popErr != nil {
		t.Errorf("Failed to pop values to list space, %v", popErr)
		return
	}
	popListSpace = popListSpace[:poppedCount]

	for idx, v := range popListSpace {
		if v != elemValues[idx] {
			t.Error("The pop-up value does not match the result")
			return
		}
	}

	if safetyQueue.GetLength() > 0 {
		t.Error("Not all queue elements popped up")
		return
	}

	queueInst := safetyQueue.GetQueueInstance().(*queue.RingQueue)
	safetyQueue.ExecutionInstanceReadMethod(func() {
		queueInst.ScanElements(
			func(value interface{}) bool {
				if value != nil {
					t.Error("There are still non empty elements in the queue")
					return false
				}
				return true
			})
	})
}

func TestSafetyRingQueuePushValues_4(t *testing.T) {
	ringQueueCapacitySize := 30
	ringQueueElemValuesLen := ringQueueCapacitySize + 5

	safetyQueue, newQueueErr := queue.NewSafetyRingDeque(func() queue.IRingQueue {
		return queue.NewRingQueue(ringQueueCapacitySize)
	})
	if newQueueErr != nil {
		t.Errorf("Failed to create a safety ring queue, %v", newQueueErr)
		return
	}

	var elemValues []interface{}
	for i := 0; i < (ringQueueElemValuesLen - 1); i++ {
		elemValues = append(elemValues, i)
	}

	queueInst := safetyQueue.GetQueueInstance().(*queue.RingQueue)

	pushedCount := queueInst.PushValuesWithoutCheck(elemValues...)
	if pushedCount != (ringQueueCapacitySize - 1) {
		t.Error("The amount of data pushed to the ring queue is incorrect")
		return
	}

	popListSpace := make([]interface{}, ringQueueElemValuesLen)
	poppedCount, popErr := safetyQueue.PopValuesToListSpace(&popListSpace)
	if popErr != nil {
		t.Errorf("Failed to pop values to list space, %v", popErr)
		return
	}
	popListSpace = popListSpace[:poppedCount]

	for idx, v := range popListSpace {
		if v != elemValues[idx] {
			t.Error("The pop-up value does not match the result")
			return
		}
	}

	if safetyQueue.GetLength() > 0 {
		t.Error("Not all queue elements popped up")
		return
	}

	safetyQueue.ExecutionInstanceReadMethod(func() {
		queueInst.ScanElements(
			func(value interface{}) bool {
				if value != nil {
					t.Error("There are still non empty elements in the queue")
					return false
				}
				return true
			})
	})

}

func TestSafetyRingQueuePushValues_5(t *testing.T) {
	ringQueueCapacitySize := 30
	ringQueueElemValuesLen := ringQueueCapacitySize - 1

	safetyQueue, newQueueErr := queue.NewSafetyRingDeque(func() queue.IRingQueue {
		return queue.NewRingQueue(ringQueueCapacitySize)
	})
	if newQueueErr != nil {
		t.Errorf("Failed to create a safety ring queue, %v", newQueueErr)
		return
	}

	var elemValues []interface{}
	for i := 0; i < ringQueueElemValuesLen; i++ {
		elemValues = append(elemValues, i)
	}

	if err := safetyQueue.PushValues(elemValues...); err != nil {
		t.Errorf("Failed to push values to the ring queue, %v", err)
		return
	}

	var poppedCount int
	poppedListSpace := make([]int, len(elemValues))
	if err := safetyQueue.PopValuesWithFilterFunction(func(value interface{}) bool {
		poppedListSpace[poppedCount] = value.(int)
		poppedCount += 1
		return poppedCount < len(poppedListSpace)
	}); err != nil {
		t.Errorf("Failed to pop values to function, %v", err)
		return
	}

	if poppedCount != len(poppedListSpace) {
		t.Error("The pop-up value does not match the result")
		return
	}

	queueInst := safetyQueue.GetQueueInstance().(*queue.RingQueue)
	queueInst.ScanElements(
		func(value interface{}) bool {
			if value != nil {
				t.Error("There are still non empty elements in the queue")
				return false
			}
			return true
		})
}
