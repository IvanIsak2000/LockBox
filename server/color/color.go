package color

import (
	"fmt"

	"github.com/helioloureiro/golorama"
)

func Print(text string, color golorama.Color) {
	fmt.Println(golorama.GetCSI(color) + text + golorama.Reset())
}