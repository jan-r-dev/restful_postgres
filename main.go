package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

type article struct {
	Pk           int
	Title        string
	Text_content []string
	Image_url    []string
	Snippet_url  []string
	Source_url   []string
}

var querystringArticle string = `
select article.pk, title, text_content, image_url, snippet_url, source_url from article

inner join title on article.title_fk = title.pk
`
var querystringProject string = `
select project.pk, title, summary, url, created_date, stack from project

inner join title on project.title_fk = title.pk
`

// TO DO
/*
1/ Create project struct
2/ Create functions for
	Querying project
	Processing project rows

*/
// TO DO ENDS
func main() {
	c := context.Background()

	conn := connect(c)
	defer conn.Close(c)

	fmt.Println("Query for article or project: ")
	var choice string
	fmt.Scanln(&choice)

	var rows pgx.Rows
	var bs [][]byte
	if choice == "a" {
		rows = queryArticle(conn, c, querystringArticle)
		bs = rowsArticle(rows)
	} else if choice == "p" {
		// currently placeholder
		rows = queryArticle(conn, c, querystringProject)
		bs = rowsProject(rows)
	}

	defer rows.Close()

	if rows.Err() != nil {
		fmt.Fprintf(os.Stderr, "Error while reading rows: %v\n", rows.Err())
	}

	fmt.Println(len(bs[0]))
}

func connect(c context.Context) *pgx.Conn {
	conn, err := pgx.Connect(c, os.Getenv("POSTGRES_DB"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return conn
}

func queryArticle(conn *pgx.Conn, c context.Context, q string) pgx.Rows {
	rows, err := conn.Query(c, q)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while executing query: %v\n", err)
	}

	return rows
}

func rowsArticle(rows pgx.Rows) [][]byte {
	bsj := [][]byte{}

	for rows.Next() {
		a := article{}
		errS := rows.Scan(
			&a.Pk,
			&a.Title,
			&a.Text_content,
			&a.Image_url,
			&a.Snippet_url,
			&a.Source_url,
		)
		if errS != nil {
			fmt.Fprintf(os.Stderr, "Error while scanning rows: %v\n", errS)
		}

		bs, errM := json.Marshal(a)
		if errM != nil {
			fmt.Fprintf(os.Stderr, "Error while marshaling json: %v\n", errM)
		}

		bsj = append(bsj, bs)
	}

	return bsj
}

func rowsProject(rows pgx.Rows) [][]byte {
	// Placeholder
	bsj := [][]byte{}

	for rows.Next() {
		a := article{}
		errS := rows.Scan(
			&a.Pk,
			&a.Title,
			&a.Text_content,
			&a.Image_url,
			&a.Snippet_url,
			&a.Source_url,
		)
		if errS != nil {
			fmt.Fprintf(os.Stderr, "Error while scanning rows: %v\n", errS)
		}

		bs, errM := json.Marshal(a)
		if errM != nil {
			fmt.Fprintf(os.Stderr, "Error while marshaling json: %v\n", errM)
		}

		bsj = append(bsj, bs)
	}

	return bsj
}
