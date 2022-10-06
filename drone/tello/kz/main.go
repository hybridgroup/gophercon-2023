package main

import (
    "fmt"
    term "github.com/nsf/termbox-go"
)

func main() {
    err := term.Init()
    if err != nil {
        panic(err)
    }
    defer term.Close()

    for {
        switch ev := term.PollEvent(); ev.Type {
        case term.EventKey:
            switch ev.Key {
            case term.KeyEsc:
                term.Sync()
                fmt.Println("ESC pressed")
            case term.KeyF1:
                term.Sync()
                fmt.Println("F1 pressed")
            case term.KeyInsert:
                term.Sync()
                fmt.Println("Insert pressed")
            case term.KeyDelete:
                term.Sync()
                fmt.Println("Delete pressed")
            case term.KeyHome:
                term.Sync()
                fmt.Println("Home pressed")
            case term.KeyEnd:
                term.Sync()
                fmt.Println("End pressed")
            case term.KeyPgup:
                term.Sync()
            case term.KeyArrowRight:
                term.Sync()
                fmt.Println("Arrow Right pressed")
            case term.KeySpace:
                term.Sync()
                fmt.Println("Space pressed")
            case term.KeyBackspace:
                term.Sync()
                fmt.Println("Backspace pressed")
            case term.KeyEnter:
                term.Sync()
                fmt.Println("Enter pressed")
            case term.KeyTab:
                term.Sync()
                fmt.Println("Tab pressed")

            default:
                term.Sync()
                fmt.Println("ASCII : ", ev.Ch)

            }
        case term.EventError:
            panic(ev.Err)
        }
    }
}

