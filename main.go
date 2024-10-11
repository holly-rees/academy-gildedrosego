package main

import (
	"fmt"
	"gildedrose/api"
	"gildedrose/repository"
	"gildedrose/service"
	"io"
	"net/http"
)

func rootHandler(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "Hello, world!")
}

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		// Continue with the next handler
		next.ServeHTTP(writer, request)
	})
}

func setUpRouter(itemsAPI *api.ItemAPI) *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("GET /", rootHandler)
	router.HandleFunc("GET /api/items", itemsAPI.GetItems)

	return router
}

func setItemRepo() (*repository.InMemoryItemRepository, error) {
	return repository.NewInMemoryItemRepository(), nil
}

func main() {
	itemRepo, err := setItemRepo()
	if err != nil {
		panic(err)
	}

	// Initialize services
	itemService := service.NewItemService(*itemRepo)
	itemAPI := api.NewItemAPI(itemService)
	router := setUpRouter(itemAPI)

	// Starting the HTTP server on port 8080
	fmt.Println("Server listening on port 8080...")
	err = http.ListenAndServe(":8080", CorsMiddleware(router))
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
