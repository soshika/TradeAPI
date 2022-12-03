package app

import "tradeAPI/controllers/ping"

func urlPatterns() {
	router.GET("/ping", ping.Ping)
}
