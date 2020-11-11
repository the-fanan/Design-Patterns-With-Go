package singleton

type Singleton interface {
	Add(amount int) int
}

type singleton struct {
	count int
}

var instance *singleton

func GetInstance() Singleton {
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}

func (s *singleton) Add(amount int) int {
	s.count += amount
	return s.count 
}

/**
* Note: asides from having a unique instance, this is used for writing the complex code for creation of the object.

* Note: this is not a thread safe implimentation of singleton pattern
*/