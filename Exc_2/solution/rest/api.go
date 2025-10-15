package rest

import (
	"encoding/json"
	"net/http"
	"ordersystem/model"
	"ordersystem/repository"

	"github.com/go-chi/render"
)

// GetMenu 			godoc
// @tags 			Menu
// @Description 	Returns the menu of all drinks
// @Produce  		json
// @Success 		200 {array} model.Drink
// @Router 			/api/menu [get]
func GetMenu(db *repository.DatabaseHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// todo
		// get slice from db
		drinks := db.GetDrinks()
		render.Status(r, http.StatusOK)
		render.JSON(w, r, drinks)
	}
}

// todo create GetOrders /api/order/all

// GetOrders 		godoc
// @tags 			Order
// @Description 	Returns all orders
// @Produce  		json
// @Success 		200 {array} model.Order
// @Router 			/api/order/all [get]
func GetOrders(db *repository.DatabaseHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get all orders from db
		orders := db.GetOrders()

		render.Status(r, http.StatusOK)
		render.JSON(w, r, orders)
	}
}

// todo create GetOrdersTotal /api/order/total

// GetOrdersTotal 	godoc
// @tags 			Order
// @Description 	Returns totalled orders per drink
// @Produce  		json
// @Success 		200 {object} map[uint64]uint64
// @Router 			/api/order/total [get]
func GetOrdersTotal(db *repository.DatabaseHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get totalled orders from db
		totalled := db.GetTotalledOrders()

		render.Status(r, http.StatusOK)
		render.JSON(w, r, totalled)
	}
}

// PostOrder 		godoc
// @tags 			Order
// @Description 	Adds an order to the db
// @Accept 			json
// @Param 			b body model.Order true "Order"
// @Produce  		json
// @Success 		200
// @Failure     	400
// @Router 			/api/order [post]
func PostOrder(db *repository.DatabaseHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// todo
		var order model.Order
		err := json.NewDecoder(r.Body).Decode(&order)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, map[string]string{"error": "invalid request"})
			return
		}
		db.AddOrder(&order)

		// add to db
		render.Status(r, http.StatusOK)
		render.JSON(w, r, "ok")
	}
}
