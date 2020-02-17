/**
 * @author tunghx
 * @email tunghx@sonek.vn
 * @create date 12/14/19 12:17 PM
 * @modify date 12/14/19 12:17 PM
 * @desc [description]
 */

package TReportStorageService

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/OpenStars/EtcdBackendService/TReportStorageService/treportstorageservice/thrift/gen-go/OpenStars/Common/TReportStorageService"

	"github.com/OpenStars/Common"
	"github.com/OpenStars/EtcdBackendService/TPostStorageService/tpoststorageservice/thrift/gen-go/OpenStars/Common/TPostStorageService"
	"github.com/OpenStars/GoEndpointManager/GoEndpointBackendManager"
	bs "github.com/OpenStars/backendclients/go/bigset/thrift/gen-go/openstars/core/bigset/generic"

	thriftpool "github.com/OpenStars/thriftpoolv2"

	transportBigset "github.com/OpenStars/backendclients/go/bigset/transports"
	transportPostService "github.com/OpenStars/backendclients/go/tpoststorageservice/transports"
)

type reportStorageService struct {
	hostEtcd   string
	portEtcd   string
	hostBigset string
	portBigset string
	sid        string
	epm        GoEndpointBackendManager.EndPointManagerIf
}

/* ==================================== PRIVATE FUNCTION ========================================================= */

func (m *reportStorageService) getBigsetDatabaseClient() *thriftpool.ThriftSocketClient {
	return transportBigset.GetBsGenericClient(m.hostBigset, m.portBigset)
}

func (m *reportStorageService) handlerEventChangeEndpoint(ep *GoEndpointBackendManager.EndPoint) {
	m.hostEtcd = ep.Host
	m.portEtcd = ep.Port
	log.Println("Change config endpoint serviceID", ep.ServiceID, m.hostEtcd, ":", m.portEtcd)
}

func makeBigsetKey(id int64) []byte {
	return []byte(string(id))
}
func makeBigsetKey2(id int64, id2 int64) []byte {
	return []byte(string(id) + "-" + string(id2))
}

