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
