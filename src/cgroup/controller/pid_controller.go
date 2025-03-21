package controller

import "strconv"

type PidController struct {
	max  int
	peak int
}

func (p PidController) GetSubControllerUpdates() []map[controllerFilename]string {
	var subControllerUpdates []map[controllerFilename]string

	subControllerUpdates = append(subControllerUpdates, map[controllerFilename]string{
		"pid.max": strconv.Itoa(p.max),
	})
	subControllerUpdates = append(subControllerUpdates, map[controllerFilename]string{
		"pid.peak": strconv.Itoa(p.peak),
	})

	return subControllerUpdates
}
