package mailbox

import "time"

// Mail orm table define
type Mail struct {
	ID         int64      `xorm:"pk autoincr"`
	Title      string     `xorm:"VARCHAR(128) notnull"`
	Content    string     `xorm:"TEXT notnull"`
	Context    string     `xorm:"TEXT"`
	CreateTime *time.Time `xorm:"TIMESTAMPZ notnull created"`
}

// Status .
type Status struct {
	ID         int64      `xorm:"pk autoincr"`
	Sender     int64      `xorm:"notnull index"`
	Receiver   int64      `xorm:"notnull index"`
	Mail       int64      `xorm:"notnull index"`
	Marked     bool       `xorm:"notnull index default false"`
	UpdateTime *time.Time `xorm:"TIMESTAMPZ notnull updated"`
}

// User .
type User struct {
	ID         int64      `xorm:"pk autoincr" json:"-"`
	PushID     string     `xorm:"notnull index" json:"pushid"`
	EMail      string     `xorm:"notnull index" json:"email"`
	Context    string     `xorm:"TEXT" json:"context"`
	CreateTime *time.Time `xorm:"TIMESTAMPZ notnull created" json:"createTime"`
}
