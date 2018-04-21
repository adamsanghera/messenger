package main

// CID -> Conversation
var convLedger map[string]*Conversation

// UID -> User
var userLedger map[string]*User

var updateWorker DBUpdateWorker
