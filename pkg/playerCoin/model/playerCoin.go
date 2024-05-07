package model

import "time"

type PlayerCoin struct {
	ID        uint64    `json:"id"`
	PlayerID  string    `json:"playerID"`
	Amount    int64     `json:"amount"`
	CreatedAt time.Time `json:"createdAt"`
}

type CoinAddingReq struct {
	PlayerID string `json:"player_id"`
	Amount   int64  `json:"amount" validate:"required,gt=0"`
}
