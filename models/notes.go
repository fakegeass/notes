package models

import (
	"fmt"
	"os"
	"strconv"
	"log"
	"encoding/json"
	"time"
)

type content struct{
	t 	string
	Val string		`json:"Content"`
}

type Note struct{
	Title string
	Date time.Time
	Uuid string
	Data content
}

var noteLink map[string]*Note

func init(){
	noteLink=make(map[string]*Note)
	if temp,err:=GetNote("","");err==nil{
		if temp_note,ok:=temp.(map[string]*Note);ok{
			noteLink=temp_note
			log.Printf("Read saved note success:%v\n",temp_note["1"])
		}
	}else{
		log.Printf("Read saved note fail for %v\n",err)
	}
	log.SetOutput(os.Stdout)
	log.Println("Notes init...")
}


func GetNotes(uuid string) *Note{
	var temp *Note
	if uuid==""{
		temp=new(Note)
		uuid=newUuid()
		noteLink[uuid]=temp
	}else{
		temp=noteLink[uuid]
	}
	return (temp)
}

func GetAll() map[string]*Note{
	return noteLink
}

func SetNotes(uuid string,data []byte) error{
	var temp *Note
	if uuid==""{
		temp=new(Note)
		uuid=newUuid()
		noteLink[uuid]=temp
	}else{
		_,ok:=noteLink[uuid]
		if !ok{
			
			temp=new(Note)
		}else{
			temp=noteLink[uuid]
		}
		
	}
	err:=json.Unmarshal(data,temp)
	if err!=nil{
		return err
	}
	noteLink[uuid]=temp
	go saveToRedis(uuid,temp)
	return nil
}

func saveToRedis(uuid string,temp *Note){
	if err:=SaveToRedis(uuid,"Title",temp.Title);err!=nil{
		log.Println(err)
	}
	if err:=SaveToRedis(uuid,"Date",temp.Date);err!=nil{
		log.Println(err)
	}
	if err:=SaveToRedis(uuid,"Data",temp.Data.Val);err!=nil{
		log.Println(err)
	}
}


func newUuid()string{
	date:=time.Now().Unix()
	return strconv.Itoa(int(date))
}

func Delete(uuid string)error{
	if _,ok:=noteLink[uuid];ok{
		go deleteForRedis(uuid)
		delete(noteLink,uuid)
		return nil
	}
	return fmt.Errorf("not exist")
}

func deleteForRedis(uuid string){
	err:=deleteInRedis(uuid)
	if err==nil{
		log.Printf("Delete %v from redis success.",uuid)
	}
}