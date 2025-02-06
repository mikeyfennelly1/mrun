package spec

// this is the specification for the runtime.
// this is NOT the specification for an image

// Process contains information to start specific application inside the container
type Process struct {
	// Terminal creates an interactive terminal for container
	Terminal bool `json:"terminal,omitempty"`
	User     User `json:"user"`
	// Env populates the process environment for the process.
	Env []string `json:"env,omitempty"`
	// Cwd is the current working directory for the process and must be
	// relative to the container's root.
	Cwd string `json:"cwd"`
	// Capabilities are Linux capabilities that are kept for the process.
	Capabilities *LinuxCapabilities `json:"capabilities,omitempty" platform:"linux"`
	// Rlimits specifies rlimit options to apply to the process.
	Rlimits []POSIXRlimit `json:"rlimits,omitempty" platform:"linux,solaris,zos"`
	// NoNewPrivileges controls whether additional privileges could be gained by processes in the container.
	NoNewPrivileges bool `json:"noNewPrivileges,omitempty" platform:"linux,zos"`
	// ApparmorProfile specifies the apparmor profile for the container.
}

type User struct {
	UID uint32 `json:"uid" platform:"linux"`
	GID uint32 `json:"gid" platform:"linux"`
	// umask is the umask for the init process
	Umask *uint32 `json:"umask,omitempty" platform:"linux"`
}

// LinuxCapabilities specifies the list of allowed capabilities that are kept for a process
// generally limit privilege escalation in process
//
// Managed in task_struct for proc
//
// You can set them with unix.Prctl to exec raw syscall
type LinuxCapabilities struct {
	Bounding []string `json:"bounding,omitempty" platform:"linux"`

	// Effective allows/disallows syscalls in kernel
	Effective []string `json:"effective,omitempty" platform:"linux"`

	// Generally not preserved across execve() when running as non-root user
	//
	// added to permitted set and remain inheritable when exec'd binary has corresponding
	// bits set in file inheritable set
	Inheritable []string `json:"inheritable,omitempty" platform:"linux"`

	Permitted []string `json:"permitted,omitempty" platform:"linux"`

	// Ambient capabilities persist across execve() syscalls
	//
	// useful if you want to run helper proc with specified caps
	//
	//execve() executes a new program within a running process.
	// when execve() is called, most of PCB is wiped
	Ambient []string `json:"ambient,omitempty" platform:"linux"`
}

type POSIXRlimit struct {
	Type string `json:"type"`
	Hard uint64 `json:"hard"`
	Soft uint64 `json:"soft"`
}

type LinuxNamespace struct {
	Type LinuxNamespaceType `json:"type"`
	Path string             `json:"path"` // path to namespace in /proc sys virtual fs
}

type LinuxNamespaceType string
