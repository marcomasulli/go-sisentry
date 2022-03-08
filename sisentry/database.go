package sisentry

import (
	"database/sql"
	"errors"

	"github.com/mattn/go-sqlite3"
)

var (
	ErrDuplicate    = errors.New("record already exists")
	ErrNotExists    = errors.New("row not exists")
	ErrUpdateFailed = errors.New("update failed")
	ErrDeleteFailed = errors.New("delete failed")
)

type SQLiteRepository struct {
	db *sql.DB
}

func NewSQLiteRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{
		db: db,
	}
}

func (r *SQLiteRepository) Migrate() error {
	query := `
    CREATE TABLE IF NOT EXISTS dbbuilds(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
        oid TEXT NOT NULL UNIQUE,
        DatamodelID TEXT NULL,
        DatamodelTitle TEXT NULL,
        InstanceID TEXT NULL,
        Created TEXT NULL,
		Started TEXT NULL,
		Completed TEXT NULL
    );
    `

	_, err := r.db.Exec(query)
	return err
}

func (r *SQLiteRepository) Create(dbBuild DbBuild) (*DbBuild, error) {
	res, err := r.db.Exec("INSERT INTO dbbuilds(oid, DataModelID, DatamodelTitle, InstanceID, Created, Started, Completed) values(?,?,?,?,?,?,?)",
		dbBuild.Oid, dbBuild.DatamodelID, dbBuild.DatamodelTitle, dbBuild.InstanceID, dbBuild.Created, dbBuild.Started, dbBuild.Completed)
	if err != nil {
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) {
			if errors.Is(sqliteErr.ExtendedCode, sqlite3.ErrConstraintUnique) {
				return nil, ErrDuplicate
			}
		}
		return nil, err
	}

	Id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	dbBuild.Id = Id

	return &dbBuild, nil
}

func (r *SQLiteRepository) GetByOid(Oid string) (*DbBuild, error) {
	row := r.db.QueryRow("SELECT * FROM dbbuilds WHERE oid = ?", Oid)

	var dbBuild DbBuild
	if err := row.Scan(&dbBuild.Id, &dbBuild.Oid, &dbBuild.DatamodelID, &dbBuild.DatamodelTitle, &dbBuild.InstanceID, &dbBuild.Created, &dbBuild.Started, &dbBuild.Completed); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotExists
		}
		return nil, err
	}
	return &dbBuild, nil
}
