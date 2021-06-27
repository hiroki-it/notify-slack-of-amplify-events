package detail

type AppId string

// NewAppId コンストラクタ
func NewAppId(value string) AppId {

	return AppId(value)
}

// Value 属性を返却します．
func (ai AppId) Value() string {
	return string(ai)
}

// Equals 等価性を検証します．
func (ai AppId) Equals(target *AppId) bool {
	return ai.Value() == target.Value()
}
