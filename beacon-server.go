package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "strings"
    "time"
)

var sep ="\n"

func response(w http.ResponseWriter, r *http.Request) {
    // create header and empty response

    w.Header().Add( "Access-Control-Allow-Origin", "*")
    w.Header().Add("Access-Control-Allow-Headers", 
                   "Origin, X-Requested-With, Content-Type, Accept");
    w.WriteHeader(204)

    // create string h 
    var h = fmt.Sprintf("{%s%q: %q,%s",sep,"time", time.Now().UTC(),sep)


    // add headers as json array with names
    for name, headers := range r.Header {
        for _, header := range headers {
              h = h + fmt.Sprintf("%q: %q,%s", name, header,sep)
        }
    }

    // create request body
    var bodyBytes []byte
    var err error

    if r.Body != nil {
	bodyBytes, err = ioutil.ReadAll(r.Body)
        if err != nil {
	  var m = fmt.Sprintf("Body reading error: %v",  err)
	  h = h + fmt.Sprintf("%q: %q%s}", "error:", m,sep)
	  fmt.Printf("%v",h)
	  return
       }
       defer r.Body.Close()
    }else{
	var m = fmt.Sprintf("Body is nil")
        h = h + fmt.Sprintf("%q: %q%s}", "error:", m,sep)

    }

    // it is assumed, that we have a valid body, which can be parsed as JSON
    if len(bodyBytes) > 0 {
		var prettyJSON bytes.Buffer
		if err = json.Indent(&prettyJSON, bodyBytes, "", ""); err != nil {

			// all double colons need to be escaped
			c := strings.ReplaceAll(string(bodyBytes), "\"", "\"")

			var m = fmt.Sprintf("JSON parse error: %v %s", err,c)
			h = h + fmt.Sprintf("%q: %q%s}", "error:", m,sep)

			fmt.Printf("%v",h)
			return
		}
		h = h + fmt.Sprintf(string(prettyJSON.Bytes())[2:])
    } else {
	    h = h + fmt.Sprintf("%q: %q%s}","error", "no body supplied",sep)
    }

    fmt.Printf("%v",h)
    return 

}

func main() {
    sep=""

    http.HandleFunc("/response", response)

    http.ListenAndServe(":8090", nil)
}
