package main

import "chapter2-challenge-sesi-2/routers"

func main() {
	PORT := ":8000"

	routers.StartServer().Run(PORT)

}
