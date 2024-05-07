package service

import (
	"github.com/MarkTBSS/096_Coin_Adding/entities"
	_playerCoinModel "github.com/MarkTBSS/096_Coin_Adding/pkg/playerCoin/model"
	_playerCoinRepository "github.com/MarkTBSS/096_Coin_Adding/pkg/playerCoin/repository"
)

type playerCoinServiceImpl struct {
	playerCoinRepository _playerCoinRepository.PlayerCoinRepository
}

func NewPlayerCoinServiceImpl(playerCoinRepository _playerCoinRepository.PlayerCoinRepository) PlayerCoinService {
	return &playerCoinServiceImpl{playerCoinRepository}
}

func (s *playerCoinServiceImpl) CoinAdding(coinAddingReq *_playerCoinModel.CoinAddingReq) (*_playerCoinModel.PlayerCoin, error) {
	playerCoinEntity := &entities.PlayerCoin{
		PlayerID: coinAddingReq.PlayerID,
		Amount:   coinAddingReq.Amount,
	}
	playerCoin, err := s.playerCoinRepository.CoinAdding(playerCoinEntity)
	if err != nil {
		return nil, err
	}
	return playerCoin.ChangeToPlayerCoinModel(), nil
}
