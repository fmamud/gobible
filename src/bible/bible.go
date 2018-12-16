package bible

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const baseUrl = "https://bible-api.com"

func errorhandling(err error) {
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(2)
	}
}

func Get(predicate string) string {
	resp, err := http.Get(fmt.Sprintf("%s/%s", baseUrl, predicate))

	errorhandling(err)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	errorhandling(err)

	return string(body)
}

func Help() {
	fmt.Println("Usage: bible <book> [<chapter>[:<verse>]]")
}
