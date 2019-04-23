package models

import (
	"strconv"
	"time"
	"github.com/go-redis/redis"
	"fmt"
)

const(
	address = ":6379"
	protocal = "tcp"
)

type Note4Redir struct{
	title string
	date int64		//uinx时间戳
	content string	//Note.Data.Val
}

var client *redis.Client = redis.NewClient(&redis.Options{
    Addr:     "localhost:6379",
    Password: "", // no password set
    DB:       0,  // use default DB
})

func SaveToRedis(uuid string,field string,val interface{})error{
	pong, err := client.Ping().Result()
	if (err!=nil)||(pong!="PONG"){
		return fmt.Errorf("SaveToRedis fail,pingpong fail:%s,result is %v",err,pong)
	}
	switch field{
	case "Title":
		err := client.HSet(uuid, "title", fmt.Sprint(val)).Err()
		if err != nil {
    		return fmt.Errorf("SaveToRedis fail,set %v as %v fail:%s",field,val,err)
		}
	case "Date":
		if temp,ok:=val.(time.Time);ok{
			err := client.HSet(uuid, "date",temp.Unix()).Err()
		if err != nil {
    		return fmt.Errorf("SaveToRedis fail,set %v as %v fail:%s",field,val,err)
		}	
		}
	case "Data":
		err := client.HSet(uuid, "content", fmt.Sprint(val)).Err()
		if err != nil {
    		return fmt.Errorf("SaveToRedis fail,set %v as %v fail:%s",field,val,err)
		}
	default :
		return fmt.Errorf("SaveToRedis fail,set %v as %v fail,wrong type",field,val)
	}
	return nil
}

func GetNote(uuid ,field string)(interface{},error){
	if uuid==""&&field==""{
		return getAllNote()
	}else if uuid==""||field==""{
		return nil,fmt.Errorf("GetNote fail,uuid(%v) or field(%v) is empty.",uuid,field)
	}
	val, err := client.HGet(uuid,field).Result()
	if err == redis.Nil {
		return nil,fmt.Errorf("GetNote fail,uuid(%v) and field(%v) is not exist.",uuid,field)
	} else if err != nil {
		return nil,fmt.Errorf("GetNote fail,uuid(%v) and field(%v) fail for reason %v.",uuid,field,err)
	}
	if field=="Date"{
		temp,_:=strconv.Atoi(val)
		return time.Unix(int64(temp),0),nil
	}
	return val,nil
}

func getAllNote()(interface{},error){
	keys, err := client.Keys("*").Result()
	if err!=nil{
		return nil,fmt.Errorf("GetNote fail,get all keys fail for %v.",err)
	}
	temp:=make(map[string]*Note)
	for _,v:=range keys{
		temp[v]=new(Note)
		temp_title,_:=GetNote(v ,"title")
		if temp_note,ok:=temp_title.(string);ok{
			temp[v].Title=temp_note
		}
		temp_date,_:=GetNote(v ,"date")
		if temp_note,ok:=temp_date.(int64);ok{
			temp[v].Date=time.Unix(temp_note,0)
		}
		temp_content,_:=GetNote(v ,"content")
		if temp_note,ok:=temp_content.(string);ok{
			temp[v].Data.Val=temp_note
		}
	}
	return temp,nil
}

func deleteInRedis(uuid string)error{
	int1, err := client.Del(uuid).Result()
	if err!=nil{
		return fmt.Errorf("Delete %v fail for %v:%v.",uuid,int1,err)
	}
	return nil
}