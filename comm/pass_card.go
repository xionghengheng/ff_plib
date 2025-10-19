package comm

import (
	"github.com/xionghengheng/ff_plib/db/model"
	"time"
)

func IsPassCardVip(userModel model.UserInfoModel) bool {
	nowTs := time.Now().Unix()
	if userModel.VipExpiredTs <= nowTs &&
		(userModel.VipPassCardType == model.Enum_VipPassCardType_Trial || userModel.VipPassCardType == model.Enum_VipPassCardType_PaidMonth) {
		return false
	}
	return false
}
