package repository

import (
	t "admin/pkg"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type RewardItemPostgres struct {
	db *sqlx.DB
}

func NewRewardItemPostgres(db *sqlx.DB) *RewardItemPostgres {
	return &RewardItemPostgres{db: db}
}

func (r *RewardItemPostgres) CreateReward(reward t.Reward) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	//TODO
	var itemId int = 1
	createItemQuery, _, err := sq.
		Insert(rewardTable).Columns("title", "desc").
		Values(reward.Title, reward.Desc).
		ToSql()

	if err != nil {
		tx.Rollback()
		return 0, err
	}

	_, err = tx.Exec(createItemQuery)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return itemId, tx.Commit()
}
