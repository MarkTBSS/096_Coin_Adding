package controller

import (
	"net/http"

	"github.com/MarkTBSS/096_Coin_Adding/pkg/custom"

	_playerCoinModel "github.com/MarkTBSS/096_Coin_Adding/pkg/playerCoin/model"
	_playerCoinService "github.com/MarkTBSS/096_Coin_Adding/pkg/playerCoin/service"
	"github.com/labstack/echo/v4"
)

type playerCoinControllerImpl struct {
	playerCoinService _playerCoinService.PlayerCoinService
}

func NewPlayerCoinControllerImpl(playerCoinService _playerCoinService.PlayerCoinService) PlayerCoinController {
	return &playerCoinControllerImpl{
		playerCoinService: playerCoinService,
	}
}

func (c *playerCoinControllerImpl) CoinAdding(pctx echo.Context) error {
	//playerID, err := validation.PlayerIDGetting(pctx)
	playerID := "1234"
	/* if err != nil {
		return custom.CustomError(pctx, http.StatusBadRequest, err)
	} */
	coinAddingReq := new(_playerCoinModel.CoinAddingReq)
	validatingContext := custom.NewCustomEchoRequest(pctx)
	err := validatingContext.Bind(coinAddingReq)
	if err != nil {
		return custom.CustomError(pctx, http.StatusBadRequest, err)
	}
	coinAddingReq.PlayerID = playerID
	playerCoin, err := c.playerCoinService.CoinAdding(coinAddingReq)
	if err != nil {
		return custom.CustomError(pctx, http.StatusInternalServerError, err)
	}
	return pctx.JSON(http.StatusCreated, playerCoin)
}
