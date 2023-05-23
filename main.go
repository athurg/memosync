package main

import (
	"flag"
	"log"
	"strings"
	"time"
)

// All flags defined here
var addr, openid, targets, interval string

func init() {
	flag.StringVar(&addr, "h", "https://usememos.com", "URL of YOUR Memos")
	flag.StringVar(&openid, "k", "", "OpenID of YOUR Memos ADMIN user")
	flag.StringVar(&targets, "targets", "", "Register targets then exit")
	flag.StringVar(&interval, "i", "10m", "Sync time interval")
}

func main() {
	flag.Parse()
	duration, err := time.ParseDuration(interval)
	if err != nil {
		log.Printf("fail to parse interval %s: %s", interval, err)
		return
	}

	if addr == "" || openid == "" {
		flag.Usage()
		return
	}

	// If `targets` is not empty, register targets and exit
	if targets != "" {
		err := registerTargets(addr, openid, strings.Split(targets, ","))
		if err != nil {
			log.Fatalf("fail to register targets: %s", err)
		}
		log.Println("Done")
		return
	}

	// always run once
	run()

	// run once and exit
	if duration == 0 {
		return
	}

	lastCheckTs = time.Now().Unix()
	for range time.Tick(duration) {
		run()
		lastCheckTs = time.Now().Unix()
	}
}

func run() {
	users, err := resetOpenIdAndFetchTargetUsers(addr, openid)
	if err != nil {
		log.Printf("fail to resetOpenIdAndGetUsers: %s", err)
		return
	}

	for _, user := range users {
		err := syncTargetToUser(user)
		if err != nil {
			log.Printf("fail to syncTargetToUser: %s", err)
		}
	}
}
