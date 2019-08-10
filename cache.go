package cache

import (
	"time"
)

func New() *Cache {
	return &Cache{cacheMap: map[string]*cacheData{}}
}

type Cache struct {
	cacheMap map[string]*cacheData
}

func (c *Cache) Set(key string, data interface{}, sec ...int) {
	var expire *time.Time
	if len(sec) == 1 {
		t := time.Now()
		oneSecond, _ := time.ParseDuration("1s")
		t = t.Add(time.Duration(sec[0]) * oneSecond)
		expire = &t
	}

	c.cacheMap[key] = &cacheData{
		expire,
		data,
	}
}

func (c *Cache) Get(key string) interface{} {
	data := c.cacheMap[key]

	// key为空
	if data == nil {
		return nil
	}

	// 没有过期时间 或 没有超时，正常返回数据
	if data.expire == nil || time.Now().Before(*data.expire) {
		return data.Value
	}

	// 超时删除数据
	delete(c.cacheMap, key)
	return nil
}

func (c *Cache) GetInt(key string) (i int, ok bool) {
	value := c.Get(key)
	if i, ok := value.(int); ok {
		return i, true
	} else {
		return 0, false
	}
}

func (c *Cache) GetFloat(key string) (f float64, ok bool) {
	value := c.Get(key)
	if f, ok := value.(float64); ok {
		return f, true
	} else if f, ok := value.(int); ok {
		return float64(f), true
	} else {
		return 0, false
	}
}

type cacheData struct {
	expire *time.Time
	Value  interface{}
}
