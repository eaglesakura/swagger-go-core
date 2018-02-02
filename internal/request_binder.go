package swagger

import (
	"fmt"
	"github.com/eaglesakura/swagger-go-core"
	"github.com/eaglesakura/swagger-go-core/errors"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

func stringToValue(value string, resultType string, result interface{}) error {
	if resultType == "string" {
		if ptr, ok := result.(**string); ok {
			*ptr = &value
			return nil
		} else {
			return errors.New(http.StatusBadRequest, "Parameter parse error")
		}
	}

	switch resultType {
	case "string":
		if ptr, ok := result.(**string); ok {
			*ptr = &value
			return nil
		}
	case "int", "int8", "int16", "int32", "int64":
		value, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}

		switch resultType {
		case "int8":
			if ptr, ok := result.(**int8); ok {
				temp := int8(value)
				*ptr = &temp
				return nil
			}
		case "int16":
			if ptr, ok := result.(**int16); ok {
				temp := int16(value)
				*ptr = &temp
				return nil
			}
		case "int32":
			if ptr, ok := result.(**int32); ok {
				temp := int32(value)
				*ptr = &temp
				return nil
			}
		case "int64":
			if ptr, ok := result.(**int64); ok {
				temp := int64(value)
				*ptr = &temp
				return nil
			}
		case "int":
			if ptr, ok := result.(**int); ok {
				temp := int(value)
				*ptr = &temp
				return nil
			}
		}
	case "float", "float32", "float64":
		value, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}

		switch resultType {
		case "float", "float32":
			if ptr, ok := result.(**float32); ok {
				temp := float32(value)
				*ptr = &temp
				return nil;
			}
		case "float64":
			if ptr, ok := result.(**float64); ok {
				temp := float64(value)
				*ptr = &temp
				return nil;
			}
		}
	case "bool":
		value, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}

		if ptr, ok := result.(**bool); ok {
			*ptr = &value
			return nil
		}
	}

	return errors.New(http.StatusBadRequest, fmt.Sprintf("Parameter parse error value[%v] type[%v]", value, resultType))
}

type BasicRequestBinder struct {
	Request         *http.Request
	BodyReader      io.Reader
	ConsumerFactory func(contentType string) swagger.Consumer
	pathValues      map[string]string
	queryValues     url.Values
}

func NewRequestBinder(req *http.Request, bodyReader io.Reader, consumerFactory func(contentType string) swagger.Consumer) swagger.RequestBinder {
	return &BasicRequestBinder{
		Request:         req,
		BodyReader:      bodyReader,
		ConsumerFactory: consumerFactory,
	}
}

func (it *BasicRequestBinder) GetConsumer(contentType string) swagger.Consumer {
	return it.ConsumerFactory(contentType)
}

func (it *BasicRequestBinder) GetContentType() string {
	return it.Request.Header.Get("Content-Type")
}

func (it *BasicRequestBinder) BindPath(key string, resultType string, result interface{}) error {
	if it.pathValues == nil {
		it.pathValues = mux.Vars(it.Request)
	}

	value, ok := it.pathValues[key]
	if !ok || len(value) == 0 {
		return errors.New(http.StatusBadRequest, fmt.Sprintf("PathValue not found key[%v] path[%v]", key, it.Request.URL))
	}

	return stringToValue(value, resultType, result)
}

func (it *BasicRequestBinder) BindQuery(key string, resultType string, result interface{}) error {
	if it.queryValues == nil {
		it.queryValues = it.Request.URL.Query()
	}

	value := it.queryValues.Get(key)
	if len(value) == 0 {
		//return errors.New(http.StatusBadRequest, fmt.Sprintf("QuryValue not found key[%v] path[%v]", key, it.Request.URL))
		return nil
	}
	return stringToValue(value, resultType, result)
}

func (it *BasicRequestBinder) BindHeader(key string, resultType string, result interface{}) error {
	value := it.Request.Header.Get(key)
	if len(value) == 0 {
		//return errors.New(http.StatusBadRequest, fmt.Sprintf("HeaderValue not found key[%v] path[%v]", key, it.Request.URL))
		return nil
	}
	return stringToValue(value, resultType, result)
}

func (it *BasicRequestBinder) BindForm(key string, resultType string, result interface{}) error {
	if it.Request.Form == nil {
		err := it.Request.ParseForm()
		if err != nil {
			return err
		}
	}

	value := it.Request.Form.Get(key)
	if len(value) == 0 {
		//return errors.New(http.StatusBadRequest, fmt.Sprintf("FormValue not found key[%v] path[%v]", key, it.Request.URL))
		return nil
	}

	return stringToValue(value, resultType, result)
}

func (it *BasicRequestBinder) BindBody(resultType string, result interface{}) error {
	consumer := it.GetConsumer(it.GetContentType())

	if consumer == nil {
		return errors.New(http.StatusBadRequest, "Body parse error[Unsupported consumer.]")
	}

	return consumer.Consume(it.BodyReader, result)
}

/*
Validatorを生成する
*/
func (it *BasicRequestBinder) NewValidator(value interface{}, isNil bool) swagger.ParameterValidator {
	return &ParameterValidatorImpl{
		Value: value,
		IsNil: isNil,
	}
}
