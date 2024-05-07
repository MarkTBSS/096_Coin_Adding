package repository

import (
	"github.com/MarkTBSS/096_Coin_Adding/databases"
	"github.com/MarkTBSS/096_Coin_Adding/entities"
	_playerCoin "github.com/MarkTBSS/096_Coin_Adding/pkg/playerCoin/exeption"
	"github.com/labstack/echo/v4"
)

type playerCoinRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewPlayerCoinRepositoryImpl(db databases.Database, logger echo.Logger) PlayerCoinRepository {
	return &playerCoinRepositoryImpl{
		db:     db,
		logger: logger,
	}
}

func (r *playerCoinRepositoryImpl) CoinAdding(playerCoinEntity *entities.PlayerCoin) (*entities.PlayerCoin, error) {
	playerCoin := new(entities.PlayerCoin)
	err := r.db.Connect().Create(playerCoinEntity).Scan(playerCoin).Error
	if err != nil {
		r.logger.Error("Player's balance recording failed:", err.Error())
		return nil, &_playerCoin.CoinAdding{}
	}
	return playerCoin, nil
}
