package main

import (
    "fmt"
    "os"
)

type Animal interface {
    Say() string
    Greet(Animal)
}

type Person struct {
}

func (p Person) Say() string {
    return "Hey there bubba!"
}

func (p Person) Greet(animalToGreet Animal) {
    fmt.Println("Hi!")
}

type Dog struct {
    age int
    breed string
    owner *Person
}

func (d Dog) Say() string {
    return "Woof woof!"
}

func (d Dog) Growl() {
    fmt.Println("Grrr!")
}

func (d *Dog) Snuggle() {
    // snuggle code...
}

func (d Dog) Sniff(animalToSniff Animal) (bool, error) {
    // sniff code...
    return true, nil
}

func (d Dog) Greet(animalToGreet Animal) {
    if _, ok := animalToGreet.(Person); ok {
        d.Snuggle()
    } else {
        friendly, err := d.Sniff(animalToGreet)
        if err != nil {
            fmt.Fprintln(os.Stderr, "Error sniffing a non-person")
        }
        if !friendly {
            d.Growl()
        }
    }
}

func main() {
    d1 := Dog{2, "shibe", &Person{}}
    d2 := Dog{3, "poodle", &Person{}}
    d2.Greet(d1)
    fmt.Println("Successfully greeted a dog.")
}
