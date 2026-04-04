package main

import (
	"html/template"
	"net/http"
	"log"
	"color/ascii"
	"strings"
)


var funcMap = template.FuncMap{
    "safeHTML": func(s string) template.HTML {
        return template.HTML(s) // converts string to HTML safely
    },
}

var tmpl = template.Must(template.New("").Funcs(funcMap).ParseFiles("templates/index.html", "templates/error.html"))

type PageData struct{
	Output string;
	Message string;
	Text string;
	Substring string;
	Color string
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{Output: ""}
	if r.Method != http.MethodGet {
		http.Error(w, "Only a GET request is allowed here", http.StatusMethodNotAllowed)
		return
	}
	

	if r.URL.Path != "/"{
		w.WriteHeader(http.StatusNotFound)
		tmpl.ExecuteTemplate(w, "error.html", nil)
		return
	}
	tmpl.ExecuteTemplate(w, "index.html", data)
}

func asciiHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost {
		http.Error(w, "Only a POST request is allowed here", http.StatusMethodNotAllowed)
		return
	}


	// allowedBanners := map[string]bool{
	// 	"standard": true,
	// 	"shadow": true,
	// 	"thinkertoy": true,
	// }
	// if !allowedBanners[banner]{
	// 	banner = r.FormValue("standard")
	// }
	
	text := r.FormValue("text")
	banner := r.FormValue("banner")
	color := r.FormValue("color")
	substring := r.FormValue("substring")

	// log.Println("Text:", text)
	// log.Println("Substring:", substring)
	// log.Println("Banner:", banner)
	// log.Println("Color:", color)

	if strings.TrimSpace(text) == ""{
		data := PageData{Output: "You have to provide a text"}
		tmpl.ExecuteTemplate(w, "index.html", data)
		return
	}

	text = strings.ReplaceAll(text, "\\n", "\n")
	text1 := strings.Split(text, "\n")
	var asciiArt strings.Builder
	for _, line := range text1{
		if strings.TrimSpace(line) == ""{
			asciiArt.WriteString("\n")
			continue
		}
		result, err := ascii.GenerateColor(line,banner, substring, color)
		if err != nil {
			data := PageData{Output: "", Message: err.Error()}
    		tmpl.ExecuteTemplate(w, "index.html", data)
			return
		} 
		asciiArt.WriteString(result)
	}

	// log.Println("ASCII output:\n", asciiArt.String())	
	data := PageData{Output: asciiArt.String(), Message: "",  Text: text, Substring: substring, Color: color}
	tmpl.ExecuteTemplate(w, "index.html", data)
}

func downloadHandler( w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost{
		http.Error(w, "Only a POST request is allowed here", http.StatusMethodNotAllowed)
		return
	}

	ascii := r.FormValue("ascii")

	if strings.TrimSpace(ascii) == "" {
		data := PageData{Message:"No Art to Download"}
		tmpl.ExecuteTemplate(w, "index.html", data)
		return
	}
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Disposition", "attachment; filename=ascii-art.txt")
	w.Write([]byte(ascii))
}

func main() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii", asciiHandler)
	http.HandleFunc("/download", downloadHandler)

	log.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}