package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Skill struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Key  string `json:"key"`
}

type PageData struct {
	Skills []Skill
}

// Helper function to generate a slice of ints (1 to 5) for radio buttons
var funcs = template.FuncMap{
	"range5": func() []int {
		return []int{1, 2, 3, 4, 5}
	},
}

func loadSkillsFromJSON(path string) ([]Skill, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var skills []Skill
	err = json.NewDecoder(file).Decode(&skills)
	if err != nil {
		return nil, err
	}

	return skills, nil
}

func submissionHandler(w http.ResponseWriter, r *http.Request) {

	username := r.FormValue("name")

	fmt.Println(username)

	skills, err := loadSkillsFromJSON("resource/skills.json")
	if err != nil {
		http.Error(w, "Unable to load skills", http.StatusInternalServerError)
		log.Println("Error loading skills:", err)
		return
	}

	for _, skill := range skills {
		f := fmt.Sprintf(`skill_%v`, skill.ID)
		fmt.Println(skill.Name)
		fmt.Println(r.FormValue(f))

	}

	tmpl, err := template.New("submission.view.html").Funcs(funcs).ParseFiles("template/submission.view.html")
	if err != nil {
		http.Error(w, "Template parsing error", http.StatusInternalServerError)
		log.Println("Error parsing template:", err)
		return
	}

	tmpl.Execute(w, nil)

}

func formHandler(w http.ResponseWriter, r *http.Request) {
	skills, err := loadSkillsFromJSON("resource/skills.json")
	if err != nil {
		http.Error(w, "Unable to load skills", http.StatusInternalServerError)
		log.Println("Error loading skills:", err)
		return
	}

	tmpl, err := template.New("form.view.html").Funcs(funcs).ParseFiles("template/form.view.html")
	if err != nil {
		http.Error(w, "Template parsing error", http.StatusInternalServerError)
		log.Println("Error parsing template:", err)
		return
	}

	data := PageData{Skills: skills}
	tmpl.Execute(w, data)
}

func main() {
	http.HandleFunc("/", formHandler)

	http.HandleFunc("/submit", submissionHandler)

	log.Println("Server started at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server error:", err)
	}
}
