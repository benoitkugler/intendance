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
type Ids []int64

func (ids Ids) AsSQL() pq.Int64Array {
	return pq.Int64Array(ids)
}

func (ids Ids) AsSet() Set {
	return NewSetFromSlice(ids)
}

func (ids Ids) AsMenuRecettes(idMenu int64) []MenuRecette {
	out := make([]MenuRecette, len(ids))
	for i, id := range ids {
		out[i] = MenuRecette{IdMenu: idMenu, IdRecette: id}
	}
	return out
}

func (ids Ids) AsRepasRecettes(idRepas int64) []RepasRecette {
	out := make([]RepasRecette, len(ids))
	for i, id := range ids {
		out[i] = RepasRecette{IdRepas: idRepas, IdRecette: id}
	}
	return out
}

// ----------------------------------------------------
const has = true

type Set map[int64]bool // on choisit bool pour l'interaction avec .js

func NewSet() Set {
	return map[int64]bool{}
}

func NewSetFromSlice(keys []int64) Set {
	out := make(Set, len(keys))
	for _, key := range keys {
		out[key] = has
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
	s[key] = has
}

// ------------------------------------------------------
type LienIngredients []LienIngredient

// AsRecetteIngredients lie les ingrédients au menu donné
func (ls LienIngredients) AsRecetteIngredients(idRecette int64) []RecetteIngredient {
	out := make([]RecetteIngredient, len(ls))
	for i, ing := range ls {
		out[i] = RecetteIngredient{IdRecette: idRecette, LienIngredient: ing}
	}
	return out
}

// AsMenuIngredients lie les ingrédients au menu donné
func (ls LienIngredients) AsMenuIngredients(idMenu int64) []MenuIngredient {
	out := make([]MenuIngredient, len(ls))
	for i, ing := range ls {
		out[i] = MenuIngredient{IdMenu: idMenu, LienIngredient: ing}
	}
	return out
}

// AsRepasIngredients lie les ingrédients au menu donné
func (ls LienIngredients) AsRepasIngredients(idRepas int64) []RepasIngredient {
	out := make([]RepasIngredient, len(ls))
	for i, ing := range ls {
		out[i] = RepasIngredient{IdRepas: idRepas, LienIngredient: ing}
	}
	return out
}
