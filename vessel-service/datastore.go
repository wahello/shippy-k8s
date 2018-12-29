package main

import (
	mgo "gopkg.in/mgo.v2"
)

func CreateSession(shard1 string, shard2 string) (*mgo.Session, error) {
	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{shard1, shard2},
		Database: "shippy",
	})
	if err != nil {
		return nil, err
	}
	session.SetMode(mgo.Monotonic, true)
	return session, nil
}
