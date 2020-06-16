package Tile38ClientService

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/OpenStars/EtcdBackendService/Tile38ClientService/common/util/distance"
	"github.com/OpenStars/EtcdBackendService/Tile38ClientService/data"
	"github.com/OpenStars/EtcdBackendService/Tile38ClientService/transports"
	"github.com/OpenStars/GoEndpointManager"
	"github.com/OpenStars/GoEndpointManager/GoEndpointBackendManager"
	"github.com/gomodule/redigo/redis"
)

type Tile38ManagerService struct {
	host        string
	port        string
	sid         string
	location    string
	etcdManager *GoEndpointManager.EtcdBackendEndpointManager
}

func NewTile38ManagerServiceModel(location string, serviceID string, etcdServers []string, defaultEnpoint GoEndpointBackendManager.EndPoint) Tile38ManagerServiceIf {

	tile38Client := &Tile38ManagerService{
		host:        defaultEnpoint.Host,
		port:        defaultEnpoint.Port,
		sid:         serviceID,
		location:    location,
		etcdManager: GoEndpointManager.GetEtcdBackendEndpointManagerSingleton(etcdServers),
	}

	if tile38Client.etcdManager != nil {
		err := tile38Client.etcdManager.SetDefaultEntpoint(serviceID, defaultEnpoint.Host, defaultEnpoint.Port)
		if err != nil {
			log.Println("SetDefaultEndpoint sid", serviceID, "err", err)
			return nil
		}
	}
	return tile38Client
}

func NewTile38ManagerServiceModel2(location string, sid string, etcdEndpoints []string, defaultHost, defaultPort string) Tile38ManagerServiceIf {
	tile38Client := &Tile38ManagerService{
		host:        defaultHost,
		port:        defaultPort,
		sid:         sid,
		location:    location,
		etcdManager: GoEndpointManager.GetEtcdBackendEndpointManagerSingleton(etcdEndpoints),
	}

	if tile38Client.etcdManager != nil {
		err := tile38Client.etcdManager.SetDefaultEntpoint(sid, defaultHost, defaultPort)
		if err != nil {
			log.Println("SetDefaultEndpoint sid", sid, "err", err)
			return nil
		}
	}
	return tile38Client
}

// DeleteLocationInTile38 delete location by key in tile38
func (r *Tile38ManagerService) DeleteLocationInTile38(keyLocation interface{}) (result bool, err error) {
	log.Printf("[DeleteLocationInTile38] keyLocation = %v \n", keyLocation)

	if r.etcdManager != nil {
		h, p, err := r.etcdManager.GetEndpoint(r.sid)
		if err != nil {
			log.Println("EtcdManager get endpoints", "err", err)
		} else {
			r.host = h
			r.port = p
		}
	}
	c, err := transports.GetTile38LocationClient(r.host, r.port)
	if c != nil {
		defer c.Close()
		deleted, err := c.Do("DEL", r.location, keyLocation)
		if err != nil {
			log.Printf("[DeleteLocationInTile38] service error : %v \n", err)
		}
		// process result
		if deleted.(int64) == 0 {
			return false, err
		} else {
			return true, err
		}
	} else {
		log.Printf("[DeleteLocationInTile38] Can not get client\n")
	}
	return
}

func (r *Tile38ManagerService) GetLocationInTile38(keyLocation interface{}) (result []float64, err error) {

	if r.etcdManager != nil {
		h, p, err := r.etcdManager.GetEndpoint(r.sid)
		if err != nil {
			log.Println("EtcdManager get endpoints", "err", err)
		} else {
			r.host = h
			r.port = p
		}
	}
	c, err := transports.GetTile38LocationClient(r.host, r.port)
	if c != nil {
		defer c.Close()

		ret, err := c.Do("GET", r.location, keyLocation)
		if err != nil {
			log.Printf("[GetLocationInTile38] service error : %v \n", err)
			return nil, err
		}
		var retmodel *data.LocationModel
		if ret == nil {
			return []float64{0, 0}, errors.New("Can not find location")
		}
		retString := fmt.Sprintf("%s", ret)
		err = json.Unmarshal([]byte(retString), &retmodel)
		if err != nil || r == nil || len(retmodel.Coordinates) == 0 {
			return nil, err
		}
		result = []float64{retmodel.Coordinates[1], retmodel.Coordinates[0]}
	} else {
		result = nil
		log.Printf("[GetLocationInTile38] Can not get client\n")
	}
	return
}

