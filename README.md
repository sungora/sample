# sampleapp
Пример (шаблон) веб или api приложения

#### get app 
Скачиваем пример (шаблон) приложения
    
    git clone git@github.com:sungora/sampleapp.git

#### use app
Собираем приложение в режиме модуля и запускаем

    go build -mod vendor -o ../bin/sample && ../bin/sample

#### developer

Конфигурационные файлы вашего модуля-приложения
должны начинаться с названия самого моудля (sample) к кторому они относятся.

#### help
Создание и загрузка зависимостей

    go mod init sample  инициализация модуля
    go mod vendor       скачивание зависимостей

