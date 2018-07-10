package main

import (
	"fmt"
	"log"
	"net/http"
	"testing"

	"github.com/pact-foundation/pact-go/dsl"
)

type Animal struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func TestConsumer(t *testing.T) {

	// Create Pact client
	pact := &dsl.Pact{
		Consumer: "MyConsumer",
		Provider: "MyProvider",
	}
	defer pact.Teardown()

	// Pass in test case
	var test = func() error {
		u := fmt.Sprintf("http://localhost:%d/api/v1/animal/1", pact.Server.Port)
		req, err := http.NewRequest("GET", u, nil)

		// NOTE: by default, request bodies are expected to be sent with a Content-Type
		// of application/json. If you don't explicitly set the content-type, you
		// will get a mismatch during Verification.
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			return err
		}
		if _, err = http.DefaultClient.Do(req); err != nil {
			return err
		}

		return err
	}

	// Set up our expected interactions.
	pact.
		AddInteraction().

		Given("Serval exists").
		UponReceiving("A request to get serval info").
		WithRequest(dsl.Request{
		Method:  "GET",
		Path:    dsl.String("/api/v1/animal/1"),
		Headers: dsl.MapMatcher{"Content-Type": dsl.Like("application/json")},
	}).
		WillRespondWith(dsl.Response{
		Status:  200,
		Headers: dsl.MapMatcher{"Content-Type": dsl.Like("application/json")},
		Body:    Animal{Id: 1, Name: "サーバル"},
	})

	// Verify
	if err := pact.Verify(test); err != nil {
		log.Fatalf("Error on Verify: %v", err)
	}
}
