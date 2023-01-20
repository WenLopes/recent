package users_history

import "github.com/WenLopes/recent/domain"

type Repository interface {
	GetAllHistories() []domain.History
}
