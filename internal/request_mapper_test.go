package swagger

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/eaglesakura/swagger-go-core/interface"
	"github.com/eaglesakura/swagger-go-core/utils"
)

func Test_MethodMapper_PutHandler(t *testing.T) {
	mapper, _ := swagger.NewHandlerMapper().(*HandleMapperImpl)
	assert.NotNil(t, mapper)

	mapper.PutHandler(swagger.HandleRequest{
		Path:"/api/v1/test",
		Method:"GET",
	})
	mapper.PutHandler(swagger.HandleRequest{
		Path:"/api/v1/test/{userId}",
		Method:"GET",
	})
	mapper.PutHandler(swagger.HandleRequest{
		Path:"/api/v1/test/{userId}",
		Method:"POST",
	})
	mapper.PutHandler(swagger.HandleRequest{
		Path:"/api/v1/debuggers",
		Method:"GET",
	})

	listMappers := mapper.ListMethodMappers()
	assert.NotNil(t, listMappers)
	assert.Equal(t, listMappers.Len(), 3)        // POST/GETで同じURLはマージされる

	for _, value := range listMappers {
		assert.NotNil(t, value)
	}

	assert.Equal(t, listMappers[0].Path, "/api/v1/test/{userId}")
	assert.Equal(t, listMappers[2].Path, "/api/v1/test")
}
