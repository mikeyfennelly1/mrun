package utils

import (
	"encoding/json"
	"fmt"
	"github.com/opencontainers/runtime-spec/specs-go"
	"github.com/sirupsen/logrus"
	"os"
)

const defaultConfigJson string = `{
	"ociVersion": "1.0.2-dev",
	"process": {
		"terminal": true,
		"user": {
			"uid": 0,
			"gid": 0
		},
		"args": [
			"sh"
		],
		"env": [
			"PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin",
			"TERM=xterm"
		],
		"cwd": "/",
		"capabilities": {
			"bounding": [
				"CAP_AUDIT_WRITE",
				"CAP_KILL",
				"CAP_NET_BIND_SERVICE"
			],
			"effective": [
				"CAP_AUDIT_WRITE",
				"CAP_KILL",
				"CAP_NET_BIND_SERVICE"
			],
			"permitted": [
				"CAP_AUDIT_WRITE",
				"CAP_KILL",
				"CAP_NET_BIND_SERVICE"
			],
			"ambient": [
				"CAP_AUDIT_WRITE",
				"CAP_KILL",
				"CAP_NET_BIND_SERVICE"
			]
		},
		"rlimits": [
			{
				"type": "RLIMIT_NOFILE",
				"hard": 1024,
				"soft": 1024
			}
		],
		"noNewPrivileges": true
	},
	"root": {
		"path": "/home/mfennelly/rootfs",
		"readonly": true
	},
	"hostname": "mrun",
	"mounts": [
		{
			"destination": "/proc",
			"type": "proc",
			"source": "proc"
		},
		{
			"destination": "/dev",
			"type": "tmpfs",
			"source": "tmpfs",
			"options": [
				"nosuid",
				"strictatime",
				"mode=755",
				"size=65536k"
			]
		},
		{
			"destination": "/dev/pts",
			"type": "devpts",
			"source": "devpts",
			"options": [
				"nosuid",
				"noexec",
				"newinstance",
				"ptmxmode=0666",
				"mode=0620",
				"gid=5"
			]
		},
		{
			"destination": "/dev/shm",
			"type": "tmpfs",
			"source": "shm",
			"options": [
				"nosuid",
				"noexec",
				"nodev",
				"mode=1777",
				"size=65536k"
			]
		},
		{
			"destination": "/dev/mqueue",
			"type": "mqueue",
			"source": "mqueue",
			"options": [
				"nosuid",
				"noexec",
				"nodev"
			]
		},
		{
			"destination": "/sys",
			"type": "sysfs",
			"source": "sysfs",
			"options": [
				"nosuid",
				"noexec",
				"nodev",
				"ro"
			]
		},
		{
			"destination": "/sys/fs/cgroup",
			"type": "cgroup",
			"source": "cgroup",
			"options": [
				"nosuid",
				"noexec",
				"nodev",
				"relatime",
				"ro"
			]
		}
	],
	"linux": {
		"resources": {
			"devices": [
				{
					"allow": false,
					"access": "rwm"
				}
			]
		},
		"namespaces": [
			{
				"type": "pid"
			},
			{
				"type": "network"
			},
			{
				"type": "ipc"
			},
			{
				"type": "uts"
			},
			{
				"type": "mount"
			},
			{
				"type": "cgroup"
			}
		],
		"maskedPaths": [
			"/proc/acpi",
			"/proc/asound",
			"/proc/kcore",
			"/proc/keys",
			"/proc/latency_stats",
			"/proc/timer_list",
			"/proc/timer_stats",
			"/proc/sched_debug",
			"/sys/firmware",
			"/proc/scsi"
		],
		"readonlyPaths": [
			"/proc/bus",
			"/proc/fs",
			"/proc/irq",
			"/proc/sys",
			"/proc/sysrq-trigger"
		]
	}
}`
const configJsonPath string = "./config.json"

func GetDefaultConfigJson() *specs.Spec {
	var configJsonStruct specs.Spec
	_ = json.Unmarshal([]byte(defaultConfigJson), &configJsonStruct)
	return &configJsonStruct
}

func GetConfigJson(pathToConfig string) (*specs.Spec, error) {
	// read the file into byte slice
	logrus.Tracef("attempting to read config: %s", pathToConfig)
	configData, err := os.ReadFile(pathToConfig)
	if err != nil {
		logrus.Fatalf("error reading config: %v", err)
		return nil, err
	}

	logrus.Infof("successfully read config")
	logrus.Tracef("ensuring config.json is correct according to schema")
	if !configIsValid(&configData) {
		err = fmt.Errorf("config.json is invalid")
		logrus.Fatalf("config.json is invalid: %v", err)
		return nil, err
	}
	logrus.Tracef("config.json valid according to OCI schema")

	var thisSpec specs.Spec
	err = json.Unmarshal(configData, &thisSpec)
	if err != nil {
		return nil, err
	}

	logrus.Infof("successfully marshalled: %s", pathToConfig)
	return &thisSpec, nil
}

// ConfigJsonExists uses Unix stat syscall to check if file
// metadata is retrievable
func ConfigJsonExists() bool {
	// fail if stat information is not found
	s, err := os.Stat(configJsonPath)
	// if there is an error, fail
	if err != nil {
		return false
	}

	if s != nil {
		return true
	}
	return false
}

// configIsValid checks if a passed config.json file
// is valid according to the config schema in the
// OCI specification.
func configIsValid(config *[]byte) bool {
	//TODO: implement me
	// this is a bit of a PIA to implement, not completely critical
	// to user functionality either just now, will come back at later
	// date
	return true
}
