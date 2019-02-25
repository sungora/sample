package servhttp

// конфигурация подгружаемая из файла конфигурации
type configFile struct {
	Http configTyp
}

// конфигурация сервера
type configTyp struct {
	Proto          string // Server Proto
	Host           string // Server Host
	Port           int    // Server Port
	ReadTimeout    int    // Время ожидания web запроса в секундах, по истечении которого соединение сбрасывается
	WriteTimeout   int    // Время ожидания окончания передачи ответа в секундах, по истечении которого соединение сбрасывается
	IdleTimeout    int    // Время ожидания следующего запроса
	MaxHeaderBytes int    // Максимальный размер заголовка получаемого от браузера клиента в байтах
}
