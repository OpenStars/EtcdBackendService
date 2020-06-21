package main

import (
	"log"

	"github.com/OpenStars/EtcdBackendService/Tile38ClientService"
	"github.com/OpenStars/GoEndpointManager/GoEndpointBackendManager"
)

func main() {
	client := Tile38ClientService.NewTile38ManagerServiceModel("location", "", []string{},
		GoEndpointBackendManager.EndPoint{
			ServiceID: "testtile38",
			Host:      "10.110.1.21",
			Port:      "9851",
		})
	err := client.DropAll()
	log.Println("drop all err", err)
	// client.SetLocationItemToTile38(1, 21.044946, 105.801253, map[string]interface{}{
	// 	"cat1": 5,
	// 	"cat2": 7,
	// })
	// client.SetLocationItemToTile38(2, 21.061043, 105.801494, map[string]interface{}{
	// 	"cat1": 5,
	// 	"cat2": 6,
	// })

	// client.SetLocationItemToTile38(3, 21.061043, 105.801584, map[string]interface{}{
	// 	"cat1": 6,
	// 	"cat2": 8,
	// })

	// r, err := client.GetLocationItemNearby(21.061043, 105.801494, 6, map[string][2]interface{}{
	// 	"cat1": [2]interface{}{5, 5},
	// }, 0, 10)
	// if err != nil {
	// 	log.Println("err", err)
	// 	return
	// }
	// log.Println("r", r)

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
	// rs, err := c.GetLocationItemNearby(1, 1, 1000, nil, 0, 10)
	// log.Println("rs", rs, "err", err)
	// // c.DeleteLocationInTile38(1)
	// fmt.Println(c.GetLocationInTile38("03f2c0e35fa86722c4e0533a37e35d5b1c73771f660f1d9081e6205fbbb950402d"))

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
