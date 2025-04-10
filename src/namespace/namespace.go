package namespace

import (
	"github.com/opencontainers/runtime-spec/specs-go"
	"golang.org/x/sys/unix"
	"log"
	"os"
	"os/exec"
	"syscall"
)

type ProcNamespaceProfile struct {
	// the binary to run when we enter the new namespace
	// if left as an empty string, process binary will be bash
	ProcessBinary string

	// these namespaces will correspond to the clone flags
	// that are used when this proc is cloned
	Namespaces []specs.LinuxNamespace
}

func (p *ProcNamespaceProfile) StartBashInNewNamespaces() {
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

func (p *ProcNamespaceProfile) getCloneFlagBitMask() uintptr {
	result := 0

	for _, ns := range p.Namespaces {
		switch ns.Type {
		case "mount":
			result |= unix.CLONE_NEWNS
		case "pid":
			result |= unix.CLONE_NEWPID
		case "network":
			result |= unix.CLONE_NEWNET
		case "uts":
			result |= unix.CLONE_NEWUTS
		case "ipc":
			result |= unix.CLONE_NEWIPC
		case "user":
			result |= syscall.CLONE_NEWUSER
		case "cgroup":
			result |= unix.CLONE_NEWCGROUP
		case "time":
			result |= unix.CLONE_NEWTIME
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
