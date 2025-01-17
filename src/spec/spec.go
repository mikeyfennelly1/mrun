package spec

import "github.com/xeipuuv/gojsonschema"

const SCHEMA_PATH = "./config-schema.json"

// GenerateSpec - tester function just to create a config.json
//
// Should be deleted.
func GenerateSpec() {
	schemLoader := gojsonschema.NewReferenceLoader(SCHEMA_PATH)

	mrunGeneratedJSON :=
}
