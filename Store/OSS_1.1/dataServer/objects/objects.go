package objects

import (
	"OSS/dataServer/conf"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func Put(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func put(w http.ResponseWriter, r *http.Request) {
	f, err := os.Create(conf.Conf.Dir + "/objects/" + strings.Split(r.URL.EscapedPath(), "/")[2])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer f.Close()
	io.Copy(f, r.Body)
}

func Get(c *gin.Context) {

}
func get(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open(conf.Conf.Dir + "/objects/" + strings.Split(r.URL.EscapedPath(), "/")[2])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	defer f.Close()
	io.Copy(w, f)
}
