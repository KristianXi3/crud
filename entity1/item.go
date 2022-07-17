package entity1

type Items struct {
	Item_id     int    `json:"Item_Id"`
	Item_code   string `json:"Item_Code"`
	Description string `json:"Description"`
	Quantity    int    `json:"Quantity"`
	Order_id    int    `json:"Order_Id"`
}
type ItemsWithoutOrderId struct {
	Item_id     int    `json:"Item_Id"`
	Item_code   string `json:"Item_Code"`
	Description string `json:"Description"`
	Quantity    int    `json:"Quantity"`
}
type ItemWithOrderID struct {
	Item_Code   string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	OrderId     int    `json:"order_id"`
}

func (i ItemWithOrderID) ToWithoutOrderID() *Items {
	retval := &Items{
		Item_code:   i.Item_Code,
		Description: i.Description,
		Quantity:    i.Quantity,
	}
	return retval
}
