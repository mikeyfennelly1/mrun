// memory_controller.go

package controller

import "strconv"

type controllerFilename string

type MemoryController struct {
	memLimits  *memLimits
	swapLimits *swapLimits
}

type memLimits struct {
	high int
	low  int
	max  int
	min  int
	peak int
}

// swapLimits
// limits for swap space
type swapLimits struct {
	high int
	max  int
	peak int
}

var DefaultMemController = MemoryController{
	memLimits:  DefaultMemLimits,
	swapLimits: DefaultSwapLimits,
}

var DefaultSwapLimits = &swapLimits{
	high: 40000,
	max:  40000,
	peak: 40000,
}

var DefaultMemLimits = &memLimits{
	high: 40000,
	low:  0,
	max:  40000,
	min:  0,
	peak: 40000,
}

func (mc *MemoryController) GetTargetWriteValKVPs() map[controllerFilename]string {
	fileWriteValMap := map[controllerFilename]string{
		"memory.high":      strconv.Itoa(mc.memLimits.high),
		"memory.low":       strconv.Itoa(mc.memLimits.low),
		"memory.max":       strconv.Itoa(mc.memLimits.max),
		"memory.min":       strconv.Itoa(mc.memLimits.min),
		"memory.peak":      strconv.Itoa(mc.memLimits.peak),
		"memory.swap.high": strconv.Itoa(mc.swapLimits.high),
		"memory.swap.max":  strconv.Itoa(mc.swapLimits.max),
		"memory.swap.peak": strconv.Itoa(mc.swapLimits.peak),
	}

	return fileWriteValMap
}
