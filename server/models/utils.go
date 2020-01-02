package models

import "time"

func (s Sejour) DateFromOffset(jour int64) time.Time {
	start := s.DateDebut
	return start.Add(time.Duration(jour*24) * time.Hour)
}
