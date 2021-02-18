package redis

import (
	"encoding/json"

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

//Updates user info in cache
func UpdateUserProfile(profile *Structure.Profile, authD map[string]string) error {
	// at := time.Unix(time.Now().Add(time.Minute*30).Unix(), 0)
	// now := time.Now()
	serialized, _ := json.Marshal(profile)
	errAccess := client.Set(authD["AccessUUID"], serialized, 0).Err()
	if errAccess != nil {
		return errAccess
	}
	return nil
}

//Adds the user's token into the redis key value pair
func CreateAuth(profile *Structure.Profile, td *Structure.TokenDetails) error {
	// at := time.Unix(td.AtExpires, 0) //converting Unix to UTC(to Time object)
	// now := time.Now()
	// context := client.Context()
	// retain readability with json
	serialized, _ := json.Marshal(profile)
	errAccess := client.Set(td.AccessUuid, serialized, 0).Err()
	if errAccess != nil {
		return errAccess
	}
	return nil
}

//Fetches auth information from redis (if not found, token may have been expired)
func FetchAuth(authD map[string]string) (*Structure.Profile, error) {
	profile, err := client.Get(authD["AccessUUID"]).Result()
	var deserialized Structure.Profile
	err = json.Unmarshal([]byte(profile), &deserialized)
	if err != nil {
		return nil, err
	}
	return &deserialized, nil
}

//The function deletes the record in redis that corresponds with the uuid passed as a parameter
func DeleteAuth(givenUUID string) (int64, error) {
	deleted, err := client.Del(givenUUID).Result()
	if err != nil {
		return 0, err
	}
	return deleted, nil
}
