package hostapi

type DemoHostInfo struct {
	HostId   string `json:"hostId"`
	HostName string `json:"hostName"`
	HostIp   string `json:"hostIp"`
}

type ImportHostRsp struct {
	HostId string `json:"hostId"`
}

type ImportHostsRsp struct {
	HostIds []string `json:"hostIds"`
}
type ListHostRsp struct {
	Hosts []HostInfo `json:"hosts"`
}

type HostDetailsRsp struct {
	Host HostInfo `json:"host"`
}

type ZoneHostStockRsp struct {
	AvailableStocks 	map[string][]ZoneHostStock
}
