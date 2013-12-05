package main

import (
	"flag"
	"fmt"
	"labix.org/v2/mgo"
	"os"
	"os/signal"
	"syscall"
)

var mgoSession *mgo.Session

func init() {
	flag.BoolVar(&fDev, "dev", false, "Dev environment; set this if not in production")
	flag.StringVar(&fLaddr, "laddr", "localhost:8765", "Address to listen on")
}

func init_DB() {
	var err error
	mgoSession, err = mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	hookCleanUp()
}

func hookCleanUp() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			if sig == syscall.SIGINT || sig == syscall.SIGTERM {
				fmt.Println("Closing mangodb session before exiting...")
				mgoSession.Close()
				os.Exit(0)
			}
		}
	}()
}

func getDB(session *mgo.Session) *mgo.Database {
	var db *mgo.Database
	if fDev {
		db = session.DB("dev")
	} else {
		db = session.DB("prod")
	}
	return db
}

func newSignature(signature *Signature) {
	session := mgoSession.Clone()
	defer session.Close()
	db := getDB(session)
	db.C("Signatures").Insert(signature)
}

func getCount() (int, error) {
	session := mgoSession.Clone()
	defer session.Close()
	db := getDB(session)
	c, err := db.C("Signatures").Count()
	return c, err
}

func getNames() []string {
	session := mgoSession.Clone()
	defer session.Close()
	db := getDB(session)
	it := db.C("Signatures").Find(nil).Iter()
	var sig Signature
	ret := make([]string, 0)
	for it.Next(&sig) {
		ret = append(ret, sig.Name)
	}
	return ret
}
