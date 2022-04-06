package configuration

type FanConfig struct {
	ID          string             `json:"id"`
	NeverStop   bool               `json:"neverStop"`
	StartPwm    *int               `json:"startPwm,omitempty"`
	Curve       string             `json:"curve"`
	HwMon       *HwMonFanConfig    `json:"hwMon,omitempty"`
	File        *FileFanConfig     `json:"file,omitempty"`
	ControlLoop *ControlLoopConfig `json:"ControlLoop,omitempty"`
}

type HwMonFanConfig struct {
	Platform  string `json:"platform"`
	Index     int    `json:"index"`
	PwmOutput string
	RpmInput  string
}

type FileFanConfig struct {
	Path string `json:"path"`
}

type ControlLoopConfig struct {
	P float64 `json:"p"`
	I float64 `json:"i"`
	D float64 `json:"d"`
}
