package go_routine_variable

import "go-routine-variable/char_variable"

func NewGoRoutineVariable[T any]() GoRoutineVariable[T] {
	return char_variable.NewCharVar[T]()
}

func NewGoRoutineVariableWithLength[T any](length int) GoRoutineVariable[T] {
	return char_variable.NewCharVarWithLength[T](length)
}
