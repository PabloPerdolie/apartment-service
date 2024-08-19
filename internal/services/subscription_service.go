package services

import (
	"apartment_search_service/internal/repositories"
	"apartment_search_service/internal/utils/sender"
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
)

type subscriptionService struct {
	houseRepo        repositories.HouseRepository
	subscriptionRepo repositories.SubscriptionRepository
	emailSender      *sender.Sender
	logger           *logrus.Logger
}

func NewSubscriptionService(h repositories.HouseRepository, s repositories.SubscriptionRepository, e *sender.Sender, l *logrus.Logger) SubscriptionService {
	return &subscriptionService{
		houseRepo:        h,
		subscriptionRepo: s,
		emailSender:      e,
		logger:           l,
	}
}

func (h *subscriptionService) AddSubscriber(houseId int32, email string) error {
	_, err := h.houseRepo.GetById(houseId)
	if err != nil {
		return err
	}
	return h.subscriptionRepo.AddSubscriber(houseId, email)
}

func (h *subscriptionService) NotifySubscribers(houseId int32, flatInfo string) {
	subscribers, err := h.subscriptionRepo.GetSubscribers(houseId)
	if err != nil {
		h.logger.Println("Error retrieving subscribers:", err)
		return
	}

	for _, subscriber := range subscribers {
		go func(email string) {
			message := fmt.Sprintf("New flat available: %s", flatInfo)
			err := h.emailSender.SendEmail(context.Background(), email, message)
			if err != nil {
				h.logger.Println("Failed to send email to", email, ":", err)
			}
		}(subscriber)
	}
}
