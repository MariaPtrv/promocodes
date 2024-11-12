package repository

import (
	t "admin/pkg"
	"encoding/json"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

type RewardItemPostgres struct {
	db *sqlx.DB
}

func NewRewardPostgres(db *sqlx.DB) *RewardItemPostgres {
	return &RewardItemPostgres{db: db}
}

func (r *RewardItemPostgres) CreateReward(reward t.Reward) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var itemId int = 0

	rj, _ := json.Marshal(reward)
	log.Printf("repository-reward: CreateReward \nreward: %s\n", string(rj))

	query := fmt.Sprintf("INSERT INTO %s (title, description) values ($1, $2) RETURNING id", rewardTable)
	row := tx.QueryRow(query, reward.Title, reward.Desc)
	err = row.Scan(&itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return itemId, tx.Commit()
}

func (r *RewardItemPostgres) DeleteReward(reward t.Reward) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	rj, _ := json.Marshal(reward)
	log.Printf("repository-reward: DeleteReward \nreward: %s\n", string(rj))

	query := fmt.Sprintf("DELETE FROM %s WHERE title = $1", rewardTable)
	_, err = tx.Exec(query, reward.Title)

	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}