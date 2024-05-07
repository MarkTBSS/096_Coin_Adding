package server

import (
	"github.com/MarkTBSS/096_Coin_Adding/pkg/itemManaging/controller"
	_itemManagingRepository "github.com/MarkTBSS/096_Coin_Adding/pkg/itemManaging/repository"
	"github.com/MarkTBSS/096_Coin_Adding/pkg/itemManaging/service"
	_itemShopRepository "github.com/MarkTBSS/096_Coin_Adding/pkg/itemShop/repository"
)

func (s *echoServer) initItemManagingRouter() {
	router := s.app.Group("/v1/item-managing")

	itemRepository := _itemShopRepository.NewItemShopRepositoryImpl(s.db, s.app.Logger)
	itemMangingRepository := _itemManagingRepository.NewItemManagingRepositoryImpl(s.db, s.app.Logger)
	itemManagingService := service.NewItemManagingServiceImpl(itemMangingRepository, itemRepository)
	itemManaging := controller.NewItemManagingControllerImpl(itemManagingService)

	router.POST("", itemManaging.Creating)
	router.PATCH("/:itemID", itemManaging.Editing)
	router.DELETE("/:itemID", itemManaging.Archiving)
}
