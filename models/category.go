package models

import (
	"time"
)

//Category modelo de los chats
type Category struct {
	Id       int       `xorm:"bigint not null autoincr pk" json:"id"`
	State    int       `xorm:"int" json:"state"`
	Nombre   string    `xorm:"varchar(200) not null" valid:"required unique" json:"nombre"`
	Lang     string    `xorm:"varchar(5) not null" valid:"required" json:"lang"`
	Date     time.Time `xorm:"date" json:"date"`
	CreateAt time.Time `xorm:"created" json:"created_at"`
	UpdateAt time.Time `xorm:"updated" json:"updated_at"`
}

//GetID Obtiene el id
func (el Category) GetID() int {
	return el.Id
}

//GetState Obtiene el estado
func (el Category) GetState() int {
	return el.State
}

//SetState Establece el estado
func (el Category) SetState(state int) {
	el.State = state
}

// TableName establece el nombre de la tabla del modelo
func (el Category) TableName() string {
	return "categorys"
}
