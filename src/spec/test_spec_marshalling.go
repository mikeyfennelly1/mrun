package spec

import (
	"encoding/json"
	"fmt"
	"log"
)

type TestSpec struct {
	Version     string            `json:"ociVersion"`
	Hostname    string            `json:"hostname,omitEmpty"`
	Domainname  string            `json:"domainname,omitempty"`
	Annotations map[string]string `json:"annotations,omitempty"`
}

func TestSpecMarshalling() {
	testSpec := TestSpec{
		Version:    "v1.2.3",
		Hostname:   "mfennellyComputer",
		Domainname: "mikeDomain",
		Annotations: map[string]string{
			"author": "john doe",
			"usage":  "test container",
		},
	}

	jsonData, err := json.MarshalIndent(testSpec, "", "  ")
	if err != nil {
		log.Fatalf("Could not marshal json data")
	}

	fmt.Println(string(jsonData))
}
