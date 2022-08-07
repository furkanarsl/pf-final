package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
	_ "github.com/joho/godotenv/autoload"
)

var seedFilePath = flag.String("f", "seed.sql", "seed file location")
var dbURL = flag.String("db", os.Getenv("DB_URL"), "seed file location")

func main() {
	conn, err := pgx.Connect(context.Background(), *dbURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	flag.Parse()
	defer conn.Close(context.Background())

	file, err := os.Open(*seedFilePath)
	if err != nil {
		log.Fatal("Failed to open seed file")
	}
	batch := &pgx.Batch{}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		batch.Queue(scanner.Text())
	}
	br := conn.SendBatch(context.Background(), batch)
	br.Close()
}
