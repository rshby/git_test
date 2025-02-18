package singleton

import (
	"github.com/sirupsen/logrus"
	"sync"
)

type Single struct {
}

// global variable (singleton)
var SingleObject *Single

var mu = &sync.Mutex{}

// GetSingleInstance retrieves single object, if not exists then create new
func GetSingleInstance() *Single {
	// check if object has been generated
	if SingleObject != nil {
		logrus.Info("Single instance already created")
		return SingleObject
	}

	// if object still nil, lets create the new one
	mu.Lock()
	defer mu.Unlock()

	if SingleObject != nil {
		logrus.Info("Single instance already created")
		return SingleObject
	}

	SingleObject = &Single{}
	logrus.Info("Creating single instance")
	return SingleObject
}
