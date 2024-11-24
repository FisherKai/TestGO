package model

type VirtualResource struct {
	ID              int    `json:"id"`               // 资源ID
	Name            string `json:"name"`             // 资源名称
	Type            string `json:"type"`             // 类型
	VCpuUsed        int    `json:"v_cpu_used"`       // 已使用CPU数
	VCpuUsable      int    `json:"v_cpu_usable"`     // 可CPU数
	RAMUsed         int    `json:"ram_used"`         // 内存使用数
	RAMUsable       int    `json:"ram_usable"`       // 内存可用数
	ROMUsed         int    `json:"rom_used"`         // 存储使用数
	ROMUsable       int    `json:"rom_usable"`       // 存储可用数
	InstanceTotal   int    `json:"instance_total"`   // 实例总数
	InstanceError   int    `json:"instance_error"`   // 实例故障数
	InstanceStop    int    `json:"instance_stop"`    // 实例停止数
	InstanceRunning int    `json:"instance_running"` // 实例运行数
	InstanceOther   int    `json:"instance_other"`   // 实例其他
}
