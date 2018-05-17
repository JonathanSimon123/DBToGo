package dcos_apps_config

import (
	"time"
)

type DcosAppsConfig struct {
	Id              int       `gorm:"column(id);pk"`            // 只增ID
	Env             string    `gorm:"column(env)"`              // 环境（生产，测试， 现在不用了）
	Cluster         string    `gorm:"column(cluster)"`          // 集群信息（eg:[{"zk_hosts": "192.168.2.14:2181", "label": "test", "quota": "1", "marathon": "m2", "name": "test"}]）
	Deploy          string    `gorm:"column(deploy)"`           // 发布信息（eg:{"app_origin": "fwqxz", "deploy_mode": "gzcxb", "app_origin_detail": {"ftp_ip": "192.168.2.63", "ftp_port": "21", "ftp_passwd": "123456", "path": "/home/cent", "ftp_username": "cent", "filename": "demo.war"}}）
	AppId           string    `gorm:"column(app_id)"`           // 应用ID
	SysName         string    `gorm:"column(sys_name)"`         // 系统名称
	AppName         string    `gorm:"column(app_name)"`         // 模块名称
	DnsRoute        string    `gorm:"column(dns_route)"`        // bamboo路由信息
	Status          string    `gorm:"column(status)"`           // 状态
	AutoscalePolicy string    `gorm:"column(autoscale_policy)"` // 是否开启扩缩策略
	Interface       string    `gorm:"column(interface)"`        // 接口人信息（eg:{"phone": "", "other": "", "name": "", "email": ""}）
	SysConf         string    `gorm:"column(sys_conf)"`         // 系统配置信息（eg:{"domain": "", "pri_ip": "", "enable_firewalll": "", "ha_addr": "", "pub_ip": "", "health_addr": "", "bamboo": "", "ha_type": ""}）
	Marathon        string    `gorm:"column(marathon)"`         // marathon上应用配置
	ImageUrl        string    `gorm:"column(image_url)"`        // 镜像地址
	NetId           string    `gorm:"column(net_id)"`           // 自定义网络（callico,对应网络表）
	VolName         string    `gorm:"column(vol_name)"`         // 共享存储名称（对应存储表）
	Instance        int       `gorm:"column(instance)"`         // 当前实例数
	CreateTime      time.Time `gorm:"column(create_time)"`
	SysId           string    `gorm:"column(sys_id)"`        // 对应sys_base 表中的 APP_SYS_NAME 中字段
	RunInstances    int       `gorm:"column(run_instances)"` // 运行实例数
	Domain          string    `gorm:"column(domain)"`
	ImageType       string    `gorm:"column(image_type)"`
	ServiceId       string    `gorm:"column(service_id)"`
	Lb              string    `gorm:"column(lb)"`
	LbUrl           string    `gorm:"column(lb_url)"`
	Request         string    `gorm:"column(request)"`
	InterfaceAId    int       `gorm:"column(interface_a_id)"`
	InterfaceBId    int       `gorm:"column(interface_b_id)"`
	BaseImage       string    `gorm:"column(base_image)"`
	Platform        string    `gorm:"column(platform)"`   // marathon or kubernetes
	Deployment      string    `gorm:"column(deployment)"` // k8s deployment 配置信息
}
