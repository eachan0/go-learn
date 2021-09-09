package computer

type Spec struct { //exported struct (导出结构体)
	Maker string //exported field (导出字段)
	model string //unexported field (未导出字段)
	Price int    //exported field (导出字段)
}
