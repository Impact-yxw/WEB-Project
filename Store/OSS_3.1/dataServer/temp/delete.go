package temp

import (
	"OSS/dataServer/conf"
	"github.com/gin-gonic/gin"
	"os"
)

func Delete(c *gin.Context) {
	uuid := c.Param("tempfile")[1:]
	infoFile := conf.Conf.Dir + "/temp/" + uuid
	dataFile := infoFile + ".dat"
	os.Remove(infoFile)
	os.Remove(dataFile)
}
