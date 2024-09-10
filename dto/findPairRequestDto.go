package dto


type FindPairRequestDto struct {
	Numbers *[]int `json:"numbers"`
	Target *int `json:"target"`
}