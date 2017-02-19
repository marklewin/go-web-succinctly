package main

type City struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	CountryCode string `json:"country"`
	District    string `json:"district"`
	Population  int    `json:"pop"`
}

type Cities []City
