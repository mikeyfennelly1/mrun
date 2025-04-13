package fs

import (
	"encoding/json"
	"github.com/mikeyfennelly1/mrun/src/namespace"
	"github.com/opencontainers/runtime-spec/specs-go"
	"github.com/stretchr/testify/require"
	"os"
	"os/exec"
	"syscall"
	"testing"
)

/**
config-schema.json

	"mounts": {
		"type": "array",
		"items": {
			"$ref": "defs.json#/definitions/Mount"
            }
        },
        "root": {
            "description": "Configures the container's root filesystem.",
            "type": "object",
            "required": [
                "path"
            ],
            "properties": {
                "path": {
                    "$ref": "defs.json#/definitions/FilePath"
                },
                "readonly": {
				"type": "boolean"
			}
		}
	},



  "root": {
      "description": "Configures the container's root filesystem.",
      "type": "object",
      "required": [
          "path"
      ],
      "properties": {
          "path": {
              "$ref": "defs.json#/definitions/FilePath"
          },
          "readonly": {
              "type": "boolean"
          }
      }
  },

config-linux.json
	"rootfsPropagation": {
		"$ref": "defs-linux.json#/definitions/RootfsPropagation"
	},

	"maskedPaths": {
		"$ref": "defs.json#/definitions/ArrayOfStrings"
	},
		"readonlyPaths": {
		"$ref": "defs.json#/definitions/ArrayOfStrings"
	},
		"mountLabel": {
		"type": "string"
	},

*/

func TestCreateFileSystem(t *testing.T) {
	// get array of bytes from config.json
	jsonContent, err := os.ReadFile("../config.json")
	require.NoError(t, err)

	var spec specs.Spec
	err = json.Unmarshal(jsonContent, &spec)
	require.NoError(t, err)

	// could be worthwhile to start a new bash in a process namespace
	// and see how we can remove mounts that may cause conflicts
	// with mount operations in new mount namespace

	// change root and mount new filesystems
	err = CreateFileSystem(spec)
	require.NoError(t, err)

	// enter new namespaces
	// at this point - can still see everything
	pNS, err := namespace.GetIsolatedProcessProfile()
	require.NoError(t, err)

	// start another process for the
	cmd := exec.Command("/proc/self/exe")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: pNS.GetCloneFlagBitMask(),
	}

}
