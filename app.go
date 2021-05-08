package app

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/pubsub"
)

type SlackJSON struct {
	Type      string
	Challenge string
}

func SlackIngester(w http.ResponseWriter, r *http.Request) {
	var j SlackJSON
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
		fmt.Printf("$GCP_PROJECT: %v, $PUBSUB_TOPIC: %v\n", os.Getenv("GCP_PROJECT"), os.Getenv("PUBSUB_TOPIC"))
		fmt.Printf("data: %s\n", data)

		// Make sure it's not a retry
		if _, ok := r.Header["X-Slack-Retry-Num"]; ok {
			fmt.Fprintf(w, "ok")
			return
		}

		client, err := pubsub.NewClient(ctx, os.Getenv("GCP_PROJECT"))
		if err != nil {
			log.Fatal(err)
		}

		topic := client.Topic(os.Getenv("PUBSUB_TOPIC"))

		res := topic.Publish(ctx, &pubsub.Message{
			Data: data,
		})
		msgID, err := res.Get(ctx)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("msgID: %s\n", msgID)

		fmt.Fprintf(w, "ok")
	default:
		fmt.Fprintf(w, "ok")
	}
}
