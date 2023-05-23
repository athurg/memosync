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
	flag.StringVar(&targets, "targets", "", "Target Memo URLs split by comma")
	flag.StringVar(&interval, "i", "10m", "Sync time interval")
}

func main() {
	flag.Parse()
	duration, err := time.ParseDuration(interval)
	if err != nil {
		log.Printf("fail to parse interval %s: %s", interval, err)
		return
	}

	if addr == "" || openid == "" || targets == "" {
		flag.Usage()
		return
	}

	// run once
	if duration == 0 {
		run()
		return
	}

	lastCheckTs = time.Now().Unix()
	for range time.Tick(duration) {
		run()
		lastCheckTs = time.Now().Unix()
	}
}

func run() {
	users, err := registerOrGetUsers(addr, openid, strings.Split(targets, ","))
	if err != nil {
		log.Printf("fail to registerOrGetUsers: %s", err)
		return
	}

	for targetUrl, user := range users {
		err := syncTargetToUser(targetUrl, user)
		if err != nil {
			log.Printf("fail to syncTargetToUser: %s", err)
		}
	}
}
