package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-playground/validator"

	"github.com/labstack/echo"
)

// CustomValidator validate api request
type CustomValidator struct {
	validator *validator.Validate
}

// Validate validate struct
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
func TestCalculateLocation(t *testing.T) {
	// Setup
	dnsServiceRequest := `{"x":"123.12","y":"456.56","z":"789.89","vel":"20.0"}`
	expectedResponse := `{"loc":1389.57}`
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	// fmt.Println(userJSON)
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(dnsServiceRequest))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if err := CalculateLocation(c); err == nil {
		if expectedResponse != rec.Body.String() {
			t.Error("wrong location")
		}
	} else {
		t.Error(err.Error())
	}
}
func TestWrongLocation(t *testing.T) {
	// Setup
	dnsServiceRequest := `{"x":"123.12","y":"456.56","z":"789.89","vel":"20.0"}`
	expectedResponse := `{"loc":189.57}`
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	// fmt.Println(userJSON)
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(dnsServiceRequest))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if err := CalculateLocation(c); err == nil {
		if expectedResponse == rec.Body.String() {
			t.Error("wrong location")
		}
	} else {
		t.Error(err.Error())
	}
}
func TestRequestParameterMissing(t *testing.T) {
	// Setup
	dnsServiceRequest := `{"x":"123.12","y":"456.56","z":"789.89"}`
	expectedResponse := `Key: 'CalculateRequestDTO.Vel' Error:Field validation for 'Vel' failed on the 'required' tag`
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	// fmt.Println(userJSON)
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(dnsServiceRequest))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if err := CalculateLocation(c); err != nil {
		if expectedResponse != rec.Body.String() {
			t.Error("Vel paremter missing")
		}
	}
}
func TestInvalidRequest(t *testing.T) {
	tests := []struct {
		name, input, expected string
	}{
		{"InValid Json", `{"x":"123.12","y":"456.56","z":"789.89",}`, "code=400, message=Syntax error: offset=41, error=invalid character '}' looking for beginning of object key string"},
		{"InValid Value", `{"x":"123.12a","y":"456.56","z":"789.89","vel":"20"}`, `{"loc":1266.45}`},
	}
	// Setup
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	for _, test := range tests {
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(test.input))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		if err := CalculateLocation(c); err == nil {
			if !strings.Contains(test.expected, rec.Body.String()) {
			}
		}
	}

}
func BenchmarkShareWith(b *testing.B) {
	dnsServiceRequest := `{"x":"123.12","y":"456.56","z":"789.89","vel":"20.0"}`
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	// fmt.Println(userJSON)
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(dnsServiceRequest))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	for i := 0; i < b.N; i++ {
		CalculateLocation(c)
	}
}
