package comm

import (
	"github.com/xionghengheng/ff_plib/db/dao"
	"github.com/xionghengheng/ff_plib/db/model"
	"github.com/xionghengheng/ff_plib/db/pass_card_dao"
	"github.com/xionghengheng/ff_plib/db/pass_card_model"
	"time"
)

func IsPassCardVip(user *model.UserInfoModel) bool {
	if user == nil {
		return false
	}
	nowTs := time.Now().Unix()
	if user.VipExpiredTs >= nowTs &&
		(user.VipPassCardType == model.Enum_VipPassCardType_Trial || user.VipPassCardType == model.Enum_VipPassCardType_PaidMonth) {
		return true
	}
	return false
}

func GetAllPassCardGym() (map[int]pass_card_model.PassCardGymInfoModel, error) {
	mapGym := make(map[int]pass_card_model.PassCardGymInfoModel)
	vecGymInfoModel, err := pass_card_dao.ImpGym.GetGymList()
	if err != nil {
		return mapGym, err
	}
	for _, v := range vecGymInfoModel {
		mapGym[v.GymID] = v
	}
	return mapGym, nil
}

func GetBindPassCardGymIdByOpenId(openId string) (int, error) {
	user, err := dao.ImpUser.GetUserByOpenId(openId)
	if err != nil || user == nil {
		return 0, err
	}
	return user.BindPassCardGymId, nil
}
