package rdshand

import (
	"fmt"
	"time"
)

// dftDuration
const dftDuration = 1 * time.Minute

// RdKey get redis key by format and id slice
func RdKey(format string, id ...interface{}) string {
	return fmt.Sprintf(format, id...)
}

// ParseKeyDuration get redis key time from string
func ParseKeyDuration(s string) (time.Duration, error) {
	dur, err := time.ParseDuration(s)
	if err != nil {
		return 0, err
	}
	return dur, nil
}

// KeyFunc get redis key by struct and its field value
// 	eg. v := i.(Type), RdKey(format1, v.id1) or RdKey(format1, v.id1, v.id2)
type KeyFunc func(i interface{}) string

// Uint64ToStringRdKeys []uint64 to []string
func Uint64ToStringRdKeys(format string, ids []uint64) []string {
	var keys []string
	for _, id := range ids {
		keys = append(keys, fmt.Sprintf(format, id))
	}
	return keys
}
