package spec

import (
	"golang.org/x/sys/unix"
	"log"
)

func setCap() {
	// This is the equivalent of runnning prctl() from <sys/prctl.h>
	// First argument describes what to do, the next describes what to do it to.
	//
	// In this case we are instructing the kernel, drop the capability (PR_CAPSET_DROP),
	// of this process, to have the CAP_NET_RAW capability.
	// This is one of 4 network capabilities; see 'man capabilities' for more info.
	//
	// arg3, arg4, and arg5 are all options whose necessity depend on preceding option selections.

	err := unix.Prctl(unix.PR_CAPBSET_DROP, unix.CAP_NET_RAW, 0, 0, 0)
	if err != nil {
		log.Fatalf("Error setting capabilities.\n")
	}
}
