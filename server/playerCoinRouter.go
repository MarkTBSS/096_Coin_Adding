package server

import (
	_playerCoinController "github.com/MarkTBSS/096_Coin_Adding/pkg/playerCoin/controller"
	_playerCoinRepository "github.com/MarkTBSS/096_Coin_Adding/pkg/playerCoin/repository"
	_playerCoinService "github.com/MarkTBSS/096_Coin_Adding/pkg/playerCoin/service"
)

func (s *echoServer) initPlayerCoinRouter() {
	router := s.app.Group("/v1/player-coin")

	playerCoinRepository := _playerCoinRepository.NewPlayerCoinRepositoryImpl(s.db, s.app.Logger)
	playerCoinService := _playerCoinService.NewPlayerCoinServiceImpl(playerCoinRepository)
	playerCoinController := _playerCoinController.NewPlayerCoinControllerImpl(playerCoinService)

	router.POST("", playerCoinController.CoinAdding)
}
