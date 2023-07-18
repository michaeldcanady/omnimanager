package policy

type Setting struct {
	DisplayName string
	Uri         string
	Value       any
}

func NewSetting(DisplayName string, Uri string, Value any) Setting {
	return Setting{
		DisplayName: DisplayName,
		Uri:         Uri,
		Value:       Value,
	}
}
