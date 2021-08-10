// +build !windows

package tee

//#cgo LDFLAGS: -L${SRCDIR}/target/debug -lsana_tee -ldl -lssl -lcrypto
//#include <stdint.h>
//extern void sev_device_id(char * result);
import "C"

func sev_device_id() (string, error) {
	id := (*C.char)(C.malloc(1024))
	C.sev_device_id(id)
	return C.GoString(id), nil
}
