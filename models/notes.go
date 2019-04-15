package models

import (
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
	return nil
}


func newUuid()string{
	date:=time.Now().Unix()
	return strconv.Itoa(int(date))
}