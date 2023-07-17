//go:build windows
// +build windows

package localmanager

import (
	"fmt"
	"reflect"
	"strconv"

	"golang.org/x/sys/windows/registry"

	"github.com/michaeldcanady/omnimanager/internal/policy"
)

const basePath = "SOFTWARE\\{CompanyName}\\PolicyManager\\Provider"

type LocalManagerWin struct {
}

func newLocalManager() (*LocalManagerWin, error) {
	return &LocalManagerWin{}, nil
}

func (m *LocalManagerWin) CacheDeviceConfiguration(policy *policy.Configuration) error {

	// Access base key
	key, err := registry.OpenKey(registry.LOCAL_MACHINE, basePath, registry.ALL_ACCESS)
	if err != nil {
		return fmt.Errorf("failed to open registry key: %w", err)
	}
	defer key.Close()

	policyKey, _, err := registry.CreateKey(key, policy.Id, registry.ALL_ACCESS)
	if err != nil {
		return fmt.Errorf("failed to create registry key %s: %w", policy.Id, err)
	}

	return m.UpdateCachedDeviceConfiguration(policy, policyKey)
}

func (m *LocalManagerWin) UpdateCachedDeviceConfiguration(policy *policy.Configuration, registryKey registry.Key) error {

	err := m.SetValue(registryKey, "DisplayName", policy.DisplayName)
	if err != nil {
		return fmt.Errorf("unable to set 'DisplayName' to %w: %w", policy.DisplayName, err)
	}

	for _, setting := range policy.Settings {
		err = m.SetValue(registryKey, setting.Uri, setting.Value)
		if err != nil {
			return fmt.Errorf("unable to set %s to %s: %w", setting.Uri, setting.Value, err)
		}
	}

	return nil
}

// RetrieveDeviceConfiguration retrieves cached device configuration
func (m *LocalManagerWin) RetrieveDeviceConfiguration(id string) (*policy.Configuration, error) {
	// Access base key
	key, err := registry.OpenKey(registry.LOCAL_MACHINE, basePath+id, registry.ALL_ACCESS)
	if err != nil {
		return nil, fmt.Errorf("failed to open registry key: %w", err)
	}
	defer key.Close()
	uris, err := key.ReadValueNames(-1)
	if err != nil {
		return nil, fmt.Errorf("failed to read value names: %w", err)
	}

	displayName, err := m.GetValue(key, "DisplayName")
	if err != nil {
		return nil, fmt.Errorf("error getting DisplayName: %w", err)
	}

	configuration := policy.Configuration{
		Id:          id,
		DisplayName: displayName.(string),
	}

	for _, uri := range uris {
		value, err := m.GetValue(key, uri)
		if err != nil {
			return nil, fmt.Errorf("error getting value: %w", err)
		}
		setting, err := policy.NewSetting("", uri, value)
		if err != nil {
			return nil, fmt.Errorf("error marshalling setting: %w", err)
		}
		configuration.Settings = append(configuration.Settings, setting)
	}
	return &configuration, nil
}

// SetValue sets a value in the Windows registry based on the value's type.
func (m *LocalManagerWin) SetValue(key registry.Key, name string, value any) error {

	var err error

	switch v := value.(type) {
	case string:
		err = key.SetStringValue(name, v)
	case uint32:
		err = key.SetDWordValue(name, v)
	case int:
		err = key.SetDWordValue(name, uint32(v))
	case uint64:
		err = key.SetQWordValue(name, v)
	case int64:
		err = key.SetQWordValue(name, uint64(v))
	case bool:
		var intValue uint32
		if v {
			intValue = 1
		}
		err = key.SetDWordValue(name, intValue)
	case []byte:
		err = key.SetBinaryValue(name, v)
	case float32:
		strValue := strconv.FormatFloat(float64(v), 'f', -1, 32)
		err = key.SetStringValue(name, strValue)
	case float64:
		strValue := strconv.FormatFloat(v, 'f', -1, 64)
		err = key.SetStringValue(name, strValue)
	default:
		return fmt.Errorf("unsupported value type: %T", value)
	}

	if err != nil {
		return fmt.Errorf("failed to set registry value: %w", err)
	}

	return nil
}

// GetValue
func (m *LocalManagerWin) GetValue(key registry.Key, name string) (any, error) {
	_, valueType, err := key.GetValue(name, nil)
	if err != nil {
		return nil, fmt.Errorf("error retriving value: %w", err)
	}

	switch valueType {
	case registry.BINARY:
		value, _, err := key.GetBinaryValue(name)
		if err != nil {
			return nil, fmt.Errorf("unable to get value: %w", err)
		}
		return value, nil
	case registry.DWORD:
		value, _, err := key.GetIntegerValue(name)
		if err != nil {
			return nil, fmt.Errorf("unable to get value: %w", err)
		}
		return value, nil
	case registry.EXPAND_SZ:
	case registry.MULTI_SZ:
	case registry.SZ:
	case registry.QWORD:
	}
	return nil, nil
}

// Contains checks if a slice contains a specific value.
func Contains(slice interface{}, value interface{}) bool {
	sliceValue := reflect.ValueOf(slice)

	if sliceValue.Kind() != reflect.Slice {
		panic("Contains function expects a slice as the first argument")
	}

	for i := 0; i < sliceValue.Len(); i++ {
		item := sliceValue.Index(i).Interface()
		if reflect.DeepEqual(item, value) {
			return true
		}
	}

	return false
}
