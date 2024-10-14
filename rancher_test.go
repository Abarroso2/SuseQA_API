package rancher

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"testing"
)

func Test_Request(t *testing.T) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://localhost/v3-public/localProviders/local?action=login", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(os.Getenv("CATTLE_ACCESS_KEY"), os.Getenv("CATTLE_SECRET_KEY"))
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
}
