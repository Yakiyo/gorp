package main

type Config struct {
	Id string
	State string
	Details string
	Timestamps TimeStamps
}

type TimeStamps struct {
	Start string
	End string
}