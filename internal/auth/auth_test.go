package auth_test

import (
	"log"
	"net/http"
	"testing"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
)

func TestNoAuthorizationKeyShouldFail(t *testing.T) {
	header := http.Header{}
	header.Add("Authorization", "ApiKey BreakIt")
	_, err := auth.GetAPIKey(header)

	if err != nil {
		log.Printf("%v has errored", err)
		return
	}
	log.Fatalf("Empty key has no error!")
}

func TestWithKeyShouldGetAValue(t *testing.T) {
	header := http.Header{}
	header.Add("Authorization", "ApiKey notEmptyKeyValue123")

	_, err := auth.GetAPIKey(header)
	if err != nil {
		log.Fatalf("an error occured: %v", err)
	}
}
