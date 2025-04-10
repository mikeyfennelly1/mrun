package namespace

import (
	"github.com/opencontainers/runtime-spec/specs-go"
	"log"
	"os"
	"os/exec"
	"syscall"
)

type procNamespaceProfile struct {
	// the binary to run when we enter the new namespace
	// if left as an empty string, process binary will be bash
	processBinary string

	// these namespaces will correspond to the clone flags
	// that are used when this proc is cloned
	Namespaces []specs.LinuxNamespaceType
}

func (p *procNamespaceProfile) startBashInNewNamespaces() {
	cmd := exec.Command("/bin/bash")

	// Set the command to run in a new mount namespace
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: p.getCloneFlagBitMask(),
	}

	// Set input/output to inherit from current process
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the command
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

func (p *procNamespaceProfile) getCloneFlagBitMask() uintptr {
	result := 0

	for _, ns := range p.Namespaces {
		switch ns {
		case "mount":
			result |= syscall.CLONE_NEWNS
		case "pid":
			result |= syscall.CLONE_NEWPID
		case "network":
			result |= syscall.CLONE_NEWNET
		case "uts":
			result |= syscall.CLONE_NEWUTS
		case "ipc":
			result |= syscall.CLONE_NEWIPC
		case "user":
			result |= syscall.CLONE_NEWUSER
		case "cgroup":
			result |= syscall.CLONE_NEWCGROUP
		case "time":
			result |= syscall.CLONE_NEWTIME
		default:
			return 0
		}
	}

	return uintptr(result)
}

/**
Containers should by default not control the terminal.
To exec into a container is to start a bash in the current
terminal session, that resides in the namespace of that container.
*/
