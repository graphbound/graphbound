package main

func main() {
	api, cleanup, err := initializeAPI()
	if err != nil {
		panic(err)
	}
	defer cleanup()

	if err := api.server.Run(); err != nil {
		panic(err)
	}
}
