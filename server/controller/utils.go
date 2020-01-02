package controller

type Set map[int64]struct{}

func NewSet() Set {
	return map[int64]struct{}{}
}

func NewSetFromSlice(keys []int64) Set {
	out := make(Set, len(keys))
	for _, key := range keys {
		out[key] = struct{}{}
	}
	return out
}

func (s Set) Keys() []int64 {
	out := make([]int64, 0, len(s))
	for k := range s {
		out = append(out, k)
	}
	return out
}

func (s Set) Has(key int64) bool {
	_, has := s[key]
	return has
}

func (s Set) Add(key int64) {
	s[key] = struct{}{}
}
