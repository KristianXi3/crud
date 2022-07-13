package entity1

type Order struct {
	Order_id      int    //`json:Order_Id`
	Customer_name string //`json:Customer_Name`
	Ordered_at    string //`json:Ordered_At`
}
