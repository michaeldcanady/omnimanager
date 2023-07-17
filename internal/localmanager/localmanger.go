package localmanager

import "github.com/michaeldcanady/omnimanager/internal/policy"

type LocalManager interface {
	CacheDeviceConfiguration(*policy.Configuration) error
}

func NewLocalManager() (LocalManager, error) {
	return newLocalManager()
}
