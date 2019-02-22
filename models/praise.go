package models

import (
	"MyBlog/syserror"
	"github.com/jinzhu/gorm"
)

type Praise struct {
	Model
	UserId   int
	User     User
	Key      string
	Table    string
	IsPraise bool
}
type TempPraise struct {
	Praise int
}

func UpdatePraise(table, key string, userid int) (pcnt int, err error) {
	d := db.Begin()
	defer func() {
		if err != nil {
			d.Rollback()
		} else {
			d.Commit()
		}
	}()
	//判断是否已经点赞
	var p Praise
	err = d.Model(&Praise{}).Where("`key` = ? and `table`= ? and user_id = ?", key, table, userid).Take(&p).Error
	if err == gorm.ErrRecordNotFound {
		p = Praise{
			Key:      key,
			Table:    table,
			UserId:   userid,
			IsPraise: false,
		}
	} else if err != nil {
		return 0, err
	}
	if p.IsPraise {
		return 0, syserror.HasPriseError{}
	}
	p.IsPraise = true
	if err = d.Save(&p).Error; err != nil {
		return 0, err
	}
	var (
		pp TempPraise
	)
	err = d.Table(table).Where("key=?", key).Select("praise").Scan(&pp).Error
	if err != nil {
		return 0, err
	}
	pcnt = pp.Praise + 1
	if err = d.Table(table).Where("key=?", key).Update("praise", pcnt).Error; err != nil {
		return 0, err
	}
	return pcnt, nil
}
