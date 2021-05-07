package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"cloud.google.com/go/pubsub"
)

type SlackJson struct {
	Type      string
	Challenge string
}

func SlackIngester(w http.ResponseWriter, r *http.Request) {
	var j SlackJson

	err := json.NewDecoder(r.Body).Decode(&j)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	switch j.Type {
	case "url_verification":
		fmt.Fprintf(w, j.Challenge)
	case "event_callback":
		client, err := pubsub.NewClient(ctx, "devs-sandbox")
		if err != nil {
			log.Fatal(err)
		}

		topic := client.Topic("ruby-slackbot")
		res := topic.Publish(ctx, &pubsub.Message{
			Data: []byte(r.Body),
		})
		msgID, err := res.Get(ctx)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprintf(w, "ok")
	default:
		fmt.Fprintf(w, "ok")
	}
}

func main() {
	fmt.Println("Starting")
	http.HandleFunc("/", SlackIngester)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
