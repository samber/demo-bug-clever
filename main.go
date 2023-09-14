package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

func warmUp(db *sql.DB, count int) {
	wg := sync.WaitGroup{}
	wg.Add(count)

	for i := 0; i < count; i++ {
		go func() {
			defer wg.Done()

			timeout, _ := context.WithTimeout(context.TODO(), 3*time.Second)

			start := time.Now()
			conn, _ := db.Conn(timeout)
			fmt.Println("Connection time: ", time.Since(start))

			_ = conn.Close()
		}()
	}

	wg.Wait()
}

func main() {
	pool, _ := sql.Open("postgres", os.Getenv("PG_URI"))

	start := time.Now()
	warmUp(pool, 100)
	log.Printf("PostgreSQL warm-up time: %s\n", time.Since(start))
}
