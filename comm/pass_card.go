package comm

import (
	"github.com/xionghengheng/ff_plib/db/model"
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
