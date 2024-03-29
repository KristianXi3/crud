package DB

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/KristianXi3/crud/entity1"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func (s *Dbstruct) GetUsers(ctx context.Context) ([]entity1.User, error) {
	var result []entity1.User

	rows, err := s.SqlDb.QueryContext(ctx, "select id, username, email, password, age, createddate, updatedate from users")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var row entity1.User
		err := rows.Scan(
			&row.Id,
			&row.Username,
			&row.Email,
			&row.Password,
			&row.Age,
			&row.CreatedAt,
			&row.UpdatedAt,
		)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		result = append(result, row)
	}
	return result, nil
}

func (s *Dbstruct) GetUserByID(ctx context.Context, userid int) (*entity1.User, error) {
	result := &entity1.User{}

	rows, err := s.SqlDb.QueryContext(ctx, "select id, username, email, password, age, createddate, updatedate from users where id = @ID", sql.Named("ID", userid))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(
			&result.Id,
			&result.Username,
			&result.Email,
			&result.Password,
			&result.Age,
			&result.CreatedAt,
			&result.UpdatedAt,
		)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
	}
	return result, nil
}

func (s *Dbstruct) CreateUser(ctx context.Context, user entity1.User) (result string, err error) {
	x, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(x)
	_, err = s.SqlDb.ExecContext(ctx, "insert into users (id, username, email, password, age, createddate, updatedate) values (@id, @username, @email, @password, @age, @now, @now)",
		sql.Named("id", user.Id),
		sql.Named("username", user.Username),
		sql.Named("email", user.Email),
		sql.Named("password", x),
		sql.Named("age", user.Age),
		sql.Named("now", time.Now()),
	)
	if err != nil {
		return "", err
	}

	result = "Inserted"

	return result, nil
}

func (s *Dbstruct) UpdateUser(ctx context.Context, userId int, user entity1.User) (result string, err error) {

	_, err = s.SqlDb.ExecContext(ctx, "update users set username = @username,email = @email, password = @password, age = @age, updatedate = @now where id = @id",
		sql.Named("id", userId),
		sql.Named("username", user.Username),
		sql.Named("email", user.Email),
		sql.Named("password", user.Password),
		sql.Named("age", user.Age),
		sql.Named("now", time.Now()))
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	result = "Updated"

	return result, nil
}

func (s *Dbstruct) DeleteUser(ctx context.Context, userId int) (result string, err error) {

	_, err = s.SqlDb.ExecContext(ctx, "delete from users where id=@id", sql.Named("id", userId))
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	result = "Deleted"

	return result, nil
}

//This part is for Order entity
func (s *Dbstruct) GetOrders(ctx context.Context) ([]entity1.OrderWithItems, error) {
	var result []entity1.OrderWithItems
	var result1 []entity1.Order
	var result21 []entity1.ItemWithOrderID

	rows, err := s.SqlDb.QueryContext(ctx, "EXEC dbo.usp_Get_Order")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var row entity1.Order
		err := rows.Scan(
			&row.Order_id,
			&row.Customer_name,
			&row.Ordered_at,
		)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		fmt.Println("Result Data1")
		result1 = append(result1, row)

	}
	if !rows.NextResultSet() {
		log.Fatal("[mssql] Expected more resultset")
		return nil, errors.New("[mssql] Expected more resultset")
	}
	for rows.Next() {
		var result2 entity1.ItemWithOrderID
		err := rows.Scan(
			&result2.Item_Code,
			&result2.Description,
			&result2.Quantity,
			&result2.OrderId,
		)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		fmt.Println("Result Data")
		fmt.Println(result2)
		result21 = append(result21, result2)
	}
	for _, o := range result1 {
		var tempOrder entity1.OrderWithItems
		var tempItems []entity1.Items
		tempOrder.Order = o
		for _, i := range result21 {
			if tempOrder.Order_id == i.OrderId {
				tempItems = append(tempItems, *i.ToWithoutOrderID())
			}
		}
		tempOrder.Item = tempItems
		result = append(result, tempOrder)
	}

	return result, nil
}

func (s *Dbstruct) GetOrderByID(ctx context.Context, orderid int) (*entity1.OrderWithItems, error) {
	result := &entity1.OrderWithItems{}
	result1 := &entity1.Order{}
	var result21 []entity1.ItemWithOrderID

	rows, err := s.SqlDb.QueryContext(ctx, "EXEC dbo.usp_Get_Order @ID = @ID", sql.Named("ID", orderid))
	//rows, err := s.SqlDb.QueryContext(ctx, "select order_id,customer_name,ordered_at from orders where @ID = @ID; select item_id", sql.Named("ID", orderid))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {

		err := rows.Scan(
			&result1.Order_id,
			&result1.Customer_name,
			&result1.Ordered_at,
		)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		fmt.Println("Result Data1")

	}
	if !rows.NextResultSet() {
		log.Fatal("[mssql] Expected more resultset")
		return nil, errors.New("[mssql] Expected more resultset")
	}
	for rows.Next() {
		var result2 entity1.ItemWithOrderID
		err := rows.Scan(
			&result2.Item_Code,
			&result2.Description,
			&result2.Quantity,
			&result2.OrderId,
		)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		fmt.Println("Result Data")
		fmt.Println(result2)
		result21 = append(result21, result2)
	}

	var tempItems []entity1.Items
	for _, i := range result21 {
		if result1.Order_id == i.OrderId {
			tempItems = append(tempItems, *i.ToWithoutOrderID())
		}
	}
	result.Order = *result1
	result.Item = tempItems

	return result, nil
}

