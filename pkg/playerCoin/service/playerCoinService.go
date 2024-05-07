package service

import _playerCoinModel "github.com/MarkTBSS/096_Coin_Adding/pkg/playerCoin/model"

type PlayerCoinService interface {
	CoinAdding(coinAddingReq *_playerCoinModel.CoinAddingReq) (*_playerCoinModel.PlayerCoin, error)
}
