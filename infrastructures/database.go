package infrastructures

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/SurgicalSteel/kvothe/interfaces"
	"github.com/SurgicalSteel/kvothe/resources"
	"github.com/SurgicalSteel/kvothe/utils"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	defaultTimeOutString = "5"
)

//Database struct
type PostgreSQLHandler struct {
	DBRead  *sqlx.DB
	DBWrite *sqlx.DB
	Tx      *sqlx.Tx
}

//SqlxTx is a wrapper struct for sqlx Tx
type SqlxTx struct {
	Tx *sqlx.Tx
}

func (d *SqlxTx) Commit() {
	_ = d.Tx.Commit()
}

func (d *SqlxTx) Rollback() {
	_ = d.Tx.Rollback()
}

func (d *PostgreSQLHandler) ConnectDB(
	dbAccRead *resources.DBAccount,
	dbAccWrite *resources.DBAccount,
) {

	if dbAccRead.Timeout == "" {
		dbAccRead.Timeout = defaultTimeOutString
	}

	dbRead, err := sqlx.Open("postgres", utils.GeneratePostgreURL(*dbAccRead))
	if err != nil {
		log.Fatalf("Failed to open connection to DB Read! Error : %s\n", err.Error())
	}

	d.DBRead = dbRead
	err = d.DBRead.Ping()
	if err != nil {
		log.Printf("postgres read username: %s, password: %s, url: %s, port: %s, dbname: %s \n", dbAccRead.Username, dbAccRead.Password, dbAccRead.URL, dbAccRead.Port, dbAccRead.DBName)
		log.Fatalf("Failed to test connection (PING) to DB Read! Error : %s\n", err.Error())
	}

	log.Println("Successfully connect to DB Read")
	if dbAccWrite.Timeout == "" {
		dbAccWrite.Timeout = defaultTimeOutString
	}

	dbWrite, err := sqlx.Open("postgres", utils.GeneratePostgreURL(*dbAccWrite))
	if err != nil {
		log.Fatalf("Failed to open connection to DB Write! Error : %s\n", err.Error())
	}

	d.DBWrite = dbWrite
	err = d.DBWrite.Ping()
	if err != nil {
		fmt.Printf("postgres write username: %s, password: %s, url: %s, port: %s, dbname: %s \n", dbAccWrite.Username, dbAccWrite.Password, dbAccWrite.URL, dbAccWrite.Port, dbAccWrite.DBName)
		log.Fatalf("Failed to test connection (PING) to Db Write! Error : ", err.Error())
	}

	log.Println("Successfully connect to DB Write")

	d.DBWrite.SetConnMaxLifetime(time.Duration(dbAccWrite.MaxLifeTime) * time.Second)
	d.DBRead.SetConnMaxLifetime(time.Duration(dbAccRead.MaxLifeTime) * time.Second)

	// max connection
	d.DBWrite.SetMaxOpenConns(dbAccWrite.MaxOpenConns)
	d.DBRead.SetMaxOpenConns(dbAccRead.MaxOpenConns)

	d.DBRead.SetMaxIdleConns(dbAccRead.MaxIdleConns)
	d.DBWrite.SetMaxIdleConns(dbAccWrite.MaxIdleConns)
}

func (d *PostgreSQLHandler) Close() {
	if d.DBRead != nil {
		if err := d.DBRead.Close(); err != nil {
			log.Printf("Failed to close connection to DB Read! Error : %s\n", err.Error())
		} else {
			log.Println("Successfuly closing connection to DB Read")
		}
	}

	if d.DBWrite != nil {
		if err := d.DBWrite.Close(); err != nil {
			log.Printf("Failed to close connection to DB Write! Error : %s\n", err.Error())
		} else {
			log.Println("Successfuly closing connection to DB Write")
		}
	}
}

func (d *PostgreSQLHandler) Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := d.DBRead.Query(query, args...)
	return rows, err
}

func (d *PostgreSQLHandler) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	result, err := d.DBWrite.ExecContext(ctx, query, args...)
	return result, err
}

func (d *PostgreSQLHandler) Exec(query string, args ...interface{}) (sql.Result, error) {
	result, err := d.DBWrite.Exec(query, args...)
	return result, err
}

func (d *PostgreSQLHandler) Get(dest interface{}, query string, args ...interface{}) error {
	err := d.DBRead.Get(dest, query, args...)
	return err
}

func (d *PostgreSQLHandler) DriverName() string {
	return d.DBRead.DriverName()
}

func (d *PostgreSQLHandler) Select(dest interface{}, query string, args ...interface{}) error {
	err := d.DBRead.Select(dest, query, args...)
	return err
}

func (d *PostgreSQLHandler) Begin() (interfaces.IDBTx, error) {
	tx, err := d.DBWrite.Beginx()
	sqlxTx := SqlxTx{
		Tx: tx,
	}
	if err != nil {
		return &sqlxTx, err
	}
	return &sqlxTx, nil
}

func (d *PostgreSQLHandler) BeginTx() (*sqlx.Tx, error) {
	tx, err := d.DBWrite.Beginx()
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (d *PostgreSQLHandler) Commit() error {
	return d.Tx.Commit()
}

func (d *PostgreSQLHandler) Rollback() error {
	return d.Tx.Rollback()
}

func (d *PostgreSQLHandler) Queryx(query string, args ...interface{}) (*sqlx.Rows, error) {
	tx, err := d.DBRead.Queryx(query, args...)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (d *PostgreSQLHandler) TransactionBlock(tx *sqlx.Tx, fc func(tx *sqlx.Tx) error) error {
	if tx == nil {
		return errors.New("DB transaction is nil")
	}
	err := fc(tx)

	if err != nil {
		errTx := tx.Rollback()
		if errTx != nil {
			return errTx
		}
		return err
	}

	errTx := tx.Commit()
	if errTx != nil {
		return errTx
	}

	return nil
}

func (d *PostgreSQLHandler) Rebind(query string) string {
	return d.DBRead.Rebind(query)
}

func (d *PostgreSQLHandler) In(query string, params ...interface{}) (string, []interface{}, error) {
	query, args, err := sqlx.In(query, params...)
	return query, args, err
}

// QueryRow executes a query that is expected to return at most one row.
// QueryRow always returns a non-nil value. Errors are deferred until
// Row's Scan method is called.
// If the query selects no rows, the *Row's Scan will return ErrNoRows.
// Otherwise, the *Row's Scan scans the first selected row and discards
// the rest.
func (d *PostgreSQLHandler) QueryRow(query string, args ...interface{}) *sql.Row {
	return d.DBWrite.QueryRowContext(context.Background(), query, args...)
}

func (d *PostgreSQLHandler) QueryRowSqlx(query string, args ...interface{}) *sqlx.Row {
	return d.DBWrite.QueryRowx(query, args...)
}

func (d *PostgreSQLHandler) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := d.DBRead.QueryContext(ctx, query, args...)
	return rows, err
}

func (d *PostgreSQLHandler) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	err := d.DBRead.GetContext(ctx, dest, query, args...)
	return err
}
