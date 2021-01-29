package redis

import (
	"time"

	"github.com/go-redis/redis"

	Structure "servers/Structure"
	Constants "servers/internal"
)

var client *redis.Client

func init() {
	//Initializing redis
	dsn := Constants.REDIS_DSN
	client = redis.NewClient(&redis.Options{
		Addr: dsn, //redis port
	})
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
}

//Adds the user's token into the redis key value pair
func CreateAuth(username string, td *Structure.TokenDetails) error {
	at := time.Unix(td.AtExpires, 0) //converting Unix to UTC(to Time object)
	now := time.Now()
	// context := client.Context()

	errAccess := client.Set(td.AccessUuid, username, at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}
	return nil
}

//Fetches auth information from redis (if not found, token may have been expired)
func FetchAuth(authD map[string]string) (string, error) {
	username, err := client.Get(authD["AccessUUID"]).Result()
	if err != nil {
		return "", err
	}
	return username, nil
}

//The function deletes the record in redis that corresponds with the uuid passed as a parameter
func DeleteAuth(givenUUID string) (int64, error) {
	deleted, err := client.Del(givenUUID).Result()
	if err != nil {
		return 0, err
	}
	return deleted, nil
}
