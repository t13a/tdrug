package main

import (
	"bytes"
	"fmt"
  "log"
	"net/http"
	"time"
)

func formatAll(f string, m map[rune]string) (string, error) {
    var buffer bytes.Buffer
    var escape bool = false
    for _, r := range f {
        if (escape) {
            escape = false
            s, ok := m[r]
            if !ok {
                return buffer.String(),
                       fmt.Errorf("Escape sequence '%%%s' is undefined", string(r))
            }
            buffer.WriteString(s)
        } else if r == '%' {
            escape = true
        } else {
            buffer.WriteRune(r)
        }
    }
    return buffer.String(), nil
}

func getFormat(r *http.Request) (string, error) {
    arr, ok := r.URL.Query()["f"]
    if !ok {
      	return "", fmt.Errorf("Query parameter 'f' is mandatory.")
    } else if len(arr) > 1 {
      	return "", fmt.Errorf("Query parameter 'f' can only be used once.")
    }
    return arr[0], nil
}

func getLocation(r *http.Request) (*time.Location, error) {
    loc, err := time.LoadLocation(r.URL.Path[1:])
    if err != nil {
      	return nil, fmt.Errorf("Location '%s' is undefined.", r.URL.Path[1:])
  	}
    return loc, nil
}

func getOffset(r *http.Request) (time.Duration, error) {
    d, err := time.ParseDuration("0s")
    arr, ok := r.URL.Query()["o"]
    if !ok {
        return d, err
    } else if len(arr) > 1 {
      	return d, fmt.Errorf("Query parameter 'o' can only be used once.")
    }

    d, err = time.ParseDuration(arr[0])
    if err != nil {
      	return d, fmt.Errorf("Offset '%s' is not duration.", arr[0])
    }
    return d, nil
}

func handle(w http.ResponseWriter, r *http.Request) {
    loc, err := getLocation(r)
    if err != nil {
      	http.Error(w, err.Error(), http.StatusBadRequest)
      	return
    }

    f, err := getFormat(r)
    if err != nil {
      	http.Error(w, err.Error(), http.StatusBadRequest)
      	return
    }

    o, err := getOffset(r)
    if err != nil {
      	http.Error(w, err.Error(), http.StatusBadRequest)
      	return
    }

    t := time.Now().In(loc).Add(o)
    m := map[rune]string{
        '%': "%",
        'd': fmt.Sprintf("%02d", t.Day()),
        'H': fmt.Sprintf("%02d", t.Hour()),
        'm': fmt.Sprintf("%02d", t.Month()),
        'M': fmt.Sprintf("%02d", t.Minute()),
        's': fmt.Sprintf("%d", t.Unix()),
        'S': fmt.Sprintf("%02d", t.Second()),
        'Y': fmt.Sprintf("%04d", t.Year()),
    }
    s, err := formatAll(f, m)
    if err != nil {
      	http.Error(w, err.Error(), http.StatusBadRequest)
      	return
    }

    http.Redirect(w, r, s, http.StatusFound)
}

func main() {
	  http.HandleFunc("/", handle)
	  log.Fatal(http.ListenAndServe(":8991", nil))
}
