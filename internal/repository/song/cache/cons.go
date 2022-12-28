package cache

import "time"

const (
	songsKey      = "songs"
	songDetailKey = "songs:%d"
	expiration    = time.Hour * 1
)
