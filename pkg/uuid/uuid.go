package uuid

import "github.com/disgoorg/snowflake/v2"

func NewUUID(key string) uint64 {
	uuid := snowflake.GetEnv(key)
	return uint64(uuid)
}
