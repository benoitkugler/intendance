package models

import (
	"time"

	"github.com/lib/pq"
)

func (s Sejour) DateFromOffset(jour int64) time.Time {
	start := s.DateDebut
	return start.Add(time.Duration(jour*24) * time.Hour)
}

// ----------------------------------------------------
type Ids struct {
	ids []int64
}

func (ids Ids) AsSQL() pq.Int64Array {
	return pq.Int64Array(ids.ids)
}

func (ids Ids) AsSet() Set {
	return NewSetFromSlice(ids.ids)
}

// ----------------------------------------------------
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
