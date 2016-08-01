package models

type Heartbeat struct {
	PCid string `json:"pc_id"`
	Ip   string `json:"ip"`
	Hb   int    `json:"hb"`
}

type PcStatus struct {
	PCid         string                 `json:"pc_id"`
	Ip           string                 `json:"ip"`
	Bank         string                 `json:"bank_name"`
	Exception    string                 `json:"exception"`
	SpiderStatus map[string]interface{} `json:"bank_status"`
}

type StepInfo struct {
	Time int64
	Step int
	Bank string
	Pcid string
}

type Email struct {
	To       string
	Subject  string
	Body     string
	MailType string
}

type CompSysStatus struct {
	Cpu []int
	Mem map[string]string
	Io  []map[string]string
	Net []map[string]string
}
