package sessions

import (
	"encoding/json"
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
	j, err := json.Marshal(state)
	if err != nil {
		return err
	}

	err = rs.Client.Set(sid.getRedisKey(), j, rs.SessionDuration).Err()

	return err
}

// Get retrieves the previously saved state data for the session id,
// and populates the `state` parameter with it. This will also
// reset the data's time to live in the store.
func (rs *RedisStore) Get(sid SessionID, state interface{}) error {
	pipe := rs.Client.Pipeline()
	sc := pipe.Get(sid.getRedisKey())
	pipe.Expire(sid.getRedisKey(), rs.SessionDuration)

	_, err := pipe.Exec()

	if err == redis.Nil {
		return ErrStateNotFound
	} else if err != nil {
		return err
	}

	data, err := sc.Bytes()
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, state)
	if err != nil {
		return err
	}

	return nil
}

// Delete deletes all state data associated with the session id from the store.
func (rs *RedisStore) Delete(sid SessionID) error {
	err := rs.Client.Del(sid.getRedisKey()).Err()
	if err != nil {
		return err
	}
	return nil
}

func (sid SessionID) getRedisKey() string {
	return redisKeyPrefix + sid.String()
}
