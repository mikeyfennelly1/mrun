package fs

import (
	"encoding/json"
	"github.com/opencontainers/runtime-spec/specs-go"
	"github.com/stretchr/testify/require"
	"os"
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
	jsonContent, err := os.ReadFile("./config.json")
	require.NoError(t, err)

	var spec specs.Spec
	err = json.Unmarshal(jsonContent, &spec)
	require.NoError(t, err)
	
	err = CreateFileSystem(spec)
	require.NoError(t, err)
}
