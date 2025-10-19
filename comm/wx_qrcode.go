package comm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type WxGetQrCodePicReq struct {
	Page       string `json:"page"`
	Scene      string `json:"scene"`
	CheckPath  bool   `json:"check_path"`
	EnvVersion string `json:"env_version"`
}

type WxGetQrCodePicRsp struct {
	Buffer  []byte `json:"buffer"`  // 图片 Buffer
	Errcode int    `json:"errcode"` // 错误码
	Errmsg  string `json:"errmsg"`  // 错误信息
}

// 二维码生成文档：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/qrcode-link/qr-code/getUnlimitedQRCode.html
func GetQrCodePic(uid int64, lessonId string) (error, WxGetQrCodePicRsp) {
	strWxQrcodePath := os.Getenv("WX_QRCODE_PATH")
	if len(strWxQrcodePath) == 0 {
		strWxQrcodePath = "pages/business/write-off/index"
	}

	envVersion := "trial"
	if IsProd() {
		envVersion = "release"
	}
	stWxGetQrCodePicReq := WxGetQrCodePicReq{
		Page:       strWxQrcodePath,
		Scene:      fmt.Sprintf("%s", lessonId),
		CheckPath:  false,
		EnvVersion: envVersion,
	}
	jsonData, err := json.Marshal(stWxGetQrCodePicReq)
	resp, err := http.Post("http://api.weixin.qq.com/wxa/getwxacodeunlimit", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		Printf("http.Post wxQrCode err, uid:%d err:%+v\n", uid, err)
		return err, WxGetQrCodePicRsp{}
	}
	defer resp.Body.Close()

	rspBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		Printf("http.Post wxQrCode ReadAll err, uid:%d err:%+v\n", uid, err)
		return err, WxGetQrCodePicRsp{}
	}
	Printf("http.Post wxQrCode succ, uid:%d req:%+v\n", uid, stWxGetQrCodePicReq)

	var stWxGetQrCodePicRsp WxGetQrCodePicRsp
	strRspBody := string(rspBody)
	if strings.Contains(strRspBody, "errcode") {
		err = json.Unmarshal(rspBody, &stWxGetQrCodePicRsp)
		if err != nil {
			Printf("Unmarshal json err, err:%+v\n", err)
			return err, WxGetQrCodePicRsp{}
		}
	} else {
		stWxGetQrCodePicRsp.Buffer = make([]byte, len(rspBody))
		copy(stWxGetQrCodePicRsp.Buffer, rspBody)
	}

	Printf("Unmarshal json succ, uid:%d errcode:%d errmsg:%s rspBody.len:%d stWxGetQrCodePicRsp.Buffer.len:%d\n",
		uid, stWxGetQrCodePicRsp.Errcode, stWxGetQrCodePicRsp.Errmsg, len(rspBody), len(stWxGetQrCodePicRsp.Buffer))
	return nil, stWxGetQrCodePicRsp
}
