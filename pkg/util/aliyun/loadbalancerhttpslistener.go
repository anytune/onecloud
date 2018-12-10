package aliyun

import "yunion.io/x/jsonutils"

type SLoadbalancerHTTPSListener struct {
	lb *SLoadbalancer

	RequestId         string //	请求ID。
	ListenerPort      int    //	负载均衡实例前端使用的端口。
	BackendServerPort int    //	负载均衡实例后端使用的端口。
	Bandwidth         int    //	监听的带宽峰值。
	Status            string //	当前监听的状态。取值：starting | running | configuring | stopping | stopped

	XForwardedFor       string //	是否开启通过X-Forwarded-For头字段获取访者真实IP。
	XForwardedFor_SLBIP string //	是否通过SLB-IP头字段获取客户端请求的真实IP。
	XForwardedFor_SLBID string //	是否通过SLB-ID头字段获取负载均衡实例ID。
	XForwardedFor_proto string //	是否通过X-Forwarded-Proto头字段获取负载均衡实例的监听协议。
	Scheduler           string //	调度算法。
	StickySession       string //	是否开启会话保持。
	StickySessionType   string //	cookie的处理方式。
	CookieTimeout       int    //	Cookie超时时间。
	Cookie              string //	服务器上配置的cookie。
	AclStatus           string //	是否开启访问控制功能。取值：on | off（默认值）

	AclType string //	访问控制类型

	AclId string //	监听绑定的访问策略组ID。当AclStatus参数的值为on时，该参数必选。

	HealthCheck            string //	是否开启健康检查。
	HealthCheckDomain      string //	用于健康检查的域名。
	HealthCheckURI         string //	用于健康检查的URI。
	HealthyThreshold       int    //	健康检查阈值。
	UnhealthyThreshold     int    //	不健康检查阈值。
	HealthCheckTimeout     int    //	每次健康检查响应的最大超时间，单位为秒。
	HealthCheckInterval    int    //	健康检查的时间间隔，单位为秒。
	HealthCheckHttpCode    string //	健康检查正常的HTTP状态码。
	HealthCheckConnectPort int    //	健康检查的端口。
	VServerGroupId         string //	绑定的服务器组ID。
	ServerCertificateId    string //	服务器证书ID。
	CACertificateId        string //	CA证书ID。
	Gzip                   string //	是否开启Gzip压缩。
	Rules                  []Rule //监听下的转发规则列表，具体请参见RuleList。
	DomainExtensions       string //	域名扩展列表，具体请参见DomainExtensions。
	EnableHttp2            string //	是否开启HTTP/2特性。取值：on（默认值）|off

	TLSCipherPolicy string //
}

func (listener *SLoadbalancerHTTPSListener) GetName() string {
	return ""
}

func (listerner *SLoadbalancerHTTPSListener) GetId() string {
	return ""
}

func (listerner *SLoadbalancerHTTPSListener) GetGlobalId() string {
	return listerner.GetId()
}

func (listerner *SLoadbalancerHTTPSListener) GetStatus() string {
	return ""
}

func (listerner *SLoadbalancerHTTPSListener) GetMetadata() *jsonutils.JSONDict {
	return nil
}

func (listerner *SLoadbalancerHTTPSListener) IsEmulated() bool {
	return false
}

func (listerner *SLoadbalancerHTTPSListener) Refresh() error {
	return nil
}

func (region *SRegion) GetLoadbalancerHTTPSListener(loadbalancerId string, listenerPort string) (*SLoadbalancerHTTPSListener, error) {
	params := map[string]string{}
	params["RegionId"] = region.RegionId
	params["LoadBalancerId"] = loadbalancerId
	params["ListenerPort"] = listenerPort
	body, err := region.lbRequest("DescribeLoadBalancerHTTPSListenerAttribute", params)
	if err != nil {
		return nil, err
	}
	listener := SLoadbalancerHTTPSListener{}
	return &listener, body.Unmarshal(&listener)
}
