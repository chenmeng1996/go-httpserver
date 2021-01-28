package main

var (
	SUCCESS = Result{
		Code:    200,
		Message: "success",
		Data:    nil,
	}
)

type Result struct {
	Code    int
	Message string
	Data    interface{}
}
