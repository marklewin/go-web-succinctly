package main

type DBUpdate struct {
	Id       int64 `json:"id"`
	Affected int64 `json:"affected"`
}
