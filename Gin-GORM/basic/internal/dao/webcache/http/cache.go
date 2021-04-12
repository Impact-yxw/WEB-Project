package http

import (
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) CacheCheck(c *gin.Context) {
	key := c.Param("key")
	if key == "" {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	m := c.Request.Method

	if m == http.MethodGet {
		b, _ := s.Get(key)
		if len(b) == 0 {
			c.Set("miss", true) // 需要查数据库
			return
		}

		c.JSON(http.StatusOK, string(b))
		c.Set("miss", false) // 不需要查数据库
		return
	}

	c.JSON(http.StatusMethodNotAllowed, nil)
}

func (s *Server) UpdateHandler(c *gin.Context) {
	key := c.Param("key")

	if key == "" {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	m := c.Request.Method
	if m == http.MethodPut {
		b, _ := io.ReadAll(c.Request.Body)
		if len(b) != 0 {
			e := s.Set(key, b)
			if e != nil {
				log.Println(e)
				c.JSON(http.StatusInternalServerError, nil)
			}
		}
		return
	}

	if m == http.MethodDelete {
		e := s.Del(key)
		if e != nil {
			log.Println(e)
			c.JSON(http.StatusInternalServerError, nil)
		}
		return
	}

	c.JSON(http.StatusMethodNotAllowed, nil)
}
