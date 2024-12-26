package define

type CommonSuccessRsp struct {
	Status int `json:"status" default:"1"`
}

type CommonSeccessDataRsp struct {
	Status int         `json:"status" default:"1"`
	Data   interface{} `json:"data"`
}

type CommonFailRsp struct {
	Status int    `json:"status" default:"0"`
	Code   int    `json:"code"`
	Errmsg string `json:"errmsg"`
}
