package comm

import (
	"fmt"
	"github.com/xionghengheng/ff_plib/db/dao"
	"github.com/xionghengheng/ff_plib/db/model"
	"os"
	"path"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"time"
)

func genOrderId(productType int, ts int64) string {
	return fmt.Sprintf("%dFUNFIT%d", productType, ts)
}

func genPackageOrderId(productType int, ts int64, gymId int, coachId int, courseId int) string {
	return fmt.Sprintf("%dFUNFIT%d_%d_%d_%d", productType, ts, gymId, coachId, courseId)
}

func parseOrderId(orderId string) (int, int64) {
	vecOrderId := strings.Split(orderId, "FUNFIT")
	if len(vecOrderId) == 2 {
		productType, _ := strconv.ParseInt(vecOrderId[0], 10, 64)
		ts, _ := strconv.ParseInt(vecOrderId[1], 10, 64)
		return int(productType), ts
	}
	return 0, 0
}

// 课包的唯一标识符（用户id_场地id_课程id_教练id_获取课包的时间戳）
func genCoursePackageId(uid int64, ts int64) string {
	return fmt.Sprintf("cp_%d_%d", uid, ts)
}

func genCoursePackageSingleLessonID(uid int64, gymid int, courseId int, coachid int, ts int64) string {
	return fmt.Sprintf("sl_%d_%d_%d_%d_%d", uid, gymid, courseId, coachid, ts)
}

func GetTodayBegTs() int64 {
	t := time.Now()
	addTime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return addTime.Unix()
}

func GetTodayBegTsByTs(givenTimestamp int64) int64 {
	// 将时间戳转换为time.Time对象
	givenTime := time.Unix(givenTimestamp, 0)

	// 获取该日期的零点时间
	midnightTime := time.Date(givenTime.Year(), givenTime.Month(), givenTime.Day(), 0, 0, 0, 0, givenTime.Location())

	// 将零点时间转换为时间戳
	midnightTimestamp := midnightTime.Unix()

	return midnightTimestamp
}

// 获取调用者的文件名和函数名
func getCallerInfo(skip int) (string, string) {
	pc, file, _, ok := runtime.Caller(skip)
	if !ok {
		return "unknown", "unknown"
	}
	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return "unknown", "unknown"
	}
	return path.Base(file), fn.Name()
}

// 包装 fmt.Printf，增加文件名和函数名打印
func Printf(format string, args ...interface{}) {
	// 这里传递 2 以获取更上层的调用者信息
	fileName, fullFuncName := getCallerInfo(2)

	var funcName string
	vecFullFuncName := strings.Split(fullFuncName, ".")
	if len(vecFullFuncName) > 0 {
		funcName = vecFullFuncName[len(vecFullFuncName)-1]
	} else {
		funcName = fullFuncName
	}
	format = fmt.Sprintf("[%s:%s] %s\n", fileName, funcName, format)
	fmt.Printf(format, args...)
}

func CalculateDaysSinceTimestamp(timestamp int64) int {
	// 将时间戳转换为 time.Time 对象
	startTime := time.Unix(timestamp, 0)

	// 获取当前时间
	currentTime := time.Now()

	// 计算两个时间之间的天数差
	days := int(currentTime.Sub(startTime).Hours() / 24)

	return days
}

func GetTodayEndTs() int64 {
	// 获取当前时间
	now := time.Now()

	// 创建当天零点的时间
	midnight := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	// 加一天，得到次日零点的时间
	nextMidnight := midnight.Add(24 * time.Hour)

	// 返回次日零点的时间戳
	return nextMidnight.Unix()
}

// 是否生产环境
func IsProd() bool {
	if os.Getenv("MiniprogramState") == "formal" {
		return true
	}
	return false
}

// 是否测试环境
func IsTrial() bool {
	if os.Getenv("MiniprogramState") == "trial" {
		return true
	}
	return false
}

// 打开多门店
func OpenMultiGym() bool {
	if os.Getenv("OpenMultiGym") == "1" {
		return true
	}
	return false
}

func uniqueVec(vecID *[]string) {
	tmpMap := make(map[string]int)
	for _, id := range *vecID {
		tmpMap[id] = 1
	}
	tmpVecID := make([]string, 0)
	for key, _ := range tmpMap {
		tmpVecID = append(tmpVecID, key)
	}
	*vecID = tmpVecID
	return
}

func uniqueVecCoach(vecID *[]model.CoachModel) {
	tmpMap := make(map[model.CoachModel]int)
	for _, id := range *vecID {
		tmpMap[id] = 1
	}
	tmpVecID := make([]model.CoachModel, 0)
	for key, _ := range tmpMap {
		tmpVecID = append(tmpVecID, key)
	}
	*vecID = tmpVecID
	return
}

func GetAllGymIds(gymIDs string) []int {
	var rsp []int
	if len(gymIDs) == 0 {
		return rsp
	}

	vecStrGymId := strings.Split(gymIDs, ",")
	for _, v := range vecStrGymId {
		nGynId, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			Printf("ParseInt err, err:%+v gymIDs:%s\n", err, gymIDs)
			continue
		}
		rsp = append(rsp, int(nGynId))
	}
	return rsp
}

func GetAllMapGymIds(gymIDs string) map[int]bool {
	rsp := make(map[int]bool)
	if len(gymIDs) == 0 {
		return rsp
	}

	vecStrGymId := strings.Split(gymIDs, ",")
	for _, v := range vecStrGymId {
		nGymId, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			Printf("ParseInt err, err:%+v gymIDs:%s\n", err, gymIDs)
			continue
		}
		rsp[int(nGymId)] = true
	}
	return rsp
}

func GetGymIdsByCoachId(coachId int) ([]int, error) {
	var rsp []int
	stCoachModel, err := dao.ImpCoach.GetCoachById(coachId)
	if err != nil {
		return rsp, err
	}
	if len(stCoachModel.GymIDs) == 0 {
		return rsp, nil
	}

	vecStrGymId := strings.Split(stCoachModel.GymIDs, ",")
	for _, v := range vecStrGymId {
		nGynId, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			Printf("ParseInt err, err:%+v gymIDs:%s\n", err, stCoachModel.GymIDs)
			continue
		}
		rsp = append(rsp, int(nGynId))
	}
	return rsp, nil
}

// GetCoachListByGymIdNew 根据健身房ID获取教练列表(不会对教练可见性做处理，需要调用方自己处理)
func GetCoachListByGymIdNew(reqGymId int) ([]model.CoachModel, error) {
	var rsp []model.CoachModel
	mapCoach := make(map[int][]model.CoachModel)
	vecCoachModel, err := dao.ImpCoach.GetCoachAll()
	if err != nil {
		return rsp, err
	}

	// 将教练按健身房ID分组
	for _, coach := range vecCoachModel {
		vecGymOfCoach := GetAllGymIds(coach.GymIDs)
		for _, gymId := range vecGymOfCoach {
			mapCoach[gymId] = append(mapCoach[gymId], coach)
		}
	}

	// 对每个健身房的教练列表按Priority降序排序
	for gymId, coaches := range mapCoach {
		sort.Slice(coaches, func(i, j int) bool {
			return coaches[i].Priority > coaches[j].Priority // 降序排列
		})
		mapCoach[gymId] = coaches
	}

	return mapCoach[reqGymId], nil
}
