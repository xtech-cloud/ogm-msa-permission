package model

import (
	"time"
)

type Scope struct {
	UUID      string `gorm:"column:uuid;type:char(32);not null;unique;primaryKey"`
	Key       string `gorm:"column:key;type:varchar(256);not null;unique`
	Name      string `gorm:"column:name;type:varchar(256);not null`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Scope) TableName() string {
	return "ogm_permission_scope"
}

type ScopeDAO struct {
	conn *Conn
}

func NewScopeDAO(_conn *Conn) *ScopeDAO {
	conn := DefaultConn
	if nil != _conn {
		conn = _conn
	}
	return &ScopeDAO{
		conn: conn,
	}
}

func (this *ScopeDAO) Count() (int64, error) {
	var count int64
	err := this.conn.DB.Model(&Scope{}).Count(&count).Error
	return count, err
}

func (this *ScopeDAO) Get(_uuid string) (*Scope, error) {
	var scope *Scope
	err := this.conn.DB.Model(&Scope{}).Where("uuid = ?", _uuid).First(scope).Error
	return scope, err
}

func (this *ScopeDAO) Insert(_entity *Scope) error {
	return this.conn.DB.Create(_entity).Error
}

func (this *ScopeDAO) Update(_entity *Scope) error {
	// 只更新非零值
	return this.conn.DB.Updates(_entity).Error
}

func (this *ScopeDAO) List(_offset int64, _count int64) (int64, []*Scope, error) {
	var entityAry []*Scope
	count := int64(0)
	db := this.conn.DB.Model(&Scope{})
	if err := db.Count(&count).Error; nil != err {
		return 0, nil, err
	}
	db = db.Offset(int(_offset)).Limit(int(_count)).Order("created_at desc")
	res := db.Find(&entityAry)
	return count, entityAry, res.Error
}

func (this *ScopeDAO) Search(_offset int64, _count int64, _key string, _name string) (int64, []*Scope, error) {
	var entityAry []*Scope
	count := int64(0)
	db := this.conn.DB.Model(&Scope{})
	if _key != "" {
		db = db.Where("`key` LIKE  ?", "%"+_key+"%")
	}
	if _name != "" {
		db = db.Where("`name` LIKE  ?", "%"+_name+"%")
	}
	if err := db.Count(&count).Error; nil != err {
		return 0, nil, err
	}
	db = db.Offset(int(_offset)).Limit(int(_count)).Order("created_at desc")
	res := db.Find(&entityAry)
	return count, entityAry, res.Error
}

func (this *ScopeDAO) Delete(_uuid string) error {
	return this.conn.DB.Where("uuid = ?", _uuid).Delete(&Scope{}).Error
}
