package enums

const (
	DAY_FORMAT				= "20060102"
    DATE_ONLY_FORMAT  		= "2006-01-02 00:00:00"
 	DATE_FORMAT  			= "2006-01-02 15:04:05"
)

//业务不相关错误
const (
	SUCCESS					= 0
	FAIL 					= 1
	DB_CONNECT_ERR 			= 2
	READ_CONFIG_ERR			= 3
	COS_GET_TOKEN_ERR		= 4
	COS_CHACHE_ERR			= 5
	COS_ENCODE_ERR			= 6
	COS_DECODE_ERR			= 7
	DECODE_ARR_ERR			= 8
)

//授权相关 1000 ~ 1999
//参数错误
const (
	AUTH_PARAMS_ERROR 				= 1000		//登录参数错误
	AUTH_LOGIN_TYPE_ERR				= 1001 		//登录类型错误
	AUTH_REQUEST_SESSION_ERR		= 1002		//请求微信session出错
	AUTH_REQUEST_SESSION_RESP_ERR	= 1003		//请求微信session返回错误异常
	AUTH_PARSE_JSON_ERR				= 1004		//解析用户json失败
	AUTH_USER_QUERY_ERR				= 1005		//用户查询错误
	AUTH_USER_SAVE_ERR				= 1006		//新增用户数据库异常
	AUTH_USER_UPDATE_ERR			= 1007		//用户数据更新失败
	AUTH_USER_PARSE_JWT_ERR			= 1008		//解析json失败
	AUTH_USER_CREATE_JWT_ERR		= 1009		//生成jwt异常
	AUTH_TOKEN_EXPIRED				= 1010		//token已过有效期
	AUTH_TOKEN_NULL					= 1011		//token为空
	AUTH_NOT_LOGIN					= 1012		//未登录
	Auth_TRANS_UID_ERR				= 1013		//userId类型转化失败
)

//活动相关 2000 ~ 2999
const (
	ACTIVITY_PARAM_ERR 				= 2000 		//参数错误
	ACTIVITY_START_DATE_ERR 		= 2001 		//活动开始如期解析错误
	ACTIVITY_END_DATE_ERR 			= 2002 		//活动截止日期解析错误
	ACTIVITY_RUN_DATE_ERR 			= 2003 		//活动开奖日期解析错误
	ACTIVITY_SAVE_ERR 				= 2004 		//活动保存失败
	ACTIVITY_PAGE_ERR				= 2005		//分页查询错误
	ACTIVITY_DETAIL_PARAM_ERR		= 2006		//详情id不能为空
	ACTIVITY_DETAIL_QUERY_ERR		= 2007		//详情查询错误
	ACTIVITY_DETAIL_NOT_FOUND		= 2008		//详情不存在
	ACTIVITY_JOIN_PARAM_ERR			= 2009		//参团参数失败，id为空
	ACTIVITY_JOIN_LIMIT				= 2010		//活动参与人数达到限制啦
	ACTIVITY_JOIN_SAVE_LOG_FAIL		= 2011		//参加活动失败
	ACTIVITY_JOIN_REPEAT			= 2012		//您已参加该活动，不可重复参加
	ACTIVITY_JOIN_QUERY_ERR			= 2013		//查询参与日志出错
	ACTIVITY_IMAGE_ENCODE_ERR		= 2014		//活动图片转换成字符串失败
	ACTIVITY_DELETE_ERR				= 2015		//活动删除失败
	ACTIVITY_UPDATE_STATUS_ERR		= 2016		//活动状态更新错误
	ACTIVITY_UPDATE_STATUS_BAD_ERR	= 2017		//活动状态更新状态不允许修改
	ACTIVITY_NOT_FOUND				= 2018		//活动不存在
	ACTIVITY_UPDATE_ERR				= 2019		//活动更新错误
	ACTIVITY_COUNT_TODAY_ERR		= 2020		//活动count今天活动发生错误
	ACTIVITY_QUERY_TOP_ERR   		= 2021		//活动查询置顶出错
)

//礼品相关 3000 ~ 3999
const (
	GIFT_SAVE_ERR					= 3000 		//礼品保存失败
	GIFT_FIRST_ERR					= 3001 		//礼品查询出错
	GIFT_NOT_FOUND					= 3002 		//礼品不存在
	GIFT_GET_DETAIL_ERR				= 3003 		//礼品详情查询错误
	GIFT_ATTACHMENTS_ENCODE_ERR		= 3004 		//attachments encode失败
	GIFT_ATTACHMENTS_DECODE_ERR		= 3005 		//attachments decode失败
	GIFT_PAGE_QUERY_FAIL			= 3006		//分页查询失败
	GIFT_FIND_ENABLE_ERR			= 3007 		//查询可用礼品出错
)

//用户相关
const (
	USER_PAGE_QUERY_ERR				= 4000		//用户分页查询出错
)

