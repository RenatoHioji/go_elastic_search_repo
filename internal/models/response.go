package models

type ResponseList[T any] struct {
	Data T `json:"data"`
}

type Response[T any] struct {
	Data T `json:"data"`
}
