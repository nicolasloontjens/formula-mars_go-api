package models

import (
	"errors"
)

type Bet struct {
	ID       uint    `gorm:"primaryKey" json:"id"`
	UserID   uint    `json:"-"`
	RaceID   int     `json:"race_id"`
	DriverId int     `json:"driver_id"`
	Amount   float64 `gorm:"type:decimal(10,2)" json:"amount"`
}

func GetBets(user_id uint) ([]Bet, error) {
	var bets []Bet

	if err := DB.Where("user_id = ?", user_id).Find(&bets).Error; err != nil {
		return bets, errors.New("bets not found")
	}

	return bets, nil
}

func (bet *Bet) AddBet() error {

	if err := DB.Create(&bet).Error; err != nil {
		return errors.New("bet not added")
	}

	return nil
}
