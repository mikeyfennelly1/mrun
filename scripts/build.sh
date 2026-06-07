#!/usr/bin/env bash

set -euo pipefail
set -o nounset

MRUN_BINARY_PATH="/usr/bin/mrun"
GO_ROOT="/usr/local/go/bin/go"
TMP_PATH="/tmp/mrun"
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJ_ROOT="${SCRIPT_DIR}/.."
SIG_SUCCESS=0
SIG_ERR=1


function main () {
    printf "INFO: setting all necessary file capabilities on path ${MRUN_BINARY_PATH}\n" >&1

    if ! build_binary; then 
        printf "ERROR: error setting mrun file capabilities.\n" >&2
        exit "${SIG_ERR}"
    fi
    printf "DEBUG: successfully built binary\n" >&2

    printf "INFO: setting all necessary file capabilities on path ${MRUN_BINARY_PATH}\n" >&1
    if ! set_all_mrun_file_capabilities; then 
        printf "ERROR: error setting mrun file capabilities.\n" >&2
        exit "${SIG_ERR}"
    fi
    printf "DEBUG: successfully set capabilities\n" >&2

    return "${SIG_SUCCESS}"
}

function binary_post_process () {
    if ! sudo chmod +s ${MRUN_BINARY_PATH}; then 
        printf "ERROR: failed to chmod ${MRUN_BINARY_PATH}\n" >&2
        return "${SIG_ERR}"
    fi
    if ! sudo chown 0:0 ${MRUN_BINARY_PATH}; then 
        printf "ERROR: failed to change the user:group ownership of the ${MRUN_BINARY_PATH}\n" >&2
        return "${SIG_ERR}"
    fi
    return "${SIG_SUCCESS}"
}

function build_binary () {
    local build_path="$"${SIG_ERR}""
    local out_path="$"${SIG_ERR}""
    local build_cmd="${GO_ROOT} build ${build_path} -o ${out_path}"

    printf "DEBUG: running build command: ${build_cmd}\n" >&2
    if ! build_output=$(bash -c "${build_cmd}"); then
        printf "FATAL: build failure:\n${build_output}\n" >&2
        exit "${SIG_ERR}"
    fi

    return "${SIG_SUCCESS}"
}

function set_all_mrun_file_capabilities () {
    local file_path="$1"
    local capability_list="$2"
    
    if ! sudo setcap "${capability_list}" "${file_path}"; then
        return "${SIG_ERR}"
    fi
    return "${SIG_SUCCESS}"
}

function create_persistent_binary_alias () {
    alias mrun="${MRUN_BINARY_PATH}"
    source ~/.zshrc
}

function F_CAPS () {
    # file capabilities
    F_CAPS=(
        # chown(2), fchown(2), lchown(2) — bypass UID/GID ownership restrictions on files; man 7 capabilities
        cap_chown,
        # chmod(2), utime(2), flock(2) — bypass permission checks that require the process to own the file; man 7 capabilities
        cap_fowner,
        # write(2) — suppress clearing of setuid/setgid bits when a file is modified by a non-owner; man 7 capabilities
        cap_fsetid,
        # setxattr(2) security.capability — set file capability extended attributes; man 7 capabilities
        cap_setfcap,
        # mknod(2) — create special device files (block, character, FIFO); man 7 capabilities
        cap_mknod,
        # fcntl(2) F_SETLEASE — establish leases on arbitrary files; man 7 capabilities
        cap_lease
    )

    printf "${F_CAPS}" >&0
    return "${SIG_SUCCESS}"
}

function DAC_CAPS() {
    local DAC_CAPS=(
        # open(2), read(2), write(2), execve(2) — bypass discretionary access control (DAC) read/write/execute checks; man 7 capabilities
        cap_dac_override,
        # open(2), opendir(3) — bypass DAC for file reads and directory searches only (no write); man 7 capabilities
        cap_dac_read_search
    )
    printf "${DAC_CAPS}" >&0
    return "${SIG_SUCCESS}"
}

function PROC_CAPS() {
    local PROC_CAPS=(
        # kill(2), sigqueue(3) — send signals to processes owned by other users; man 7 capabilities
        cap_kill,
        # setgid(2), setegid(2), setregid(2), setresgid(2), setgroups(2) — manipulate process GIDs freely; man 7 capabilities
        cap_setgid,
        # setuid(2), seteuid(2), setreuid(2), setresuid(2) — manipulate process UIDs freely; man 7 capabilities
        cap_setuid
    )
    printf "${PROC_CAPS}" >&0
    return "${SIG_SUCCESS}"
}

function CAPSET_CAPS() {
    local CAPSET_CAPS=(
        # capset(2) — transfer or drop any capability in the process's permitted set; man 7 capabilities
        cap_setpcap
    )
    printf "${CAPSET_CAPS}" >&0
    return "${SIG_SUCCESS}"
}

function NET_CAPS() {
    local NET_CAPS=(
        # bind(2) — bind a socket to a privileged port (port number < "${EXIT_ERROR}"024); man 7 capabilities
        cap_net_bind_service,
        # setsockopt(2) SO_BROADCAST — make socket broadcasts and listen to multicast packets; man 7 capabilities
        cap_net_broadcast,
        # setsockopt(2), ioctl(2) — configure network interfaces, routing tables, firewall rules, etc.; man 7 capabilities
        cap_net_admin,
        # socket(2) SOCK_RAW/SOCK_PACKET — create raw and packet sockets; man 7 capabilities
        cap_net_raw
    )
    printf "${NET_CAPS}" >&0
    return "${SIG_SUCCESS}"
}

