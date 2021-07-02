package core

// ID ゲッターとの名前の重複を防ぐために，名前を大文字にしています．
type ID interface {
	Id() string
	Equals(target ID) bool
}
