package model

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/sungora/app/connect"

	"github.com/sungora/sample/internal/model/scenario"
	"github.com/sungora/sample/internal/model/sql"
)

// Модель
type User struct {
	ID         uint64     ``
	InvoiceID  *uint64    ``
	Nam        *string    ``
	Age        int        ``
	Credit     float64    ``
	IsOnline   bool       `gorm:"not null;default:1;"`
	Status     string     `gorm:"type:enum('Актив','Пассив','Универсал');not null;default:'Пассив';"`
	Hobby      *string    `gorm:"type:set('music','sport','reading', 'stamps', 'travel');"`
	SampleJson *string    `gorm:"type:json;"`
	Address    *string    `gorm:"type:text;"`
	CreatedAt  time.Time  ``
	UpdatedAt  time.Time  ``
	DeletedAt  *time.Time ``
}

// New создание модели
func NewUser(ID uint64) *User {
	u := new(User)
	if ID == 0 {
		u.IsOnline = true
		u.Status = "Пассив"
	}
	return u
}

// TableName определение таблицы источника обьектов
func (u *User) TableName() string {
	return "users"
}

// Load загрузка модели
func (u *User) Load() error {
	return connect.GetDB().Where(u).First(u).Error
	// model.DB.Find(u, ID)
	// attr := User{}
	// attr.IsOnline = true
	// attr.Status = "Пассив"
	// return model.DB.Attrs(attr).FirstOrInit(u, *u).Error
	// if isCreate {
	// 	return model.DB.FirstOrCreate(u, *u).Error
	// } else {
	// 	attr := User{}
	// 	attr.IsOnline = true
	// 	attr.Status = "Пассив"
	// 	return model.DB.Attrs(attr).FirstOrInit(u, *u).Error
	// }
}

// Save сохранение модели
func (u *User) Save() error {
	if u.ID > 0 {
		return connect.GetDB().Save(u).Error
	} else {
		return connect.GetDB().Create(u).Error
	}
}

// Delete удаление модели
func (u *User) Delete() error {
	return connect.GetDB().Where(u).Delete(u).Error
	// if u.ID == 0 {
	// 	return errors.New("cancel delete - ID is 0")
	// }
	// err := model.DB.Delete(u).Error
	// if err == nil {
	// 	u.ID = 0
	// }
	// return err
}

// Валидация при охранении в БД

// BeforeSave функция - хук вызовется перед сохранением записи
// Также будет вызван перед хуком при создании
func (u *User) BeforeSave(scope *gorm.Scope) error {
	return nil
}

// AfterSave функция - хук вызовется после сохранения записи
// Также будет вызван после хука при создании
func (u *User) AfterSave(scope *gorm.Scope) error {
	return nil
}

// GetScenario получение сценариев модели
func GetScenario() *scenario.Config {
	return scenario.Scenario
}

// custom query

// GetListFilter получение списка пользователей
func (u *User) GetListFilter(limit int) (userList []*User, err error) {
	err = connect.GetDB().Raw(sql.Sql.GetListFilter, limit).Scan(&userList).Error
	return
}

func sampleOther() {
	// sample custom orm query
	var users []*User
	var count int
	err := connect.GetDB().
		Select("id, name").
		Table("users").
		Joins("...", "...").
		Where("...", "...").
		Group("...").
		Having("...", "...").
		Order("id ASC").
		Limit(5).
		Find(&users).
		Count(&count).Error
	fmt.Println(users, count, err)
	// sample slice one column
	var names []string
	connect.GetDB().Model(&User{}).Pluck("names", &names)
	//
}
