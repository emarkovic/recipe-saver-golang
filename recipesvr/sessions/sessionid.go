package sessions

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
)

// InvalidSessionID represents an empty, invalid session ID
const InvalidSessionID SessionID = ""

const idLength = 32
const signedLength = idLength + sha256.Size

// SessionID represents a valid, digitally-signed sessionID
type SessionID string

// ErrInvalidID is returned when an invalid session id is passed to ValidateID
var ErrInvalidID = errors.New("Invalid Session ID")

// NewSessionID creates a new digitally signed session id using
// 'signingKey' as the HMAC signing key. Returns an error if there was
// a problem generating random bytes.
func NewSessionID(signingKey string) (SessionID, error) {
	/*
		Making a session id:
			make a byte slice of signedLength (32 + sha length), this will be the buffer 'buf'
				[] of length 32 + sha length
			make a byte slice of idLength (32), this will be the 'sessionID'
				[] of length 32
			fill the sessionID byte sclice with random numbers
				[23, 56, 19, ...] of length 32
			copy the random numbers into the buf sclice
				[23, 56, 19, ...] of length 32 + sha length
			sha265 a byte slice of the signing key, this will be the 'hash'
			add sessionID to the hash
			sum the hash + the session id
				byte slice of signing key + [23, 56, 19, ...] = b94d27b9934d3e08a52...
				in variable 'sig'
			copy the sig into buf after the id length
				[(id for 32 spaces), (sig for sha length spaces)]
			return a base 64 encoding of buf
	*/
	buf := make([]byte, signedLength)

	sessionID := make([]byte, idLength)
	_, err := rand.Read(sessionID)
	if err != nil {
		return InvalidSessionID, err
	}
	copy(buf, sessionID)

	hash := hmac.New(sha256.New, []byte(signingKey))
	hash.Write(sessionID)
	sig := hash.Sum(nil)
	copy(buf[idLength:], sig)

	return SessionID(base64.URLEncoding.EncodeToString(buf)), nil
}

// ValidateID validates the 'id' parameter using the 'signingKey' and returns
// an error if invalid, or a SignedID if valid
func ValidateID(id string, signingKey string) (SessionID, error) {
	/*
		Validating a session id:
			base64 decode the 'id' string
				result: [(random nums for 32 spaces), (sessionID sig for sha length spaces)]
			grab the session id from the first part of the buff
			grab the 'sig' from the second part of the buff
			create a new hash of the signing key
			add the session id to the hash 'buffer'
			sum up the hash
				b94d27b9934d3e08a52... as sig2
			compare sig and sig 2
	*/
	buf, err := base64.URLEncoding.DecodeString(id)
	if err != nil {
		return InvalidSessionID, ErrInvalidID
	}

	if len(buf) < signedLength {
		return InvalidSessionID, ErrInvalidID
	}

	sessionID := buf[:idLength]
	sig := buf[idLength:]
	// creating a new hash with the signing key to compare it with the passed in id
	hash := hmac.New(sha256.New, []byte(signingKey))
	// adds sessionID to the hash 'buffer'
	hash.Write(sessionID)

	sig2 := hash.Sum(nil)
	if hmac.Equal(sig, sig2) {
		return SessionID(id), nil
	}
	return InvalidSessionID, ErrInvalidID
}

// String returns a strings representation of SessionID
func (sid SessionID) String() string {
	return string(sid)
}
