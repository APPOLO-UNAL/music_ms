package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/olivere/elastic/v7"
)

func main() {
	// Reemplaza con tu URL de Elasticsearch y API key
	elasticURL := "https://660529f167b9476490072c3dae728b4d.us-central1.gcp.cloud.es.io:443"
	apiKey := "YW9JSzRJOEI3MndYRF9jQ20zRUU6M1VPcEc3WnZRNXVUVklyRml6YThHZw==" // Este es el valor de "encoded" de tu API key

	// Crear un cliente de Elasticsearch
	client, err := elastic.NewClient(
		elastic.SetURL(elasticURL),
		elastic.SetSniff(false), // Desactiva Sniffing en un entorno gestionado como GCP
		elastic.SetHeaders(http.Header{
			"Authorization": []string{fmt.Sprintf("ApiKey %s", apiKey)},
		}),
	)
	if err != nil {
		log.Fatalf("Error creando el cliente: %s", err)
	}

	// Pinging al cluster
	info, code, err := client.Ping(elasticURL).Do(context.Background())
	if err != nil {
		log.Fatalf("Error haciendo ping al cluster: %s", err)
	}
	fmt.Printf("Elasticsearch version %s, status code %d\n", info.Version.Number, code)
	fmt.Println("Acabo")
}
