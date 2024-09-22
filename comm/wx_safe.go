package comm

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// 微信官方文字安全打击
// 参考文档：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/sec-center/sec-check/msgSecCheck.html
type WxCheckContentSafeReq struct {
	OpenID  string `json:"openid"`
	Scene   int    `json:"scene"`
	Version int    `json:"version"`
	Content string `json:"content"`
}

type WxCheckContentSafeRsp struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	Result  struct {
		Suggest string `json:"suggest"`
		Label   int    `json:"label"`
	} `json:"result"`
	Detail []struct {
		Strategy string `json:"strategy"`
		ErrCode  int    `json:"errcode"`
		Suggest  string `json:"suggest"`
		Label    int    `json:"label"`
		Prob     int    `json:"prob,omitempty"`
		Level    int    `json:"level,omitempty"`
		Keyword  string `json:"keyword,omitempty"`
	} `json:"detail"`
	TraceID string `json:"trace_id"`
}

func WxCheckContentSafe(uid int64, stWxCheckContentSafeReq WxCheckContentSafeReq) (error, bool) {
	if len(stWxCheckContentSafeReq.Content) == 0 {
		return nil, true
	}
	jsonData, err := json.Marshal(stWxCheckContentSafeReq)
	resp, err := http.Post("https://api.weixin.qq.com/wxa/msg_sec_check", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		Printf("http.Post msg_sec_check err, uid:%d err:%+v\n", uid, err)
		return err, false
	}
	defer resp.Body.Close()

	rspBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		Printf("http.Post msg_sec_check ReadAll err, uid:%d err:%+v\n", uid, err)
		return err, false
	}
	Printf("http.Post msg_sec_check succ, uid:%d req:%+v\n", uid, stWxCheckContentSafeReq)

	var stWxCheckContentSafeRsp WxCheckContentSafeRsp
	err = json.Unmarshal(rspBody, &stWxCheckContentSafeRsp)
	if err != nil {
		Printf("Unmarshal json err, err:%+v\n", err)
		return err, false
	}

	Printf("Unmarshal json succ, uid:%d req:%+v rsp:%+v\n", uid, stWxCheckContentSafeReq, stWxCheckContentSafeRsp)

	//测试环境，微信不会做检测，直接返回成功
	if os.Getenv("MiniprogramState") == "trial"{
		return nil, true
	}

	if stWxCheckContentSafeRsp.ErrCode != 0 {
		return errors.New(fmt.Sprintf("msg_sec_check err, code:%d", stWxCheckContentSafeRsp.ErrCode)), false
	}

	//Suggest = risky、pass、review
	//Label = 命中标签枚举值，100 正常；10001 广告；20001 时政；20002 色情；20003 辱骂；20006 违法犯罪；20008 欺诈；20012 低俗；20013 版权；21000 其他
	bCheckPass := false
	if stWxCheckContentSafeRsp.Result.Suggest == "pass" {
		bCheckPass = true
	}

	return nil, bCheckPass
}

// 微信官方图片、音频安全打击
// 参考文档：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/sec-center/sec-check/mediaCheckAsync.html
// RequestData 定义请求数据的结构
type WxCheckPicSafeReq struct {
	OpenID    string `json:"openid"`
	Scene     int    `json:"scene"`
	Version   int    `json:"version"`
	MediaURL  string `json:"media_url"`
	MediaType int    `json:"media_type"`
}

// ResponseData 定义响应数据的结构
type WxCheckPicSafeRsp struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	TraceID string `json:"trace_id"`
}

func WxCheckPicSafe(uid int64, stWxCheckPicSafeReq WxCheckPicSafeReq) (error,string) {
	if len(stWxCheckPicSafeReq.MediaURL) == 0 {
		return nil, ""
	}
	jsonData, err := json.Marshal(stWxCheckPicSafeReq)
	resp, err := http.Post("https://api.weixin.qq.com/wxa/media_check_async", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		Printf("http.Post media_check_async err, uid:%d err:%+v\n", uid, err)
		return err, ""
	}
	defer resp.Body.Close()

	rspBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		Printf("http.Post media_check_async ReadAll err, uid:%d err:%+v\n", uid, err)
		return err, ""
	}
	Printf("http.Post media_check_async succ, uid:%d req:%+v rspBody:%s\n", uid, stWxCheckPicSafeReq, string(rspBody))

	var stWxCheckPicSafeRsp WxCheckPicSafeRsp
	err = json.Unmarshal(rspBody, &stWxCheckPicSafeRsp)
	if err != nil {
		Printf("Unmarshal json err, err:%+v\n", err)
		return err, ""
	}

	Printf("Unmarshal json succ, uid:%d req:%+v rsp:%+v\n", uid, stWxCheckPicSafeReq, stWxCheckPicSafeRsp)

	if stWxCheckPicSafeRsp.ErrCode != 0 {
		return errors.New(fmt.Sprintf("media_check_async err, code:%d", stWxCheckPicSafeRsp.ErrCode)), ""
	}

	//suggest=建议，有risky、pass、review三种值
	//label=命中标签枚举值，100 正常；20001 时政；20002 色情；20006 违法犯罪；21000 其他

	return nil, stWxCheckPicSafeRsp.TraceID
}
