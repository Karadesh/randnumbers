package main

import (
	"bufio"
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	api "grpcserverclient/api/proto"
	"log"
	"os"
	"time"
)

const address = "localhost:8080"

func main() {
	//запускаем сервер
	var conn *grpc.ClientConn
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("could not close connection: %v", err)
		}
	}(conn)

	c := api.NewRandomNumbersClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	for {
		//вечный цикл запросов и получения ответа
		log.Printf("Type 1 to generate numbers\n " +
			"Type 2 to get generated numbers\n " +
			"Type 3 to get encrypted numbers\n " +
			"Type sign and number to get numbers generated with params. Like '-2' or '+4'\n" +
			"Type 'exit' to quit")
		request := makeRequest()
		if request == "exit" {
			break
		}
		//отправка запроса
		r, err := c.GenerateRequest(ctx, &api.GenRequest{Request: request})
		if err != nil {
			log.Fatalf("could not generate request: %v", err)
		}
		//получение результата
		log.Println(r.GetResult())
		//возможность проверить зашифрованный код
		if request == "encrypt" {
			var answer string
			log.Printf("Wanna check encryption? Type yes or no")
			fmt.Scan(&answer)
			if answer == "yes" {
				log.Printf(decryptNumbers(r.GetResult()))
			}
		}
	}
}

// функция, которая дает нужный запрос серверу в зависимости от выбора пользователя
func makeRequest() string {
	var request string
	fmt.Scan(&request)
	if request == "1" {
		return "generate"
	} else if request == "2" {
		return "getGenerated"
	} else if request == "3" {
		return "encrypt"
	} else if request == "exit" {
		return "exit"
	} else {
		return request
	}
}

// дешифратор
func decryptNumbers(r string) string {
	log.Printf("Please type numbers to check encrypted numbers")
	var answer string
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	answer = in.Text()
	bytesObject := []byte(r)
	err := bcrypt.CompareHashAndPassword(bytesObject, []byte(answer))
	if err == nil {
		return "True"
	} else {
		return "False"
	}
}
