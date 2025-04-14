package process

import (
	"github.com/opencontainers/runtime-spec/specs-go"
	"github.com/syndtr/gocapability/capability"
	"os"
)

func setCapabilities(spec specs.Spec) {
	procCaps, err := capability.NewPid2(os.Getpid())
	if err != nil {
		return
	}

	for _, capStr := range spec.Process.Capabilities.Ambient {
		thisCap := getCap(capStr)
		procCaps.Set(capability.AMBIENT, thisCap)
	}

	for _, capStr := range spec.Process.Capabilities.Inheritable {
		thisCap := getCap(capStr)
		procCaps.Set(capability.INHERITABLE, thisCap)
	}

	for _, capStr := range spec.Process.Capabilities.Effective {
		thisCap := getCap(capStr)
		procCaps.Set(capability.EFFECTIVE, thisCap)
	}

	for _, capStr := range spec.Process.Capabilities.Bounding {
		thisCap := getCap(capStr)
		procCaps.Set(capability.BOUNDING, thisCap)
	}

	for _, capStr := range spec.Process.Capabilities.Permitted {
		thisCap := getCap(capStr)
		procCaps.Set(capability.PERMITTED, thisCap)
	}
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
		return 0
	}
}
