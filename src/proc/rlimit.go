package proc

import (
	"github.com/opencontainers/runtime-spec/specs-go"
	"github.com/sirupsen/logrus"
	"golang.org/x/sys/unix"
)

func SetRLIMITsForProcess(processRlimits []specs.POSIXRlimit) {
	for _, rlimit := range processRlimits {
		var thisRlim unix.Rlimit
		thisRlim.Cur = rlimit.Soft
		thisRlim.Max = rlimit.Hard
		err := unix.Setrlimit(getRlimit(rlimit.Type), &thisRlim)
		if err != nil {
			logrus.Warnf("error setting rate limit '%s' to container process: %v", rlimit.Type, err)
		}
	}
}

func getRlimit(rlimitStr string) int {
	switch rlimitStr {
	case "RLIMIT_NOFILE":
		return unix.RLIMIT_NOFILE
	case "RLIMIT_AS":
		return unix.RLIMIT_AS
	case "RLIMIT_CORE":
		return unix.RLIMIT_CORE
	case "RLIMIT_CPU":
		return unix.RLIMIT_CPU
	case "RLIMIT_DATA":
		return unix.RLIMIT_DATA
	case "RLIMIT_FSIZE":
		return unix.RLIMIT_FSIZE
	case "RLIMIT_LOCKS":
		return unix.RLIMIT_LOCKS
	case "RLIMIT_MEMLOCK":
		return unix.RLIMIT_MEMLOCK
	case "RLIMIT_MSGQUEUE":
		return unix.RLIMIT_MSGQUEUE
	case "RLIMIT_NICE":
		return unix.RLIMIT_NICE
	case "RLIMIT_NPROC":
		return unix.RLIMIT_NPROC
	case "RLIMIT_RSS":
		return unix.RLIMIT_RSS
	case "RLIMIT_RTPRIO":
		return unix.RLIMIT_RTPRIO
	case "RLIMIT_RTTIME":
		return unix.RLIMIT_RTTIME
	case "RLIMIT_SIGPENDING":
		return unix.RLIMIT_SIGPENDING
	case "RLIMIT_STACK":
		return unix.RLIMIT_STACK
	default:
		return 0
	}
}
