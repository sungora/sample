### Структура проекта

    /assets     статичные данные проекта (выкладываются на сервер)
    /cmd        точки входа в приложения (main, golang)
    /data       служебные данные для ведения проекта
    /docs       всевозможная документация на проект
    /internal   код проекта (golang)
      /model    модели
      /api      контроллеры обработчики запросов
      /middl    middleware промежуточные обработчики запросов
      /worker   воркеры приложения
      /config   конфигурация приложения
    /migration  миграции БД
    /pkg        вспомогательный модульный код проекта (golang)
      /modname  модуль приложения
      ...

### Структура репозитория и деплой

Ветки

	master  бой
	stage   релиз кандидат
	dev     разработка

Регламент деплоя:

<pre>
Бой
http://domain.ru/
Выкладывается по обнаружению пуша и нового тега 

Релиз кандидат
http://stage.domain.ru/
Выкладывается по обнаружению пуша в ветку stage и из ветки stage

Разработка
http://dev.domain.ru/
Выкладывается по обнаружению пуша в ветку dev и из ветки dev
</pre>

**Информация по миграциям**

Информация по версии БД (номер актуальной миграции) будет содержаться в метаданных тега
*  пример: `v2.0.0+migrate-YYYYMMDDhhmmss`

### Миграция БД
библиотека: https://github.com/webnice/migrate
* миграция пишется в обе стороны и строго на локальной БД.
* вспомогательный инструментарий в make

### Документирование api
библиотека: https://github.com/swaggo/swag#api-operation

Описание документирования api:
<pre>
//+funcName godoc
//+@Summary Авторизация пользователя по логину и паролю (ldap).     пишем кратко о чем речь и что принимает на входе
// @Description Возвращается токен авторизации и пользователья      пишем что возвращает и возможно подробности
//+@Tags tagName                                                    группировка api запросов
//+@Router /page/page [post]                                        относительный роутинг от базового и метод
//+@Param name TARGET TYPE-PARAM true "com"                         входящие параметры
//+@Success 200 {TYPE} string "com"                                 положительный ответ
//+@Failure 400 {TYPE} request.Error "com"                          отрицательный ответ
//+@Failure 401 {TYPE} request.Error "user unauthorized"            пользователь не авторизован
// @Accept json                                                     тип принимаемых данных
// @Produce json                                                    тип возвращаемых данных
// @Security ApiKeyAuth                                             запрос авторизованный по ключу или токену
</pre>

<pre>
+ Обязательные теги и теги по контексту (параметров может и не быть...)
TARGET      = header | path | query  | body | formData
TYPE-PARAM  = string | int  | number | bool | file | userGolangStruct
TYPE        = string | int  | number | bool | file | object | array
</pre>

Пример:
<pre>
// Login авторизация пользователя по логину и паролю ldap
// @Summary авторизация пользователя по логину и паролю (ldap).
// @Description возвращается токен авторизации
// @Tags Auth
// @Router /auth/login [post]
// @Param credentials body models.Credentials true "реквизиты доступа"
// @Success 200 {string} string "успешная авторизация"
// @Failure 400 {object} request.Error "operation error"
// @Failure 401 {object} request.Error "unauthorized"
// @Failure 403 {object} request.Error "forbidden"
// @Failure 404 {object} request.Error "not found"
// @Accept json                                                    
// @Produce json                                                   
// @Security ApiKeyAuth
</pre>

Формирование документации в make

Документация доступа на сервере после выкладки по адресу `/api/vN/index.html`

Проблемы:

* Не умеет работать со встроенными типами.
* Не умеет работать с алиасами в импортах.
* Типы slice, map не поддерживаются нужно оборачивать в отдельные типы
