package in_memory

import (
	"apartment_search_service/internal/repositories"
	"errors"
	"sync"
)

type SubscriptionRepo struct {
	data map[int32][]string
	mu   sync.RWMutex
}

func NewSubscriptionRepo() repositories.SubscriptionRepository {
	return &SubscriptionRepo{
		data: make(map[int32][]string),
	}
}

func (r *SubscriptionRepo) AddSubscriber(houseId int32, email string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.data[houseId]; !exists {
		r.data[houseId] = []string{}
	}
	for _, subscriber := range r.data[houseId] {
		if subscriber == email {
			return errors.New("already subscribed")
		}
	}
	r.data[houseId] = append(r.data[houseId], email)
	return nil
}

func (r *SubscriptionRepo) GetSubscribers(houseId int32) ([]string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.data[houseId], nil
}
