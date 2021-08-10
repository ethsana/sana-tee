// +build windows

package tee

func sev_device_id() (string, error) {
	return "", ErrNotSupport
}
