package service

import (
	_entities "github.com/MarkTBSS/096_Coin_Adding/entities"
	_models "github.com/MarkTBSS/096_Coin_Adding/pkg/itemShop/models"
	"github.com/MarkTBSS/096_Coin_Adding/pkg/itemShop/repository"
)

type itemShopServiceImpl struct {
	itemShopRepository repository.ItemShopRepository
}

func (s *itemShopServiceImpl) Listing(itemFilter *_models.ItemFilter) (*_models.ItemResult, error) {
	itemEntityList, err := s.itemShopRepository.Listing(itemFilter)
	if err != nil {
		return nil, err
	}

	totalItems, err := s.itemShopRepository.Counting(itemFilter)
	if err != nil {
		return nil, err
	}

	size := itemFilter.Paginate.Size
	page := itemFilter.Paginate.Page
	totalPage := s.totalPageCalculation(totalItems, size)
	result := s.toItemResultsResponse(itemEntityList, page, totalPage)

	return result, nil
}

func (s *itemShopServiceImpl) totalPageCalculation(totalItems, size int64) int64 {
	totalPage := totalItems / size
	if totalItems%size != 0 {
		totalPage++
	}
	return totalPage
}

func (s *itemShopServiceImpl) toItemResultsResponse(itemEntityList []*_entities.Item, page, totalPage int64) *_models.ItemResult {
	items := make([]*_models.Item, 0)
	for _, itemEntity := range itemEntityList {
		items = append(items, itemEntity.ChangeToItemModel())
	}
	return &_models.ItemResult{
		Items: items,
		Paginate: _models.PaginateResult{
			Page:      page,
			TotalPage: totalPage,
		},
	}
}

func NewItemShopServiceImpl(itemShopRepository repository.ItemShopRepository) ItemShopService {
	return &itemShopServiceImpl{itemShopRepository}
}
