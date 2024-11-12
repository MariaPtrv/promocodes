package repository

import (
	t "admin/pkg"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

type RewardsPostgres struct {
	db *sqlx.DB
}

func NewRewardsPostgres(db *sqlx.DB) *RewardsPostgres {
	return &RewardsPostgres{db: db}
}

func (r *RewardsPostgres) GetRewards() ([]t.Reward, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return []t.Reward{}, err
	}

	log.Printf("repository-rewards: GetRewards\n")

	query := fmt.Sprintf("SELECT * FROM %s", rewardTable)

	var rewards []t.Reward

	err = r.db.Select(&rewards, query)
	if err != nil {
		tx.Rollback()
		return []t.Reward{}, err
	}

	return rewards, tx.Commit()
}
