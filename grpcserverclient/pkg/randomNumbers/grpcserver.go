package randomNumbers

import (
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	api "grpcserverclient/api/proto"
	"math/rand"
	"strconv"
	"strings"
)

type Values struct {
}

// Server ...
type Server struct {
	api.UnimplementedRandomNumbersServer
}

func (s *Server) Generate(ctx context.Context, req *api.GenRequest) (*api.GenResponse, error) {
	return &api.GenResponse{Result: req.GetRequest()}, nil
}
func (s *Server) GenerateCrypto(ctx context.Context, req *api.CryptRequest) (*api.CryptResponse, error) {
	return &api.CryptResponse{Result: req.GetRequest()}, nil
}

func (s *Server) GenerateNumbers(ctx context.Context, req *api.GenNumRequest) (*api.GenNumResponse, error) {
	return &api.GenNumResponse{Result: req.GetUserNumbers() + req.GetUserCommand()}, nil
}

func (s *Server) SendNumbers(ctx context.Context, req *api.SendNumRequest) (*api.SendNumResponse, error) {
	return &api.SendNumResponse{Response: req.GetRequest()}, nil
}

func generateHash(strObject string) string {
	bytesObject := []byte(strObject)
	encryptedObject, _ := bcrypt.GenerateFromPassword(bytesObject, bcrypt.MinCost)
	return string(encryptedObject)
}

func generatedNumbers(command string) string {
	var generatedNumbers string
	var encryptedNumbers string
	var answer string
	var genError string
	generatedNumbersLink := &generatedNumbers
	generateEncryptedNumbersLink := &encryptedNumbers
	if command == "generate" {
		*generatedNumbersLink = generateNumbers()
		answer = *generatedNumbersLink
	} else if command == "getGenerated" {
		if *generatedNumbersLink == "" {
			*generatedNumbersLink = generateNumbers()
		}
		answer = *generatedNumbersLink
	} else if command == "encrypt" {
		if *generatedNumbersLink == "" {
			*generatedNumbersLink = generateNumbers()
		}
		*generateEncryptedNumbersLink = generateHash(*generatedNumbersLink)
		answer = *generateEncryptedNumbersLink
	} else {
		answer, genError = generateWithParams(command)
		if genError != "" {
			*generatedNumbersLink = answer
			answer = fmt.Sprintf("Generated numbers: %s; Numbers with params: %s", genError, answer)
		}
	}
	return answer
}

func generateNumbers() string {
	var randSlice []string
	for i := 0; i < rand.Intn(10); {
		m := rand.Intn(100)
		mString := strconv.Itoa(m)
		randSlice = append(randSlice, mString)
	}
	return strings.Join(randSlice, " ")
}

func generateWithParams(strObject string) (string, string) {
	numbers := generateNumbers()
	numSlice := strings.Split(numbers, " ")
	var intSlice []int
	var answer []string
	var stringAnswer string
	var numberString string
	var numberInt int
	var err error
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
	}
	for _, value := range intSlice {
		answer = append(answer, strconv.Itoa(value))
	}
	stringAnswer = strings.Join(answer, " ")
	return stringAnswer, numbers
}
