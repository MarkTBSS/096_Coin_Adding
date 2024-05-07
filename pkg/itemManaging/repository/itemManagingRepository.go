package repository

import (
	"github.com/MarkTBSS/096_Coin_Adding/entities"
	"github.com/MarkTBSS/096_Coin_Adding/pkg/itemManaging/models"
)

type ItemManagingRepository interface {
	Creating(itemEntity *entities.Item) (*entities.Item, error)
	Editing(itemID uint64, itemEditingReq *models.ItemEditingReq) (uint64, error)
	Archiving(itemID uint64) error
}
