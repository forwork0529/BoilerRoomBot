package main

import (
	//"AquaBot/packages/comPort"
	"AquaBot/packages/computer"
	"AquaBot/packages/myBot"
	"AquaBot/packages/structs"
	"AquaBot/packages/tcpServer"
	"fmt"
	"log"
	"os"
	"os/signal"
	//"path/filepath"
	"time"
)


func main(){
	bot := myBot.New(getToken(),&structs.Vars) // Создали бота передали общие переменные

	bot.Start() // Запустили бота в работу
	//input := comPort.New("COM16", 9600)  // Запустили чтение из com порта
	input := tcpServer.New()  // Запустили чтение из com порта
	computer.New(input, &structs.Vars)	// Запустили обработку общих переменных
	fmt.Println("All functions started..")
	endFunc() // функция для прерывания работы программы на ctrl + c
}


func getToken()string{


	tokenB, err := os.ReadFile("./token.txt")

	if err != nil{
		pwd, err := os.Getwd()
		tokenB, err = os.ReadFile(pwd + `/token.txt`)
			if err != nil{
				log.Fatalf("cant read token file: %v", err)
			}
	}
	return string(tokenB)
}

func endFunc(){
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<- c
	fmt.Println("the application is being terminated..")
	time.Sleep(time.Second * 2)
}
