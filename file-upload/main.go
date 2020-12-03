package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// MaxUploadSize is a file limits size.
const MaxUploadSize = 1024 * 1024 //1MB

// Progress is used to track the progress of a file upload.
// It implements the io.Writer interface so it can be passed
// to an io.TeeReader()
type Progress struct {
	TotalSize	int64
	BytesRead	int64
}

// Write is used to sarisfy the io.Writer interface.
// Instead of writing somewhere, it simply aggregates
// the total bytes on each read
func (pr *Progress) Write(p []byte) (n int, err error) {
	n, err = len(p), nil
	pr.BytesRead += int64(n)
	pr.Print()
	return
}

// Print displays the current progress of the file upload
// each time Write is called
func (pr *Progress) Print() {
	if pr.BytesRead == pr.TotalSize {
		fmt.Println("DONE!")
		return
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	http.ServeFile(w, r, "index.html")
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	//-----Handle multiple files------------------------------------
	// ======================================================================
	// ここが追加::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
	// 32MB is the default used by FormFile()
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get a reference to the fileHeaders.
	// They are accessible only after ParseMultipartForm is called.
	files := r.MultipartForm.File["file"]

	for _, fileHeader := range files {
		// Restrict the size of each uploaded file to 1MB.
		// To prevent the aggregate size from exceeding
		// a specified value, use the http.MaxBytesReader() method
		// before calling ParseMultipartForm()
		if fileHeader.Size > MaxUploadSize {
			http.Error(w, fmt.Sprintf(
				"The uploaded image is too big: %s. Please use an image less than 1MB in size",
				fileHeader.Filename,
				),
				http.StatusBadRequest,
			)
			return
		}

		// Open the file
		file, err := fileHeader.Open()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		defer file.Close()

		// :::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::

		buff := make([]byte, 512)
		_, err = file.Read(buff)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	
		filetype := http.DetectContentType(buff)
		if filetype != "image/jpeg" && filetype != "image/png" {
			http.Error(
				w,
				"The provided file format is not allowed. Please upload a JPEG or PNG image",
				http.StatusBadRequest,
			)
			return
		}
	
		_, err = file.Seek(0, io.SeekStart)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	
	
		// Create the uploads folder if it doesn't
		// already exist
		err = os.Mkdir("./uploads", os.ModePerm)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	
		// Create a new file in the uploades directory
		dst, err := os.Create(fmt.Sprintf(
			"./uploads/%d%s",
			time.Now().UnixNano(),
			filepath.Ext(fileHeader.Filename),
		))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	
		defer dst.Close()

		pr := &Progress{
			TotalSize: fileHeader.Size,
		}
	
		// Copy the uploaded file to the filesystem
		// at the specified destination
		_, err = io.Copy(dst, io.TeeReader(file, pr))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}
	// ==============================================================

	// ===================================================
	// r.Body = http.MaxBytesReader(w, r.Body, MaxUploadSize)
	// if err := r.ParseMultipartForm(MaxUploadSize); err != nil {
	// 	http.Error(
	// 		w,
	// 		"The uploaded file is too big. Please choose an file that's less than 1MB in size",
	// 		http.StatusBadRequest,
	// 	)
	// 	return
	// }

	// The argument to FormFile must match the name attribute
	// of the file input on the frontend.
	// file, fileHandler, err := r.FormFile("file")
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	// defer file.Close()

	// 個々に追加
	// ========================================

	fmt.Fprintf(w, "Upload successful")
}

func setupRoutes() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/upload", uploadHandler)

	// Here we kick off our server on localhost:8080
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Println("File Upload")
	fmt.Println("Access the localhost:8080!")
	setupRoutes()
}