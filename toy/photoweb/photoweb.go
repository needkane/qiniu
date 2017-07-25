package main

import (
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"runtime/debug"
)

const (
	ListDir      = 0x0001
	UPLOAD_DIR   = "./uploads"
	TEMPLATE_DIR = "./views"
)

//cache storage all template
var templates = make(map[string]*template.Template)

func init() {
	fileInfoArr, err := ioutil.ReadDir(TEMPLATE_DIR)
	check(err)
	var templateName, templatePath string
	for _, fileInfo := range fileInfoArr {
		templateName = fileInfo.Name()
		if ext := path.Ext(templateName); ext != ".html" {
			continue
		}
		templatePath = TEMPLATE_DIR + "/" + templateName
		//log.Println(templateName)
		log.Println("Loading template:", templatePath)
		t := template.Must(template.ParseFiles(templatePath))
		templates[templateName] = t
	}

	//if const template
	/*for _, tmpl := range []string{"upload", "list"} {
		//Must ensure in case can't analytic will do error operate,If the template loading is not successful, the program will exit
		t := template.Must(template.ParseFiles(tmpl + ".html"))
		templates[tmpl] = t
	}*/
}
func check(err error) {
	if err != nil {
		panic(err)
	}
}
func renderHtml(w http.ResponseWriter, tmpl string, locals map[string]interface{}) {
	/*tt := templates[tmpl]
	log.Println("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx", templates, tmpl, tt)*/
	log.Println(locals)
	err := templates[tmpl].Execute(w, locals)

	check(err)
}
func isExists(path string) bool {
	_, err := os.Stat(path)
	//log.Println(err)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}

//Callback method,
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	//log.Println(r.Method)
	if r.Method == "GET" {
		renderHtml(w, "upload.html", nil)
	}
	if r.Method == "POST" {
		f, h, err := r.FormFile("image")
		//log.Println(f, h, err)
		check(err)
		filename := h.Filename
		defer f.Close()
		//log.Println(UPLOAD_DIR, filename)
		t, err := ioutil.TempFile(UPLOAD_DIR, filename)
		//log.Println(t)
		check(err)
		defer t.Close()
		_, err = io.Copy(t, f) //dst src
		check(err)
		http.Redirect(w, r, "/view?id="+filename, http.StatusFound)

	}
}
func viewHandler(w http.ResponseWriter, r *http.Request) {
	imageId := r.FormValue("id")
	imagePath := UPLOAD_DIR + "/" + imageId
	log.Println(imageId, imagePath)
	if exists := isExists(imagePath); !exists {
		http.NotFound(w, r)
		return
	}
	//log.Println("222222222222")
	w.Header().Set("Content-Type", "image")
	http.ServeFile(w, r, imagePath)
}
func listHandler(w http.ResponseWriter, r *http.Request) {
	fileInfoArr, err := ioutil.ReadDir("./uploads")
	//log.Println(fileInfoArr)
	check(err)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	locals := make(map[string]interface{})
	images := []string{}
	for _, fileInfo := range fileInfoArr {
		//log.Println(fileInfo)
		images = append(images, fileInfo.Name())
	}
	locals["images"] = images //这里map的key是images，对应list.html里面的$.images
	renderHtml(w, "list.html", locals)

}

func safeHandler(fn http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			log.Println("can't past execute or finished execute")
			if e, ok := recover().(error); ok {
				log.Println("50x error")
				//50x error
				http.Error(w, e.Error(), http.StatusInternalServerError)
				log.Println("Warn : panic in %v. - %v", fn, e)
				log.Println(string(debug.Stack()))
			}
		}()

		log.Println("if no panic then first execute")
		fn(w, r)
	}
}
func staticDirHandler(mux *http.ServeMux, prefix string, staticDir string, flags int) {
	mux.HandleFunc(prefix, func(w http.ResponseWriter, r *http.Request) {
		file := staticDir + r.URL.Path[len(prefix)-1:]
		if (flags & ListDir) == 0 {
			if exists := isExists(file); !exists {
				http.NotFound(w, r)
				return
			}
		}
		http.ServeFile(w, r, file)
	})

}
func main() {
	mux := http.NewServeMux()
	staticDirHandler(mux, "/assets/", "./public", 0)
	mux.HandleFunc("/", safeHandler(listHandler))
	mux.HandleFunc("/view", safeHandler(viewHandler))
	mux.HandleFunc("/upload", safeHandler(uploadHandler))
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("ListenAndServe:", err.Error())
	}
}
