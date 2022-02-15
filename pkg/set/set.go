package set

import "sync"

type StringSet struct {
	flags map[string]struct{}
	mux   sync.Mutex
}

func NewStringSet() *StringSet {
	return &StringSet{
		flags: make(map[string]struct{}, 16),
		mux:   sync.Mutex{},
	}
}

func (s *StringSet) Add(element string) bool {
	s.mux.Lock()
	defer s.mux.Unlock()

	if _, ok := s.flags[element]; ok {
		return false
	}

	s.flags[element] = struct{}{}
	return true
}

func (s *StringSet) Remove(key string) bool {
	s.mux.Lock()
	defer s.mux.Unlock()

	_, result := s.flags[key]
	delete(s.flags, key)
	return result
}

func (s *StringSet) Elements() []string {
	var result = make([]string, 0, len(s.flags))
	for element, _ := range s.flags {
		result = append(result, element)
	}
	return result
}
