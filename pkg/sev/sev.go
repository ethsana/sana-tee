package sev

import (
	"fmt"
	"net/http"
)

const defaultVerifyUrl = `https://kdsintf.amd.com/cek/id/`

func Verify(device string) (bool, error) {
	uri := fmt.Sprint(defaultVerifyUrl, device)
	resp, err := http.Get(uri)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	return resp.StatusCode == http.StatusOK, nil
}
