# Создание DLL на GO

Этот репозиторий содержит простой пример создавания dll библиотеки на Go, для дальнейшего подключения в других программах.

# Чтобы создать dll, делаем следующее:

## 1) Создать проект Go
    go mod init Go_DLL_Create_Example

## 2) Вставляем следующий код:
    package main

    import "C"
    import "fmt"

    //export sayHello
    func sayHello() {
        fmt.Println("Hello world!")
    }

    func main() {}

### Важно
    func main - должен быть пустым
    Над методами, которые должны быть доступны для вызова обязательно должен быть комментарий
    //export {имя метода}

## 3) Подготовка среды
Проверяем переменные Go
    Go env

По умолчанию Go имеет настройку CGO_ENABLED=0, надо установить в 1 (проверить можно командой )
    go env -w CGO_ENABLED=1
    
Установливаем gcc
    choco install mingw -y

* Как установить choco думаю найдете сами :)

## 4) Запуск команды создания dll
    go build -buildmode=c-shared -o example_from_go.dll main.go

## 5) Проверка работы dll
### Проверка 1:
    Выпонить команды в терминале:
        Rundll32.exe .\example_from_go.dll,sayHello
    Если никакой ошибки не возникло, значит все ок.
    fmt.Println("Hello world!") - выводится на экран не будет!!!

### Проверка 2:
    Создаем файл checkDll.go и встаяляем след код
        package main

        import "C"
        import (
            "fmt"
            "syscall"
        )

        func main() {
            dll := syscall.MustLoadDLL("example_from_go.dll")
            proc := dll.MustFindProc("sayHello")
            proc.Call()
            fmt.Println("Function executed successfully.")
        }
    Вот тут уже, должно вывестить 
        Hello world!
        Function executed successfully.
