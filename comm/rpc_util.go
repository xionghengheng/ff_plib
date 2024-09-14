package comm

import (
	"github.com/xionghengheng/ff_plib/db/dao"
	"github.com/xionghengheng/ff_plib/db/model"
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

func GetAllCouse() (map[int]model.CourseModel, error) {
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
	if err != nil || user == nil{
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
