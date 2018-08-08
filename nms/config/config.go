package nms

type Config struct {
	ProcessName *string                  `json:"proc-name"`
	Pointers    map[string]ConfigPointer `json:"pointers"`
}

type ConfigPointer struct {
	Module  *string   `json:"module-base"`
	Offsets []uintptr `json:"offsets"`
}
