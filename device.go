package tee

import (
	"encoding/binary"

	"github.com/ethsana/sana-tee/pkg/sev"
)

type Platform uint32

const (
	AMD = iota
	Intel
	Unknown
)

type Device struct {
	Id       string
	Platform Platform
}

func NewDevice(byts []byte) *Device {
	if len(byts) < 4 {
		return nil
	}
	platform := Platform(binary.BigEndian.Uint32(byts[:4]))
	switch platform {
	case AMD:
		return &Device{Id: string(byts[4:132]), Platform: AMD}
	case Intel:
		return &Device{Id: ``, Platform: Intel}
	default:
		return &Device{Id: ``, Platform: Unknown}
	}
}

func (d *Device) Bytes() []byte {
	byts := make([]byte, 4)
	binary.BigEndian.PutUint32(byts, uint32(d.Platform))
	return append(byts, []byte(d.Id)...)
}

func (d *Device) Verify() (bool, error) {
	switch d.Platform {
	case AMD:
		return sev.Verify(d.Id)

	}

	return false, ErrNotSupport
}
