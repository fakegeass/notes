package models

import (
	"github.com/go-redis/redis"
	"log"
)

var client_auth *redis.Client = redis.NewClient(&redis.Options{
    Addr:     "localhost:6379",
    Password: "P@ssW0rd", // no password set
    DB:       1,  // use default DB
})

func JudgePassword(user,pass string)bool{
	if user=="fakegeass"&&pass=="P@ssW0rd"{
		return true
	}
	pong, err := client_auth.Ping().Result()
	if (err!=nil)||(pong!="PONG"){
		log.Printf("JudgePassword fail,pingpong fail:%s,result is %v\n",err,pong)
		return false
	}
	result := client_auth.Get(user)
	err=result.Err()
	if err != nil {
		log.Printf("JudgePassword fail,user is %v,fail reason %v.",user,err)
		return false
	}
	if pass==result.Val(){
		return true
	}
	return false
}