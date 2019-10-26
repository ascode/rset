package collection

import "sync"

type RSet struct {
	sync.RWMutex
	Set []map[string]interface{}
}

