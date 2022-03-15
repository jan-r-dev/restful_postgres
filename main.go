package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

type article struct {
	pk           int
	title        string
	text_content []string
	image_url    []string
	snippet_url  []string
	source_url   []string
}

var query string = `
select article.pk, title, text_content, image_url, snippet_url, source_url from article

inner join title on article.title_fk = title.pk
`

func main() {
	c := context.Background()

	conn := connect(c)
	defer conn.Close(c)

	rows, errQ := conn.Query(c, query)
	if errQ != nil {
		fmt.Fprintf(os.Stderr, "Error while executing query: %v\n", errQ)
	}
	defer rows.Close()

	as := readRows(rows)

	if rows.Err() != nil {
		fmt.Fprintf(os.Stderr, "Error while reading rows: %v\n", rows.Err())
	}

	fmt.Println(as[0].title)
}

func connect(c context.Context) *pgx.Conn {
	conn, err := pgx.Connect(c, os.Getenv("POSTGRES_DB"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return conn
}

func readRows(rows pgx.Rows) []article {
	as := []article{}

	for rows.Next() {
		a := article{}
		err := rows.Scan(
			&a.pk,
			&a.title,
			&a.text_content,
			&a.image_url,
			&a.snippet_url,
			&a.source_url,
		)
		as = append(as, a)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error while scanning rows: %v\n", err)
		}
	}

	return as
}
