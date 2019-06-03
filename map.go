package stringset

// Map is the lower-level underlying type for storing a string set
// Note: This is not thread-safe. If you need thread-safety, please use StringSet type
type Map map[string]struct{}

// Set will place a key
func (m Map) Set(key string) {
	m[key] = struct{}{}
}

// Unset will remove a key
func (m Map) Unset(key string) {
	delete(m, key)
}

// Has will return whether or not a key exists
func (m Map) Has(key string) (has bool) {
	_, has = m[key]
	return
}

// Slice will return the keys as a slice
func (m Map) Slice() (keys []string) {
	for key := range m {
		keys = append(keys, key)
	}

	return
}
