package services

import (
	"context"
	"github.com/go-redis/redis/v8"
	sts "github.com/tencentyun/qcloud-cos-sts-sdk/go"
	"luck-admin/enums"
	"luck-admin/util"
	"time"
)

const COS_TOKEN  = "luck_cos_token"

var ctx = context.Background()
var second int64
var redisClient *redis.Client = util.Redis()

func CosToken() (string,*enums.ErrorInfo) {
	result := redisClient.Get(ctx,COS_TOKEN)
	if len(result.Val()) <= 0 {
		token,errInfo := GetCosToken()
		if errInfo != nil {
			return "",errInfo
		}

		cacheResult := CacheCosToken(token)
		if cacheResult.Err() != nil {
			return "",&enums.ErrorInfo{enums.CosCacheErr,enums.COS_CHACHE_ERR}
		}

		return token,nil
	}

	return result.Val(),nil
}

func GetCosToken() (string,*enums.ErrorInfo) {
	secretId,_ 	:= util.GetCosIni("cos_secret_id")
	secretKey,_ := util.GetCosIni("cos_secret_key")
	appId,_ 	:= util.GetCosIni("cos_app_id")
	bucket,_ 	:= util.GetCosIni("cos_bucket")
	region,_ 	:= util.GetCosIni("cos_region")

	c := sts.NewClient(
		secretId,
		secretKey,
		nil,
	)
	opt := &sts.CredentialOptions{
		DurationSeconds: int64(time.Hour.Seconds()),
		Region:          region,
		Policy: &sts.CredentialPolicy{
			Statement: []sts.CredentialPolicyStatement{
				{
					Action: []string{
						"name/cos:PostObject",
						"name/cos:PutObject",
					},
					Effect: "allow",
					Resource: []string{
						//这里改成允许的路径前缀，可以根据自己网站的用户登录态判断允许上传的具体路径，例子： a.jpg 或者 a/* 或者 * (使用通配符*存在重大安全风险, 请谨慎评估使用)
						"qcs::cos:" + region + ":uid/" + appId + ":" + bucket,
					},
				},
			},
		},
	}
	res, err := c.GetCredential(opt)
	if err != nil {
		util.ErrDetail(enums.COS_GET_TOKEN_ERR,err.Error(),nil)
		return "",&enums.ErrorInfo{enums.CosGetTokenErr,enums.COS_GET_TOKEN_ERR}
	}

	return res.Credentials.SessionToken,nil
}

func CacheCosToken(token string) *redis.StatusCmd {
	result := redisClient.SetEX(ctx,COS_TOKEN,token,time.Hour)
	return result
}
