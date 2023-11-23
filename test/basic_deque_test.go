package test

import (
	"github.com/akley-MK4/go-data-structure/queue"
	"testing"
)

const (
	dequeCapacitySize  = 20
	dequeElemValuesLen = 30
)

// back - front  first-in first-out
func TestSafetyLinkListDequePushValues_1(t *testing.T) {
	safetyQueue, newQueueErr := queue.NewSafetyDeque(func() queue.IDeque {
		return queue.NewLinkListDeque(-1)
	})
	if newQueueErr != nil {
		t.Errorf("Failed to create a safety dual end queue, %v", newQueueErr)
		return
	}

	var elemValues []interface{}
	for i := 0; i < dequeElemValuesLen; i++ {
		elemValues = append(elemValues, i)
	}

	if err := safetyQueue.PushValuesToBack(elemValues...); err != nil {
		t.Errorf("Failed to push the value to queue back, %v", err)
		return
	}

	maxTestNum := 3
	for i := 0; i < maxTestNum; i++ {
		oncePopCount := dequeElemValuesLen / maxTestNum
		poppedValues := safetyQueue.PopValuesFromFront(oncePopCount)
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

// back - back last-in first-out
func TestSafetyLinkListDequePushValues_2(t *testing.T) {
	safetyQueue, newQueueErr := queue.NewSafetyDeque(func() queue.IDeque {
		return queue.NewLinkListDeque(-1)
	})
	if newQueueErr != nil {
		t.Errorf("Failed to create a safety dual end queue, %v", newQueueErr)
		return
	}

	var elemValues []interface{}
	for i := 0; i < dequeElemValuesLen; i++ {
		elemValues = append(elemValues, i)
	}

	if err := safetyQueue.PushValuesToBack(elemValues...); err != nil {
		t.Errorf("Failed to push the value to queue back, %v", err)
		return
	}

	maxTestNum := 3
	for i := maxTestNum; i > 0; i-- {
		oncePopCount := dequeElemValuesLen / maxTestNum
		poppedValues := safetyQueue.PopValuesFromBack(oncePopCount)
		if len(poppedValues) != oncePopCount {
			t.Error("The pop-up value does not match the result")
			return
		}

		for idx, v := range poppedValues {
			if v != elemValues[i*oncePopCount-idx-1] {
				t.Error("The pop-up value does not match the result")
				return
			}
		}

		if safetyQueue.GetLength() != (dequeElemValuesLen - oncePopCount*(maxTestNum-i+1)) {
			t.Error("Wrong number of remaining elements")
			return
		}
	}

	if safetyQueue.GetLength() > 0 {
		t.Error("Not all queue elements popped up")
		return
	}
}

// front - back  first-in first-out
func TestSafetyLinkListDequePushValues_3(t *testing.T) {
	safetyQueue, newQueueErr := queue.NewSafetyDeque(func() queue.IDeque {
		return queue.NewLinkListDeque(-1)
	})
	if newQueueErr != nil {
		t.Errorf("Failed to create a safety dual end queue, %v", newQueueErr)
		return
	}

	var elemValues []interface{}
	for i := 0; i < dequeElemValuesLen; i++ {
		elemValues = append(elemValues, i)
	}

	if err := safetyQueue.PushValuesToFront(elemValues...); err != nil {
		t.Errorf("Failed to push the value to queue back, %v", err)
		return
	}

	maxTestNum := 3
	for i := 0; i < maxTestNum; i++ {
		oncePopCount := dequeElemValuesLen / maxTestNum
		poppedValues := safetyQueue.PopValuesFromBack(oncePopCount)
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

// front - front last-in first-out
func TestSafetyLinkListDequePushValues_4(t *testing.T) {
	safetyQueue, newQueueErr := queue.NewSafetyDeque(func() queue.IDeque {
		return queue.NewLinkListDeque(-1)
	})
	if newQueueErr != nil {
		t.Errorf("Failed to create a safety dual end queue, %v", newQueueErr)
		return
	}

	var elemValues []interface{}
	for i := 0; i < dequeElemValuesLen; i++ {
		elemValues = append(elemValues, i)
	}

	if err := safetyQueue.PushValuesToFront(elemValues...); err != nil {
		t.Errorf("Failed to push the value to queue back, %v", err)
		return
	}

	maxTestNum := 3
	for i := maxTestNum; i > 0; i-- {
		oncePopCount := dequeElemValuesLen / maxTestNum
		poppedValues := safetyQueue.PopValuesFromFront(oncePopCount)
		if len(poppedValues) != oncePopCount {
			t.Error("The pop-up value does not match the result")
			return
		}

		for idx, v := range poppedValues {
			if v != elemValues[i*oncePopCount-idx-1] {
				t.Error("The pop-up value does not match the result")
				return
			}
		}

		if safetyQueue.GetLength() != (dequeElemValuesLen - oncePopCount*(maxTestNum-i+1)) {
			t.Error("Wrong number of remaining elements")
			return
		}
	}

	if safetyQueue.GetLength() > 0 {
		t.Error("Not all queue elements popped up")
		return
	}
}

func TestSafetyLinkListDequePushValues_5(t *testing.T) {
	safetyQueue, newQueueErr := queue.NewSafetyDeque(func() queue.IDeque {
		return queue.NewLinkListDeque(-1)
	})
	if newQueueErr != nil {
		t.Errorf("Failed to create a safety dual end queue, %v", newQueueErr)
		return
	}

	var elemValues []interface{}
	for i := 0; i < dequeElemValuesLen; i++ {
		elemValues = append(elemValues, i)
	}

	if err := safetyQueue.PushValuesToBack(elemValues...); err != nil {
		t.Errorf("Failed to push the value to queue back, %v", err)
		return
	}

	queueInst := safetyQueue.GetQueueInstance().(*queue.LinkListDeque)
	popListSpace := make([]interface{}, dequeElemValuesLen)
	safetyQueue.ExecutionInstanceWriteMethod(func() {
		poppedCount, popErr := queueInst.PopValuesFromFrontToListSpace(&popListSpace)
		if popErr != nil {
			t.Errorf("Failed to pop values to list space, %v", popErr)
			return
		}
		popListSpace = popListSpace[:poppedCount]
	})

	for idx, v := range popListSpace {
		if v != elemValues[idx] {
			t.Error("The pop-up value does not match the result")
			return
		}
	}

	if err := safetyQueue.PushValuesToBack(elemValues...); err != nil {
		t.Errorf("Failed to push the value to queue back, %v", err)
		return
	}
	popListSpace = make([]interface{}, 0, dequeElemValuesLen)
	safetyQueue.ExecutionInstanceWriteMethod(func() {
		poppedCount, popErr := queueInst.PopValuesFromFrontToListSpace(&popListSpace)
		if popErr != nil {
			t.Errorf("Failed to pop values to list space, %v", popErr)
			return
		}
		popListSpace = popListSpace[:poppedCount]
	})

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
}
