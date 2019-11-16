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

func (c *Cache) Delete(key string) {
	delete(c.cacheMap, key)
}

func (c *Cache) GetInt(key string) (i int, ok bool) {
	value := c.Get(key)
	switch i := value.(type) {
	case int:
		return i, true
	case int32:
		return int(i), true
	default:
		return 0, false
	}
}

func (c *Cache) GetInt64(key string) (i int64, ok bool) {
	value := c.Get(key)
	switch i := value.(type) {
	case int64:
		return i, true
	case int:
		return int64(i), true
	case int32:
		return int64(i), true
	default:
		return 0, false
	}
}

func (c *Cache) GetFloat(key string) (f float64, ok bool) {
	value := c.Get(key)
	switch f := value.(type) {
	case float64:
		return f, true
	case float32:
		return float64(f), true
	case int:
		return float64(f), true
	case int64:
		return float64(f), true
	case int32:
		return float64(f), true
	default:
		return 0, false
	}
}

func (c *Cache) GetFloat64(key string) (f float64, ok bool) {
	return c.GetFloat(key)
}

func (c *Cache) GetString(key string) (s string, ok bool) {
	value := c.Get(key)
	switch s := value.(type) {
	case string:
		return s, true
	default:
		return "", false
	}
}

type cacheData struct {
	expire *time.Time
	Value  interface{}
}
