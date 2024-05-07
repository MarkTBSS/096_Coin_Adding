package repository

import "github.com/MarkTBSS/096_Coin_Adding/entities"

type PlayerCoinRepository interface {
	CoinAdding(playerCoinEntity *entities.PlayerCoin) (*entities.PlayerCoin, error)
}
