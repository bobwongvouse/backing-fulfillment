package service

import(
	"net/http"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

//getFulfillmentStatusHandler simulates actual fulfillment by supplying
//bogus values for QuantityInStock
//and ShipsWithin for any given SKU. Used to demonstrate a backing service
//supporting a primary service.

func getFulfillmentStatusHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

	}
}
