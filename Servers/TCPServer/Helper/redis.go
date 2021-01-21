package helper

import (
	"os"
	"time"

	"github.com/go-redis/redis"

	Structure "../Structure"
)

var client *redis.Client

func init() {
	//Initializing redis
	dsn := os.Getenv("REDIS_DSN")
	if len(dsn) == 0 {
		dsn = "localhost:6379"
	}
	client = redis.NewClient(&redis.Options{
		Addr: dsn, //redis port
	})
	_, err := client.Ping(client.Context()).Result()
	if err != nil {
		panic(err)
	}
}

//Adds the user's token into the redis key value pair
func CreateAuth(username string, td *Structure.TokenDetails) error {
	at := time.Unix(td.AtExpires, 0) //converting Unix to UTC(to Time object)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()
	context := client.Context()

	errAccess := client.Set(context, td.AccessUuid, username, at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}
	errRefresh := client.Set(context, td.RefreshUuid, username, rt.Sub(now)).Err()
	if errRefresh != nil {
		return errRefresh
	}
	return nil
}

//Fetches auth information from redis (if not found, token may have been expired)
func FetchAuth(authD map[string]string) (string, error) {
	username, err := client.Get(client.Context(), authD["AccessUUID"]).Result()
	if err != nil {
		return "", err
	}
	return username, nil
}

//The function deletes the record in redis that corresponds with the uuid passed as a parameter
func DeleteAuth(givenUUID string) (int64, error) {
	deleted, err := client.Del(client.Context(), givenUUID).Result()
	if err != nil {
		return 0, err
	}
	return deleted, nil
}
