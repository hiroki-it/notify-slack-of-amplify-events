package core

// ID ゲッターとの名前の重複を防ぐために，名前を大文字にしています．
type ID struct {
	id string
}

// Id idを返却します．
func (i *ID) Id() string {
	return i.id
}
