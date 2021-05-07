package app

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	ctx := context.Background()

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal([]byte(data), &j)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	switch j.Type {
	case "url_verification":
		fmt.Fprintf(w, j.Challenge)
	case "event_callback":
		fmt.Println("data: %s", data)

		client, err := pubsub.NewClient(ctx, "devs-sandbox")
		if err != nil {
			log.Fatal(err)
		}

		topic := client.Topic("ruby-slackbot")

		res := topic.Publish(ctx, &pubsub.Message{
			Data: data,
		})
		msgID, err := res.Get(ctx)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("msgID: %~v", msgID)

		fmt.Fprintf(w, "ok")
	default:
		fmt.Fprintf(w, "ok")
	}
}

/*
func main() {
	fmt.Println("Starting")
	http.HandleFunc("/", SlackIngester)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
*/
