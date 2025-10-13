package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"school-management-system/internal/config"

	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
)

func main() {
	cnf := config.GetConfig()
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [up|down|status]")
		return
	}
	actions := os.Args[1]
	db, err := sql.Open("postgres", cnf.DBCnf.DBUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	migrations := &migrate.FileMigrationSource{
		Dir: "migrate/migrations",
	}
	switch actions {
	case "up":
		n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Applied %d up migrations\n", n)
	case "down":
		n, err := migrate.Exec(db, "postgres", migrations, migrate.Down)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Applied %d down migrations\n", n)
	case "status":
		records, err := migrate.GetMigrationRecords(db, "postgres")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Migration Status:")
		for _, r := range records {
			fmt.Println("-", r.Id, r.AppliedAt)
		}

	default:
		fmt.Println("Unknown action. Use up, down, or status")
	}

}
