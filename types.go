package main

type userInfo struct {
	ID            int64   `json:"id"`
	Username      *string `json:"username"`
	Alias         *string `json:"alias"`
	ResellerAlias *string `json:"reseller_alias"`
	AuthMax       *int    `json:"auth_max"`
	AuthLeft      *int    `json:"auth_left"`
	DeauthLeft    *int    `json:"deauth_left"`
	Reseller      *int64  `json:"reseller"`
}

type liveSession struct {
	ID      *int64    `json:"id"`
	UserID  int64     `json:"user_id"`
	URL     *string   `json:"url"`
	Title   *string   `json:"title"`
	Host    *string   `json:"host"`
	Comment *string   `json:"comment"`
	Begin   *float64  `json:"begin"`
	End     *float64  `json:"end"`
	Tags    *[]string `json:"tags"`
}

type customerInfo struct {
	ID           *int64    `json:"id"`
	UserID       int64     `json:"user_id"`
	CustomerName *string   `json:"customer_name"`
	Mobile       *string   `json:"mobile"`
	Status       *string   `json:"status"`
	Tags         *[]string `json:"tags"`
}

type liveActivity struct {
	ID           *int64        `json:"id"`
	UserID       int64         `json:"user_id"`
	LiveID       *int64        `json:"live_id"`
	Time         *float64      `json:"time"`
	CustomerID   *int64        `json:"customer_id"`
	CustomerInfo *customerInfo `json:"customer_info"`
	Activity     *string       `json:"activity"`
}
