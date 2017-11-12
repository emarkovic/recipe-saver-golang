package sessions

import "errors"
import "net/http"
import "strings"

// these are taken and returned in the request
const authorizationHeader = "Authorization"
const schemeBearer = "Bearer "

// ErrNoSessionID is returned when a session id is not provided in the authorization header
var ErrNoSessionID = errors.New("no session ID found in " + authorizationHeader + " header")

// ErrInvalidScheme is returned when the autorization scheme is not supported
var ErrInvalidScheme = errors.New("scheme used in Authorization header is not supported")

// BeginSession creates a new session ID, saves the state to the store, adds a
// header to the response with the session ID, and returns the new session ID
// sigingKey (string) 			- used to create the digitally signed session ID
// store (Store -> redisStore)	- saves the SessionID to the store
// state (interface)			- saves other information to the store under the SessionID
// w (http.responseWriter)		- helps write to response
func BeginSession(signingKey string, store Store, state interface{}, w http.ResponseWriter) (SessionID, error) {
	sid, err := NewSessionID(signingKey)
	if err != nil {
		return InvalidSessionID, err
	}

	err = store.Save(sid, state)
	if err != nil {
		return InvalidSessionID, err
	}

	w.Header().Add(authorizationHeader, schemeBearer+sid.String())

	return sid, nil
}

// GetSessionID extracts and validates the SessionID from the request headers
func GetSessionID(r *http.Request, signingKey string, authQueryString string) (SessionID, error) {
	ah := ""
	if r != nil {
		ah = r.Header.Get(authorizationHeader)
	}

	if len(ah) == 0 {
		if len(authQueryString) == 0 {
			return InvalidSessionID, ErrNoSessionID
		}
		ah = authQueryString
	} else if !strings.HasPrefix(ah, schemeBearer) {
		return InvalidSessionID, ErrInvalidScheme
	}

	id := strings.TrimPrefix(ah, schemeBearer)
	sid, err := ValidateID(id, signingKey)
	if err != nil {
		return InvalidSessionID, err
	}

	return sid, nil
}

// GetState extracts the SessionID from the request,
// and gets the associated state from the provided store
func GetState(r *http.Request, signingKey string, store Store, state interface{}, authQueryString string) (SessionID, error) {
	sid, err := GetSessionID(r, signingKey, authQueryString)
	if err != nil {
		return sid, err
	}

	err = store.Get(sid, state)
	if err != nil {
		return sid, err
	}

	return sid, nil
}

// EndSession extracts the SessionID from the request,
// and deletes the associated data in the provided store
func EndSession(r *http.Request, signingKey string, store Store) (SessionID, error) {
	sid, err := GetSessionID(r, signingKey, "")
	if err != nil {
		return sid, err
	}

	err = store.Delete(sid)
	if err != nil {
		return sid, err
	}

	return sid, nil
}
