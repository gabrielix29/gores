package gores

import (
	"encoding/json"
	"net/http"
)

// HTTP Methods
const (
	CONNECT = "CONNECT"
	DELETE  = "DELETE"
	GET     = "GET"
	HEAD    = "HEAD"
	OPTIONS = "OPTIONS"
	PATCH   = "PATCH"
	POST    = "POST"
	PUT     = "PUT"
	TRACE   = "TRACE"
)

// Media Types
const (
	ApplicationJSON            = "application/json"
	ApplicationJSONCharsetUTF8 = ApplicationJSON + "; " + CharsetUTF8
	ApplicationForm            = "application/x-www-form-urlencoded"
	TextHTML                   = "text/html"
	TextHTMLCharsetUTF8        = TextHTML + "; " + CharsetUTF8
	TextPlain                  = "text/plain"
	TextPlainCharsetUTF8       = TextPlain + "; " + CharsetUTF8
	MultipartForm              = "multipart/form-data"
)

// Headers
const (
	AcceptEncoding     = "Accept-Encoding"
	Authorization      = "Authorization"
	ContentDisposition = "Content-Disposition"
	ContentEncoding    = "Content-Encoding"
	ContentLength      = "Content-Length"
	ContentType        = "Content-Type"
	Location           = "Location"
	Upgrade            = "Upgrade"
	Vary               = "Vary"
	WWWAuthenticate    = "WWW-Authenticate"
	XForwardedFor      = "X-Forwarded-For"
	XRealIP            = "X-Real-IP"
)

// Charsets
const (
	// CharsetUTF8 utf8 character set
	CharsetUTF8 = "charset=utf-8"
)

func JSON(w http.ResponseWriter, code int, i interface{}) error {
	b, err := json.Marshal(i)
	if err != nil {
		return err
	}
	_json(w, code, b)
	return nil
}

func _json(w http.ResponseWriter, code int, b []byte) {
	w.Header().Set(ContentType, ApplicationJSONCharsetUTF8)
	w.WriteHeader(code)
	_, _ = w.Write(b)
}
