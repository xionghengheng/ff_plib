package comm

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// 微信官方订阅消息发送
// MsgDataField 定义了模板数据中的字段
// 云托管参考文档：https://developers.weixin.qq.com/miniprogram/dev/wxcloudrun/src/scene/deploy/subscribe.html
// 参考文档：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/mp-message-management/subscribe-message/sendMessage.html#%E8%AE%A2%E9%98%85%E6%B6%88%E6%81%AF%E5%8F%82%E6%95%B0%E5%80%BC%E5%86%85%E5%AE%B9%E9%99%90%E5%88%B6%E8%AF%B4%E6%98%8E
type MsgDataField struct {
	Value string `json:"value"`
}

// TemplateRequest 定义了请求的结构体
type WxSendMsg2UserReq struct {
	ToUser           string                  `json:"touser"`
	TemplateID       string                  `json:"template_id"`
	Page             string                  `json:"page"`
	MiniprogramState string                  `json:"miniprogram_state"`
	Lang             string                  `json:"lang"`
	Data             map[string]MsgDataField `json:"data"`
}

type WxSendMsg2UserRsp struct {
	Errcode int    `json:"errcode"` // 错误码
	Errmsg  string `json:"errmsg"`  // 错误信息
}

func SendMsg2User(uid int64, stWxSendMsg2UserReq WxSendMsg2UserReq) error {
	jsonData, err := json.Marshal(stWxSendMsg2UserReq)
	resp, err := http.Post("https://api.weixin.qq.com/cgi-bin/message/subscribe/send", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		Printf("http.Post wxSendMsg err, uid:%d err:%+v\n", uid, err)
		return err
	}
	defer resp.Body.Close()

	rspBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		Printf("http.Post wxSendMsg ReadAll err, uid:%d err:%+v\n", uid, err)
		return err
	}
	Printf("http.Post wxSendMsg succ, uid:%d req:%+v\n", uid, stWxSendMsg2UserReq)

	var stWxSendMsg2UserRsp WxSendMsg2UserRsp
	err = json.Unmarshal(rspBody, &stWxSendMsg2UserRsp)
	if err != nil {
		Printf("Unmarshal json err, err:%+v\n", err)
		return err
	}

	Printf("Unmarshal json succ, uid:%d req:%+v rsp:%+v\n", uid, stWxSendMsg2UserReq, stWxSendMsg2UserRsp)
	return nil
}