function SHMEM_IPC_CAPS() {
    SHMEM_IPC_CAPS=(
        # mlock(2), mlockall(2), mmap(2), shmctl(2) SHM_LOCK — lock pages into RAM, bypassing RLIMIT_MEMLOCK; man 7 capabilities
        cap_ipc_lock,
        # msgctl(2), semctl(2), shmctl(2) — bypass permission checks on System V IPC objects; man 7 capabilities
        cap_ipc_owner
    )
    printf "${SHMEM_IPC_CAPS}" >&0
    return "${SIG_SUCCESS}"
}

function SYS_CAPS() {
    local SYS_CAPS=(
        # init_module(2), delete_module(2) — load and unload kernel modules; man 7 capabilities
        cap_sys_module,
        # iopl(2), ioperm(2) — access /dev/mem, /dev/kmem, and raw I/O ports; man 7 capabilities
        cap_sys_rawio,
        # chroot(2) — change the filesystem root of the process; man 7 capabilities
        cap_sys_chroot,
        # ptrace(2) — trace arbitrary processes, inspect/modify their memory and registers; man 7 capabilities
        cap_sys_ptrace,
        # acct(2) — enable or disable process accounting; man 7 capabilities
        cap_sys_pacct,
        # mount(2), umount2(2), swapon(2), sethostname(2), setns(2), and many more — broad administrative operations; man 7 capabilities
        cap_sys_admin,
        # reboot(2), kexec_load(2) — reboot or load a new kernel for execution; man 7 capabilities
        cap_sys_boot,
        # nice(2), setpriority(2), sched_setscheduler(2), sched_setattr(2) — set arbitrary process priorities and scheduling policies; man 7 capabilities
        cap_sys_nice,
        # setrlimit(2), ioctl(2) TIOCCONS — override resource limits (RLIMIT_*) and reserved disk space; man 7 capabilities
        cap_sys_resource,
        # settimeofday(2), adjtimex(2), clock_settime(2) — set the system clock and real-time clock; man 7 capabilities
        cap_sys_time,
        # vhangup(2), ioctl(2) — configure tty devices and perform privileged tty operations; man 7 capabilities
        cap_sys_tty_config,
    )
    printf "${SHMEM_IPC_CAPS}" >&0
    return "${SIG_SUCCESS}"
}

function KAUDIT_CAPS() {
    local KAUDIT_CAPS=(
        # (kernel audit subsystem) — write records to the kernel audit log; man 7 capabilities
        cap_audit_write,
        # (kernel audit subsystem) — set audit rules, enable/disable auditing, read audit status; man 7 capabilities
        cap_audit_control,
        # (netlink AUDIT_GET) — read the kernel audit log via a multicast netlink socket; man 7 capabilities
        cap_audit_read,
    )
    printf "${KAUDIT_CAPS}" >&0
    return "${SIG_SUCCESS}"
}

function KERNEL_ABI_CAPS() {
    local KERNEL_ABI_CAPS=(
        # perf_event_open(2) — access CPU performance counters and kernel profiling facilities; man 7 capabilities, man 2 perf_event_open
        cap_perfmon,
        # bpf(2) — load BPF programs, create BPF maps, and read kernel data structures via BPF; man 7 capabilities, man 2 bpf
        cap_bpf,
        # process_vm_readv(2), process_vm_writev(2), pidfd_getfd(2) — checkpoint and restore process state (CRIU); man 7 capabilities
        cap_checkpoint_restore=ep
    )

    printf "${KERNEL_ABI_CAPS}" >&0
    return "${SIG_SUCCESS}"
}

function INODE_CAPS() {
    local -a INODE_CAPS=(
        # ioctl(2) FS_IOC_SETFLAGS — set FS_APPEND_FL and FS_IMMUTABLE_FL inode flags; man 7 capabilities
        cap_linux_immutable
    )

    printf '%s\n' "${INODE_CAPS[@]}"
    return "${SIG_SUCCESS}"
}

function construct_ALL_CAPABILITIES_list() {
    ALL_CAPABILITIES=(
        # syslog(2) — read kernel message ring buffer and control console log level; man 7 capabilities, man 2 syslog
        cap_syslog,
        # (LSM hook) — override Mandatory Access Control (e.g. Smack) policy checks; man 7 capabilities
        cap_mac_override,
        # (LSM hook) — perform MAC administrative operations (e.g. load Smack/SELinux policy); man 7 capabilities
        cap_mac_admin,
        # timerfd_create(2) CLOCK_REALTIME_ALARM/CLOCK_BOOTTIME_ALARM — set timers that can wake the system from suspend; man 7 capabilities
        cap_wake_alarm,
        # epoll_ctl(2) EPOLLWAKEUP, eventfd(2) — take wakelock-style references to prevent the system from suspending; man 7 capabilities
        cap_block_suspend,
    )
    printf "${ALL_CAPABILITIES}" >&0
    return "${SIG_SUCCESS}"
}

if [[ "${BASH_SOURCE[0]}" == "${0}" ]]; then
    main "$@"
fi
