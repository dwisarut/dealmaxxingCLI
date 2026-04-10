package api

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/kv"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/dwisarut/dealmaxxingCLI/internal/model"
	"github.com/joho/godotenv"
)

func ReadEnv() model.Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file!")
	}

	apiKey := os.Getenv("CLOUDFLARE_API_TOKEN")
	namespaceID := os.Getenv("CLOUDFLARE_NAMESPACE_ID")
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	config := model.Config{
		APIToken:    apiKey,
		NamespaceID: namespaceID,
		AccountID:   accountID,
	}

	return config
}

func NewService(env model.Config) *cloudflare.Client {
	client := cloudflare.NewClient(
		option.WithAPIToken(env.APIToken),
	)

	return client
}

func GetKV(config model.Config, client *cloudflare.Client) model.ProcessedKV {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	page, err := client.KV.Namespaces.Keys.List(
		ctx,
		config.NamespaceID,
		kv.NamespaceKeyListParams{
			AccountID: cloudflare.F(config.AccountID),
		},
	)

	if err != nil {
		log.Fatal("KV list failed!")
	}

	var data model.ProcessedKV

	for _, item := range page.Result {
		processed := model.Keys{
			Name: item.Name,
		}
		data.Result = append(data.Result, processed)
	}

	data.Count = int(page.ResultInfo.Count)
	return data
}
