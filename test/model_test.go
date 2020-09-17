package test

import (
	"jupiter/application/utils"
	"reflect"
	"testing"
	"time"
)

func init() {

}

func TestCache(t *testing.T) {
	utils.MemCache.Set(`name`, `yangsen`, time.Second)
	val, ok := utils.MemCache.Get(`name`)
	t.Log(val, ok)
	time.Sleep(2 * time.Second)
	val2, ok2 := utils.MemCache.Get(`name`)
	t.Log(val2, ok2)
}

func TestInt(t *testing.T) {
	var m int
	m = 6553500
	t.Log(m)
	t.Log(reflect.TypeOf(m))
}

func TestUuid(t *testing.T) {
	uuid := utils.GetUUID()
	t.Log(uuid)
	ip, err := utils.GetLocalIp()
	t.Log(ip, err)
}
