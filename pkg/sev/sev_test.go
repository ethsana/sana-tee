package sev_test

import (
	"fmt"
	"testing"

	"github.com/ethsana/sana-tee/pkg/sev"
)

func TestOutput(t *testing.T) {
	fmt.Println(sev.Output())
	fmt.Println(sev.Ok())
}
