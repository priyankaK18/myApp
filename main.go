package main
import (
   "net/http"
   "fmt"
   "time"
   "html/template"
)

//Create a struct that holds information to be displayed in our HTML file
type Welcome struct {
   Name string
   Time string
}

//Go application entrypoint
func main() {
   welcome := Welcome{"Anonymous", time.Now().Format(time.Stamp)}
   templates := template.Must(template.ParseFiles("template/template.html")) 
   
   http.Handle("/static/", //final url can be anything
      http.StripPrefix("/static/",
         http.FileServer(http.Dir("static")))) //Go looks in the relative "static" directory first using http.FileServer(), then matches it to a
         //url of our choice as shown in http.Handle("/static/"). This url is what we need when referencing our css files
         //once the server begins. Our html code would therefore be <link rel="stylesheet"  href="/static/stylesheet/...">
         //It is important to note the url in http.Handle can be whatever we like, so long as we are consistent.

   //This method takes in the URL path "/" and a function that takes in a response writer, and a http request.
   http.HandleFunc("/" , func(w http.ResponseWriter, r *http.Request) {

      //Takes the name from the URL query e.g ?name=Martin, will set welcome.Name = Martin.
      if name := r.FormValue("name"); name != "" {
         welcome.Name = name;
      }
      //If errors show an internal server error message
      //I also pass the welcome struct to the welcome-template.html file.
      if err := templates.ExecuteTemplate(w, "template.html", welcome); err != nil {
         http.Error(w, err.Error(), http.StatusInternalServerError)
      }
   })

   //Start the web server, set the port to listen to 8080. Without a path it assumes localhost
   //Print any errors from starting the webserver using fmt
   fmt.Println("Listening");
   fmt.Println(http.ListenAndServe(":8082", nil));
}