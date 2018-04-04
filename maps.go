package gostdlib

// Check that the key 'v' exists in the map 'm'
func In(v interface{}, m map[interface{}]interface{}) bool {
	_, ok := m[v]
	return ok
}
