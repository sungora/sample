// Контроллер главной страницы
package page

import (
	"context"
	"fmt"
	"github.com/sungora/app/core"
	"sample/model/mdlusers"
	"time"
)

// GET действие по умолчанию
func SampleAll(ctx context.Context, rw *core.RW) (context.Context, *core.RW) {
	// сессия
	var count int = 10
	// if control.Session.Get("count") != nil {
	// 	count, _ = control.Session.Get("count").(int)
	// }
	// count += 1
	// control.Session.Set("count", count)

	// работа с моделью
	u := mdlusers.New(0)
	invoiceID := uint64(8697)
	u.InvoiceID = &invoiceID
	name := "Вася пупкин"
	u.Nam = &name
	u.Age = count

	rw.Variables["Header"] = "Head ControlSample"
	rw.Variables["User"] = u
	rw.TplController += "/page/sample.html"
	rw.TplLayout += "/index.html"

	time.Sleep(5 * time.Second)
	fmt.Println("OK")

	ctx, rw = core.MidlResponseDefault(ctx, rw)

	return ctx, rw
}
