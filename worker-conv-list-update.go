package main

type ConvListUpdateWorker interface {
	refreshLists() // maybe export???
	MarkAsStale(UID string)
}
