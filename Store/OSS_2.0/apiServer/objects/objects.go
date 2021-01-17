package objects

import (
	"OSS/apiServer/es"
	"OSS/apiServer/heartbeat"
	"OSS/apiServer/locate"
	"OSS/apiServer/objectstream"
	"OSS/apiServer/utils"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

//func get(w http.ResponseWriter, r *http.Request) {
//	object := strings.Split(r.URL.EscapedPath(), "/")[3]
//	log.Println(object)
//	stream, e := getStream(object)
//	if e != nil {
//		log.Println(e)
//		w.WriteHeader(http.StatusNotFound)
//		return
//
//	}
//	io.Copy(w, stream)
//
//}

func get(w http.ResponseWriter, r *http.Request) {
	name := strings.Split(r.URL.EscapedPath(), "/")[3]
	versionId := r.URL.Query()["version"]
	version := 0
	var e error
	if len(versionId) != 0 {
		version, e = strconv.Atoi(versionId[0])
		if e != nil {
			log.Println(e)
			w.WriteHeader(http.StatusBadRequest)
			return

		}

	}
	meta, e := es.GetMetadata(name, version)
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
		return

	}
	if meta.Hash == "" {
		w.WriteHeader(http.StatusNotFound)
		return

	}
	object := url.PathEscape(meta.Hash)
	stream, e := getStream(object)
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusNotFound)
		return

	}
	io.Copy(w, stream)

}

func getStream(object string) (io.Reader, error) {
	server := locate.Locate(object)
	if server == "" {
		return nil, fmt.Errorf("object %s locate fail", object)

	}
	return objectstream.NewGetStream(server, object)

}

//func put(w http.ResponseWriter, r *http.Request) {
//	object := strings.Split(r.URL.EscapedPath(), "/")[3]
//	c, e := storeObject(r.Body, object)
//	if e != nil {
//		log.Println(e)
//
//	}
//	w.WriteHeader(c)
//
//}

func put(w http.ResponseWriter, r *http.Request) {
	hash := utils.GetHashFromHeader(r.Header)
	if hash == "" {
		log.Println("missing object hash in digest header")
		w.WriteHeader(http.StatusBadRequest)
		return

	}

	c, e := storeObject(r.Body, url.PathEscape(hash))
	if e != nil {
		log.Println(e)
		w.WriteHeader(c)
		return

	}
	if c != http.StatusOK {
		w.WriteHeader(c)
		return

	}

	name := strings.Split(r.URL.EscapedPath(), "/")[3]
	size := utils.GetSizeFromHeader(r.Header)
	e = es.AddVersion(name, hash, size)
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)

	}

}

func putStream(object string) (*objectstream.PutStream, error) {
	server := heartbeat.ChooseRandomDataServer()
	if server == "" {
		return nil, fmt.Errorf("cannot find any dataServer")

	}

	return objectstream.NewPutStream(server, object), nil

}

func storeObject(r io.Reader, object string) (int, error) {
	stream, e := putStream(object)
	if e != nil {
		return http.StatusServiceUnavailable, e

	}

	io.Copy(stream, r)
	e = stream.Close()
	if e != nil {
		return http.StatusInternalServerError, e

	}
	return http.StatusOK, nil

}

func del(w http.ResponseWriter, r *http.Request) {
	name := strings.Split(r.URL.EscapedPath(), "/")[3]
	version, e := es.SearchLatestVersion(name)
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
		return

	}
	e = es.PutMetadata(name, version.Version+1, 0, "")
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

}
func Handler(w http.ResponseWriter, r *http.Request) {
	m := r.Method
	if m == http.MethodPut {
		put(w, r)
		return
	} else if m == http.MethodGet {
		get(w, r)
		return
	} else if m == http.MethodDelete {
		del(w, r)
		return
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
