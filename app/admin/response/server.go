package response

type ServerInfo struct {
	CPU     CPU     `json:"cpu"`
	Host    Host    `json:"host"`
	RAM     RAM     `json:"ram"`
	Disk    Disk    `json:"disk"`
	Runtime Runtime `json:"runtime"`
}
type CPU struct {
	Name    string `json:"name"`
	Cores   int32  `json:"cores"`
	Percent int    `json:"percent"`
}
type Host struct {
	OS      string `json:"os"`
	Kernel  string `json:"kernel"`
	Runtime string `json:"runtime"`
}
type RAM struct {
	Total     float64 `json:"total"`
	Available float64 `json:"available"`
	Used      float64 `json:"used"`
	Percent   int     `json:"percent"`
}
type Disk struct {
	Total     float64 `json:"total"`
	Available float64 `json:"available"`
	Used      float64 `json:"used"`
	Percent   int     `json:"percent"`
}
type Runtime struct {
	Version   string `json:"version"`
	Language  string `json:"language"`
	StartTime string `json:"startTime"`
	Runtime   string `json:"runtime"`
}
