package main

import (
	"fmt"

	"github.com/DustiasTheGuy/servman/service"
)

func main() {
	s := service.Service{
		ProccessID: 0,
		Debug:      false,
		Path:       "main.exe",
		WorkingDir: "D:/Development/GO/isak_tech/server",
	}

	// if err := s.StartService(); err != nil {
	// 	log.Fatal(err)
	// }

	fmt.Println(s.IsAlive())
}
