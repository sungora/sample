package model

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi"
	"github.com/sungora/app/connect"
)

type Access struct {
	ID     uint64  `json:"id" gorm:"primary_key"`
	Urn    string  `json:"urn" gorm:"UNIQUE_INDEX:target"`
	Method string  `json:"method" gorm:"UNIQUE_INDEX:target"`
	Name   string  `json:"name"`
	Use    bool    `json:"use"`
	Roles  []*Role `json:"roles,omitempty" gorm:"many2many:role_to_access;"`
}

// TableName определение таблицы источника обьектов
func (m *Access) TableName() string {
	return "access"
}

var access map[string]map[string]map[uint64]struct{}

// CheckAccess проверка прав на запрос для указанного списка ролей
func CheckAccess(r *http.Request, ids []uint64) bool {
	var (
		ok     bool
		path   = chi.RouteContext(r.Context()).RoutePattern()
		method = r.Method
	)
	fmt.Println(path, method)

	// если запрошен метод который не контролируется, разрешаем
	if _, ok = controlMethods[method]; ok == false {
		return true
	}
	// запрос на несуществующий адрес (ddos, spam, etc...)
	if _, ok = access[path][method]; ok == false {
		return false
	}
	for i := range ids {
		if ids[i] == 1 {
			return true
		}
		if _, ok = access[path][method][ids[i]]; ok == true {
			return true
		}
	}
	return false
}

var controlMethods = map[string]bool{
	http.MethodGet:    true,
	http.MethodPost:   true,
	http.MethodPut:    true,
	http.MethodDelete: true,
}

// LoadRouteAccess загрузка роутинга в БД и инциализация прав на него
func LoadRouteAccess(router *chi.Mux) {
	connect.GetDB().AutoMigrate(&Access{})
	access = make(map[string]map[string]map[uint64]struct{})
	connect.GetDB().Model(&Access{}).UpdateColumn("use", false)
	loadAccessRoute(router.Routes(), "")
	connect.GetDB().Where("use = ?", false).Delete(&Access{})
}

func loadAccessRoute(routes []chi.Route, path string) {
	var ac *Access
	var ok bool
	var i int
	var urn, method string
	for i, _ = range routes {
		if routes[i].SubRoutes != nil {
			loadAccessRoute(routes[i].SubRoutes.Routes(), path+strings.Replace(routes[i].Pattern, "/*", "", -1))
			continue
		}
		urn = path + strings.TrimRight(routes[i].Pattern, "/")
		//
		access[urn] = make(map[string]map[uint64]struct{})
		for method, _ = range routes[i].Handlers {
			if _, ok = controlMethods[method]; ok == false {
				continue
			}
			ac = &Access{
				Urn:    urn,
				Method: method,
				Use:    true,
			}
			connect.GetDB().
				Where(Access{Urn: ac.Urn, Method: ac.Method}).
				Assign(Access{Use: ac.Use}).
				FirstOrCreate(ac)
			access[urn][method] = make(map[uint64]struct{})

			connect.GetDB().Preload("Roles").Find(ac)
			for i, _ = range ac.Roles {
				access[urn][method][ac.Roles[i].ID] = struct{}{}
			}
		}
		if len(access[urn]) == 0 {
			delete(access, urn)
			fmt.Fprintf(os.Stderr, "not control method from urh: [%s]\n", urn)
		}
	}
}
