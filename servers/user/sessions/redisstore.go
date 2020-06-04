package sessions

import (
	"encoding/json"
	"time"

	"github.com/go-redis/redis"
)

// RedisStore represents a session.Store backed by redis.
type RedisStore struct {
	//Redis client used to talk to redis server.
	Client *redis.Client
	//Used for key expiry time on redis.
	SessionDuration time.Duration
}

// NewRedisStore constructs a new RedisStore
func NewRedisStore(client *redis.Client, sessionDuration time.Duration) *RedisStore {
	// Initialize and return a new RedisStore struct
	return &RedisStore{
		Client: client, SessionDuration: sessionDuration,
	}

}

//Store implementation

/**
Save saves the provided `sessionState` and associated SessionID to the store.
The `sessionState` parameter is typically a pointer to a struct containing
all the data you want to associated with the given SessionID.
*/
func (rs *RedisStore) Save(sid SessionID, sessionState interface{}) error {

	// Marshal the `sessionState` to JSON and save it in the redis database
	data, err := json.Marshal(sessionState)
	if nil != err {
		return err
	}

	// Set the duration
	rs.Client.Set(sid.getRedisKey(), data, rs.SessionDuration)
	return nil
}

/**
Get populates `sessionState` with the data previously saved
for the given SessionID
*/
func (rs *RedisStore) Get(sid SessionID, sessionState interface{}) error {

	// Get the previously-saved session state data from redis
	key := sid.getRedisKey()

	pipe := rs.Client.Pipeline()
	defer pipe.Close()

	data := pipe.Get(key)
	pipe.Expire(key, rs.SessionDuration)

	_, err := pipe.Exec()
	if err == redis.Nil {
		return ErrStateNotFound
	} else if err != nil {
		return err
	}

	state, err := data.Result()
	if err != nil {
		return err
	}

	// Unmarshal it back into the `sessionState` parameter
	err = json.Unmarshal([]byte(state), sessionState)
	if err != nil {
		return err
	}

	return nil
}

/**
Delete deletes all state data associated with the SessionID from the store
*/
func (rs *RedisStore) Delete(sid SessionID) error {
	rs.Client.Del(sid.getRedisKey())
	return nil
}

/**
getRedisKey() returns the redis key to use for the SessionID
*/
func (sid SessionID) getRedisKey() string {
	//convert the SessionID to a string and add the prefix "sid:" to keep
	//SessionID keys separate from other keys that might end up in this
	//redis instance
	return "sid:" + sid.String()
}
