package main

import (
	"fmt"

	"github.com/rs/zerolog/log"
)



func main() {
	fmt.Println("hello world");

	log.Print(fmt.Errorf("errou aqui"));
}


