package main

import (
	"os"
	"os/signal"

	"github.com/PbChuy/tink/events"
	"github.com/PbChuy/tink/services"
	"github.com/joho/godotenv"
)

func init() {
	//load .env
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}
}

// im jesus and i am such a bad programmer its actually insane
func main() {
	services.ConnectDiscord(events.Events)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	println("Press Ctrl+C to exit")

	//handle shutdown
	<-stop
	println("Shutting Down")
	services.DisconnectDiscord()
	println("Goodbye :)")

}
