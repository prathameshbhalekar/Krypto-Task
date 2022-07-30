package schemas

type Alert struct {
	AlertUUID   string  `json:"alert_id"`
	Email       string  `json:"email"`
	GreaterThan bool    `json:"greater_than"`
	AlertValue  float32 `json:"alert_value"`
	Status      string  `json:"status"`
	UserUUID    string  `json:"user_uuid"`
}

type CoinGeckoResponse struct {
	Name         string  `json:"name"`
	CurrentPrice float32 `json:"current_price"`
}

type CacheKey struct {
	UserUUID string `json:"user_uuid"`
	PageNo   int    `json:"page_no"`
	PageSize int    `json:"page_size"`
	Status   string `json:"status"`
}

type WebsocketResponse struct {
	C string `json:"c"`
}
