package magicgate

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", handler)
}

const content = `<?xml version="1.0" encoding="UTF-8"?>
	<Response>
		<Say>Hello Barry!</Say>
		<Play digits="9"></Play>
	</Response>`

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/xml")
	fmt.Fprint(w, content)
}
