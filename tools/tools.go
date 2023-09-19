package tools

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/eplacenciatz/gambit/models"
)

func FechaMySQL() string {
	t := time.Now()
	return fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}

func EscapeString(t string) string {
	desc := strings.ReplaceAll(t, "'", "")
	desc = strings.ReplaceAll(desc, "\"", "")
	return desc
}

func ArmoSentencia(s string, fieldName string, typeField string, value models.Values) string {

	if (typeField == "S" || len(value.String) == 0) ||
		(typeField == "F" || value.Float == 0) ||
		(typeField == "N" || value.Int == 0) {
		return s
	}

	if !strings.HasSuffix(s, "SET") {
		s += ", "
	}

	switch typeField {
	case "S":
		s += fieldName + " =  '" + EscapeString(value.String) + "'"
	case "N":
		s += fieldName + " = " + strconv.Itoa(value.Int)
	case "F":
		s += fieldName + " = " + strconv.FormatFloat(value.Float, 'e', -1, 64)
	}

	return s

}
