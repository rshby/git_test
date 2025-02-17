package product

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

type file struct {
	Name string
}

// NewFile creates new instance of file. it implements from interface INode
func NewFile(name string) INode {
	return &file{
		Name: name,
	}
}

func (f *file) Print(indentation string) {
	logrus.Infof("%s %s", indentation, f.Name)
}

func (f *file) Clone() INode {
	// creata new object
	return NewFile(fmt.Sprintf("%s_close", f.Name))
}
