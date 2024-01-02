package service

import (
	"github.com/MrRytis/key-store/internal/storage"
	"time"
)

type ExpirationChecker struct {
	Storage *storage.Storage
}

func NewExpirationChecker(storage *storage.Storage) *ExpirationChecker {
	return &ExpirationChecker{
		Storage: storage,
	}
}

func (e *ExpirationChecker) Start() {
	for {
		for _, v := range e.Storage.Values {
			currentTime := time.Now()

			if v.ExpireAt > 0 && v.ExpireAt < currentTime.Unix() {
				e.Storage.DeleteValue(v.Key)
			}
		}
	}
}