func (r *Tile38ManagerService) GetLocationItemNearby(lat, long, radius float64, fields map[string][2]interface{}, pageNumber, pageSize int64) (result map[string]float64, err error) {
	log.Printf("[GetLocationNearby] lat:%f long:%f radius:%f, pageNumber = %d, pageSize = %d, fields=%v \n", lat, long, radius, pageNumber, pageSize, fields)
	result = make(map[string]float64)
	if r.etcdManager != nil {
		h, p, err := r.etcdManager.GetEndpoint(r.sid)
		if err != nil {
			log.Println("EtcdManager get endpoints", "err", err)
		} else {
			r.host = h
			r.port = p
		}
	}
	c, err := transports.GetTile38LocationClient(r.host, r.port)
	if c != nil {
		defer c.Close()

		if pageNumber <= 0 || pageSize <= 0 {
			pageNumber = 1
			pageSize = 50
		}

		beginItem := (pageNumber - 1) * pageSize

		args := redis.Args{}.Add(r.location)
		for k, v := range fields {
			if len(v) < 2 {
				continue
			}
			args = args.Add("where").Add(k).AddFlat(v[0]).AddFlat(v[1])
		}
		// , "CURSOR", beginItem, "LIMIT", pageSize, "POINT", lat, long, radius*1000)
		args = args.Add("CURSOR").Add(beginItem).Add("LIMIT").Add(pageSize).Add("POINT").Add(lat).Add(long).Add(radius * 1000)
		fmt.Println(args)
		//radius m -> convert km
		vals, err := redis.Values(c.Do("NEARBY", args...))
		if err != nil {
			log.Printf("[GetLocationNearby] could not NEARBY: %v\n", err)
		}
		// the first element is the cursor
		if len(vals) < 2 {
			log.Printf("[GetLocationNearby] invalid value")
		}
		vals, err = redis.Values(vals[1], nil)
		if err != nil {
			log.Printf("[GetLocationNearby] invalid value")
		}
		for _, val := range vals {

			strs, _ := redis.Strings(val, nil)
			// if err != nil || len(strs) < 2 {
			// 	log.Printf("[GetLocationNearby] invalid value: %v\n", err)
			// }
			fmt.Printf("%s >> %s\n", strs[0], val)
			//convert result
			loc := &data.LocationModel{}
			e := json.Unmarshal([]byte(strs[1]), loc)
			if e != nil {
				continue
			}

			// log.Printf("[Location - get location nearby] = %v \n", loc)
			result[strs[0]] = distance.DistanceBetween2Points(lat, long, loc.Coordinates[1], loc.Coordinates[0])
		}

	} else {
		log.Printf("[GetLocationNearby] Can not get client\n")
	}

	return
}

func (r *Tile38ManagerService) SetLocationItemToTile38(keyLocation interface{}, lat, lng float64, fields map[string]interface{}) (err error) {
	log.Printf("SetLocationItemToTile38 keyLocation: %v, %f, %f, fields: %v\n", keyLocation, lat, lng, fields)
	if r.etcdManager != nil {
		h, p, err := r.etcdManager.GetEndpoint(r.sid)
		if err != nil {
			log.Println("EtcdManager get endpoints", "err", err)
		} else {
			r.host = h
			r.port = p
		}
	}
	c, err := transports.GetTile38LocationClient(r.host, r.port)
	if c != nil {
		defer c.Close()

		// fieldString := ""
		args := redis.Args{}.Add(r.location).Add(keyLocation)
		for k, v := range fields {
			//"FIELD", "categoryid", categoryid
			args = args.Add("FIELD").AddFlat(k).AddFlat(v)
		}

		// _, err = c.Do("SET", r.location, keyLocation, "POINT", lat, lng)

		args = args.Add("POINT").AddFlat(lat).AddFlat(lng)
		fmt.Println(args)
		_, err = c.Do("SET", args...)

		if err != nil {
			log.Printf("[SetLocationItemToTile38] service error : %v \n", err)
		}
		log.Println("[SetLocationItemToTile38] done")

	} else {
		log.Printf("[SetLocationItemToTile38] Can not get client\n")
	}
	return
}

func (r *Tile38ManagerService) DropAll() error {
	log.Printf("[DropAll] key:%v", r.location)

	if r.etcdManager != nil {
		h, p, err := r.etcdManager.GetEndpoint(r.sid)
		if err != nil {
			log.Println("EtcdManager get endpoints", "err", err)
		} else {
			r.host = h
			r.port = p
		}
	}
	c, err := transports.GetTile38LocationClient(r.host, r.port)
	if err != nil {
		return err
	}
	defer c.Close()
	_, err = c.Do("DROP " + r.location)
	if err != nil {
		return err
	}
	return nil
}
