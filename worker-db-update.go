package main

type DBUpdateWorker interface {
	sendUpdates() // maybe want an exported version?
}
