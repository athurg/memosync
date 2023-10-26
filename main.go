package main

import (
	"flag"
	"log"
	"time"
)

// All flags defined here
var addr, password, interval string

func init() {
	flag.StringVar(&addr, "h", "https://usememos.com", "URL of YOUR Memos")
	flag.StringVar(&password, "p", "secret", "Share password of external users")
	flag.StringVar(&interval, "i", "10m", "Sync time interval")
}

func main() {
	flag.Parse()
	duration, err := time.ParseDuration(interval)
	if err != nil {
		log.Printf("fail to parse interval %s: %s", interval, err)
		return
	}

	if addr == "" || password == "" {
		flag.Usage()
		return
	}

	// Force syncMemos once
	if duration == 0 {
		syncMemos()
		return
	}

	lastCheckTs = time.Now().Unix()
	syncMemos()
	for range time.Tick(duration) {
		syncMemos()
		lastCheckTs = time.Now().Unix()
	}
}
