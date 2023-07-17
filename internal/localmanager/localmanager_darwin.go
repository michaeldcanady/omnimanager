//go:build darwin
// +build darwin

package localmanager

type LocalManagerDarwin struct {
}

func newLocalManager() (*LocalManagerDarwin, error) {
	return &LocalManagerDarwin{}, nil
}
