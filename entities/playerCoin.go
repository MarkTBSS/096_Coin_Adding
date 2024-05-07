package entities

import (
	"time"

	_playerCoinModel "github.com/MarkTBSS/096_Coin_Adding/pkg/playerCoin/model"
)

type (
	PlayerCoin struct {
		ID        uint64    `gorm:"primaryKey;autoIncrement;"`
		PlayerID  string    `gorm:"type:varchar(64);not null;"`
		Amount    int64     `gorm:"not null;"`
		CreatedAt time.Time `gorm:"not null;autoCreateTime;"`
	}
)

func (p *PlayerCoin) ChangeToPlayerCoinModel() *_playerCoinModel.PlayerCoin {
	return &_playerCoinModel.PlayerCoin{
		ID:        p.ID,
		PlayerID:  p.PlayerID,
		Amount:    p.Amount,
		CreatedAt: p.CreatedAt,
	}
}
