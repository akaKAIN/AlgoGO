package main

import (
	"http_server"
	"time"
)

/*
 Написать два алгоритма нахождения i-го по счёту простого числа.
Без использования «Решета Эратосфена»;
Используя алгоритм «Решето Эратосфена»
*/

func main() {
	//task_1.Sieve()
	go http_server.RunServer()
	time.Sleep(time.Second * 25)
}
