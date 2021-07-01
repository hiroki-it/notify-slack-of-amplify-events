package core

type Entity interface {
	Equals(target Entity) bool
}
