package repository

import (
	"fmt"
	t "promocodes"

	"log"

	"github.com/jmoiron/sqlx"
)

type PromocodesPostgres struct {
	db *sqlx.DB
}

func NewPromocodesPostgres(db *sqlx.DB) *PromocodesPostgres {
	return &PromocodesPostgres{db: db}
}

func (p *PromocodesPostgres) GetPromocode(promocode t.Promocode) (t.Promocode, error) {
	tx, err := p.db.Begin()
	if err != nil {
		return t.Promocode{}, err
	}

	log.Printf("repository-promocodes: GetPromocode promocode %d\n", promocode.Promocode)

	query := fmt.Sprintf("SELECT * FROM %s WHERE promocode = $1", promocodeTable)

	var prwcd t.Promocode

	err = p.db.Get(&prwcd, query, promocode.Promocode)
	if err != nil {
		tx.Rollback()
		return t.Promocode{}, err
	}

	return prwcd, tx.Commit()
}

func (p *PromocodesPostgres) GetRewardsRecordByUserId(record t.RewardsRecord) (t.RewardsRecord, error) {
	tx, err := p.db.Begin()
	if err != nil {
		return t.RewardsRecord{}, err
	}

	log.Printf("repository-promocodes: GetRewardsRecordByUserId\n")

	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1 AND promocode_id = $2", rewardsTable)

	var rdb t.RewardsRecord

	err = p.db.Get(&rdb, query, record.User_id, record.Promocode_id)
	if err != nil {
		tx.Rollback()
		return t.RewardsRecord{}, err
	}
	return rdb, tx.Commit()
}

func (p *PromocodesPostgres) GetRewardById(reward t.Reward) (t.Reward, error) {
	tx, err := p.db.Begin()
	if err != nil {
		return t.Reward{}, err
	}

	log.Printf("repository-promocodes: GetRewardById reward id: %d\n", reward.Id)

	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", rewardTable)

	var rdb t.Reward

	err = p.db.Get(&rdb, query, reward.Id)
	if err != nil {
		tx.Rollback()
		return t.Reward{}, err
	}

	return rdb, tx.Commit()
}

func (p *PromocodesPostgres) ApplyPromocodeAction(record t.RewardsRecord, promocode t.Promocode) error {
	tx, err := p.db.Begin()
	if err != nil {
		return err
	}

	log.Printf("repository-promocodes: ApplyPromocodeAction")

	query := fmt.Sprintf("INSERT INTO %s (promocode_id, user_id, \"timestamp\") VALUES ($1, $2, $3)", rewardsTable)
	_, err = tx.Exec(query, record.Promocode_id, record.User_id, record.Timestamp)
	if err != nil {
		tx.Rollback()
		return err
	}

	query = fmt.Sprintf(`UPDATE %s promocode SET remain_uses=$1 WHERE promocode.id = $2`,
		promocodeTable)

	_, err = tx.Exec(query, *promocode.Remain_uses, *promocode.Id)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
