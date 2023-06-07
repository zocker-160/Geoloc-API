package impl

import (
	"database/sql"
	"errors"
	"strings"

	_ "github.com/glebarez/go-sqlite"
)


type DBRow struct {
	Ipfrom int64
	Ipto int64
	Countrycode string
	Country string
	State string
	City string
	Latitude float32
	Longitude float32
	Id int64
}

func FindEntryDB(ip string, db *sql.DB) (*DBRow, error) {
	if strings.HasPrefix(ip, "127.0.") {
		return nil, errors.New("requested loalhost lol")
	}

	ipDec, err := IpToDecimal(ip)
	if err != nil {
		return nil, err
	}

	query := `
	SELECT * FROM ips
	WHERE ? BETWEEN ipfrom AND ipto;`

	data := DBRow{}

	row := db.QueryRow(query, ipDec)
	err = row.Scan(
		&data.Ipfrom, &data.Ipto,
		&data.Countrycode, &data.Country,
		&data.State, &data.City,
		&data.Latitude, &data.Longitude,
		&data.Id,
	)

	return &data, err
}
