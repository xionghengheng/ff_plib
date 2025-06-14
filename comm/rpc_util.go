package comm

import (
	"github.com/xionghengheng/ff_plib/db/dao"
	"github.com/xionghengheng/ff_plib/db/model"
	"sort"
	"strconv"
	"strings"
)

func GetAllGym() (map[int]model.GymInfoModel, error) {
	mapGym := make(map[int]model.GymInfoModel)
	vecGymInfoModel, err := dao.ImpGym.GetGymList()
	if err != nil {
		return mapGym, err
	}
	for _, v := range vecGymInfoModel {
		mapGym[v.GymID] = v
	}
	return mapGym, nil
}

func GetAllCoach() (map[int]model.CoachModel, error) {
	mapCoach := make(map[int]model.CoachModel)
	vecCoachModel, err := dao.ImpCoach.GetCoachAll()
	if err != nil {
		return mapCoach, err
	}
	for _, v := range vecCoachModel {
		mapCoach[v.CoachID] = v
	}
	return mapCoach, nil
}

func GetAllCoachByGym() (map[int][]model.CoachModel, error) {
	tmp := make(map[int][]model.CoachModel)
	vecCoachModel, err := dao.ImpCoach.GetCoachAll()
	if err != nil {
		return tmp, err
	}
	for _, v := range vecCoachModel {
		if v.GymIDs == "" {
			continue
		}

		vecGymIdList := strings.Split(v.GymIDs, ",")
		uniqueVec(&vecGymIdList)
		if len(vecGymIdList) > 0 {
			for _, gym := range vecGymIdList {
				tmpGym, _ := strconv.ParseInt(gym, 10, 64)
				nGym := int(tmpGym)
				tmp[nGym] = append(tmp[nGym], v)
			}
		}
	}

	mapGymId2CoachList := make(map[int][]model.CoachModel)
	for k, v := range tmp {

		//去重
		uniqueVecCoach(&v)

		//按照优先级从大到小排序
		sort.Slice(v, func(i, j int) bool {
			return v[i].Priority > v[j].Priority
		})

		mapGymId2CoachList[k] = v
	}

	return mapGymId2CoachList, nil
}

func GetCoachListByGymId(gymId int) ([]model.CoachModel, error) {
	mapGymId2CoachList, err := GetAllCoachByGym()
	if err != nil {
		return nil, err
	}
	return mapGymId2CoachList[gymId], nil
}

func GetAllUser() (map[int64]model.UserInfoModel, error) {
	mapUser := make(map[int64]model.UserInfoModel)
	vecAllUserModel, err := dao.ImpUser.GetAllUser()
	if err != nil {
		return mapUser, err
	}
	for _, v := range vecAllUserModel {
		mapUser[v.UserID] = v
	}
	return mapUser, nil
}

func GetAllCourse() (map[int]model.CourseModel, error) {
	mapCourse := make(map[int]model.CourseModel)
	vecCourseInfoModel, err := dao.ImpCourse.GetCourseList()
	if err != nil {
		return mapCourse, err
	}
	for _, v := range vecCourseInfoModel {
		mapCourse[v.CourseID] = v
	}
	return mapCourse, nil
}

// getUserInfo 查询用户信息
func GetUserInfoByOpenId(openId string) (*model.UserInfoModel, error) {
	user, err := dao.ImpUser.GetUserByOpenId(openId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GetCoachIdByOpenId 查询用户绑定的教练id
func GetCoachIdByOpenId(openId string) (int, error) {
	user, err := dao.ImpUser.GetUserByOpenId(openId)
	if err != nil || user == nil {
		return 0, err
	}
	return user.CoachId, nil
}

func GetUserInfoByUid(uid int64) (*model.UserInfoModel, error) {
	user, err := dao.ImpUser.GetUser(uid)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetLoginUid(strOpenId string) (int64, string, error) {
	if len(strOpenId) == 0 {
		return 0, "", nil
	}
	stUserInfoModel, err := GetUserInfoByOpenId(strOpenId)
	if err != nil {
		return 0, "", err
	}
	return stUserInfoModel.UserID, stUserInfoModel.Nick, nil
}
