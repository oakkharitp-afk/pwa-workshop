package handler

import (
	"net/http"
	"pwa-workshop/project-structure/errs"
	"strings"

	"github.com/labstack/echo/v4"
)

// Error handles errors in Echo HTTP handlers by returning a JSON response.
// If the error implements the errs.Error type, it responds with the error's code and detail.
// Otherwise, it returns a 500 Internal Server Error with the error message as the detail.
//
// ฟังก์ชัน Error ใช้สำหรับจัดการข้อผิดพลาดใน Echo HTTP handler โดยจะคืนค่าเป็น JSON response
// หาก error ที่รับเข้ามาเป็นชนิด errs.Error จะตอบกลับด้วย code และ detail ของ error นั้น
// หากไม่ใช่ จะตอบกลับด้วย 500 Internal Server Error และแสดงข้อความ error ใน detail
func Error(c echo.Context, err error) error {
	if e, ok := err.(*errs.Exception); ok {
		return c.JSON(e.Code, &errs.Exception{
			Code:   e.Code,
			Detail: upperPrefix(e.Detail),
		})
	}
	return c.JSON(http.StatusInternalServerError, errs.Exception{
		Code:   http.StatusInternalServerError,
		Detail: upperPrefix(err.Error()),
	})
}

func upperPrefix(s string) string {
	if len(s) == 0 {
		return s
	}
	runes := []rune(s)
	first := strings.ToUpper(string(runes[0]))
	remain := string(runes[1:])
	return first + remain
}
