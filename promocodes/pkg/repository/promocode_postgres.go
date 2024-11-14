package repository

import (
	"encoding/json"
	"fmt"
	t "promocodes"
	"strings"
	"time"

	"log"

	"github.com/jmoiron/sqlx"
)

type PromocodePostgres struct {
	db *sqlx.DB
}

func NewPromocodePostgres(db *sqlx.DB) *PromocodePostgres {
	return &PromocodePostgres{db: db}
}

func (p *PromocodePostgres) GetPromocode(promocode t.Promocode) (t.Promocode, error) {
	tx, err := p.db.Begin()
	if err != nil {
		return t.Promocode{}, err
	}

	log.Printf("repository-promocode: GetPromocode promocode %d\n", promocode.Promocode)

	query := fmt.Sprintf("SELECT * FROM %s WHERE promocode = $1", promocodeTable)

	var prwcd t.Promocode

	err = p.db.Get(&prwcd, query, promocode.Promocode)
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
