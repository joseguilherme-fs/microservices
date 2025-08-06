package api
3
4 import (
5 "github.com/joseguilherme-fs/microservices/order/internal/application/core/domain"
6 "github.com/joseguilherme-fs/microservices/order/internal/ports"
7 )
8
9 type Application struct {
10 	db ports . DBPort
11 }
12
13 func NewApplication ( db ports . DBPort ) * Application {
14 	return & Application {
15 		db : db ,
16 	}
17 }
18
19 func ( a Application ) PlaceOrder ( order domain . Order ) ( domain . Order , error ) {
20 	err := a . db . Save (& order )
21 	if err != nil {
22 		return domain . Order {} , err
23 	}
24 	return order, nil 
}