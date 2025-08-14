package init

import (
	"fmt"
	"github.com/mikeyfennelly1/mrun/state"
	"github.com/mikeyfennelly1/mrun/utils"
	"github.com/sirupsen/logrus"
	"os"
)

// Init is used to initialize the hub
//
// It determines the current state of the container (whether)
// it be non-existent or in a partially created state, and initializes
// based on that.
func Init(state state.StateSubsytem) (HubInterface, error) {
	// check if the config.json exists in the pwd
	if !utils.ConfigJsonExists() {
		logrus.Fatal("./config.json does not exist")
		os.Exit(1)
	}
	// read contents of config.json
	contents, err := utils.GetConfigJsonContents()
	if err != nil {
		logrus.Fatal("./config.json does not exist")
		os.Exit(1)
	}
	// validate that it is structurally correct

	return nil, nil
}
