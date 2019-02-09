package mdlusers

// Справочник типов полей:
// Link - Связь с другой моделью, список - в форме Select
// Text - Текст - в форме Text
// Number - Целое число - в форме Text
// / Float - Число с плавающей точкой - в форме Text
// Check - Проверочное поле - в форме один элемент Checkbox
// Select - Выбор одного из многих, список - в форме Select
// Radio - Выбор одного из многих, список - в форме Radio
// Checkbox - Выбор многих из многих, список - в форме Checkbox
// SelectSize - Выбор многих из многих, список - в форме SelectSize
// Textarea - Большре текстовое поле Проверочное поле - в форме Textarea

// сценарии по работе с моделью для виджетов
func init() {
	Scenario.AdminGrid = map[string]map[string]string{
		"Nam": map[string]string{
			"Name": "ФИО",
			"Typ":  "Text",
		},
		"Age": map[string]string{
			"Name": "Возраст",
			"Typ":  "Number",
		},
		"Credit": map[string]string{
			"Name": "Кредит",
			"Typ":  "Float",
		},
	}
	Scenario.AdminForm = map[string]map[string]string{
		"InvoiceID": map[string]string{
			"Name": "Заявка",
			"Typ":  "Link",
		},
		"Nam": map[string]string{
			"Name": "ФИО",
			"Typ":  "Text",
		},
		"Age": map[string]string{
			"Name": "Возраст",
			"Typ":  "Number",
		},
		"Credit": map[string]string{
			"Name": "Кредит",
			"Typ":  "Float",
		},
		"IsOnline": map[string]string{
			"Name": "Пользователь онлайн",
			"Typ":  "Check",
		},
		"Status": map[string]string{
			"Name":    "Статус",
			"Typ":     "Select",
			"Default": "Актив",
			"NotNull": "1",
		},
		"Hobby": map[string]string{
			"Name": "Хобби",
			"Typ":  "Checkbox",
		},
		"SampleJson": map[string]string{
			"Name": "Пример поля json",
			"Typ":  "Checkbox|Radio|Select",
		},
		"Address": map[string]string{
			"Name": "Адресс",
			"Typ":  "Textarea",
		},
	}
}

type scenarioTyp struct {
	AdminGrid map[string]map[string]string
	AdminForm map[string]map[string]string
}

// хранилище пользовательских запросов
var Scenario = new(scenarioTyp)
