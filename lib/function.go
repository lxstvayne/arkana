package lib

type Function interface {
	Execute(args []Value) Value
}
