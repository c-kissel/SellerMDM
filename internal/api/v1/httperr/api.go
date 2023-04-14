package httperr

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

// Takes response writer, HTTP status code, formatted message, params and sends it to logs and to writer
// Example: sendError(w, http.StatusBadRequest, "invalid params: %s and %s", param1, param2)
func Send(w http.ResponseWriter, code int, message string, params ...interface{}) {
	output := fmt.Sprintf(message, params...)
	logrus.Error(output)
	http.Error(w, output, code)
}
