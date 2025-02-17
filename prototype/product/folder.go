package product

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

type Folder struct {
	Children []INode
	Name     string
}

// NewFolder creates new instance Folder. it implements from interface INode
func NewFolder(name string) *Folder {
	return &Folder{
		Name: name,
	}
}

func (f *Folder) Print(indentation string) {
	logrus.Infof("%s %s", indentation, f.Name)
	if f.Children != nil {
		for _, c := range f.Children {
			c.Print(indentation + indentation)
		}
	}
}

func (f *Folder) Clone() INode {
	newFolder := NewFolder(fmt.Sprintf("%s_clone", f.Name))

	var tempChild []INode
	for _, c := range f.Children {
		cp := c.Clone()
		tempChild = append(tempChild, cp)
	}

	newFolder.Children = tempChild
	return newFolder
}
