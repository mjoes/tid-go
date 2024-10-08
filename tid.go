package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"path/filepath"
  "fmt"
)

var db_path string = "~/.local/share/tid/tid.db"

func main() {
	app := &cli.App{
		Usage: "App for tidsregistrering",
		Commands: []*cli.Command{
			{
        Name:    "add",
        Aliases: []string{"a"},
        Usage:   "Add a new AO to the list",
        Action: func(cCtx *cli.Context) error {
          insertAO := `INSERT INTO A0 (code, name) VALUES (?, ?)`
          sqliteDatabase, _ := sql.Open("sqlite3", db_path)
          defer sqliteDatabase.Close()
          stmt, _ := sqliteDatabase.Prepare(insertAO)
          defer stmt.Close()
          _, _ = stmt.Exec("test", "asdf")
          fmt.Println("added task: ", cCtx.Args().First(), "AND", cCtx.String("full-name"))
          return nil
        },
        Flags: []cli.Flag{
            &cli.StringFlag{
                Name:  "full-name",
                Usage:  "Full name of AO",
                Aliases: []string{"f"},
                Value: "",
                Required: false,
            },
        },
        Before: func(cCtx *cli.Context) error {
          if cCtx.Args().Len() != 1 {
            return fmt.Errorf("Expecting exactly 1 argument")
          }
          return nil
        },
      },
		},
	}
	if fileExists(db_path) {
		os.MkdirAll(filepath.Dir(db_path), 0700)
		create_database()
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func create_database() {
	sqliteDatabase, _ := sql.Open("sqlite3", db_path)
	defer sqliteDatabase.Close()
	initAO(sqliteDatabase)
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return os.IsNotExist(err)
}

func initAO(db *sql.DB) {
	sqlAO := `CREATE TABLE AO (
		"code" TEXT NOT NULL PRIMARY KEY,		
		"name" TEXT
  );`

	_, err := db.Exec(sqlAO)
	if err != nil {
		log.Fatal(err.Error())
	}
}
