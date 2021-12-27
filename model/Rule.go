package model

import (
	"time"
)

type Rule struct {
	UUID      string `gorm:"column:uuid;type:char(32);not null;unique;primaryKey"`
	Scope     string `gorm:"column:scope;type:char(32);not null`
	Key       string `gorm:"column:key;type:varchar(256);not null`
	Name      string `gorm:"column:name;type:varchar(256);not null`
	State     int32  `gorm:"column:state;type:tinyint;not null;default:0`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Rule) TableName() string {
	return "ogm_permission_rule"
}

type RuleDAO struct {
	conn *Conn
}

func NewRuleDAO(_conn *Conn) *RuleDAO {
	conn := DefaultConn
	if nil != _conn {
		conn = _conn
	}
	return &RuleDAO{
		conn: conn,
	}
}

func (this *RuleDAO) Count() (int64, error) {
	var count int64
	err := this.conn.DB.Model(&Rule{}).Count(&count).Error
	return count, err
}

func (this *RuleDAO) Get(_uuid string) (*Rule, error) {
	var rule *Rule
	err := this.conn.DB.Model(&Scope{}).Where("uuid = ?", _uuid).First(rule).Error
	return rule, err
}

func (this *RuleDAO) Insert(_entity *Rule) error {
	return this.conn.DB.Create(_entity).Error
}

func (this *RuleDAO) Update(_entity *Rule) error {
	// 只更新非零值
	return this.conn.DB.Updates(_entity).Error
}

func (this *RuleDAO) List(_offset int64, _count int64, _scope string) (int64, []*Rule, error) {
	var entityAry []*Rule
	count := int64(0)
	db := this.conn.DB.Model(&Rule{})
	db = db.Where("scope = ?", _scope)
	if err := db.Count(&count).Error; nil != err {
		return 0, nil, err
	}
	db = db.Offset(int(_offset)).Limit(int(_count)).Order("created_at desc")
	res := db.Find(&entityAry)
	return count, entityAry, res.Error
}

func (this *RuleDAO) Search(_offset int64, _count int64, _scope string, _key string, _name string, _state int32) (int64, []*Rule, error) {
	var entityAry []*Rule
	count := int64(0)
	db := this.conn.DB.Model(&Rule{})
	db = db.Where("scope = ?", _scope)
	if "" != _key {
		db = db.Where("`key` LIKE ?", "%"+_key+"%")
	}
	if "" != _name {
		db = db.Where("`name` LIKE ?", "%"+_name+"%")
	}
	if 0 != _state {
		db = db.Where("`state` = ?", _state)
	}
	if err := db.Count(&count).Error; nil != err {
		return 0, nil, err
	}
	db = db.Offset(int(_offset)).Limit(int(_count)).Order("created_at desc")
	res := db.Find(&entityAry)
	return count, entityAry, res.Error
}

func (this *RuleDAO) Delete(_uuid string) error {
	return this.conn.DB.Where("uuid = ?", _uuid).Delete(&Rule{}).Error
}
