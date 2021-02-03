package enums

import "errors"

//数据库连接
var (
	ConnectErr		 				= errors.New("系统异常")							//数据库连接异常
	CosGetTokenErr		 			= errors.New("cos获取token失败")					//cos获取token失败
	CosCacheErr			 			= errors.New("cos缓存失败")						//cos缓存失败
)

//登录授权
var (
	LoginTypeErr					= errors.New("登录类型错误")
	LoginRequestSessionErr			= errors.New("请登录")
	LoginFail						= errors.New("登录失败")  						//请求微信数据异常
	LoginParseUserJsonErr			= errors.New("解析数据异常，请重试")  			//解析数据异常，请重试
	LoginQueryUserErr				= errors.New("用户查询错误")  					//用户查询错误
	LoginSaveUserDbErr				= errors.New("新增用户异常")  					//新增用户数据库异常
	LoginInsertUserErr				= errors.New("用户数据保存失败")  				//用户数据保存失败
	UpdateNicknameAvatarErr			= errors.New("用户数据更新失败")  				//用户数据更新失败
	//JwtParseErr						= errors.New("解析数据失败")
	UnKownSignMethod				= errors.New("授权异常")  						//Unexpected signing method
	LoginCreateTokenErr				= errors.New("授权错误")  						//生成token出错
	TokenNotValid					= errors.New("token非法")  					//生成token出错
	TokenExpired					= errors.New("token已过期")  					//token已过期
	TokenNull						= errors.New("token不能为空")  				//token不能为空
	UserIdTransErr					= errors.New("系统异常")  						//userId转换异常
)

//读取配置
var (
	ReadConfigErr					= errors.New("配置信息错误")
)

//用户相关
var (
	UserPageQueryErr				= errors.New("查询失败")							//用户分页查询出错
)

//礼品相关
var (
	GiftNotFound error 				= errors.New("礼品不存在")
 	GiftSaveErr error 				= errors.New("数据异常，保存失败")
 	GiftAttachmentsEncodeErr error 	= errors.New("图片数据异常，保存失败")
 	GiftAttachmentsDecodeErr error 	= errors.New("图片数据异常")
 	GiftPageQueryFail error 		= errors.New("查询出错")
 	GiftFindEndableErr error 		= errors.New("查询可用礼品错误")
)

//活动相关
var (
	StartDateErr 			error 		= errors.New("活动开始日期格式错误")
 	EndDateErr 				error 		= errors.New("活动截止日期格式错误")
 	RunDateErr 				error 		= errors.New("活动开奖日期格式错误")
 	ActivityDetailNotFound 	error 		= errors.New("活动详情不存在")
 	JoinLimit 				error 		= errors.New("活动参与人数达到限制啦")
 	SaveJoinLogFail 		error 		= errors.New("参加活动失败")
 	ExistsJoinLog	 		error 		= errors.New("您已参加该活动，不可重复参加")
 	QueryJoinLogDbErr	 	error 		= errors.New("查询出错")
 	CreateLDFail 			error 		= errors.New("数据保存失败")
 	ActivityEncodeImageErr  error 		= errors.New("图片转化失败")
 	ActivityDeleteErr					= errors.New("活动删除失败")
 	ActivityUpdateStatusErr				= errors.New("活动状态更新错误")
 	ActivityAcountTodayErr				= errors.New("统计今天活动数量发生错误")
 	ActivityQueryTopErr					= errors.New("活动查询置顶出错")
)

//公共错误
var (
	DecodeErr						= errors.New("数据解析失败")
	SystemErr						= errors.New("系统异常")
)
