// Контроллер главной страницы
package api

import (
    "strconv"

    "github.com/sungora/app/core"
    "github.com/sungora/app/lg"
    "sample/config"
    "sample/model"
    "sample/model/mdlusers"
)

func NewControlModel() core.Controller {
    return new(ControlModel)
}

type ControlModel struct {
    core.ControllerApi
}

// GET действие по умолчанию
func (control *ControlModel) GET() {
    var err error
    // сессия
    var count int
    if control.Session.Get("count") != nil {
        count, _ = control.Session.Get("count").(int)
    }
    count += 1
    control.Session.Set("count", count)

    // работа с моделью
    u := mdlusers.New(10)
    control.Data = u

    model.DB.AutoMigrate(u)

    if err = u.Load(); err != nil {
        lg.Error(err)
    }

    invoiceID := uint64(1234)
    u.InvoiceID = &invoiceID
    u.Age = count
    u.IsOnline = true
    u.Status = "Пассив"

    if err = u.Save(); err != nil {
        lg.Error(err)
        return
    }
}

// POST работа с моделью создание
func (control *ControlModel) POST() {
    var err error
    u := mdlusers.New(0)
    control.Data = u

    invoiceID := uint64(8697)
    u.InvoiceID = &invoiceID
    name := "Вася пупкин"
    u.Nam = &name
    u.Age = 67
    json := `{"ServiceID": 24,"ClientID": 24,"InvoiceID": 24,"BillingCode": "NL","LocationCode": "NL"}`
    u.SampleJson = &json

    if err = u.Save(); err != nil {
        lg.Error(err)
        return
    }
}

// PUT работа с моделью изменение
func (control *ControlModel) PUT() {
    var err error
    u := mdlusers.New(0)
    control.Data = u

    if err = control.RW.RequestBodyDecodeJson(u); err != nil {
        lg.Error(err)
        return
    }

    if err = u.Save(); err != nil {
        lg.Error(err)
        return
    }
}

// DELETE работа с моделью удаление
func (control *ControlModel) DELETE() {
    var err error
    var ID uint64
    ID, _ = strconv.ParseUint(control.RW.RequestParams["userID"][0], 10, 64)
    u := mdlusers.New(ID)
    control.Data = u

    if err = u.Delete(); err != nil {
        lg.Error(err)
        return
    }
}

// OPTIONS получение различных опций для текущего контроллера
func (control *ControlModel) OPTIONS() {
    // Выводимые данные
    control.Data = map[string]interface{}{
        "scenario form": mdlusers.Scenario.AdminForm, // сценарии для виджетов
        "config access": config.SampleCustomConfig,   // какая-то конфигурация
    }
}
