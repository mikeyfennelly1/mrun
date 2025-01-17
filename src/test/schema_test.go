package test

import (
	"github.com/xeipuuv/gojsonschema"
	"os"
	"testing"
)

func TestSchema(t *testing.T) {
	schemaLoader := gojsonschema.NewReferenceLoader("/home/mfennelly/projects/mrun/src/test/config-schema.json")

	// Read the configuration file
	configData, err := os.ReadFile("./exampleConfig.json")
	if err != nil {
		t.Fatalf("Failed to read config file: %v", err)
	}

	// Load the configuration JSON into a loader
	configLoader := gojsonschema.NewStringLoader(string(configData))

	// Validate the configuration against the schema
	result, err := gojsonschema.Validate(schemaLoader, configLoader)
	if err != nil {
		t.Fatalf("Schema validation failed: %v", err)
	}

	// Assert validation result
	if !result.Valid() {
		t.Errorf("Validation failed. The document is not schema-compliant.")
		for _, err := range result.Errors() {
			t.Errorf("- %s", err)
		}
	} else {
		t.Log("Validation succeeded. The document is schema-compliant.")
	}
}
