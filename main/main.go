package main

import (
	"errors"
	"log"

	"github.com/DustiasTheGuy/servman/service"
)

// Services contains all running services
var services []*service.Service

func main() {
	s := service.Service{
		Label:      "isak_tech",
		ProccessID: nil,
		Debug:      true,
		Path:       "main.exe",
		WorkingDir: "D:/Development/GO/isak_tech/server",
	}

	for i := 0; i < len(services); i++ {
		if services[i].Label == s.Label {
			log.Fatal(errors.New("Service already running"))
		}
	}

	services = append(services, &s)
	if err := s.StartService(); err != nil {
		// the pid will be updated if proccess
		// launches successfully
		panic(err)
	}
}