func (s *Dbstruct) CreateOrder(ctx context.Context, order entity1.Order) (string, error) {
	var result string
	data, err := s.SqlDb.QueryContext(ctx, "EXEC dbo.usp_Create_Order @CheckUpdate = @checkupdate, @CustomerName = @customer_name,@OrderedAt =  @ordered_at",
		sql.Named("checkupdate", 0),
		sql.Named("customer_name", order.Customer_name),
		sql.Named("ordered_at", time.Now()))
	if err != nil {
		log.Fatal(err)
		return fmt.Sprintf("Internal Server Error: %s", err.Error()), err
	}

	var lastOrderId int
	for data.Next() {
		err := data.Scan(&lastOrderId)
		if err != nil {
			log.Fatal(err)
			return fmt.Sprintf("Internal Server Error: %s", err.Error()), err
		}
	}
	for i := 0; i < len(order.Item); i++ {
		_, err = s.SqlDb.QueryContext(ctx, "EXEC dbo.usp_Create_Order @CheckUpdate = @checkupdate2 , @ItemCode = @code, @IDesc =  @description, @IQty = @quantity, @OrderId = @order_id",
			sql.Named("checkupdate2", 1),
			sql.Named("code", order.Item[i].Item_code),
			sql.Named("description", order.Item[i].Description),
			sql.Named("quantity", order.Item[i].Quantity),
			sql.Named("order_id", lastOrderId))
		if err != nil {
			log.Fatal(err)
			return fmt.Sprintf("Internal Server Error: %s", err.Error()), err
		}

	}

	result = "Order created successfully"
	return result, nil
}

// func (s *Dbstruct) UpdateOrder(ctx context.Context, orderid int, order entity1.Order, item entity1.Items) (result entity1.Order, err error) {
func (s *Dbstruct) UpdateOrder(ctx context.Context, orderid int, order entity1.Order2) (string, error) {
	end := ""
	_, err := s.SqlDb.ExecContext(ctx, "EXEC dbo.usp_Update_Order @CheckUpdate = 0, @CustomerName = @customer_name, @OrderedAt = @ordered_at,@OrderId = @order_id",
		sql.Named("customer_name", order.Customer_name),
		sql.Named("ordered_at", time.Now()),
		sql.Named("order_id", orderid))

	if err != nil {
		log.Fatal(err)
		return fmt.Sprintf("Internal Server Error: %s", err.Error()), err
	}

	for i := 0; i < len(order.Item); i++ {
		var item entity1.ItemsWithoutOrderId
		exist := false
		data, err := s.SqlDb.QueryContext(ctx, "EXEC dbo.usp_Get_Item @ID = @itemId, @OrderId = @order_id ",
			sql.Named("order_id", orderid),
			sql.Named("itemId", order.Item[i].Item_id))

		if err != nil {
			log.Fatal(err)
			return fmt.Sprintf("Internal Server Error: %s", err.Error()), err
		}
		defer data.Close()

		for data.Next() {
			err := data.Scan(&item.Item_id, &item.Item_code, &item.Description, &item.Quantity)
			if err != nil {
				log.Fatal(err)
				return fmt.Sprintf("Internal Server Error: %s", err.Error()), err
			}
		}

		if (item != entity1.ItemsWithoutOrderId{}) {
			_, err := s.SqlDb.ExecContext(ctx, "EXEC dbo.usp_Update_Order @CheckUpdate = 1, @ItemCode = @item_code ,@IDesc = @description , @IQty = @quantity , @ItemID = @item_id",
				sql.Named("item_code", order.Item[i].Item_code),
				sql.Named("description", order.Item[i].Description),
				sql.Named("quantity", order.Item[i].Quantity),
				sql.Named("item_Id", item.Item_id))

			if err != nil {
				log.Fatal(err)
				return fmt.Sprintf("Internal Server Error: %s", err.Error()), err
			}
			exist = true
		}

		if !exist {
			end += fmt.Sprintf("Item id %d not found \n", order.Item[i].Item_id)
		}
	}
	end += "Updated successfully"
	return end, nil
}

func (s *Dbstruct) DeleteOrder(ctx context.Context, orderId int) (result string, err error) {

	_, err = s.SqlDb.ExecContext(ctx, "EXEC dbo.usp_Delete_Order @OrderId = @id", sql.Named("id", orderId))
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	result = "Deleted"

	return result, nil
}
func (s *Dbstruct) LoginsUser(ctx context.Context, cred entity1.Credentials) (string, error) {
	var user entity1.User
	// var cred entity1.Credentials
	var jwtKey = []byte("my_secret_key")

	rows, err := s.SqlDb.QueryContext(ctx, "SELECT username,password FROM users WHERE username = @username",
		sql.Named("username", cred.Username))
	if err != nil {
		log.Fatal(err)
		return fmt.Sprintf("Internal Server Error: %s", err.Error()), err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(
			&user.Username,
			&user.Password)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println(user.Password, cred.Password)
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(cred.Password))
	if err != nil {
		log.Fatal(err)
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &entity1.Claims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {

		return "", err
	}
	return tokenString, nil
}
