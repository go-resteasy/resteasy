package client

import (
	"fmt"
	"testing"

	"github.com/bmizerany/assert"
)

// func TestGet(t *testing.T) {
// 	res := &Response{}
// 	data := make(map[string]string)
// 	res.ReadEntity(data)
// }
func TestNewClient(t *testing.T) {
	target := RestClient.Target("http://172.17.122.124:9000/api/v1/location/regions")
	res, err := target.NewRequest().Get()
	if err != nil {
		//todo
	}
	rp := ReturnPackage{}
	res.ReadEntity(&rp)
	defer res.Body.Close()
	assert.Equal(t, "1000", rp.Code)
	fmt.Printf("data:%v\n", rp)

}

func TestPost(t *testing.T) {
	target := RestClient.Target("http://111.9.116.133:10080/services/app/v1/smscode")
	data := make(map[string]string)
	data["phone"] = "18502842115"
	resp, err := target.NewRequest().Post(NewEntity(&data, "application/json"))
	if err != nil {
		//todo
	}
	defer resp.Body.Close()
	rp := ReturnPackage{}
	resp.ReadEntity(&rp)
	assert.Equal(t, "1016", rp.Code)
	fmt.Printf("data:%v\n", rp)
}

type Region struct {
	Code string
	Name string
}
type ReturnPackage struct {
	Code string
	Data []Region
	Msg  string
	Des  string
}
