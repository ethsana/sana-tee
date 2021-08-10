package sgx_test

import (
	"fmt"
	"testing"

	"github.com/ethsana/sana-tee/pkg/sgx"
)

func TestOutput(t *testing.T) {
	fmt.Println(sgx.Output())
	fmt.Println(sgx.Ok())
}
