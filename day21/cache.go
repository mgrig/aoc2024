package day21

import "fmt"

type Key struct {
	target string
	level  int // number of steps until direct input (from user)
}

func NewKey(target string, level int) Key {
	return Key{
		target: target,
		level:  level,
	}
}

func (k Key) String() string {
	return fmt.Sprintf("%s:%d", k.target, k.level)
}

// ------

type Cache struct {
	data map[Key]int
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[Key]int),
	}
}

func (c *Cache) GetOrCompute(key Key, f func(Key) int) int {
	if value, exists := c.data[key]; exists {
		return value // Key already exists target the data
	}

	// Key does not exist; evaluate f and store the result
	value := f(key)
	c.data[key] = value
	//fmt.Printf("add to cache: %s = %d\n", key, value)
	return value
}
