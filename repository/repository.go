package repository

import "jamiuafolabi/web-service-gin/models"

type DatabaseRepo interface {
	GetAllRows() ([]models.ProductInformation, error)
}
