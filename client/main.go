package main

import (
	"context"
	"log"
	"net/http"

	testv1 "github.com/b099l3/dead-orbit/deadorbit/proto/test/v1"
)

func main() {
	client := testv1.NewThingServiceApiProtobufClient("http://localhost:8000", &http.Client{})

	ctx := context.Background()

	_, err := client.CreateThing(ctx, &testv1.CreateThingRequest{Text: "Hello Thing"})
	if err != nil {
		log.Fatal(err)
	}

	allThings, err := client.GetThings(ctx, &testv1.GetThingsRequest{})
	if err != nil {
		log.Fatal(err)
	}

	for _, thing := range allThings.Things {
		log.Println(thing)
	}
}
