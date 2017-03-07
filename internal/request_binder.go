package swagger

import (
	"net/http"
	"strconv"
	"net/url"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/eaglesakura/swagger-go-core"
	"github.com/eaglesakura/swagger-go-core/errors"
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
	}

	return errors.New(http.StatusBadRequest, fmt.Sprintf("Parameter parse error value[%v] type[%v]", value, resultType))
}

type BasicRequestBinder struct {
	Request         *http.Request
	ConsumerFactory func(contentType string) swagger.Consumer
	pathValues      map[string]string
	queryValues     url.Values
}

func (it *BasicRequestBinder)GetConsumer(contentType string) swagger.Consumer {
	return it.ConsumerFactory(contentType)
}

func (it *BasicRequestBinder)GetContentType() string {
	return it.Request.Header.Get("Content-Type")
}

func (it *BasicRequestBinder)BindPath(key string, resultType string, result interface{}) error {
	if it.pathValues == nil {
		it.pathValues = mux.Vars(it.Request)
	}

	value, ok := it.pathValues[key]
	if !ok {
		return errors.New(http.StatusBadRequest, fmt.Sprintf("PathValue not found key[%v] path[%v]", key, it.Request.URL))
	}

	return stringToValue(value, resultType, result)
}

func (it *BasicRequestBinder)BindQuery(key string, resultType string, result interface{}) error {
	if it.queryValues == nil {
		it.queryValues = it.Request.URL.Query()
	}

	return stringToValue(it.queryValues.Get(key), resultType, result)
}

func (it *BasicRequestBinder)BindHeader(key string, resultType string, result interface{}) error {
	return stringToValue(it.Request.Header.Get(key), resultType, result)
}

func (it *BasicRequestBinder)BindForm(key string, resultType string, result interface{}) error {

	if it.Request.Form == nil {
		err := it.Request.ParseForm()
		if err != nil {
			return err
		}
	}

	return stringToValue(it.Request.Form.Get(key), resultType, result)
}

func (it *BasicRequestBinder)BindBody(resultType string, result interface{}) error {
	consumer := it.GetConsumer(it.GetContentType())

	if consumer == nil {
		return errors.New(http.StatusBadRequest, "Body parse error[Unsupported consumer.]")
	}

	return consumer.Consume(it.Request.Body, result)
}
