package spring

type ReqCreateDto struct {
	Eid     string            `json:"e_id"`
	Method  string            `json:"method"`
	Content *CreateContentDto `json:"content"`
}

type ReqQueryDto struct {
	Eid     string           `json:"e_id"`
	Method  string           `json:"method"`
	Content *QueryContentDto `json:"content"`
}

type ReqCancelDto struct {
	Eid     string            `json:"e_id"`
	Method  string            `json:"method"`
	Content *CancelContentDto `json:"content"`
}

type RespPushQueryDto struct {
	Success bool                     `json:"success"`
	Result  []RespPushQueryResultDto `json:"result"`
	Error   struct {
		Code    int         `json:"code,omitempty"`
		Message string      `json:"message,omitempty"`
		Details interface{} `json:"details,omitempty"`
	} `json:"error"`
}

type RespCancelDto struct {
	Success bool     `json:"success"`
	Result  []string `json:"result"`
	Error   struct {
		Code    int         `json:"code,omitempty"`
		Message string      `json:"message,omitempty"`
		Details interface{} `json:"details,omitempty"`
	} `json:"error"`
}

type RespQueryDto struct {
	Success bool            `json:"success"`
	Result  RespQueryResult `json:"result"`
	Error   struct {
		Code    int         `json:"code,omitempty"`
		Message string      `json:"message,omitempty"`
		Details interface{} `json:"details,omitempty"`
	} `json:"error"`
}

type RespCreateDto struct {
	Success bool             `json:"success"`
	Result  RespCreateResult `json:"result"`
	Error   struct {
		Code    int         `json:"code,omitempty"`
		Message string      `json:"message,omitempty"`
		Details interface{} `json:"details,omitempty"`
	} `json:"error"`
}

//dto

type RespPushQueryResultDto struct {
	BillNo    string `json:"billno"`
	OrderNo   string `json:"orderno"`
	Time      string `json:"time"`
	Type      string `json:"type"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CancelContentDto struct {
	BillNo  string `json:"billno"`
	OrderNo string `json:"orderno"`
}

type QueryContentDto struct {
	BillNo  string `json:"billno"`
	OrderNo string `json:"orderno"`
}

type RespQueryResult struct {
	Result    string `json:"result"`
	ErrorCode string `json:"error_code"`
	Remark    string `json:"remark"`
	Info      []Info `json:"info"`
}

type Info struct {
	Time      string `json:"time"`
	BillNo    string `json:"billno"`
	StateInfo string `json:"state_info"`
	State     string `json:"state"`
}

type CreateContentDto struct {
	BillNo   string       `json:"billno"`
	Quantity int64        `json:"quantity"`
	Package  int64        `json:"package"`
	Volume   float64      `json:"volume"`
	Receiver *ReceiverDto `json:"receiver"`
	Sender   *SenderDto   `json:"sender"`
}
type ReceiverDto struct {
	Code    string `json:"code"`
	Name    string `json:"name"`
	Company string `json:"company"`
	City    string `json:"city"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	Mobile  string `json:"mobile"`
}
type SenderDto struct {
	Code    string `json:"code"`
	Name    string `json:"name"`
	Company string `json:"company"`
	City    string `json:"city"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	Mobile  string `json:"mobile"`
}
type RespCreateResult struct {
	Result    string `json:"result"`
	ErrorCode string `json:"error_code"`
	Remark    string `json:"remark"`
	Info      struct {
		OrderNo string `json:"orderno"`
		BillNo  string `json:"billno"`
	}
}
