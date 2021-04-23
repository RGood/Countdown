package helpers

import "sync"

type StringSet struct {
	data  map[string]struct{}
	mutex sync.Mutex
}

func NewStringSet() *StringSet {
	return &StringSet{
		data:  map[string]struct{}{},
		mutex: sync.Mutex{},
	}
}

func (s *StringSet) Add(val string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.data[val] = struct{}{}
}

func (s *StringSet) Contains(val string) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	_, ok := s.data[val]
	return ok
}

func (s *StringSet) Delete(val string) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	_, ok := s.data[val]
	delete(s.data, val)
	return ok
}

func (s *StringSet) Values() []string {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	values := []string{}
	for k, _ := range s.data {
		values = append(values, k)
	}

	return values
}

func (s *StringSet) Size() int {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return len(s.data)
}
