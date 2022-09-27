package history

import "github.com/WenLopes/recent/domain"

type Repository interface {
	GetAll() []domain.History
}
