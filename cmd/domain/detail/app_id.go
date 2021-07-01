package detail

type AppId struct {
	value string
}

// NewAppId コンストラクタ
func NewAppId(value string) *AppId {

	return &AppId{
		value: value,
	}
}

// Value 属性を返却します．
func (ai AppId) Value() string {
	return ai.value
}

// Equals 等価性を検証します．
func (ai AppId) Equals(target *AppId) bool {
	return ai.value == target.Value()
}
