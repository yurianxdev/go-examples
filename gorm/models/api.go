package models

type ResponseError struct {
	Errors []string
}

type ResponseSucceed struct {
	Message string
}
