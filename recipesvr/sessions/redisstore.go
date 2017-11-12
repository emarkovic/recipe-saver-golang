package sessions

import (
	"time"

	redis "gopkg.in/redis.v5"
)

// redisKeyPrefix is the prefix we will use for keys related to session IDs.
// this keeps session ID keys separate from other keys in the shared redis key namespace.
const redisKeyPrefix = "sid:"

// RedisStore represents a session.Store backed by redis
type RedisStore struct {
	// Redis client used to talk to redis server
	Client *redis.Client
	// Used for key expiry time in redis
	SessionDuration time.Duration
}

// NewRedisStore constructs a new RedisStore, using the provided client session duration.
// If the 'client' is nil, it will be set to redis.NewClient() pointing at a local redis instance.
// If 'sessionDuration' is negative it will be set to 'DefaultSessionDuration'
func NewRedisStore(client *redis.Client, sessionDuration time.Duration) *RedisStore {
	if client == nil {
		client = redis.NewClient(&redis.Options{
			Addr:     "127.0.0.1:6379",
			Password: "",
			DB:       0,
		})
	}

	if sessionDuration < 0 {
		sessionDuration = DefaultSessionDuration
	}

	return &RedisStore{
		Client:          client,
		SessionDuration: sessionDuration,
	}
}

// Save associates the provided state data with the provided sid in the store
func (rs *RedisStore) Save(sid SessionID, state interface{}) error {
	return nil
}

// Get retrieves the previously saved state data for the session id,
// and populates the `state` parameter with it. This will also
// reset the data's time to live in the store.
func (rs *RedisStore) Get(sid SessionID, state interface{}) error {
	return nil
}

// Delete deletes all state data associated with the session id from the store.
func (rs *RedisStore) Delete(sid SessionID) error {
	return nil
}
