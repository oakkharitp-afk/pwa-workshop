package main

import (
	"log"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// ตรวจ method (basic error handling)
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// set header station
	w.WriteHeader(http.StatusOK)

	// response เป็น text
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("hello world, this is simple string"))

	// resp := Response{
	// 	Message: "Hello from Go HTTP Server",
	// }

	// // encode json
	// if err := json.NewEncoder(w).Encode(resp); err != nil {
	// w.Header().Set("Content-Type", "application/json")
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }
}

func main() {

	// สร้าง http server แบบปกติ
	http.HandleFunc("/hello", helloHandler)

	log.Println("server started at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

	// สร้าง http server แบบใช้ framework
	// e := echo.New()
	// // route: GET /hello
	// e.GET("/hello", func(c echo.Context) error {
	// 	resp := Response{
	// 		Message: "Hello from Echo server",
	// 	}

	// 	return c.JSON(http.StatusOK, resp)
	// })

	// // start server
	// e.Logger.Fatal(e.Start(":8080"))
}
