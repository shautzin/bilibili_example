package main

import (
	"github.com/go-redis/redis"
)

func GetRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "redis.yy:6379",
		Password: "87660543",
		DB:       0,
	})
}

//func main() {
//	client := GetRedisClient()
//
//	err := client.Ping().Err()
//	if err != nil {
//		panic(err)
//	}
//
//	err = client.Set("user:1:name", "Jack Ma", 0).Err()
//	if err != nil {
//		panic(err)
//	}
//
//	result := client.Get("user:1:name").String()
//	fmt.Println(result)
//
//	client.HSet("users", "user1", "Jack Ma")
//
//	result2 := client.HGet("users", "user1")
//	fmt.Println(result2.String())
//
//	client.LPush("userIds", 1,2,3,4,5)
//
//	for {
//		result3 := client.RPop("userIds")
//		if result3.Err() == nil {
//			fmt.Println(result3.String())
//		} else {
//			break;
//		}
//	}
//
//	result4 := client.Do("get", "user:1:name")
//	fmt.Println(result4.String())
//}
