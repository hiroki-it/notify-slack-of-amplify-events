package core

type ValueObject interface {
	Equals(target ValueObject) bool
}
