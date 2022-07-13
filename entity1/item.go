package entity1

type Items struct {
	Item_id     int    `json:Item_Id`
	Item_code   string `json:Item_Code`
	Description string `json:Description`
	Quantity    int    `json:Quantity`
	Order_id    int    `json:Order_Id`
}
