package controllers

type pidController struct {
	max  int64
	peak int64
}

var DefaultPidController = pidController{
	max:  20,
	peak: 20,
}
