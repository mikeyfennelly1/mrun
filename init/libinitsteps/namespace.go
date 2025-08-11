package libinitsteps

import (
	"encoding/json"
	"github.com/mikeyfennelly1/mrun/init"
	"github.com/opencontainers/runtime-spec/specs-go"
	"golang.org/x/sys/unix"
	"log"
	"os"
	"os/exec"
	"syscall"
)

type namespaceLink struct {
	next init.Step
}

func (nci *namespaceLink) execute(spec *specs.Spec) {

}

type processNamespaceProfile struct {
	// the binary to run when we enter the new namespace
	// if left as an empty string, process binary will be bash
	ProcessBinary string

	// these namespaces will correspond to the clone flags
	// that are used when this proc is cloned
	Namespaces []specs.LinuxNamespace
}

func getIsolatedProcessProfile() (*processNamespaceProfile, error) {
	jsonNamespaces := `[
			{ "type": "pid" },
			{ "type": "network" },
			{ "type": "ipc" },
			{ "type": "uts" },
			{ "type": "mount" },
			{ "type": "cgroup" }
		]`

	var testNamespaces []specs.LinuxNamespace
	err := json.Unmarshal([]byte(jsonNamespaces), &testNamespaces)
	if err != nil {
		return nil, err
	}

	var testNamespaceProfile processNamespaceProfile
	testNamespaceProfile.Namespaces = testNamespaces
	testNamespaceProfile.ProcessBinary = ""

	return &testNamespaceProfile, nil
}

func (p *processNamespaceProfile) startShellInNewNamespaces() {
	cmd := exec.Command("/bin/sh")

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

func (p *processNamespaceProfile) getCloneFlagBitMask() uintptr {
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

// restartInNewNS execs the current program, but in new namespaces
// according to the config file namespaces.
func restartInNewNS(args ...string) error {
	p, err := getIsolatedProcessProfile()
	if err != nil {
		return err
	}

	cmd := exec.Command("/proc/self/exe", args...)

	// Set the command to run in a new mount namespace
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: p.getCloneFlagBitMask(),
	}

	// Set input/output to standard io options
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the command
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
