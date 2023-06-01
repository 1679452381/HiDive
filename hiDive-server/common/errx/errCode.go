package errx

// 成功
const OK uint32 = 200

// 全局错误码
// 服务错误
const SERVER_COMMON_ERROR uint32 = 100001

// 参数错误
const REQUEST_PARAM_ERROR uint32 = 100002

// token超时
const TOKEN_EXPIRE_ERROR uint32 = 100003

// 生成token错误
const TOKEN_GENERATE_ERROR uint32 = 100004

// 数据库错误
const DB_ERROR uint32 = 100005
