// memory.go

package controllers

type memController struct {
	memLimits  *memLimits
	swapLimits *swapLimits
}

type memLimits struct {
	high int64
	low  int64
	max  int64
	min  int64
	peak int64
}

// swapLimits
// limits for swap space
type swapLimits struct {
	high int64
	max  int64
	peak int64
}

var DefaultMemController = memController{
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

func (mc *memController) write(parentDirPath string) {

}
