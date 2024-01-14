package main

import (
	"database/sql"
	"log"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/go-testfixtures/testfixtures/v3"
)

var (
	fixtures *testfixtures.Loader
)

func TestMain(m *testing.M) {
	c := mysql.Config{
		DBName:    "db_test",
		User:      "root",
		Passwd:    "root",
		Addr:      "localhost:3306",
		Net:       "tcp",
		ParseTime: true,
		Collation: "utf8mb4_unicode_ci",
		Loc:       time.UTC,
	}
	var err error
	db, err = sql.Open("mysql", c.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(m.Run())
}

func prepareTestDatabase(t *testing.T, filePaths ...string) {
	t.Helper()

	tf, err := testfixtures.New(
		testfixtures.Database(db),
		testfixtures.Dialect("mysql"),
		testfixtures.Files(filePaths...),
	)
	if err != nil {
		t.Fatal(err)
	}

	if err = tf.Load(); err != nil {
		t.Fatal(err)
	}
}

func TestSelectUsers(t *testing.T) {
	tests := []struct {
		name    string
		want    []User
		files   []string
		wantErr bool
	}{
		{
			name: "found",
			want: []User{
				{
					ID:        1,
					Name:      "test_user_1",
					CreatedAt: time.Date(2024, 1, 14, 04, 14, 15, 0, time.UTC),
					UpdatedAt: time.Date(2024, 1, 14, 04, 14, 15, 0, time.UTC),
				},
			},
			files: []string{
				"./testdata/fixtures/users.yml",
				"./testdata/fixtures/posts.yml",
			},
		},
		{
			name: "found2",
			want: []User{
				{
					ID:        2,
					Name:      "test_user_2",
					CreatedAt: time.Date(2024, 1, 14, 04, 14, 15, 0, time.UTC),
					UpdatedAt: time.Date(2024, 1, 14, 04, 14, 15, 0, time.UTC),
				},
			},
			files: []string{
				"./testdata/fixtures2/users.yml",
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		prepareTestDatabase(t, tt.files...)

		t.Run(tt.name, func(t *testing.T) {
			got, err := SelectUsers()
			if (err != nil) != tt.wantErr {
				t.Errorf("SelectUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}
