package pdnsmodel

import "database/sql"

// Domain model
type Domain struct {
	ID             uint `gorm:"primaryKey"`
	Name           string
	Master         sql.NullString
	LastCheck      sql.NullInt32
	Type           string
	NotifiedSerial sql.NullInt64
	Account        sql.NullString
}

// Record model
type Record struct {
	ID         uint `gorm:"primaryKey"`
	DomainID   int
	Domain     Domain
	Name       sql.NullString
	Type       sql.NullString
	Content    sql.NullString
	TTL        sql.NullInt32
	Prio       sql.NullInt32
	ChangeDate sql.NullInt64
	Disabled   sql.NullBool
	Ordername  sql.NullString
	Auth       sql.NullBool
}
