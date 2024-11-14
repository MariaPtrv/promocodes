package repository

import (
	t "admin/pkg"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
)

type PromocodePostgres struct {
	db *sqlx.DB
}

func NewPromocodePostgres(db *sqlx.DB) *PromocodePostgres {
	return &PromocodePostgres{db: db}
}

func (p *PromocodePostgres) CreatePromocode(promocode t.Promocode) (int, error) {
	tx, err := p.db.Begin()
	if err != nil {
		return 0, err
	}

	var itemId int = 0

	rj, _ := json.Marshal(promocode)
	log.Printf("repository-promocode: CreatePromocode promocode: %s\n", string(rj))

	query := fmt.Sprintf("INSERT INTO %s (promocode, reward_id, expires, max_uses, remain_uses) VALUES ($1, $2, $3, $4, $5) RETURNING id", promocodeTable)
	row := tx.QueryRow(query, promocode.Promocode, promocode.Reward_id, promocode.Expires, promocode.Max_uses, promocode.Remain_uses)
	err = row.Scan(&itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	var s time.Time
	if promocode.Expires != nil && s.Unix() == promocode.Expires.Unix() {
		query = fmt.Sprintf("UPDATE %s SET expires = NULL WHERE id = $1", promocodeTable)
		_, err = tx.Exec(query, itemId)

		if err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	return itemId, tx.Commit()
}

func (p *PromocodePostgres) GetPromocodeById(promocode t.Promocode) (t.Promocode, error) {
	tx, err := p.db.Begin()
	if err != nil {
		return t.Promocode{}, err
	}

	log.Printf("repository-promocode: GetPromocodeById promocode id: %d\n", promocode.Id)

	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", promocodeTable)

	var prwcd t.Promocode

	err = p.db.Get(&prwcd, query, promocode.Id)
	if err != nil {
		tx.Rollback()
		return t.Promocode{}, err
	}

	return prwcd, tx.Commit()
}

func (p *PromocodePostgres) UpdatePromocode(promocode t.Promocode) (int, error) {
	tx, err := p.db.Begin()
	if err != nil {
		return 0, err
	}

	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if promocode.Promocode != nil {
		setValues = append(setValues, fmt.Sprintf("promocode=$%d", argId))
		args = append(args, *promocode.Promocode)
		argId++
	}

	if promocode.Reward_id != nil {
		setValues = append(setValues, fmt.Sprintf("reward_id=$%d", argId))
		args = append(args, *promocode.Reward_id)
		argId++
	}

	if promocode.Expires != nil {
		setValues = append(setValues, fmt.Sprintf("expires=$%d", argId))
		args = append(args, *promocode.Expires)
		argId++
	}

	if promocode.Max_uses != nil {
		setValues = append(setValues, fmt.Sprintf("max_uses=$%d", argId))
		args = append(args, *promocode.Max_uses)
		argId++
	}

	if promocode.Remain_uses != nil {
		setValues = append(setValues, fmt.Sprintf("remain_uses=$%d", argId))
		args = append(args, *promocode.Remain_uses)
		argId++
	}

	rj, _ := json.Marshal(promocode)
	log.Printf("repository-promocode: UpdatePromocode promocode: %s\n", string(rj))

	setQuery := strings.Join(setValues, ", ")

	var itemId int
	query := fmt.Sprintf(`UPDATE %s promocode SET %s WHERE promocode.id = $%d RETURNING id`,
		promocodeTable, setQuery, argId)

	args = append(args, *promocode.Id)
	row := tx.QueryRow(query, args...)
	err = row.Scan(&itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	var s time.Time
	if promocode.Expires != nil && s.Unix() == promocode.Expires.Unix() {
		query = fmt.Sprintf("UPDATE %s SET expires = NULL WHERE id = $1", promocodeTable)
		_, err = tx.Exec(query, itemId)

		if err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	return itemId, tx.Commit()
}

func (p *PromocodePostgres) DeletePromocode(promocode t.Promocode) error {
	tx, err := p.db.Begin()
	if err != nil {
		return err
	}

	rj, _ := json.Marshal(promocode)
	log.Printf("repository-promocode: DeletePromocode promocode: %s\n", string(rj))

	query := fmt.Sprintf("DELETE FROM %s WHERE promocode = $1", promocodeTable)
	_, err = tx.Exec(query, promocode.Promocode)

	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (p *PromocodePostgres) GetPromocodes() ([]t.Promocode, error) {
	tx, err := p.db.Begin()
	if err != nil {
		return []t.Promocode{}, err
	}

	log.Printf("repository-promocode: GetPromocodes\n")

	query := fmt.Sprintf("SELECT * FROM %s", promocodeTable)

	var rewards []t.Promocode

	err = p.db.Select(&rewards, query)
	if err != nil {
		tx.Rollback()
		return []t.Promocode{}, err
	}

	return rewards, tx.Commit()
}
