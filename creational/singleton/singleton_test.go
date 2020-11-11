package singleton

import "testing"

func TestGetInstance(t *testing.T){
	counter1 := GetInstance()

	if counter1 == nil {
		t.Error("Expected pointer to singleton after calling GetInstance(), nil returned")
	}
	expectedCounter := counter1
	currentCount := counter1.Add(1)
	if currentCount != 1 {
		t.Errorf("After calling for the first time to count, the count must be 1 but %d was gotten\n", currentCount)
	}

	counter2 := GetInstance()
	if counter2 != expectedCounter {
		t.Error("Expected same instance in second counter but a different one was returned")
	}

	currentCount = counter2.Add(1)
	if currentCount != 2 {
		t.Errorf("After calling for the first time to count, the count must be 2 but %d was gotten\n", currentCount)
	}
}

