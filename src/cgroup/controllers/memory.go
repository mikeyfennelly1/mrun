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

type swapLimits struct {
	high int64
	max  int64
	peak int64
}
