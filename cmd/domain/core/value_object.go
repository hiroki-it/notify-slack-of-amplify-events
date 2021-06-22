package core

type ValueObject interface {
	Equals(value ValueObject) bool
}
