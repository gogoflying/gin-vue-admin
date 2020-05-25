package servers

import (
	"sync"
)

var g_tmpMap *tmpMap

type tmpMap struct {
	mapTmp map[string]interface{}
	lock   sync.RWMutex
}

func newTmpMap() *tmpMap {
	return &tmpMap{
		mapTmp: make(map[string]interface{}, 0),
	}
}

func AddMap(k string, v interface{}) {
	if g_tmpMap == nil {
		g_tmpMap = newTmpMap()
	}
	g_tmpMap.lock.Lock()
	defer g_tmpMap.lock.Unlock()
	g_tmpMap.mapTmp[k] = v
}

func GetMap(k string) interface{} {
	if g_tmpMap == nil {
		return nil
	}
	g_tmpMap.lock.RLock()
	defer g_tmpMap.lock.RUnlock()
	v, ok := g_tmpMap.mapTmp[k]
	if !ok {
		return nil
	}
	return v
}

func DelMap(k string) {
	if g_tmpMap == nil {
		return
	}
	g_tmpMap.lock.Lock()
	defer g_tmpMap.lock.Unlock()
	delete(g_tmpMap.mapTmp, k)
}
