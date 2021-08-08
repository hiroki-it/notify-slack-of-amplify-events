package domain

type Value interface {

	// Equals 等価性を検証します．
	Equals(target Value) bool
}
