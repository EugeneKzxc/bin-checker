package scope

import "backend/server/internal/models"

type Storage struct {
	Mapping [50000000]string
}

func NewStorage(data []models.BinBank) *Storage {
	s := Storage{}
	for i, _ := range data {
		s.Mapping[data[i].Bin] = data[i].Bank
	}
	return &s
}

func (s *Storage) Add(bin int, bank string) {
	s.Mapping[bin] = bank
}
