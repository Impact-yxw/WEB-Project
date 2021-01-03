package Objects

import "io"
import "log"
import "net/http"
import "os"
import "path/filepath"
import "strings"

func Handler(w http.ResponseWriter, r *http.Request) {
	m := r.Method
	if m == http.MethodPut {
		put(w, r)
		return
	}
	if m == http.MethodGet {
		get(w, r)
		return
	}
	if m == http.MethodPost {
		post(w, r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func put(w http.ResponseWriter, r *http.Request) {
	res := strings.Join(strings.Split(r.URL.EscapedPath(), "/")[2:], "/")
	dirName := filepath.Dir(res)
	if dirName != "." {
		e := os.Mkdir(os.Getenv("STORAGE_ROOT")+"/objects/"+dirName, 0755)
		if e != nil {
			log.Fatalln(e)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	f, e := os.Create(os.Getenv("STORAGE_ROOT") + "/objects/" + res)
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer f.Close()
	io.Copy(f, r.Body)
}

func get(w http.ResponseWriter, r *http.Request) {
	res := strings.Join(strings.Split(r.URL.EscapedPath(), "/")[2:], "/")
	f, e := os.Open(os.Getenv("STORAGE_ROOT") + "/objects/" + res)
	log.Println(res)
	if e != nil {
		log.Fatalln(e)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer f.Close()
	io.Copy(w, f)
}
func post(w http.ResponseWriter, r *http.Request) {
	res := strings.Join(strings.Split(r.URL.EscapedPath(), "/")[2:], "/")
	dirName := filepath.Dir(res)
	if dirName != "." {
		e := os.Mkdir(os.Getenv("STORAGE_ROOT")+"/objects/"+dirName, 0755)
		if e != nil {
			log.Fatalln(e)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	f, e := os.Create(os.Getenv("STORAGE_ROOT") + "/objects/" + res)
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer f.Close()
	io.Copy(f, r.Body)
}
