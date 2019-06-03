package stringset

import "sync"

// New will create a new instance of StringSet
func New() *StringSet {
	var s StringSet
	s.m = make(Map)
	return &s
}

// StringSet is a thread-safe wrapper for the stringset.Map
type StringSet struct {
	mux sync.RWMutex
	m   Map
}

// Set will place a key
func (s *StringSet) Set(key string) {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.m.Set(key)
}

// Unset will remove a key
func (s *StringSet) Unset(key string) {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.m.Unset(key)
}

// Has will return whether or not a key exists
func (s *StringSet) Has(key string) (has bool) {
	s.mux.RLock()
	defer s.mux.RUnlock()
	return s.m.Has(key)
}

// Slice will return the keys as a slice
func (s *StringSet) Slice() (keys []string) {
	s.mux.RLock()
	defer s.mux.RUnlock()
	return s.m.Slice()
}
