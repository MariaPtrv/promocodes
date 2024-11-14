package repository

import (
	"fmt"
	"log"
	t "promocodes"

	"github.com/jmoiron/sqlx"
)

type RewardsPostgres struct {
	db *sqlx.DB
}

func NewRewardsPostgres(db *sqlx.DB) *RewardPostgres {
	return &RewardPostgres{db: db}
}

func (r *RewardPostgres) NewRewardsRecord(record t.RewardsRecord) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	log.Printf("repository-rewards-record: NewRewardsRecord\n")

	query := fmt.Sprintf("INSERT INTO %s (promocode_id, user_id, \"timestamp\") VALUES ($1, $2, $3)", rewardsTable)
	_, err = tx.Exec(query, record.Promocode_id, record.User_id, record.Timestamp)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *RewardPostgres) GetRewardsRecordByUserId(record t.RewardsRecord) (t.RewardsRecord, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return t.RewardsRecord{}, err
	}

	log.Printf("repository-rewards-record: GetRewardsRecordByUserId\n")

	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1 AND promocode_id = $2", rewardsTable)

	var rdb t.RewardsRecord

	err = r.db.Get(&rdb, query, record.User_id, record.Promocode_id)
	if err != nil {
		tx.Rollback()
		return t.RewardsRecord{}, err
	}
	return rdb, tx.Commit()
}
