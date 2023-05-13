package main

func main() {
	api, cleanup, err := initializeAPI()
	if err != nil {
		panic(err)
	}
	defer cleanup()

	api.server.GET("/", api.quoteController.GetQuote)
	if err := api.server.Run(); err != nil {
		panic(err)
	}
}
