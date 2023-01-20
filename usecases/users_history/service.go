package users_history

import "github.com/WenLopes/recent/domain"

type Service interface {
	GetByUserId(userId string) domain.UsersHistory
	GetAll() []domain.History
}

type usersHistoryService struct {
	repo Repository
}

func NewService(repo Repository) *usersHistoryService {
	service := usersHistoryService{
		repo: repo,
	}
	return &service
}

func (service *usersHistoryService) GetByUserId(userId string) domain.UsersHistory {
	data := domain.UsersHistory{}
	return data
}

func (service *usersHistoryService) GetAll() []domain.History {
	data := service.repo.GetAllHistories()
	return data
}
