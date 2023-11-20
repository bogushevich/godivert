package main

import (
	"fmt"

	"github.com/bogushevich/godivert"
)

func main() {
	winDivert, err := godivert.OpenHandleWithFilter("true")
	if err != nil {
		panic(err)
	}
	defer winDivert.Close()

	packet, err := winDivert.Recv()
	if err != nil {
		panic(err)
	}

	fmt.Println(packet)

	packet.Send(winDivert)
}
