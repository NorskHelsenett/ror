package kvcachehelper

type CacheInterface interface {
	Add(key string, value string)
	Get(key string) (string, bool)
	Remove(key string) bool
}
