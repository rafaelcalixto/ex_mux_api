package endpoint

import (
  // Core libraries
  "fmt"
  "net/http"
  "io/ioutil"
)

// This function returns for the Browsers some informations about the API
// This is a mandatory function to some Browsers allows access to the API
func enableCors(w *http.ResponseWriter) {
    (*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
    enableCors(&w)
    fmt.Fprint(w, "This is the API for the SRC project")
    fmt.Println("File Name")
}

func ReceiveFiles(w http.ResponseWriter, r *http.Request) {
    enableCors(&w)

    // 1. parse input
    r.ParseMultipartForm(10 * 1024 * 1024)

    // 2. retive file form posted form-data
    files := r.MultipartForm.File["myfiles"]

    for _, file := range files {
        fmt.Println("File Name: ", file.Filename)
        fmt.Println("File Size: ", file.Size)
        fmt.Println("File Type: ", file.Header.Get("Content-Type"))
        fmt.Println("------------------------------------")

        // 3. Write temporary file on our server
        f, _ := file.Open()

        tempFile, err := ioutil.TempFile("uploads", "upload-*.pdf")
        if err != nil {
            fmt.Println(err)
            return
        }
        defer tempFile.Close()

        fileBytes, err2 := ioutil.ReadAll(f)
        if err != nil {
            fmt.Println(err2)
            return
        }
        tempFile.Write(fileBytes)
    }
    // 4. return wheter or not this has been sucessful
    fmt.Fprintf(w, "Successfully, Uploaded File\n")
}
