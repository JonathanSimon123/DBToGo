package users

import (
	"time"
)

type Users struct {
	Id                  int       `orm:"column(id);pk"` // 用户id
	Code                string    `orm:"column(code)"`
	Account             string    `orm:"column(account)"`        // 帐号
	LoginPassword       string    `orm:"column(login_password)"` // 登录密码
	State               int       `orm:"column(state)"`          // 用户状态  0：正常，1：冻结
	CreateTime          time.Time `orm:"column(create_time)"`    // 注册时间
	MobileType          string    `orm:"column(mobile_type)"`    // 手机类型
	MobileVersion       string    `orm:"column(mobile_version)"` // 手机版本
	LoginTime           time.Time `orm:"column(login_time)"`
	HisDevicecode       string    `orm:"column(his_devicecode)"`        // 手机设备标识
	MobileTypeRecent    string    `orm:"column(mobile_type_recent)"`    // 最近操作的结果，手机型号
	MobileVersionRecent string    `orm:"column(mobile_version_recent)"` // 最近操作的结果，手机app 版本号
	OperationTime       time.Time `orm:"column(operation_time)"`        // 最近操作的时间（1.2.1 版本加入的获取用户接口）
	App                 int       `orm:"column(app)"`                   // 1 ios,2 android,3 wx,4pc5 wd 6 wap
	Token               string    `orm:"column(token)"`                 // 服务器保存一致信息
	PkgType             int       `orm:"column(pkg_type)"`              // 包的标识
	Source              string    `orm:"column(source)"`                // 渠道
	AppVersion          string    `orm:"column(app_version)"`           // app版本
	AppVersionRecent    string    `orm:"column(app_version_recent)"`    // 最近操作app版本
	UserImg             string    `orm:"column(user_img)"`              // 用户头像
	Isemulator          string    `orm:"column(isemulator)"`            // 是否是模拟器
	Jpushid             string    `orm:"column(jpushid)"`               // 手机推送标识
	LoanControl         int       `orm:"column(loan_control)"`          // 借款控制: 0进行中有一笔不能申请第二笔,1:不做限制
	Ip                  string    `orm:"column(ip)"`
	HisDevicecodeRecent string    `orm:"column(his_devicecode_recent)"` // 最近手机设备标识
	Location            string    `orm:"column(location)"`
	Address             string    `orm:"column(address)"` // 地址
	IpLocation          string    `orm:"column(ip_location)"`
	IpAddress           string    `orm:"column(ip_address)"` // 详细地址
	TagType             string    `orm:"column(tag_type)"`   // 标记欺诈类型
	TagSysId            int       `orm:"column(tag_sys_id)"` // 标记人账号id
	Alipayno            string    `orm:"column(alipayno)"`   // 支付宝账号
	IsTag               int       `orm:"column(is_tag)"`     // 0 未标记欺诈 1 家人代偿,2 本人还款意愿差,3 其他,4盗办，5死亡，6坐牢
	IsToutiao           int       `orm:"column(is_toutiao)"` // 0 默认注册 1来自头条
	IsBlack             int       `orm:"column(is_black)"`   // 0默认,是否为黑名单 1-不是黑名单 2-是黑名单
	BlackTime           time.Time `orm:"column(black_time)"` // 标记欺诈的时间
	Remark              string    `orm:"column(remark)"`     // 备注
}
