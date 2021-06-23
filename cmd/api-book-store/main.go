package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/renatospaka/api-book-store/gqlgen" // update the username
	"github.com/renatospaka/api-book-store/pg"     // update the username
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// username, password, database :=
	// 		os.Getenv("POSTGRES_USER"),
	// 		os.Getenv("POSTGRES_PASSWORD"),
	// 		os.Getenv("POSTGRES_DB")
	username, _ := os.LookupEnv("POSTGRES_USER")
	password, _ := os.LookupEnv("POSTGRES_PASSWORD")
	database, _ := os.LookupEnv("POSTGRES_DB")

	// initialize the db
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s",
					username, password, database)
	db, err := pg.Open(dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// initialize the repository
	repo := pg.NewRepository(db)

	// configure the server
	mux := http.NewServeMux()
	mux.Handle("/", gqlgen.NewPlaygroundHandler("/query"))
	mux.Handle("/query", gqlgen.NewHandler(repo))

	// run the server
	port := ":8080"
	fmt.Fprintf(os.Stdout, "🚀 Server ready at http://localhost%s\n", port)
	fmt.Fprintln(os.Stderr, http.ListenAndServe(port, mux))
}