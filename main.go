package main

import (
	"context"
	"encoding/json"
	"fmt"
	"gamedatacollectorcron/dto"
	"gamedatacollectorcron/model"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func fetchGames() ([]model.Game, error) {
	apiKey := os.Getenv("APIKEY")
	url := os.Getenv("URL")
	endpoint := fmt.Sprintf(url+"?key=%s&page=1&page_size=10", apiKey)

	resp, err := http.Get(endpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result dto.ResultGameDTO

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	var games []model.Game
	for _, item := range result.Results {
		genero := ""
		if len(item.Genres) > 0 {
			genero = item.Genres[0].Name
		}
		games = append(games, model.Game{
			ID:     fmt.Sprintf("%d", item.ID),
			Nome:   item.Name,
			Genero: genero,
		})
	}

	return games, nil
}

func saveGames(games []model.Game, collection *mongo.Collection) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var docs []interface{}
	for _, game := range games {
		docs = append(docs, game)
	}

	_, err := collection.InsertMany(ctx, docs)
	return err
}

func run(collection *mongo.Collection) {
	log.Println("Iniciando busca de jogos...")
	games, err := fetchGames()
	if err != nil {
		log.Printf("Erro ao buscar jogos: %v", err)
		return
	}
	log.Printf("Foram encontrados %d jogos.", len(games))

	err = saveGames(games, collection)
	if err != nil {
		log.Printf("Erro ao salvar jogos no MongoDB: %v", err)
	} else {
		log.Printf("Jogos salvos com sucesso.")
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error ao carregar o arquivo .env")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGOURI")))
	if err != nil {
		log.Fatalf("Error ao conectar no MongoDB: %v", err)
	}
	defer client.Disconnect(ctx)

	collection := client.Database("gamerview").Collection("games")

	crn := cron.New()
	_, err = crn.AddFunc("@every 10s", func() {
		run(collection)
	})

	if err != nil {
		log.Fatal(err)
	}
	crn.Start()

	log.Println("Cron inciado. Pressione Ctrl+C para sair.")
	select {}
}
