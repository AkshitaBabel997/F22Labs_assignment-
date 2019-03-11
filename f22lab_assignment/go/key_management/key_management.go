package keymanagement

import (
	"fmt"
	"math/rand"

	"github.com/go-redis/redis"
)

var letterRunes = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

var client *redis.Client

// StartRedis ...
func StartRedis() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
}

// CreateAPIKey ...
func CreateAPIKey(user string) (string, error) {
	var runeWithKey = randStringRunes(16)
	err := client.Set(runeWithKey, user, 0).Err()
	if err != nil {
		return "", err
	}
	return runeWithKey, nil
}

// GetValueFromKey ...
func GetValueFromKey(apikey string) (string, error) {
	user, err := client.Get(apikey).Result()
	if err != nil {
		return "", nil
	}
	return user, nil
}

// DeleteKeyValuePair ...
func DeleteKeyValuePair(apikey string) {
	client.Del(apikey)
}
