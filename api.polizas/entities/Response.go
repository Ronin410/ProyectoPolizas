package entities

type Response struct {
	Meta Meta
	Data any
}

type ResponseSSO struct {
	Meta MetaSSO `json:"meta"`
	Data string  `json:"data"`
}
