package users_history

import "github.com/WenLopes/recent/domain"

//public function getByType(string $userId, string $type, int $quantity = 0): Collection

type UsersHistory interface {
	GetByUserId(userId string) domain.UsersHistory
}

type usersHistoryService struct {
	repo Repository
}

func NewUsersHistoryService(repo Repository) *usersHistoryService {
	service := usersHistoryService{
		repo: repo,
	}
	return &service
}

func (service *usersHistoryService) GetByUserId(userId string) domain.UsersHistory {
	data := domain.UsersHistory{}
	return data
}
