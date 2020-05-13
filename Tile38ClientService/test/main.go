package main

import (
	"fmt"

	"github.com/OpenStars/EtcdBackendService/Tile38ClientService"
	"github.com/OpenStars/GoEndpointManager/GoEndpointBackendManager"
)

func main() {
	c := Tile38ClientService.NewTile38ManagerServiceModel("location", "", []string{},
		GoEndpointBackendManager.EndPoint{
			ServiceID: "testtile38",
			Host:      "127.0.0.1",
			Port:      "9851",
		})

	// var e error

	// e = c.SetLocationItemToTile38("1", 10.11111, 10.11111, map[string]interface{}{})
	// fmt.Printf("set 1 %v \n", e)
	// e = c.SetLocationItemToTile38("2", 10.11112, 10.11111, map[string]interface{}{})
	// e = c.SetLocationItemToTile38("3", 10.11113, 10.11112, map[string]interface{}{})
	// e = c.SetLocationItemToTile38("10", 10.11111, 10.111118, map[string]interface{}{
	// 	"age":    22,
	// 	"areaid": 1,
	// })
	// fmt.Printf("set 4 %v \n", e)
	// e = c.SetLocationItemToTile38("5", 10.111115, 10.111122, map[string]interface{}{
	// 	"age":    20,
	// 	"areaid": 2,
	// })
	// e = c.SetLocationItemToTile38("6", 10.111116, 10.11111, map[string]interface{}{
	// 	"age":    18,
	// 	"areaid": 1,
	// })
	// fmt.Printf("set 6 %v \n", e)

	// fmt.Println(c.GetLocationInTile38(1))
	fmt.Println(c.GetLocationInTile38(1))

	// c.DeleteLocationInTile38(1)
	// fmt.Println(c.GetLocationInTile38(1))

	// // 10 km, lay 5 phan tu dau tien
	// fmt.Println(c.GetLocationItemNearby(10.11111, 10.11111, 10, map[string][2]interface{}{}, 1, 5))
	// fmt.Println(c.GetLocationItemNearby(10.11111, 10.11111, 10, map[string][2]interface{}{
	// 	"age": [2]interface{}{18, 20},
	// }, 1, 5))
	// fmt.Println(c.GetLocationItemNearby(10.11111, 10.11111, 10, map[string][2]interface{}{
	// 	"age":    [2]interface{}{18, 20},
	// 	"areaid": [2]interface{}{1, 1},
	// }, 1, 5))
}
