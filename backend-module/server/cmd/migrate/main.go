package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/jackc/pgx/v5/stdlib"
)

const (
	defaultHost     = "47.104.208.126"
	defaultPort     = "5432"
	defaultUser     = "postgres"
	defaultPassword = "PgTest_9Kx7mL2Qv8Tn"
	defaultDBName   = "ovo_gaming_platform"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	adminDSN := buildDSN("postgres")
	adminDB, err := sql.Open("pgx", adminDSN)
	if err != nil {
		return err
	}
	defer adminDB.Close()

	if err := adminDB.Ping(); err != nil {
		return fmt.Errorf("ping admin db: %w", err)
	}

	if _, err := adminDB.Exec(`CREATE DATABASE ` + defaultDBName); err != nil {
		if !strings.Contains(strings.ToLower(err.Error()), "already exists") {
			return fmt.Errorf("create database: %w", err)
		}
	}

	targetDSN := buildDSN(defaultDBName)
	targetDB, err := sql.Open("pgx", targetDSN)
	if err != nil {
		return err
	}
	defer targetDB.Close()

	if err := targetDB.Ping(); err != nil {
		return fmt.Errorf("ping target db: %w", err)
	}

	files := []string{"sql/001_init.sql", "sql/002_seed.sql"}
	for _, path := range files {
		raw, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		if err := execSQLScript(targetDB, string(raw)); err != nil {
			return fmt.Errorf("execute %s: %w", path, err)
		}
	}

	log.Println("database migration completed")
	return nil
}

func buildDSN(dbname string) string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		defaultHost, defaultPort, defaultUser, defaultPassword, dbname,
	)
}

func execSQLScript(db *sql.DB, script string) error {
	scanner := bufio.NewScanner(strings.NewReader(script))
	var builder strings.Builder
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "\\c") {
			continue
		}
		builder.WriteString(scanner.Text())
		builder.WriteByte('\n')
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	statements := splitSQL(builder.String())
	for _, stmt := range statements {
		if strings.TrimSpace(stmt) == "" {
			continue
		}
		if _, err := db.Exec(stmt); err != nil {
			return err
		}
	}
	return nil
}

func splitSQL(script string) []string {
	var (
		statements []string
		current    strings.Builder
		inSingle   bool
		inDouble   bool
	)

	for _, r := range script {
		switch r {
		case '\'':
			if !inDouble {
				inSingle = !inSingle
			}
			current.WriteRune(r)
		case '"':
			if !inSingle {
				inDouble = !inDouble
			}
			current.WriteRune(r)
		case ';':
			if inSingle || inDouble {
				current.WriteRune(r)
				continue
			}
			stmt := strings.TrimSpace(current.String())
			if stmt != "" {
				statements = append(statements, stmt)
			}
			current.Reset()
		default:
			current.WriteRune(r)
		}
	}

	if tail := strings.TrimSpace(current.String()); tail != "" {
		statements = append(statements, tail)
	}
	return statements
}
