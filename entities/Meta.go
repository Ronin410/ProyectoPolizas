package entities

type Meta struct {
	Status string
}

type MetaSSO struct {
	Status string `json:"status"`
	Count  int
	Error  Error
}

type Error struct {
	ErrorCode   int    `json:"errorCode"`
	UserMessage string `json:"userMessage"`
}
