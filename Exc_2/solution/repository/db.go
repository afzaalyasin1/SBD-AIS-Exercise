package repository

import "ordersystem/model"

type DatabaseHandler struct {
	// drinks represent all available drinks
	drinks []model.Drink
	// orders serves as order history
	orders []model.Order
}

// todo
func NewDatabaseHandler() *DatabaseHandler {
	// Init the drinks slice with some test data
	drinks := []model.Drink{
		{ID: 1, Name: "Coke", Price: 1.5, Description: "Cold Soda"},
		{ID: 2, Name: "Pepsi", Price: 1.4, Description: "Refreshing drink"},
		{ID: 3, Name: "Water", Price: 1.0, Description: "Pure water"},
	}
	// Init orders slice with some test data
	orders := []model.Order{
		{DrinkID: 1, Amount: 2},
		{DrinkID: 2, Amount: 1},
	}

	return &DatabaseHandler{
		drinks: drinks,
		orders: orders,
	}
}

func (db *DatabaseHandler) GetDrinks() []model.Drink {
	return db.drinks
}

func (db *DatabaseHandler) GetOrders() []model.Order {
	return db.orders
}

// todo
func (db *DatabaseHandler) GetTotalledOrders() map[uint64]uint64 {
	totalledOrders := make(map[uint64]uint64)
	// calculate total orders
	for _, order := range db.orders {
		totalledOrders[order.DrinkID] += uint64(order.Amount)
	}
	return totalledOrders
}

func (db *DatabaseHandler) AddOrder(order *model.Order) {
	db.orders = append(db.orders, *order)

}
