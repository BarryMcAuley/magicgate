package magicgate

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"google.golang.org/appengine/urlfetch"
)

func init() {
	http.HandleFunc("/", handler)
}

const (
	welcomeMsg = `<?xml version="1.0" encoding="UTF-8"?>
	<Response>
		<Say>Hello Barry, how are you?</Say>
		<Record timeout="5"/>
	</Response>`
	echoMsg = `<?xml version="1.0" encoding="UTF-8"?>
	<Response>
		<Play>%s</Play>
	</Response>`
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/xml")

	rec := r.FormValue("RecordingUrl")
	if rec == "" {
		fmt.Fprint(w, welcomeMsg)
		return
	}

	fmt.Fprintf(w, echoMsg, rec)
}

func transcribe(c context.Context, url string) ([]byte, error) {
	client := urlfetch.Client(c)
	res, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Could not fetch %v, %v", url, err)
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Fetched status: %v", res.Status)
	}

	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("Unable to read contents")
	}

	return b, nil
}
