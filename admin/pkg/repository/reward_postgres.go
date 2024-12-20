package repository

import (
	t "admin/pkg"
	"encoding/json"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

type RewardPostgres struct {
	db *sqlx.DB
}

func NewRewardPostgres(db *sqlx.DB) *RewardPostgres {
	return &RewardPostgres{db: db}
}

func (r *RewardPostgres) CreateReward(reward t.Reward) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var itemId int = 0

	rj, _ := json.Marshal(reward)
	log.Printf("repository-reward: CreateReward reward: %s\n", string(rj))

	query := fmt.Sprintf("INSERT INTO %s (title, description) values ($1, $2) RETURNING id", rewardTable)
	row := tx.QueryRow(query, reward.Title, reward.Description)
	err = row.Scan(&itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return itemId, tx.Commit()
}

func (r *RewardPostgres) DeleteReward(reward t.Reward) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	rj, _ := json.Marshal(reward)
	log.Printf("repository-reward: DeleteReward reward: %s\n", string(rj))

	query := fmt.Sprintf("DELETE FROM %s WHERE title = $1", rewardTable)
	_, err = tx.Exec(query, reward.Title)

	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
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

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return rdb, err
	}

	return rdb, nil
}

func (r *RewardPostgres) GetRewards() ([]t.Reward, error) {
	log.Printf("repository-rewards: GetRewards\n")

	query := fmt.Sprintf("SELECT * FROM %s", rewardTable)

	var rewards []t.Reward

	err := r.db.Select(&rewards, query)

	if err != nil {
		return []t.Reward{}, err
	}

	return rewards, nil
}
