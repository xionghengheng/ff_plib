package comm

import (
	"fmt"
	"strconv"
	"strings"
)

// 课包的唯一标识符（用户id_获取课包的时间戳_课包类型）
func GenCoursePackageId(uid int64, ts int64, packageType int) string {
	return fmt.Sprintf("cp_%d_%d_%d", uid, ts, packageType)
}

// 课包的唯一标识符（用户id_获取课包的时间戳_课包类型）
func ParseCoursePackageId(strPackageId string) (int64, int64, int) {
	vecPackageId := strings.Split(strPackageId, "_")
	if len(vecPackageId) >= 4 {
		uid, _ := strconv.ParseInt(vecPackageId[1], 10, 64)
		ts, _ := strconv.ParseInt(vecPackageId[2], 10, 64)
		packageType, _ := strconv.ParseInt(vecPackageId[3], 10, 64)
		return uid, ts, int(packageType)
	}
	return 0, 0, 0
}

// GenPhysicalAssessmentReportID 生成体测报告唯一标识符。
// 格式：par_{package_id}_{report_type}，例如 pay_cp_10001_1722470400_2_1。
func GenPhysicalAssessmentReportID(packageID string, reportType int) string {
	return fmt.Sprintf("par_%s_%d", packageID, reportType)
}

// ParsePhysicalAssessmentReportID 解析体测报告唯一标识符。
// 因 package_id 本身包含下划线，报告阶段从最后一个下划线后解析。
// 返回 false 表示 ID 格式不合法。
func ParsePhysicalAssessmentReportID(reportID string) (packageID string, reportType int, ok bool) {
	const reportIDPrefix = "par_"
	if !strings.HasPrefix(reportID, reportIDPrefix) {
		return "", 0, false
	}

	payload := strings.TrimPrefix(reportID, reportIDPrefix)
	lastSeparator := strings.LastIndex(payload, "_")
	if lastSeparator <= 0 || lastSeparator == len(payload)-1 {
		return "", 0, false
	}

	reportType64, err := strconv.ParseInt(payload[lastSeparator+1:], 10, 0)
	if err != nil || reportType64 <= 0 {
		return "", 0, false
	}

	return payload[:lastSeparator], int(reportType64), true
}

func GenCoursePackageSingleLessonID(uid int64, gymid int, courseId int, coachid int, ts int64) string {
	return fmt.Sprintf("sl_%d_%d_%d_%d_%d", uid, gymid, courseId, coachid, ts)
}

func ParseCoursePackageSingleLessonID(strLessonId string) (int64, int, int, int, int64) {
	vecPackageId := strings.Split(strLessonId, "_")
	if len(vecPackageId) >= 6 {
		uid, _ := strconv.ParseInt(vecPackageId[1], 10, 64)
		gymid, _ := strconv.ParseInt(vecPackageId[2], 10, 64)
		courseId, _ := strconv.ParseInt(vecPackageId[3], 10, 64)
		coachid, _ := strconv.ParseInt(vecPackageId[4], 10, 64)
		ts, _ := strconv.ParseInt(vecPackageId[5], 10, 64)
		return uid, int(gymid), int(courseId), int(coachid), ts
	}
	return 0, 0, 0, 0, 0
}

func GetUidFromCoursePackageSingleLessonID(strLessonId string) int64 {
	vecPackageId := strings.Split(strLessonId, "_")
	if len(vecPackageId) >= 6 {
		uid, _ := strconv.ParseInt(vecPackageId[1], 10, 64)
		return uid
	}
	return 0
}

func GenPassCardLessonID(uid int64, gymid int, cardType int, ts int64) string {
	return fmt.Sprintf("pcl_%d_%d_%d_%d", uid, gymid, cardType, ts)
}

func ParsePassCardLessonID(strLessonId string) (int64, int, int, int64) {
	vecPackageId := strings.Split(strLessonId, "_")
	if len(vecPackageId) >= 5 {
		uid, _ := strconv.ParseInt(vecPackageId[1], 10, 64)
		gymid, _ := strconv.ParseInt(vecPackageId[2], 10, 64)
		cardType, _ := strconv.ParseInt(vecPackageId[3], 10, 64)
		ts, _ := strconv.ParseInt(vecPackageId[4], 10, 64)
		return uid, int(gymid), int(cardType), ts
	}
	return 0, 0, 0, 0
}
