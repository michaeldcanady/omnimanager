package policy

type SettingString struct {
	DisplayName string
	Uri         string
	Value       string
}

func newSettingString(DisplayName string, Uri string, Value string) Setting {
	return SettingString{
		DisplayName: DisplayName,
		Uri:         Uri,
		Value:       Value,
	}
}
