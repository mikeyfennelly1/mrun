package src

import (
	"fmt"
	"github.com/opencontainers/runtime-spec/specs-go"
	"github.com/sirupsen/logrus"
	"github.com/syndtr/gocapability/capability"
	"os"
	"strings"
)

// SetFileCapabilities for the binary for the init process of
// the container. This requires usage of a file system that
// supports extended file attributes.
//
// https://man7.org/linux/man-pages/man7/xattr.7.html
func SetFileCapabilities(spec specs.Spec, filename string) error {
	// apply permitted capabilities to /bin/sh
	setAndApplyCapsetToFile(capability.PERMITTED,
		spec.Process.Capabilities.Permitted,
		filename)

	// apply permitted inheritable to /bin/sh
	setAndApplyCapsetToFile(capability.INHERITABLE,
		spec.Process.Capabilities.Inheritable,
		filename)

	// apply effective inheritable to /bin/sh
	setAndApplyCapsetToFile(capability.EFFECTIVE,
		spec.Process.Capabilities.Permitted,
		filename)

	return nil
}

func setAndApplyCapsetToFile(capset capability.CapType, capsetCaps []string, file string) {
	panic("Implement setAndApplyCapsetToFile")
}

func SetAndApplyCapsetToCurrentPid(capabilitySet capability.CapType, capabilities []string) {
	for _, thisCap := range capabilities {
		err := SetAndApplyCapToCurrentPid(capabilitySet, getCap(thisCap))
		if err != nil {
			logrus.Errorf("%v", err)
		}
	}
}

func SetAndApplyCapToCurrentPid(capabilitySet capability.CapType, which capability.Cap) error {
	pid := os.Getpid()
	procCaps, err := capability.NewPid2(pid)
	if err != nil {
		return fmt.Errorf("error getting process capabilities: %v\n", err)
	}
	err = procCaps.Load()
	if err != nil {
		return err
	}

	procCaps.Set(capabilitySet, which)

	err = procCaps.Apply(capability.CAPS)
	if err != nil {
		capName := strings.ToUpper(fmt.Sprintf("CAP_%v", which))
		capSetName := strings.ToUpper(fmt.Sprintf("%v", capabilitySet))
		bold, reset := "\033[1m", "\033[0m"
		return fmt.Errorf("error applying capability %s%v%s to the %s%v%s capability set: %v\n", bold, capName, reset, bold, capSetName, reset, err)
	}

	return nil
}

func PrintCapabilityStatus() error {
	filename := fmt.Sprintf("/proc/%d/status", os.Getpid())
	contents, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	fmt.Printf("/proc/pid/status: %s\n\n", string(contents))

	return nil
}

func getCap(which string) capability.Cap {
	switch which {
	case "CAP_AUDIT_CONTROL":
		return capability.CAP_AUDIT_CONTROL
	case "CAP_AUDIT_READ":
		return capability.CAP_AUDIT_READ
	case "CAP_AUDIT_WRITE":
		return capability.CAP_AUDIT_WRITE
	case "CAP_BLOCK_SUSPEND":
		return capability.CAP_BLOCK_SUSPEND
	case "CAP_BPF":
		return capability.CAP_BPF
	case "CAP_CHECKPOINT_RESTORE":
		return capability.CAP_CHECKPOINT_RESTORE
	case "CAP_CHOWN":
		return capability.CAP_CHOWN
	case "CAP_DAC_OVERRIDE":
		return capability.CAP_DAC_OVERRIDE
	case "CAP_DAC_READ_SEARCH":
		return capability.CAP_DAC_READ_SEARCH
	case "CAP_FOWNER":
		return capability.CAP_FOWNER
	case "CAP_FSETID":
		return capability.CAP_FSETID
	case "CAP_IPC_LOCK":
		return capability.CAP_IPC_LOCK
	case "CAP_IPC_OWNER":
		return capability.CAP_IPC_OWNER
	case "CAP_KILL":
		return capability.CAP_KILL
	case "CAP_LEASE":
		return capability.CAP_LEASE
	case "CAP_LINUX_IMMUTABLE":
		return capability.CAP_LINUX_IMMUTABLE
	case "CAP_MAC_ADMIN":
		return capability.CAP_MAC_ADMIN
	case "CAP_MAC_OVERRIDE":
		return capability.CAP_MAC_OVERRIDE
	case "CAP_MKNOD":
		return capability.CAP_MKNOD
	case "CAP_NET_ADMIN":
		return capability.CAP_NET_ADMIN
	case "CAP_NET_BIND_SERVICE":
		return capability.CAP_NET_BIND_SERVICE
	case "CAP_NET_BROADCAST":
		return capability.CAP_NET_BROADCAST
	case "CAP_NET_RAW":
		return capability.CAP_NET_RAW
	case "CAP_PERFMON":
		return capability.CAP_PERFMON
	case "CAP_SETGID":
		return capability.CAP_SETGID
	case "CAP_SETFCAP":
		return capability.CAP_SETFCAP
	case "CAP_SETPCAP":
		return capability.CAP_SETPCAP
	case "CAP_SETUID":
		return capability.CAP_SETUID
	case "CAP_SYS_ADMIN":
		return capability.CAP_SYS_ADMIN
	case "CAP_SYS_BOOT":
		return capability.CAP_SYS_BOOT
	case "CAP_SYS_CHROOT":
		return capability.CAP_SYS_CHROOT
	case "CAP_SYS_MODULE":
		return capability.CAP_SYS_MODULE
	case "CAP_SYS_NICE":
		return capability.CAP_SYS_NICE
	case "CAP_SYS_PACCT":
		return capability.CAP_SYS_PACCT
	case "CAP_SYS_PTRACE":
		return capability.CAP_SYS_PTRACE
	case "CAP_SYS_RAWIO":
		return capability.CAP_SYS_RAWIO
	case "CAP_SYS_RESOURCE":
		return capability.CAP_SYS_RESOURCE
	case "CAP_SYS_TIME":
		return capability.CAP_SYS_TIME
	case "CAP_SYS_TTY_CONFIG":
		return capability.CAP_SYS_TTY_CONFIG
	case "CAP_SYSLOG":
		return capability.CAP_SYSLOG
	case "CAP_WAKE_ALARM":
		return capability.CAP_WAKE_ALARM
	default:
		// log an error
		err := fmt.Errorf("unknown value capability value, given: %s", which)
		logrus.Warn("%v", err)

		// return 0 to have no resulting bitmask effects
		return 0
	}
}
