package main

import (
	"log"
	"os"
)

const (
	engineMongo = 101
	engineRedis = 102
)

//Storage an abstract storage interface to make it easy
// to work with several storage systems
type Storage interface {
	Save(key string, val string) bool
	SaveWithTTL(key string, val string, ttl int) bool
	Read(key string) string
	Del(key string) int
	//Exists(key string) bool
}

//Model a set of funcs which makes it easy to work with data writing and reading
type Model struct {
	redis   Storage
	mongo   Storage
	current int
	ttl     int
}

//TTL sets the ttl for all save operations of the model.
//If you want just to use it once, then set it back to 0
//after you usage.
//For ease of use, you can right after calling  this method,
//call another method of the Model. Method
func (m *Model) TTL(ttl int) *Model {
	m.ttl = ttl
	return m
}

//Save saves the data to current active storage
//It saves it with the given key. If the storage does not require
//a key, then it gets ignored. If two args are passed, the first arg is
// treated as key and the second as value, if 1 is passed, it is treated
// as value.
func (m *Model) Save(args ...string) bool {
	var key string
	var val string
	if len(args) == 1 {
		key = ""
		val = args[0]
	} else if len(args) == 2 {
		key = args[0]
		val = args[1]
	}
	if m.ttl != 0 {
		return m.getEngine().SaveWithTTL(key, val, m.ttl)
	}
	return m.getEngine().Save(key, val)
}

func (m *Model) getEngine() Storage {
	switch m.current {
	case engineMongo:
		return m.mongo
		break
	case engineRedis:
		return m.redis
		break
	}
	log.Panic("Cannot find a proper storage engine. Process abroted.")
	os.Exit(1)
	return nil
}

func (m *Model) createEngineObject() {
	switch m.current {
	// case engineMongo:
	// 	// @todo for Mongo adapter
	// 	m.mongo = Storage{}
	// 	return
	case engineRedis:
		m.redis = NewRedis()
		return
	}
	log.Panic("Cannot find a proper storage engine. Process abroted.")
	os.Exit(1)
}

func (m *Model) SetEngine(engineCode int) *Model {
	m.current = engineCode
	m.createEngineObject()
	return m
}
