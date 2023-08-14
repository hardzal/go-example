package productcontroller

import (
	"net/http"

	"github.com/hardzal/go-example/go-books/pkg/utils"
)

func Index(w http.ResponseWriter, r *http.Request) {
	data := []map[string]interface{}{
		{
			"id":    1,
			"name":  "Baju 1",
			"stock": 100,
		},
		{
			"id":    2,
			"name":  "Baju 2",
			"stock": 150,
		},
		{
			"id":    3,
			"name":  "Baju 3",
			"stock": 200,
		},
	}

	utils.ResponseJSON(w, http.StatusOK, data)
}
