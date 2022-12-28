package cache

import "time"

const (
	artistsKey      = "artists"
	artistDetailKey = "artists:%d"
	expiration      = time.Hour * 1
)
