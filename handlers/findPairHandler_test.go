package services

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/singhpranshu/streak-assignment/dto"
)

func TestFindPairHandler(t *testing.T)  {
	tt := []struct{
        input dto.FindPairRequestDto
        result [][2]int
    }{
        {dto.FindPairRequestDto{Numbers: &[]int{1,2,3,4,5}, Target: getTargetPointer(6)}, [][2]int{{3,1}, {4,0}}},
        {dto.FindPairRequestDto{Numbers: &[]int{1,2,3,4,5}, Target: getTargetPointer(40)}, [][2]int{}},
        
    }

    for tcIndex, tc := range tt {
		input, _ := json.Marshal(tc.input)
        req, err := http.NewRequest("POST", "/find-pairs", bytes.NewReader(input))
        if err != nil {
            t.Fatal(err)
        }

        rr := httptest.NewRecorder()
	
	// To add the vars to the context, 
	// we need to create a router through which we can pass the request.
		router := mux.NewRouter()
        router.HandleFunc("/find-pairs", FindPairHandler)
        router.ServeHTTP(rr, req)

        // In this case, our MetricsHandler returns a non-200 response
        // for a route variable it doesn't know about.
		var responseBody [][2]int
		json.Unmarshal([]byte(rr.Body.Bytes()), &responseBody)
        if rr.Code != http.StatusOK || len(responseBody) != len(tc.result)  {
            t.Errorf("wrong response %v for input %v", responseBody, tcIndex)
        }

		for index,v  := range responseBody {
			if v[0] != tc.result[index][0]  || v[1] != tc.result[index][1]{
				t.Errorf("wrong response %v for input index %v", responseBody, tcIndex)
				break
			}
		}
    }
	
}

func getTargetPointer(target int) *int {
	return &target
	
}