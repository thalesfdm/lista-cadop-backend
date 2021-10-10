package main

import (
	"log"
	"os"
	api "server/internal/api/v1"
	. "server/internal/models"
)

func main() {
	f, err := os.Open("data/relatorio_cadop.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var operadoras Operadoras
	operadoras.LoadFromCSV(f)

	api.Serve(&operadoras)
}
