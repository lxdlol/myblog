package models

import (
	"fmt"
	"reflect"
)

type Note struct {
	Model
	Key     string `gorm:"unique;not null" json:"key"`
	Title   string `gorm:"type:text" json:"title"`
	Comtent string `gorm:"type:text" json:"comtent"`
	UserId  int    `json:"user_id"`
	User    User   `json:"user"`
	Summary string `gorm:"type:text" json:"summary"`
	Visit   int    `gorm:"default:0" json:"visit"`
	Praise  int    `gorm:"default:0"json:"praise"`
}

func QuerryByKey(key string) (note Note, err error) {
	//redisdb = RedisDb()
	//defer redisdb.Close()
	//reply, e := redis.String(redisdb.Do("get", key))
	//if e != nil {
	err = db.Preload("User").Where("key = ?", key).Take(&note).Error
	//	bytes, _ := json.Marshal(note)
	//	redisdb.Do("set", key, string(bytes))
	//	return note, err
	//}
	//err = json.Unmarshal([]byte(reply), &note)
	return note, err
}
func SaveNote(note *Note) error {

	return db.Save(note).Error
}
func QuerrByPageTitle(page, limit int, title string) (note []Note, err error) {
	return note, db.Where("title like ?", fmt.Sprintf("%%%s%%", title)).Order("Updated_at desc").Offset((page - 1) * limit).Limit(limit).Find(&note).Error
}
func QuerryCount(title string) (count int, err error) {
	return count, db.Model(&Note{}).Where("title like ?", fmt.Sprintf("%%%s%%", title)).Count(&count).Error
}
func QuerryByKeyAndUserId(key string, userid int) (note Note, err error) {
	return note, db.Where("key = ? and user_id = ?", key, userid).Take(&note).Error
}
func DeleteByKeyAndId(key string, id int) error {
	return db.Delete(&Note{}, "key = ? and user_id = ?", key, id).Error
}

func StructToMap(obj interface{}) map[string]interface{} {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		data[obj1.Field(i).Name] = obj2.Field(i).String()
	}
	return data

}

//func RedisInit() (*redis.Client, error) {
//	redisdb := redis.NewClient(&redis.Options{
//		Addr:     "127.0.0.1:6379",
//		Password: "",
//		DB:       0,
//		PoolSize: 10,
//	})
//	_, err := redisdb.Ping().Result()
//	if err != nil {
//		return nil, err
//	} else {
//		return redisdb, nil
//	}
//}

//toMap := StructToMap(note)
//client.HMSet(key, toMap)
////beego.Info(1)
////beego.Info(toMap)
//f := map[string]interface{}
//f["id"] = note.ID
//f["created_at"] = note.CreatedAt
//f["updated_at"] = note.UpdatedAt
//f["deleted_at"] = note.DeletedAt
//f["key"] = note.Key
//f["title"] = note.Title
//f["comtent"] = note.Comtent
//f["user_id"] = note.UserId
//f["user"] = note.User
//f["summary"] = note.Summary
//f["visit"] = note.Visit
//f["praise"] = note.Praise
//beego.Info(f)
//client.HMSet(key, f)
//beego.Info(err.Error())
//beego.Info(note)
//}
//json.Unmarshal([]byte(i), &note)
//err = mapstructure.Decode(result, &note)
//return note, err
//e := mapstructure.Decode(result, &note)
//beego.Info(e.Error())
//return note, e
