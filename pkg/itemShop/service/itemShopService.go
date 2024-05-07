package service

import _models "github.com/MarkTBSS/096_Coin_Adding/pkg/itemShop/models"

type ItemShopService interface {
	Listing(itemFilter *_models.ItemFilter) (*_models.ItemResult, error)
}
