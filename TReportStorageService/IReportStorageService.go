/**
 * @author tunghx
 * @email tunghx@sonek.vn
 * @create date 12/14/19 12:17 PM
 * @modify date 12/14/19 12:17 PM
 * @desc [description]
 */

package TReportStorageService

import (
	"github.com/OpenStars/backendclients/go/tpoststorageservice/thrift/gen-go/OpenStars/Common/TPostStorageService"
	"github.com/OpenStars/backendclients/go/treportstorageservice/thrift/gen-go/OpenStars/Common/TReportStorageService"
)

type IReportStorageService interface {
	GetReportById(idReport int64) (*TReportStorageService.TReportItem, error)
	PutReport(idReport int64, data *TReportStorageService.TReportItem) error
	RemoveReport(idReport int64) error
	GetAll(idReport []int64) ([]*TReportStorageService.TReportItem, error)
	GetAllFromStartReportId(idReport int64, count int32) ([]*TReportStorageService.TReportItem, error)
	GetAllFromPosition(start int32, count int32) ([]*TReportStorageService.TReportItem, error)
	GetPostById(idPost int64) (*TPostStorageService.TPostItem, error)
}
