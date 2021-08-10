package main

import (
	"fmt"

	tee "github.com/ethsana/sana-tee"
)

func main() {
	id, err := tee.DeviceID()
	if err != nil {
		panic(err)
	}

	fmt.Println(id.Platform, id.Id)
	fmt.Print(id.Verify())

	byts := id.Bytes()

	id2 := tee.NewDevice(byts)
	fmt.Println(id2)
}
