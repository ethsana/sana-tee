package tee

import (
	"bytes"
	"errors"
	"runtime"

	"github.com/ethsana/sana-tee/pkg/sev"
	"github.com/ethsana/sana-tee/pkg/sgx"
)

var ErrNotSupport = errors.New("The current platform is not supported")

func Ok() bool {
	return sev.Ok() || sgx.Ok()
}

func Output() string {
	if runtime.GOOS == `windows` {
		return ErrNotSupport.Error()
	}

	buffer := bytes.NewBufferString("")
	buffer.WriteString("Platform: AMD\n")
	buffer.WriteString(sev.Output())

	buffer.WriteString("Platform: Intel\n")
	buffer.WriteString(sgx.Output())
	return buffer.String()
}

func DeviceID() (*Device, error) {
	if sev.Ok() {
		device, err := sev_device_id()
		if err != nil {
			return nil, err
		}

		return &Device{Id: device, Platform: AMD}, nil
	}
	return nil, ErrNotSupport
}
