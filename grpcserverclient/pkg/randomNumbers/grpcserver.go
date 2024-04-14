package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	api "grpcserverclient/api/proto"
	"log"
	"math/rand"
	"net"
	"strconv"
	"strings"
)

// Server ...
type Server struct {
	api.UnimplementedRandomNumbersServer
}

func (s *Server) GenerateRequest(ctx context.Context, req *api.GenRequest) (*api.GenResponse, error) {
	log.Println("got request")
	return &api.GenResponse{Result: generatedNumbers(req.GetRequest())}, nil
}

// функция-шифровальщик
func generateHash(strObject string) string {
	bytesObject := []byte(strObject)
	encryptedObject, _ := bcrypt.GenerateFromPassword(bytesObject, bcrypt.MinCost)
	log.Printf(decryptNumbers(string(encryptedObject), strObject))
	return string(encryptedObject)
}

// функция-дешифратор
func decryptNumbers(hashedObject string, unhashedObject string) string {
	log.Printf("Please type numbers to check encrypted numbers")
	bytesObject := []byte(hashedObject)
	err := bcrypt.CompareHashAndPassword(bytesObject, []byte(unhashedObject))
	if err == nil {
		return "True"
	} else {
		return "False"
	}
}

// функция, которая вызывает другие функции исходя из запроса
func generatedNumbers(command string) string {
	var genNumbers string
	var encNumbers string
	var answer string
	var genError string
	if command == "generate" {
		genNumbers = DbGetter()
		if genNumbers == "" {
			//добавление в бд результат функции-генерации чисел
			genNumbers = DbInserter(generateNumbers())
		} else {
			genNumbers = DbUpdater(generateNumbers())
		}
		answer = "Done!"
	} else if command == "getGenerated" {
		genNumbers = DbGetter()
		if genNumbers == "" {
			genNumbers = DbInserter(generateNumbers())
			genNumbers = DbGetter()
		}
		answer = genNumbers
	} else if command == "encrypt" {
		genNumbers = DbGetter()
		if genNumbers == "" {
			genNumbers = DbInserter(generateNumbers())
			genNumbers = DbGetter()
		}
		encNumbers = generateHash(genNumbers)
		answer = encNumbers
	} else {
		answer, genError = generateWithParams(command)
		if genError != "" {
			genNumbers = DbGetter()
			if genNumbers == "" {
				genNumbers = DbInserter(answer)
			} else {
				genNumbers = DbUpdater(answer)
			}
			answer = fmt.Sprintf("Generated numbers: %s; Numbers with params: %s", genError, answer)
		}
	}
	return answer
}

// функция генерации чисел
func generateNumbers() string {
	var randSlice []string
	for i := 0; i < rand.Intn(10); {
		m := rand.Intn(100)
		mString := strconv.Itoa(m)
		randSlice = append(randSlice, mString)
	}
	return strings.Join(randSlice, " ")
}

// функция генерации с параметром от пользователя
func generateWithParams(strObject string) (string, string) {
	numbers := generateNumbers()
	numSlice := strings.Split(numbers, " ")
	var intSlice []int
	var answer []string
	var stringAnswer string
	var numberString string
	var numberInt int
	var err error
	//ищем знак по префиксу, если не находим - в реквест отправляем params are unavailable
	if strings.HasPrefix(strObject, "+") == true {
		numberString = strings.Trim(strObject, "+")
		numberInt, err = strconv.Atoi(numberString)
		if err != nil {
			return "Params are unavailable", ""
		} else {
			for _, value := range numSlice {
				intValue, _ := strconv.Atoi(value)
				intSlice = append(intSlice, intValue+numberInt)
			}
		}
	} else if strings.HasPrefix(strObject, "-") == true {
		numberString = strings.Trim(strObject, "-")
		numberInt, err = strconv.Atoi(numberString)
		if err != nil {
			return "Params are unavailable", ""
		} else {
			for _, value := range numSlice {
				intValue, _ := strconv.Atoi(value)
				intSlice = append(intSlice, intValue-numberInt)
			}
		}
	} else if strings.HasPrefix(strObject, "*") == true {
		numberString = strings.Trim(strObject, "*")
		numberInt, err = strconv.Atoi(numberString)
		if err != nil {
			return "Params are unavailable", ""
		} else {
			for _, value := range numSlice {
				intValue, _ := strconv.Atoi(value)
				intSlice = append(intSlice, intValue*numberInt)
			}
		}
	} else if strings.HasPrefix(strObject, "/") == true {
		numberString = strings.Trim(strObject, "/")
		numberInt, err = strconv.Atoi(numberString)
		if err != nil {
			return "Params are unavailable", ""
		} else if numberInt == 0 {
			return "Params are unavailable", ""
		} else {
			for _, value := range numSlice {
				intValue, _ := strconv.Atoi(value)
				intSlice = append(intSlice, intValue/numberInt)
			}
		}
	} else if strings.HasPrefix(strObject, "%") == true {
		numberString = strings.Trim(strObject, "%")
		numberInt, err = strconv.Atoi(numberString)
		if err != nil {
			return "Params are unavailable", ""
		} else if numberInt == 0 {
			return "Params are unavailable", ""
		} else {
			for _, value := range numSlice {
				intValue, _ := strconv.Atoi(value)
				intSlice = append(intSlice, intValue%numberInt)
			}
		}
	} else {
		return "Params are unavailable", ""
	}
	//Полученные значения переводим в строку для удобства получения ответа сервера
	for _, value := range intSlice {
		answer = append(answer, strconv.Itoa(value))
	}
	stringAnswer = strings.Join(answer, " ")
	return stringAnswer, numbers
}

// функция запуска сервера
func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	api.RegisterRandomNumbersServer(s, &Server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// DbInserter функция добавления в бд
func DbInserter(numbers string) string {
	db, err := sql.Open("postgres", "user=admin password=admin host=localhost dbname=randnumbers sslmode=disable")
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer db.Close()
	inserter := fmt.Sprintf("INSERT INTO randomnumbers(number_id, number_value) VALUES(0, '%s')", numbers)
	result, err := db.Exec(inserter)
	if err != nil {
		fmt.Printf("failed to insert into randomnumbers: %v", err)
	}
	affectedRows, err := result.RowsAffected()
	log.Printf("Inserted %d rows\n", affectedRows)
	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}
	return "Done!"
}

// DbGetter функция получения данных из бд
func DbGetter() string {
	db, err := sql.Open("postgres", "user=admin password=admin host=localhost dbname=randnumbers sslmode=disable")
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer db.Close()
	selector := fmt.Sprintf("SELECT number_value FROM randomnumbers WHERE number_id=0")
	someRow, err := db.Query(selector)
	if err != nil {
		return ""
	}
	var someValue string
	someValue = ""
	for someRow.Next() {
		something := someRow.Scan(&someValue)
		fmt.Println(something)
	}
	return someValue
}

// DbUpdater функция, которая обновляет данные в бд, если запись уже есть
func DbUpdater(numbers string) string {
	db, err := sql.Open("postgres", "user=admin password=admin host=localhost dbname=randnumbers sslmode=disable")
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer db.Close()
	updater := fmt.Sprintf("UPDATE randomnumbers SET number_value ='%s' WHERE number_id = 0", numbers)
	result, err := db.Exec(updater)
	if err != nil {
		fmt.Printf("failed to insert into randomnumbers: %v", err)
	}
	affectedRows, err := result.RowsAffected()
	log.Printf("Updated %d rows\n", affectedRows)
	return "Done!"
}
