package items

import (
	"database/sql"
	"fmt"
	"testmock/internal/domain"
)

// constructor
func NewStorageSQL(db *sql.DB) Storage {
	return &storageSQL{db: db}
}

// controller
const (
	QUERY_GET_BY_NAME = "SELECT id, name, weight, price, release_date FROM items WHERE name = ?;"
	QUERY_UPDATE = "UPDATE items SET weight=?, price=?, release_date=? WHERE name = ?;"
)

type storageSQL struct {
	db *sql.DB
}
// read
func (st *storageSQL) GetByName(name string) (i domain.Item, err error) {
	var stmt *sql.Stmt
	stmt, err = st.db.Prepare(QUERY_GET_BY_NAME)
	if err != nil {
		err = fmt.Errorf("%w. %s", ErrStorageInternal, err.Error())
		return
	}

	row := stmt.QueryRow(name)
	err = row.Scan(&i.ID, &i.Name, &i.Weight, &i.Price, &i.Release)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			err = fmt.Errorf("%w. %s", ErrStorageItemNotFound, err.Error())
		default:
			err = fmt.Errorf("%w. %s", ErrStorageInternal, err.Error())
		}
		return
	}
	return
}
// write
func (st *storageSQL) UpdateByName(name string, i domain.Item) (err error) {
	var stmt *sql.Stmt
	stmt, err = st.db.Prepare(QUERY_UPDATE)
	if err != nil {
		err = fmt.Errorf("%w. %s", ErrStorageInternal, err.Error())
		return
	}

	var result sql.Result
	result, err = stmt.Exec(i.Weight, i.Price, i.Release, name)
	if err != nil {
		err = fmt.Errorf("%w. %s", ErrStorageInternal, err.Error())
		return
	}

	var rows int64
	rows, err = result.RowsAffected()
	if err != nil {
		err = fmt.Errorf("%w. %s", ErrStorageInternal, err.Error())
	}
	
	if rows != 1 {
		err = ErrStorageRowsAffected
	}

	return
}