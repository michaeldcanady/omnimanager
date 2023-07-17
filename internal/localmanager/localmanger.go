package localmanager

import "github.com/michaeldcanady/omnimanageragent/internal/policy"

type LocalManager interface {
	CacheDeviceConfiguration(*policy.Configuration) error
}

func NewLocalManager() (LocalManager, error) {
	return newLocalManager()
}
