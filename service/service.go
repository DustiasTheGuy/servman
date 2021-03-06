package service

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"time"
)

// Service contains nessecary information about a proccess
type Service struct {
	Label      string
	ProcessID  *int      // A reference to the proccess that can be used to terminate or modify the proccess
	Path       string    // where is the binary located?
	WorkingDir string    // what directory should the program launch from? very useful if the program is serving static files from a relative path
	Debug      bool      // output useful messages
	Cmd        *exec.Cmd // the parent
	Started    time.Time // when did the process start?
}

// ConstructCmd creates a pointer to a cmd which can be used to execute commands
func (s *Service) ConstructCmd() *exec.Cmd {
	return &exec.Cmd{
		Path: s.Path,
		Args: []string{"START"},
		Dir:  s.WorkingDir,
	}
}

// StartService executes a start commandand prints the pid to stdout
func (s *Service) StartService() error {
	s.Cmd = s.ConstructCmd()
	err := s.Cmd.Start()

	if err != nil {
		return err
	}

	s.ProcessID = &s.Cmd.Process.Pid
	s.Started = time.Now()
	if s.Debug {
		fmt.Printf("Started Proccess: %d\n", s.Cmd.Process.Pid)
	}

	return nil
}

// KillService kills the proccess and requires proccess id to be populated
func (s *Service) KillService() error {
	if s.IsAlive() {
		p, err := os.FindProcess(*s.ProcessID)

		if err != nil {
			return err
		}

		if err := p.Kill(); err != nil {
			return err
		}
	}

	return errors.New("Proccess May Have Already Been Closed")
}

// IsAlive checkes if a proccess is currently alive
func (s *Service) IsAlive() bool {
	_, err := os.FindProcess(*s.ProcessID)

	if err != nil {
		if s.Debug {
			fmt.Fprint(s.Cmd.Stdout, fmt.Sprintf("%v\n", err))
		}

		return false
	}

	return true
}
