package database

import (
	"time"
)

// DBValidateAccount 验证用户权限
func (d *DataBaseMgr) DBValidateAccount(id uint64, pwd string) int {
	var user DBUser
	err := d.accountDao.Get(&user, "select * from user where id=?", id)
	if err != nil || user.PWD != pwd {
		return 0
	}
	// 不要用PermissionType，注意解耦
	return int(user.Permission)
}

// DBCreateAccount 创建用户
func (d *DataBaseMgr) DBCreateAccount(id uint64, pwd string, permission uint8) bool {
	_, err := d.accountDao.Exec("insert into user(id,pwd,permission) values(?,?,?)", id, pwd, permission)
	if err != nil {
		return false
	}
	return true
}

// DBInsertContent 插入内容
func (d *DataBaseMgr) DBInsertContent(name string, tag string, data string) uint64 {
	// ID 自增长
	result, err := d.contentDao.Exec("insert into content(name,tag,data) values(?,?,?)", name, tag, data)
	if err != nil {
		return -1
	}
	id, err := result.LastInsertId()
	if err != nil {
		return -1
	}
	return uint64(id)
}

// DBInsertContentTag 插入内容标签
func (d *DataBaseMgr) DBInsertContentTag(tag string, contentID uint64) bool {
	_, err := d.contentDao.Exec("insert into content_tag(tag,content_id) values(?,?)", tag, contentID)
	if err != nil {
		return false
	}
	return true
}

// DBExInsertContent 高级插入内容
func (d *DataBaseMgr) DBExInsertContent(name string, tags []string, content string) bool {
	if len(tags) == 0 {
		tags = append(tags, "无标签")
	}
	tagParse := ""
	for _, tag := range tags {
		tagParse += tag + ","
	}
	tagParse = tagParse[:len(tagParse)-1]
	contentID := d.DBInsertContent(name, tagParse, content)
	if contentID == -1 {
		return false
	}
	for _, tag := range tags {
		d.DBInsertContentTag(tag, contentID)
	}
	return true
}

// DBGetContentFromTime 根据时间获取内容
func (d *DataBaseMgr) DBGetContentFromTime(beginTime time.Time, endTime time.Time) []*DBContent {
	var contents []*DBContent
	err := d.contentDao.Select(&contents, "select * from content where time>=? and time<=?", beginTime, endTime)
	if err != nil {
		return nil
	}
	return contents
}

// DBGetContentYears 获取内容年份
func (d *DataBaseMgr) DBGetContentYears(tag string) []int {
	var query string
	if tag == "" {
		query = "select distinct year(time) from content"
	} else {
		query = "select distinct year(time) from content full join content_tag on content_tag.content_id=content.id where tag=?"
	}
	var years []int
	var err error
	if tag == "" {
		err = d.contentDao.Select(&years, query)
	} else {
		err = d.contentDao.Select(&years, query, tag)
	}
	if err != nil {
		return nil
	}
	return years
}

// DBGetContentMonthFromYear 获取当前年的所有月份
func (d *DataBaseMgr) DBGetContentMonthFromYear(year int, tag string) []int {
	var query string
	if tag == "" {
		query = "select distinct month(time) from content where year(time)=?"
	} else {
		query = "select distinct month(time) from content full join content_tag on content_tag.content_id=content.id where tag=? and year(time)=?"
	}
	var months []int
	var err error
	if tag == "" {
		err = d.contentDao.Select(&months, query, year)
	} else {
		err = d.contentDao.Select(&months, query, tag, year)
	}
	if err != nil {
		return nil
	}
	return months
}


// DBGetContentDayFromYearMonth 获取当前年月的所有日期
func (d *DataBaseMgr) DBGetContentDayFromYearMonth(year int, month int, tag string) []int {
	var query string
	if tag == "" {
		query = "select distinct day(time) from content where year(time)=? and month(time)=?"
	} else {
		query = "select distinct day(time) from content full join content_tag on content_tag.content_id=content.id where tag=? and year(time)=? and month(time)=?"
	}
	var days []int
	var err error
	if tag == "" {
		err = d.contentDao.Select(&days, query, year, month)
	} else {
		err = d.contentDao.Select(&days, query, tag, year, month)
	}
	if err != nil {
		return nil
	}
	return days
}

// DBGetContentFromYearMonthDay 获取当前年月日的所有内容
func (d *DataBaseMgr) DBGetContentFromYearMonthDay(year int, month int, day int, tag string) []*DBContent {
	var query string
	if tag == "" {
		query = "select * from content where year(time)=? and month(time)=? and day(time)=?"
	} else {
		query = "select * from content full join content_tag on content_tag.content_id=content.id where tag=? and year(time)=? and month(time)=? and day(time)=?"
	}
	var contents []*DBContent
	var err error
	if tag == "" {
		err = d.contentDao.Select(&contents, query, year, month, day)
	} else {
		err = d.contentDao.Select(&contents, query, tag, year, month, day)
	}
	if err != nil {
		return nil
	}
	return contents
}

// DBGetContent 获取内容
func (d *DataBaseMgr) DBGetContent(id uint64) *DBContent {
	var content DBContent
	err := d.contentDao.Get(&content, "select * from content where id=?", id)
	if err != nil {
		return nil
	}
	return &content
}

// DBGetContents 获取内容
func (d *DataBaseMgr) DBGetContents(ids []uint64) []*DBContent {
	var contents []*DBContent
	err := d.contentDao.Select(&contents, "select * from content where id in (?)", ids)
	if err != nil {
		return nil
	}
	return contents
}
