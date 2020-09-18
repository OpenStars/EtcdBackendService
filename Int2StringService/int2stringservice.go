package Int2StringService

import (
	"context"
	"errors"
	"log"

	"github.com/OpenStars/EtcdBackendService/Int2StringService/i2skv/thrift/gen-go/OpenStars/Common/I2SKV"
	"github.com/OpenStars/EtcdBackendService/Int2StringService/i2skv/transports"
	"github.com/OpenStars/GoEndpointManager"
	"github.com/OpenStars/GoEndpointManager/GoEndpointBackendManager"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Int2StringService struct {
	host        string
	port        string
	sid         string
	epm         GoEndpointBackendManager.EndPointManagerIf
	etcdManager *GoEndpointManager.EtcdBackendEndpointManager

	bot_token  string
	bot_chatID int64
	botClient  *tgbotapi.BotAPI
}

func (m *Int2StringService) notifyEndpointError() {
	if m.botClient != nil {
		msg := tgbotapi.NewMessage(m.bot_chatID, "Hệ thống kiểm soát phát hiện service int2string có địa chỉ "+m.host+":"+m.port+" đang không hoạt động")
		m.botClient.Send(msg)
	}
}

func (m *Int2StringService) PutData(key int64, value string) error {
	if m.etcdManager != nil {
		h, p, err := m.etcdManager.GetEndpoint(m.sid)
		if err != nil {
			log.Println("EtcdManager get endpoints", "err", err)
		} else {
			m.host = h
			m.port = p
		}
	}

	client := transports.GetTI2StringServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		go m.notifyEndpointError()
		return errors.New("Can not connect to model")
	}

	tkey := I2SKV.TKey(key)
	tvalue := &I2SKV.TStringValue{
		Value: value,
	}
	_, err := client.Client.(*I2SKV.TI2StringServiceClient).PutData(context.Background(), tkey, tvalue)
	if err != nil {
		go m.notifyEndpointError()
		return errors.New("Int2StringService sid:" + m.sid + " addresss: " + m.host + ":" + m.port + " err: " + err.Error())
	}
	defer client.BackToPool()

	return nil

}

func (m *Int2StringService) GetData(key int64) (string, error) {
	if m.etcdManager != nil {
		h, p, err := m.etcdManager.GetEndpoint(m.sid)
		if err != nil {
			log.Println("EtcdManager get endpoints", "err", err)
		} else {
			m.host = h
			m.port = p
		}
	}
	client := transports.GetTI2StringServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		go m.notifyEndpointError()
		return "", errors.New("Can not connect to model")
	}

	tkey := I2SKV.TKey(key)
	r, err := client.Client.(*I2SKV.TI2StringServiceClient).GetData(context.Background(), tkey)
	// log.Println(r, err)
	if err != nil {
		go m.notifyEndpointError()
		return "", errors.New("Int2StringService sid:" + m.sid + " addresss: " + m.host + ":" + m.port + " err: " + err.Error())
	}
	defer client.BackToPool()

	if r.Data == nil || r.ErrorCode != I2SKV.TErrorCode_EGood || r.Data.Value == "" {
		return "", errors.New("Int2StringService sid:" + m.sid + " addresss: " + m.host + ":" + m.port + " err: " + r.ErrorCode.String())
	}
	return r.Data.Value, nil
}

func (m *Int2StringService) PutData2(key int64, value string) (bool, error) {
	if m.etcdManager != nil {
		h, p, err := m.etcdManager.GetEndpoint(m.sid)
		if err != nil {
			log.Println("EtcdManager get endpoints", "err", err)
		} else {
			m.host = h
			m.port = p
		}
	}

	client := transports.GetTI2StringServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		go m.notifyEndpointError()
		return false, errors.New("Can not connect to model")
	}

	tkey := I2SKV.TKey(key)
	tvalue := &I2SKV.TStringValue{
		Value: value,
	}
	r, err := client.Client.(*I2SKV.TI2StringServiceClient).PutData(context.Background(), tkey, tvalue)
	if err != nil {
		go m.notifyEndpointError()
		return false, errors.New("Int2StringService sid:" + m.sid + " addresss: " + m.host + ":" + m.port + " err: " + err.Error())
	}
	defer client.BackToPool()

	if r != I2SKV.TErrorCode_EGood {
		return false, nil
	}
	return true, nil

}

