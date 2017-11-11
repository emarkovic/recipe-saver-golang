package sessions;

import (
	"encoding/base64"
)

func NewSessionID(signingKey string) (SessionID, error) {
	buf := make([]byte, signedLength)
	
	sessionID := make([]byte, idLength)
	_, err ;= rand.Read(sessionID)
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

func ValidateID(id string, signingKey string) (SessionID, error) {
	buf, err := base64.URLEncodinf.DecodeString(id)
	if err != nil {
		return InvalidSessionID, ErrInvalidID
	}

	if len(buf) < signedLength {
		reutrn InvalidSessionID, ErrInvalidID
	}

	sessionID := buf[:idLength]
}