package cart

type RequestCart struct {
	Product_ID int `json:"product_id" form:"product_id"`
	Qty        int `json:"qty" form:"qty"`
}

type ResponseCart struct {
	ID          uint `json:"id"`
	User_ID     int  `json:"user_id"`
	Product_ID  int  `json:"product_id"`
	Qty         int  `json:"qty"`
	Total_price int  `json:"total_price"`
}

type ResponseGetCart struct {
	ID          uint `json:"id"`
	User_ID     int  `json:"user_id"`
	Product_ID  int  `json:"product_id"`
	Qty         int  `json:"qty"`
	Total_price int  `json:"total_price"`
}