/* ==================================== END PRIVATE ============================================================== */
/* ==================================== PUBLISH FUNCTION ========================================================= */
func (m *reportStorageService) GetReportById(idReport int64) (*TReportStorageService.TReportItem, error) {
	client := m.getBigsetDatabaseClient()
	bsKey := makeBigsetKey(idReport)
	if client != nil {
		defer client.BackToPool()

		res, err := client.Client.(*bs.TStringBigSetKVServiceClient).BsGetItem(
			context.Background(),
			bs.TStringKey(Common.REPORT_APP_ID),
			bs.TItemKey(bsKey))
		fmt.Println("Get result: ", res, " with err: ", err)

		if res != nil && err == nil {
			var result TReportStorageService.TReportItem
			result.FromBytes(res.Item.Value)

			return &result, nil
		} else if err != nil {
			return nil, errors.New("Backend service:" + m.sid + " err:" + err.Error())
		} else {
			return nil, errors.New("Backend service:" + m.sid + " key not found")
		}
	}
	return nil, errors.New("Cannot connect to bigset database: " + m.sid + "host: " + m.hostBigset + "port: " + m.portBigset)
}
func (m *reportStorageService) PutReport(idReport int64, data *TReportStorageService.TReportItem) error {
	client := m.getBigsetDatabaseClient()
	bsKey := makeBigsetKey(idReport)

	if client != nil {
		defer client.BackToPool()

		res, err := client.Client.(*bs.TStringBigSetKVServiceClient).BsPutItem(
			context.Background(),
			bs.TStringKey(Common.REPORT_APP_ID),
			&bs.TItem{
				Key:   bsKey,
				Value: data.ToBytes(),
			})
		fmt.Println("Get result: ", res, " with err: ", err, " value: ", data)
		if res != nil && res.Error == bs.TErrorCode_EGood && err == nil {
			if res.IsSetOldItem() {
				var oldItem TReportStorageService.TReportItem
				oldItem.FromBytes(res.GetOldItem().Value)

				fmt.Println("Update old item: ", res, " value: ", oldItem.ReportId)
			}

			if data.CommentId != nil {
				commentBSKey := makeBigsetKey2(*data.CommentId, data.UID)
				res, err := client.Client.(*bs.TStringBigSetKVServiceClient).BsPutItem(
					context.Background(),
					bs.TStringKey(Common.REPORT_BY_COMMENT_ID),
					&bs.TItem{
						Key:   commentBSKey,
						Value: data.ToBytes(),
					})
				if res == nil || res.Error != bs.TErrorCode_EGood || err != nil {
					fmt.Println("Cannot put report data with comment id: ", *data.CommentId)
				}
			} else {
				postBSKey := makeBigsetKey2(*data.PostId, data.UID)
				res, err := client.Client.(*bs.TStringBigSetKVServiceClient).BsPutItem(
					context.Background(),
					bs.TStringKey(Common.REPORT_BY_POST_ID),
					&bs.TItem{
						Key:   postBSKey,
						Value: data.ToBytes(),
					})

				if res == nil || res.Error != bs.TErrorCode_EGood || err != nil {
					fmt.Println("Cannot put report data with post id: ", *data.PostId)
				}
			}
			return nil
		} else if err != nil {
			return errors.New("Backend service:" + m.sid + " err:" + err.Error())
		}
	}
	return errors.New("Cannot connect to bigset database: " + m.sid + "host: " + m.hostBigset + "port: " + m.portBigset)
}
func (m *reportStorageService) RemoveReport(idReport int64) error {
	client := m.getBigsetDatabaseClient()
	bsKey := makeBigsetKey(idReport)
	if client != nil {
		defer client.BackToPool()

		res, err := client.Client.(*bs.TStringBigSetKVServiceClient).BsRemoveItem(
			context.Background(),
			bs.TStringKey(Common.REPORT_APP_ID),
			bs.TItemKey(bsKey))

		fmt.Println("Get result: ", res, " with err: ", err)
		if res == true && err == nil {
			fmt.Println("Remove item: ", idReport, " successful")
			return nil
		} else if err != nil {
			return errors.New("Backend service:" + m.sid + " err:" + err.Error())
		}
	}
	return errors.New("Cannot connect to bigset database: " + m.sid + "host: " + m.hostBigset + "port: " + m.portBigset)
}
func (m *reportStorageService) GetAll(idReport []int64) ([]*TReportStorageService.TReportItem, error) {
	client := m.getBigsetDatabaseClient()
	if client != nil {
		defer client.BackToPool()

		var results []*TReportStorageService.TReportItem
		for i := 0; i < len(idReport); i++ {
			bsKey := makeBigsetKey(idReport[i])
			res, err := client.Client.(*bs.TStringBigSetKVServiceClient).BsGetItem(
				context.Background(),
				bs.TStringKey(Common.REPORT_APP_ID),
				bs.TItemKey(bsKey))

			if res != nil && err == nil {
				var result TReportStorageService.TReportItem
				result.FromBytes(res.Item.Value)
				results = append(results, &result)
			} else if err != nil {
				return nil, errors.New("Backend service:" + m.sid + " err:" + err.Error())
			}
		}
		return results, nil
	}
	return nil, errors.New("Cannot connect to bigset database: " + m.sid + "host: " + m.hostBigset + "port: " + m.portBigset)
}
func (m *reportStorageService) GetAllFromStartReportId(idReport int64, count int32) ([]*TReportStorageService.TReportItem, error) {
	client := m.getBigsetDatabaseClient()
	bsKey := makeBigsetKey(idReport)
	if client != nil {
		defer client.BackToPool()

		res, err := client.Client.(*bs.TStringBigSetKVServiceClient).BsGetSliceFromItem(
			context.Background(),
			bs.TStringKey(Common.REPORT_APP_ID),
			bs.TItemKey(bsKey),
			count)

		if res != nil && err == nil {
			if res.IsSetItems() {
				var results []*TReportStorageService.TReportItem
				for _, item := range res.GetItems().Items {
					var result TReportStorageService.TReportItem
					result.FromBytes(item.Value)
					results = append(results, &result)
				}
				return results, nil
			}
		} else {
			return nil, err
		}
	}
	return nil, errors.New("Cannot connect to bigset database: " + m.sid + "host: " + m.hostBigset + "port: " + m.portBigset)
}
func (m *reportStorageService) GetAllFromPosition(start int32, count int32) ([]*TReportStorageService.TReportItem, error) {
	client := m.getBigsetDatabaseClient()
	if client != nil {
		defer client.BackToPool()

		res, err := client.Client.(*bs.TStringBigSetKVServiceClient).BsGetSliceR(
			context.Background(),
			bs.TStringKey(Common.REPORT_APP_ID),
			start,
			count)

		if res != nil && err == nil {
			if res.IsSetItems() {
				var results []*TReportStorageService.TReportItem
				for _, item := range res.GetItems().Items {
					var result TReportStorageService.TReportItem
					result.FromBytes(item.Value)
					results = append(results, &result)
				}
				return results, nil
			}
		} else {
			return nil, err
		}
	}
	return nil, errors.New("Cannot connect to bigset database: " + m.sid + "host: " + m.hostBigset + "port: " + m.portBigset)
}
func (m *reportStorageService) GetPostById(idPost int64) (*TPostStorageService.TPostItem, error) {
	client := transportPostService.GetTPostStorageServiceCompactClient(m.hostEtcd, m.portEtcd)
	if client != nil {
		defer client.BackToPool()

		res, err := client.Client.(*TPostStorageService.TPostStorageServiceClient).GetData(
			context.Background(),
			TPostStorageService.TKey(idPost))

		if res != nil && err == nil {
			return res.Data, nil
		} else if err != nil {
			return nil, errors.New("Backend service:" + m.sid + " err:" + err.Error())
		} else {
			return nil, errors.New("Backend service:" + m.sid + " key not found")
		}
	}
	return nil, errors.New("Cannot connect to backend service: " + m.sid + "host: " + m.hostEtcd + "port: " + m.portEtcd)
}
func (m *reportStorageService) CheckReportByPostAndUId(idPost int64, uId int64) (*bool, error) {
	client := m.getBigsetDatabaseClient()
	bsKey := makeBigsetKey2(idPost, uId)
	if client != nil {
		defer client.BackToPool()

		res, err := client.Client.(*bs.TStringBigSetKVServiceClient).BsGetItem(
			context.Background(),
			bs.TStringKey(Common.REPORT_BY_POST_ID),
			bs.TItemKey(bsKey))
		fmt.Println("Get result: ", res, " with err: ", err)
		result := res.Item != nil
		if res != nil && err == nil {
			return &result, nil
		} else if err != nil {
			return nil, err
		} else {
			result = false
			return &result, nil
		}
	}
	return nil, errors.New("Cannot connect to bigset database: " + m.sid + "host: " + m.hostBigset + "port: " + m.portBigset)
}
func (m *reportStorageService) CheckReportByCommentAndUId(idComment int64, uId int64) (*bool, error) {
	client := m.getBigsetDatabaseClient()
	bsKey := makeBigsetKey2(idComment, uId)
	if client != nil {
		defer client.BackToPool()

		res, err := client.Client.(*bs.TStringBigSetKVServiceClient).BsGetItem(
			context.Background(),
			bs.TStringKey(Common.REPORT_BY_COMMENT_ID),
			bs.TItemKey(bsKey))
		fmt.Println("Get result: ", res, " with err: ", err)

		result := res.Item != nil
		if res != nil && err == nil {
			return &result, nil
		} else if err != nil {
			return nil, err
		} else {
			result = false
			return &result, nil
		}
	}
	return nil, errors.New("Cannot connect to bigset database: " + m.sid + "host: " + m.hostBigset + "port: " + m.portBigset)
}
func (m *reportStorageService) GetAllUnProcessReport(start int32, count int32) ([]*TReportStorageService.TReportItem, error) {
	client := m.getBigsetDatabaseClient()
	if client != nil {
		defer client.BackToPool()

		res, err := client.Client.(*bs.TStringBigSetKVServiceClient).BsGetSliceR(
			context.Background(),
			bs.TStringKey(Common.REPORT_APP_ID),
			start,
			99999999)

		if res != nil && err == nil {
			if res.IsSetItems() {
				var results []*TReportStorageService.TReportItem
				var countIndex int32 = 0
				for _, item := range res.GetItems().Items {
					var result TReportStorageService.TReportItem
					result.FromBytes(item.Value)
					if result.Action == 0 {
						if countIndex < start {
							countIndex++
						} else {
							results = append(results, &result)
						}
						if int32(len(results)) == count {
							break
						}
					}
				}
				return results, nil
			}
		} else {
			return nil, err
		}
	}
	return nil, errors.New("Cannot connect to bigset database: " + m.sid + "host: " + m.hostBigset + "port: " + m.portBigset)
}

/* ==================================== END PUBLISH ============================================================== */