func (m *Int2StringService) GetData2(key int64) (string, error) {
	if m.etcdManager != nil {
		h, p, err := m.etcdManager.GetEndpoint(m.sid)
		if err != nil {
			log.Println("EtcdManager get endpoints", "err", err)
		} else {
			m.host = h
			m.port = p
		}
	}
	client := transports.GetTI2StringServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		go m.notifyEndpointError()
		return "", errors.New("Can not connect to model")
	}

	tkey := I2SKV.TKey(key)
	r, err := client.Client.(*I2SKV.TI2StringServiceClient).GetData(context.Background(), tkey)
	if err != nil {
		go m.notifyEndpointError()
		return "", errors.New("Int2StringService sid:" + m.sid + " addresss: " + m.host + ":" + m.port + " err: " + err.Error())
	}
	defer client.BackToPool()

	if r.Data == nil || r.ErrorCode != I2SKV.TErrorCode_EGood || r.Data.Value == "" {
		return "", nil
	}
	return r.Data.Value, nil
}

func (m *Int2StringService) handleEventChangeEndpoint(ep *GoEndpointBackendManager.EndPoint) {
	m.host = ep.Host
	m.port = ep.Port
	log.Println("Change config endpoint serviceID", ep.ServiceID, m.host, ":", m.port)
}

func NewInt2StringService(serviceID string, etcdServers []string, defaultEndpoint GoEndpointBackendManager.EndPoint) Int2StringServiceIf {

	// aepm := GoEndpointBackendManager.NewEndPointManager(etcdServers, serviceID)
	// err, ep := aepm.GetEndPoint()
	// if err != nil {
	// 	// log.Println("Load endpoit ", serviceID, "err", err.Error())
	// 	log.Println("Init Local Int2StringService sid:", defaultEndpoint.ServiceID, "host:", defaultEndpoint.Host, "port:", defaultEndpoint.Port)
	// 	return &Int2StringService{
	// 		host: defaultEndpoint.Host,
	// 		port: defaultEndpoint.Port,
	// 		sid:  defaultEndpoint.ServiceID,
	// 	}
	// }

	i2ssv := &Int2StringService{
		host:        defaultEndpoint.Host,
		port:        defaultEndpoint.Port,
		sid:         defaultEndpoint.ServiceID,
		etcdManager: GoEndpointManager.GetEtcdBackendEndpointManagerSingleton(etcdServers),

		bot_chatID: -1001469468779,
		bot_token:  "1108341214:AAEKNbFf6PO7Y6UJGK-xepDDOGKlBU2QVCg",
		botClient:  nil,
	}
	bot, err := tgbotapi.NewBotAPI(i2ssv.bot_token)
	if err == nil {
		i2ssv.botClient = bot
	}
	if i2ssv.etcdManager == nil {
		return i2ssv
	}
	err = i2ssv.etcdManager.SetDefaultEntpoint(serviceID, defaultEndpoint.Host, defaultEndpoint.Port)
	if err != nil {
		log.Println("SetDefaultEndpoint sid", serviceID, "err", err)
		return i2ssv
	}
	// i2ssv.etcdManager.GetAllEndpoint(serviceID)
	return i2ssv

	// sv := &Int2StringService{
	// 	host: ep.Host,
	// 	port: ep.Port,
	// 	sid:  ep.ServiceID,
	// }
	// go aepm.EventChangeEndPoints(sv.handleEventChangeEndpoint)
	// sv.epm = aepm
	// log.Println("Init From Etcd Int2StringService sid:", sv.sid, "host:", sv.host, "port:", sv.port)
	// return sv
}

func NewClient(etcdServers []string, sid, host, port string) Client {
	i2ssv := &Int2StringService{
		host:        host,
		port:        port,
		sid:         sid,
		etcdManager: GoEndpointManager.GetEtcdBackendEndpointManagerSingleton(etcdServers),

		bot_chatID: -1001469468779,
		bot_token:  "1108341214:AAEKNbFf6PO7Y6UJGK-xepDDOGKlBU2QVCg",
		botClient:  nil,
	}
	bot, err := tgbotapi.NewBotAPI(i2ssv.bot_token)
	if err == nil {
		i2ssv.botClient = bot
	}
	if i2ssv.etcdManager == nil {
		return i2ssv
	}
	err = i2ssv.etcdManager.SetDefaultEntpoint(sid, host, port)
	if err != nil {
		log.Println("SetDefaultEndpoint sid", sid, "err", err)
		return i2ssv
	}
	// i2ssv.etcdManager.GetAllEndpoint(serviceID)
	return i2ssv
}
