package comm

import (
	"fmt"
	"path"
	"runtime"
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