package database

import (
	"mini-redis-go/interface/resp"
	"mini-redis-go/resp/reply"
)

func Ping(db *DB, args [][]byte) resp.Reply {
	return reply.MakePongReply()
}

// PING
func init() {
	RegisterCommand("ping", Ping, 1)
}
