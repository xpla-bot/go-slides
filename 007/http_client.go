package main

import (
	// importujemy paczkę pozwalającą na czytanie ze strumienia
	"io/ioutil"
	// paczkę do obsługi http
	"net/http"
	// paczka do obsługi logowania
	"log"
)

func main() {
	// wyślij zapytanie post
	resp, err := http.Post("http://127.0.0.1:8080/insert/test", "text/plain", nil)

	// jeżeli jest błąd to zaloguj go i zakończ program
	if err != nil {
		log.Fatalf("post request failed: %v", err)
	}

	// defer - wykonaj funkcje na samym końcu akutalnej funkcji
	defer resp.Body.Close()

	// wczytaj response body
	data, err := ioutil.ReadAll(resp.Body)

	// jeśli jest bład przy wczytywaniu to zaloguj go i zakończ program
	if err != nil {
		log.Fatalf("read post response failed: %v", err)
	}

	// wyśliwet otrzymaną odpowieć i http status code
	log.Printf("response\n\tstatus=%d\n\tbody=%s\n", resp.StatusCode, string(data))

	// tworzyenie requestu DELETE
	// wrappery na metody HEAD, GET, POST są domyślnie
	// zaimplementowane w paczce net/http. Do pozostałych metod
	// musimy tworzyć Requesty.
	req, err := http.NewRequest(http.MethodDelete, "http://127.0.0.1:8080/remove", nil)

	// jeśli jest bład przy tworzeniu requesta to zaloguj go i zakończ program
	if err != nil {
		log.Fatalf("create delete request failed: %v", err)
	}

	// Wyślij request za pomocą standardowego klienta http
	dresp, err := http.DefaultClient.Do(req)

	// jeżeli jest błąd to zaloguj go i zakończ program
	if err != nil {
		log.Fatalf("delete request failed: %v", err)
	}

	// zamknij body na samym końcu
	defer dresp.Body.Close()

	// i dalej tak samo jak w przypadku POST
	data, err = ioutil.ReadAll(dresp.Body)
	if err != nil {
		log.Fatalf("read delete response failed: %v", err)
	}
	log.Printf("response\n\tstatus=%d\n\tbody=%s\n", resp.StatusCode, string(data))
}
