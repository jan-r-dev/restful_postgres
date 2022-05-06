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

type fakeRow struct {
	one   int
	two   string
	three []string
	four  []string
	five  []string
	six   []string
	seven []string
}

func main() {
	a := article{}

	fakerows := []fakeRow{{
		one:   1,
		two:   "hello",
		three: []string{"three1", "three2"},
		four:  []string{"four1", "four2"},
		five:  []string{"five1", "five2"},
		six:   []string{"six1", "six2"},
		seven: []string{"seven1", "seven2"},
	}}

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

	readRows(&a, fakerows)

	//readRows(a)
}

func readRows(a interface{}, rows []fakeRow) error {
	type1 := reflect.TypeOf(a)
	//value1 := reflect.ValueOf(a)

	s := reflect.ValueOf(&a).Elem()

	fmt.Println(a, s)

	if type1 == reflect.TypeOf(article{}) {

		art := a.(article)

		fmt.Println(art)
	}
	/*
		for _, row := range rows {
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

	// Read section titled Structs: https://go.dev/blog/laws-of-reflection

	return nil
}

func unveilStruct(str interface{}) {
	fileType := reflect.TypeOf(str)

	fs := reflect.VisibleFields(fileType)

	for i, el := range fs {
		fmt.Println(i, el.Name)
	}
}
