package domain
2
3 import " time "
4
5 type OrderItem struct {
6 ProductCode string ‘json :" product_code "‘
7 UnitPrice float32 ‘json :" unit_price "‘
8 Quantity int32 ‘json :" quantity "‘
9 }
10
11 type Order struct {
12 ID int64 ‘json : "id"‘
13 CustomerID int64 ‘json :" customer_id "‘
14 Status string ‘json :" status "‘
15 OrderItems [] OrderItem ‘json :" order_items "‘
16 CreatedAt int64 ‘json :" created_at "‘
17 }
18
19 func NewOrder ( customerId int64 , orderItems [] OrderItem ) Order {
20 return Order {
21 CreatedAt : time . Now () . Unix () ,
22 Status : " Pending ",
23 CustomerID : customerId ,
24 OrderItems : orderItems ,
25 }
26 }