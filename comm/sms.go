package comm

import (
	"encoding/json"
	"errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	tencentclouderrors "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
	"os"
)

// 定义字符串常量
const (
	SmsTemplateId_LoginVerifyCode         = "2272690" // 登录验证码
	SmsTemplateId_CoachCancelLesson       = "2272688" // 课程取消通知
	SmsTemplateId_CoachSchedueLesson      = "2272687" // 排课成功的预约提醒
	SmsTemplateId_LessonStartRemind       = "2272686" // 课程开始前提醒
	SmsTemplateId_LessonFirstBook         = "2296175" // 用户首次约课提醒
	SmsTemplateId_UserBeVip               = "2291779" // 用户订阅提醒
	SmsTemplateId_CoachSetAvaTimeRemind   = "2291778" // 教练设置可约时间提醒
	SmsTemplateId_LessonMissedRemindCoach = "2305573" // 课程旷课提醒

	//用户约课相关
	SmsTemplateId_UserBook       = "2306223" // 学员约课通知
	SmsTemplateId_UserCancelBook = "2306242" // 学员取消预约提醒

	//购课相关
	SmsTemplateId_UserBuyPackage                 = "2306213" // 学员购课通知，给教练通知
	SmsTemplateId_UserBuyPackageNotifyConsultant = "2306244" // 学员购课通知，给顾问通知
	SmsTemplateId_PushUserBuy                    = "2381955" // 学员体验课已完成，通知学员购课优惠
	SmsTemplateId_PushUserBuyWhenFirstOver       = "2410925" // 第一节体验课完成后的购课提醒
)

