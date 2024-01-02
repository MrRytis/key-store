package storage

import (
	"errors"
	"time"
)

type Value struct {
	Key      string `json:"key"`
	Value    string `json:"value"`
	ExpireAt int64  `json:"expire_at"`
}

type Storage struct {
	Values map[string]Value
}

func NewStorage() *Storage {
	return &Storage{
		Values: make(map[string]Value),
	}
}

func (s *Storage) GetAllValues() *[]Value {
	values := make([]Value, 0)

	for _, v := range s.Values {
		if v.ExpireAt > time.Now().Unix() {
			values = append(values, v)
		}
	}

	return &values
}

func (s *Storage) GetValueByKey(key string) (*Value, error) {
	v, exists := s.Values[key]

	if !exists || v.ExpireAt < time.Now().Unix() {
		return nil, errors.New("value not found")
	}

	return &v, nil
}

func (s *Storage) StoreValue(key string, value string, expireAt int64) *Value {
	s.Values[key] = Value{
		Key:      key,
		Value:    value,
		ExpireAt: expireAt,
	}

	v := s.Values[key]

	return &v
}

func (s *Storage) UpdateValue(key string, value string, expireAt int64) *Value {
	s.Values[key] = Value{
		Key:      key,
		Value:    value,
		ExpireAt: expireAt,
	}

	v := s.Values[key]

	return &v
}

func (s *Storage) DeleteValue(key string) {
	delete(s.Values, key)
}
