package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type Message struct {
	chats   []string
	friends []string
}

func main() {
	now := time.Now()
	id := getUserByName("john")
	println(id)

	ch := make(chan *Message, 2)
	// defer close(ch) // this does not resolve deadlock issue why. defer works before go routine exits here main is blocked by for never exists
	// The goroutine that creates and manages the channel owns it, and therefore must close it. main should close it
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go getUserChats(id, ch, wg)
	go getUserFriends(id, ch, wg)

	wg.Wait()
	close(ch)
	for msg := range ch {
		log.Println(msg)
	}
	unblock()
	block()
	log.Println(time.Since(now))
}

func getUserFriends(id string, ch chan<- *Message, wg *sync.WaitGroup) {
	time.Sleep(time.Second)

	ch <- &Message{
		friends: []string{
			"john",
			"hane",
			"jane",
			"joe",
			"tiago",
		},
	}

	wg.Done()
}

func getUserChats(id string, ch chan<- *Message, wg *sync.WaitGroup) {
	time.Sleep(time.Second * 2)
	ch <- &Message{
		chats: []string{
			"john",
			"jane",
			"joe",
		},
	}

	wg.Done()
}

func getUserByName(name string) string {
	time.Sleep(time.Second * 1)

	return fmt.Sprintf("%s-2", name)
}

func unblock() {
	ch := make(chan int)
	defer close(ch)
	log.Println("Namaste")
	go func() {
		msg := <-ch
		log.Println(msg)
	}()
	ch <- 5
	log.Println("Namha")
}

func block() {
	ch := make(chan int)
	defer close(ch)

	log.Println("Namaste")
	ch <- 5
	go func() {
		msg := <-ch
		log.Println(msg)
	}()

	log.Println("Namaha")
}
