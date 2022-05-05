package main

import (
	"fmt"
	"reflect"
)

type article struct {
	Pk           int      `sql:"pk"`
	Title        string   `sql:"title"`
	Text_content []string `sql:"text"`
	Image_url    []string `sql:"images"`
	Snippet_url  []string `sql:"snippets"`
	Source_url   []string `sql:"sources"`
}

func main() {
	a := article{}

	/*
		ctx := context.Background()

		conn, err := pgx.Connect(ctx, os.Getenv("POSTGRES_DB"))
		if err != nil {
			conn.Close(ctx)
		}
		defer conn.Close(ctx)

		qs := "select pk, title, text_content, image_url, snippet_url, source_url from article where pk = 1"
		rows, err := conn.Query(ctx, qs)
		if err != nil {
			rows.Close()
		}
		defer rows.Close()

	*/

	// Next step: Attempted to assign to the fields as per reflection
	fileType := reflect.TypeOf(a)
	fmt.Println(fileType)

	//readRows(a)
}

func unveilStruct(str interface{}) {
	fileType := reflect.TypeOf(str)
	fmt.Println(fileType)

	/*
		fs := reflect.VisibleFields(fileType)

		for i, el := range fs {
			fmt.Println(i, el.Name)
		}
		fmt.Println(len(fs))
	*/
}

func readRows(str interface{}) error {
	fileType := reflect.TypeOf(str)
	fmt.Println(fileType)

	/*
		for rows.Next() {
			errScan := rows.Scan(
				&a.Pk,
				&a.Title,
				&a.Text,
				&a.Image_url,
				&a.Snippet_url,
				&a.Source_url,
			)
			if errScan != nil {
				fmt.Fprintf(os.Stderr, "Error while scanning rows in the Article table: %v\n", errScan)
				return a, errScan
			}
		}
	*/

	return nil
}
