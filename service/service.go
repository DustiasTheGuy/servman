package service

import (
	"fmt"
	"os"
	"os/exec"
)

// Service contains nessecary information about a proccess
type Service struct {
	ProccessID int    // A reference to the proccess that can be used to terminate or modify the proccess
	Path       string // where is the binary located?
	WorkingDir string // what directory should the program launch from? very useful if the program is serving static files from a relative path
	Debug      bool   // output useful messages
	Cmd        *exec.Cmd
}

// ConstructCmd creates a pointer to a cmd which can be used to execute commands
func (s *Service) ConstructCmd() *exec.Cmd {
	return &exec.Cmd{
		Path:   s.Path,
		Args:   []string{"START"},
		Dir:    s.WorkingDir,
		Stderr: os.Stderr,
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
	}
}

// StartService executes a start commandand prints the pid to stdout
func (s *Service) StartService() error {
	s.Cmd = s.ConstructCmd()

	if err := s.Cmd.Start(); err != nil {
		return err
	}

	s.ProccessID = s.Cmd.Process.Pid
	if s.Debug {
		fmt.Fprint(s.Cmd.Stdout, fmt.Sprintf("Started Proccess: %d\n", s.Cmd.Process.Pid))
	}

	return nil
}

// IsAlive checkes if a proccess is currently alive
func (s *Service) IsAlive() bool {
	_, err := os.FindProcess(s.ProccessID)

	if err != nil {
		if s.Debug {
			fmt.Fprint(s.Cmd.Stdout, fmt.Sprintf("%v\n", err))
		}

		return false
	}

	return true
}
