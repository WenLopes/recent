package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type UsersHistory struct {
	UserId        string
	HasManualData bool
	CreatedAt     primitive.DateTime
	UpdatedAt     primitive.DateTime
	Histories     []History
}

type History struct {
	UserId            string
	Id                int
	HistoryType       string
	KeyAddressingType string
	ReceiverData      receiverData
}

type receiverData struct {
	key    string
	bank   string
	ispb   string
	name   string
	userId string
}
