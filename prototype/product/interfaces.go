package product

type INode interface {
	Print(indentation string)
	Clone() INode
}
