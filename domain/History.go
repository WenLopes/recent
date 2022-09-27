package domain

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
