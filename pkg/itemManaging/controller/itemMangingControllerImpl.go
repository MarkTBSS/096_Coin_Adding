package controller

import (
	"net/http"
	"strconv"

	"github.com/MarkTBSS/096_Coin_Adding/pkg/custom"
	_itemManagingService "github.com/MarkTBSS/096_Coin_Adding/pkg/itemManaging/service"
	"github.com/labstack/echo/v4"

	_itemManagingModel "github.com/MarkTBSS/096_Coin_Adding/pkg/itemManaging/models"
)

type itemManagingImpl struct {
	itemMangingService _itemManagingService.ItemManagingService
}

func NewItemManagingControllerImpl(itemMangingService _itemManagingService.ItemManagingService) ItemManagingController {
	return &itemManagingImpl{itemMangingService: itemMangingService}
}

func (c *itemManagingImpl) Creating(pctx echo.Context) error {
	itemCreatingReq := new(_itemManagingModel.ItemCreatingReq)
	validatingContext := custom.NewCustomEchoRequest(pctx)
	if err := validatingContext.Bind(itemCreatingReq); err != nil {
		return custom.CustomError(pctx, http.StatusBadRequest, err)
	}
	item, err := c.itemMangingService.Creating(itemCreatingReq)
	if err != nil {
		return custom.CustomError(pctx, http.StatusInternalServerError, err)
	}
	return pctx.JSON(http.StatusCreated, item)
}

func (c *itemManagingImpl) Editing(pctx echo.Context) error {
	itemID, err := c.getItemID(pctx)
	if err != nil {
		return custom.CustomError(pctx, http.StatusBadRequest, err)
	}
	editItemReq := new(_itemManagingModel.ItemEditingReq)
	validatingContext := custom.NewCustomEchoRequest(pctx)
	err = validatingContext.Bind(editItemReq)
	if err != nil {
		return custom.CustomError(pctx, http.StatusBadRequest, err)
	}
	item, err := c.itemMangingService.Editing(itemID, editItemReq)
	if err != nil {
		return custom.CustomError(pctx, http.StatusInternalServerError, err)
	}
	return pctx.JSON(http.StatusOK, item)
}

func (c *itemManagingImpl) getItemID(pctx echo.Context) (uint64, error) {
	itemID := pctx.Param("itemID")
	itemIDUint64, err := strconv.ParseUint(itemID, 10, 64)
	if err != nil {
		return 0, err
	}
	return itemIDUint64, nil
}

func (c *itemManagingImpl) Archiving(pctx echo.Context) error {
	itemID, err := c.getItemID(pctx)
	if err != nil {
		return custom.CustomError(pctx, http.StatusBadRequest, err)
	}
	err = c.itemMangingService.Archiving(itemID)
	if err != nil {
		return custom.CustomError(pctx, http.StatusInternalServerError, err)
	}
	return pctx.NoContent(http.StatusNoContent)
}
