package services

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/singhpranshu/streak-assignment/dto"
	exception "github.com/singhpranshu/streak-assignment/errors"
)

func FindPairHandler(w http.ResponseWriter, r *http.Request) {
	var body dto.FindPairRequestDto;
	err := validateRequest(r.Body, &body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		validationError, _ :=json.Marshal(exception.CreateBadRequestRestError(err.Error()))
		w.Write([]byte(validationError))
		return
	}
	result := findPair(&body)
	jsonResult, _ := json.Marshal(*result)
	w.WriteHeader(http.StatusOK)
    w.Write(jsonResult)
}

func findPair(body *dto.FindPairRequestDto) *[][2]int {
	result := [][2]int{}
	cache := make(map[int][]int, 0)
	for firstNumberIndex, firstNumber := range *body.Numbers {
		value, ok := cache[*body.Target - firstNumber]
		if ok {
			for _, secondNumberIndex := range value {
				result = append(result, [2]int{firstNumberIndex,secondNumberIndex})
			}
		} else {
			cache[firstNumber] = []int{firstNumberIndex}
		}
	}
	return &result

}

func validateRequest(requestbody io.ReadCloser, body *dto.FindPairRequestDto) error {
	err := json.NewDecoder(requestbody).Decode(&body)
	if err != nil {
		errorField := strings.Split(err.Error(), " ")
		if len(errorField) >= 8 {
			errorKey := strings.Split(errorField[8], ".")
			if len(errorKey) >= 1 {
				return errors.New("invalid value for field " + errorKey[1])
			}
		}
		return err
		
	}
	return nil
}