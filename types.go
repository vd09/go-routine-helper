package go_routine_variable

import "time"

type GoRoutineVariable[T any] interface {
	WriteValue(value T)
	WriteAllValue(values []T)
	ReadValue() (T, bool)
	ReadAllAvailableValues() ([]T, bool)
	ReadFirstNValues(count int) ([]T, bool)
	ReadAllValues() []T
	ReadAllValuesWithTimeout(timeout time.Duration) ([]T, bool)
	StopWriting()
}
