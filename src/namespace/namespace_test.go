package namespace

import (
	"encoding/json"
	"github.com/opencontainers/runtime-spec/specs-go"
	"github.com/stretchr/testify/require"
	"testing"
)

/**

config.linux
            "namespaces": {
                "type": "array",
                "items": {
                    "anyOf": [
                        {
                            "$ref": "defs-linux.json#/definitions/NamespaceReference"
                        }
                    ]
                }
            },

defs-linux.json
  "NamespaceReference": {
        "type": "object",
        "properties": {
		"type": {
		"$ref": "#/definitions/NamespaceType"
		},
		"path": {
			"$ref": "defs.json#/definitions/FilePath"
			}
		},
		"required": [
			"type"
			]
		},
*/

func TestNSEnter(t *testing.T) {
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
		return
	}

	var testNamespaceProfile ProcNamespaceProfile
	testNamespaceProfile.Namespaces = testNamespaces
	testNamespaceProfile.ProcessBinary = ""

	testNamespaceProfile.StartShellInNewNamespaces()

	require.NoError(t, err)
}
