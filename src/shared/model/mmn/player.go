package mmn

import (
	"fmt"

	"github.com/eframework-cn/EP.GO.CORE/xorm"

	"eframe/src/shared/model/mdf"
)

type Player struct {
	xorm.Table `orm:"-" json:"-"`
	ID         int    `orm:"column(id);pk"`
	Account    string `orm:"column(account)"`  // 账号
	Password   string `orm:"column(password)"` // 密码
	Online     int    `orm:"column(online)"`   // 0-不在线 1-在线
	ConnUrl    string `orm:"column(conn_url)"` // 连接地址
	ConnID     int64  `orm:"column(conn_id)"`  // 连接ID
}

func RegPlayer() *Player {
	this := NewPlayer()
	xorm.RegisterModel(this)
	return this
}

func NewPlayer() *Player {
	this := new(Player)
	this.CTOR(this)
	return this
}

func (this *Player) DBName() string {
	return mdf.TYPE_DB_MAIN
}

func (this *Player) TableName() string {
	return "player"
}

func (this *Player) RedisName() string {
	if this.ID == xorm.ORM_INT_UNDEFINED {
		return fmt.Sprintf("%v_%v", this.DBName(), this.TableName())
	} else {
		return fmt.Sprintf("%v_%v_%v", this.DBName(), this.TableName(), this.ID)
	}
}

func (this *Player) RedisPrefix() string {
	return fmt.Sprintf("%v_%v", this.DBName(), this.TableName())
}
