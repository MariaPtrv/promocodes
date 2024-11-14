package repository

import (
	"fmt"
	"log"
	t "promocodes"

	"github.com/jmoiron/sqlx"
)

type RewardPostgres struct {
	db *sqlx.DB
}

func NewRewardPostgres(db *sqlx.DB) *RewardPostgres {
	return &RewardPostgres{db: db}
}

func (r *RewardPostgres) GetRewardById(reward t.Reward) (t.Reward, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return t.Reward{}, err
	}

	log.Printf("repository-reward: GetRewardById reward id: %d\n", reward.Id)

	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", rewardTable)

	var rdb t.Reward

	err = r.db.Get(&rdb, query, reward.Id)
	if err != nil {
		tx.Rollback()
		return t.Reward{}, err
	}

	return rdb, tx.Commit()
}