func SendSmsMsg2User(templateId string, uid int64, vecTemplateParam []string, phone string) error {

	/* 必要步骤：
	 * 实例化一个认证对象，入参需要传入腾讯云账户密钥对secretId，secretKey。
	 * 这里采用的是从环境变量读取的方式，需要在环境变量中先设置这两个值。
	 * 您也可以直接在代码中写死密钥对，但是小心不要将代码复制、上传或者分享给他人，
	 * 以免泄露密钥对危及您的财产安全。
	 * SecretId、SecretKey 查询: https://console.cloud.tencent.com/cam/capi */
	credential := common.NewCredential(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"),
	)

	/* 非必要步骤:
	 * 实例化一个客户端配置对象，可以指定超时时间等配置 */
	cpf := profile.NewClientProfile()

	/* SDK默认使用POST方法。
	 * 如果您一定要使用GET方法，可以在这里设置。GET方法无法处理一些较大的请求 */
	cpf.HttpProfile.ReqMethod = "POST"

	/* SDK有默认的超时时间，非必要请不要进行调整
	 * 如有需要请在代码中查阅以获取最新的默认值 */
	// cpf.HttpProfile.ReqTimeout = 5
	/* 指定接入地域域名，默认就近地域接入域名为 sms.tencentcloudapi.com ，也支持指定地域域名访问，例如广州地域的域名为 sms.ap-guangzhou.tencentcloudapi.com */
	cpf.HttpProfile.Endpoint = "sms.tencentcloudapi.com"

	/* SDK默认用TC3-HMAC-SHA256进行签名，非必要请不要修改这个字段 */
	cpf.SignMethod = "HmacSHA1"

	/* 实例化要请求产品(以sms为例)的client对象
	 * 第二个参数是地域信息，可以直接填写字符串ap-guangzhou，支持的地域列表参考 https://cloud.tencent.com/document/api/382/52071#.E5.9C.B0.E5.9F.9F.E5.88.97.E8.A1.A8 */
	client, _ := sms.NewClient(credential, "ap-guangzhou", cpf)

	/* 实例化一个请求对象，根据调用的接口和实际情况，可以进一步设置请求参数
	 * 您可以直接查询SDK源码确定接口有哪些属性可以设置
	 * 属性可能是基本类型，也可能引用了另一个数据结构
	 * 推荐使用IDE进行开发，可以方便的跳转查阅各个接口和数据结构的文档说明 */
	request := sms.NewSendSmsRequest()

	/* 基本类型的设置:
	 * SDK采用的是指针风格指定参数，即使对于基本类型您也需要用指针来对参数赋值。
	 * SDK提供对基本类型的指针引用封装函数
	 * 帮助链接：
	 * 短信控制台: https://console.cloud.tencent.com/smsv2
	 * 腾讯云短信小助手: https://cloud.tencent.com/document/product/382/3773#.E6.8A.80.E6.9C.AF.E4.BA.A4.E6.B5.81 */
	/* 短信应用ID: 短信SdkAppId在 [短信控制台] 添加应用后生成的实际SdkAppId，示例如1400006666 */
	// 应用 ID 可前往 [短信控制台](https://console.cloud.tencent.com/smsv2/app-manage) 查看
	request.SmsSdkAppId = common.StringPtr("1400911457")

	/* 短信签名内容: 使用 UTF-8 编码，必须填写已审核通过的签名 */
	// 签名信息可前往 [国内短信](https://console.cloud.tencent.com/smsv2/csms-sign) 或 [国际/港澳台短信](https://console.cloud.tencent.com/smsv2/isms-sign) 的签名管理查看
	//request.SignName = common.StringPtr("趣私教FunCoach")
	request.SignName = common.StringPtr(os.Getenv("TENCENTCLOUD_SMS_SIGN_NAME"))

	/* 模板 ID: 必须填写已审核通过的模板 ID */
	// 模板 ID 可前往 [国内短信](https://console.cloud.tencent.com/smsv2/csms-template) 或 [国际/港澳台短信](https://console.cloud.tencent.com/smsv2/isms-template) 的正文模板管理查看
	request.TemplateId = common.StringPtr(templateId)

	/* 模板参数: 模板参数的个数需要与 TemplateId 对应模板的变量个数保持一致，若无模板参数，则设置为空*/
	request.TemplateParamSet = common.StringPtrs(vecTemplateParam) //验证码5分钟内过期

	/* 下发手机号码，采用 E.164 标准，+[国家或地区码][手机号]
	 * 示例如：+8613711112222， 其中前面有一个+号 ，86为国家码，13711112222为手机号，最多不要超过200个手机号*/
	request.PhoneNumberSet = common.StringPtrs([]string{phone})

	/* 用户的 session 内容（无需要可忽略）: 可以携带用户侧 ID 等上下文信息，server 会原样返回 */
	request.SessionContext = common.StringPtr("")

	/* 短信码号扩展号（无需要可忽略）: 默认未开通，如需开通请联系 [腾讯云短信小助手] */
	request.ExtendCode = common.StringPtr("")

	/* 国内短信无需填写该项；国际/港澳台短信已申请独立 SenderId 需要填写该字段，默认使用公共 SenderId，无需填写该字段。注：月度使用量达到指定量级可申请独立 SenderId 使用，详情请联系 [腾讯云短信小助手](https://cloud.tencent.com/document/product/382/3773#.E6.8A.80.E6.9C.AF.E4.BA.A4.E6.B5.81)。 */
	request.SenderId = common.StringPtr("")

	// 通过client对象调用想要访问的接口，需要传入请求对象
	response, err := client.SendSms(request)
	// 处理异常
	if _, ok := err.(*tencentclouderrors.TencentCloudSDKError); ok {
		Printf("An API error has returned, templateId:%d phone:%s uid:%d err:%+v", templateId, phone, uid, err)
		return err
	}
	// 非SDK异常，直接失败。实际代码中可以加入其他的处理。
	if err != nil {
		Printf("非SDK异常，直接失败。实际代码中可以加入其他的处理, templateId:%d phone:%s uid:%d err:%+v\n", templateId, phone, uid, err)
		return err
	}
	b, _ := json.Marshal(response.Response)
	// 打印返回的json字符串
	Printf("send msCode succ, templateId:%s phone:%s uid:%d rsp:%s\n", templateId, phone, uid, b)
	if len(response.Response.SendStatusSet) == 0 || response.Response.SendStatusSet[0].Code == nil || response.Response.SendStatusSet[0].Message == nil {
		Printf("rspData format error, templateId:%s phone:%s uid:%d response:%+v", templateId, phone, uid, response)
		return err
	}

	if (*response.Response.SendStatusSet[0].Code) != "Ok" {
		Printf("rspData format error, templateId:%s phone:%s uid:%d response:%+v", templateId, phone, uid, response)
		return errors.New(*response.Response.SendStatusSet[0].Message)
	}
	return nil
}
