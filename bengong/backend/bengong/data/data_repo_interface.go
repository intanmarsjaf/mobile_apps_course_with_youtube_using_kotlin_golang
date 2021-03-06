package data

import "bengong/model"

type DataRepo interface {
	GetDataByUser(userID int) (*[]model.Data, error)
	Insert(data *model.Data) (*model.Data, error)
}
