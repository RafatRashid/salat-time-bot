package cache

import (
	"telegram-bot/utils"
	"time"
)

func SetString(key string, member string, ttl time.Duration) error {
	client := GetClient()

	res := client.Set(key, member, ttl)
	if err := res.Err(); err != nil {
		return err
	}

	return nil
}

func GetString(key string) (string, error) {
	client := GetClient()

	res := client.Get(key)
	utils.LogDebug("redis.GetString: ", res.Val(), res.Err())
	return res.Val(), res.Err()
}

func GetFolderElements(key string) ([]string, error) {
	client := GetClient()

	res := client.Scan(0, key, 1000)
	elements, _ := res.Val()
	utils.LogDebug("redis.GetFolderElements: ", elements, res.Err())

	return elements, res.Err()

}

func RemoveString(key string) error {
	client := GetClient()

	res := client.Del(key)
	if err := res.Err(); err != nil {
		return err
	}

	return nil
}
