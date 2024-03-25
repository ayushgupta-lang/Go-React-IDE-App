package handler

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"

	"github.com/ayush/ide/model"
)

func HandleRunCode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	var req model.CodeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var cmd *exec.Cmd
	var output bytes.Buffer
	switch req.Language {
	case "python":
		cmd = exec.Command("python", "-c", req.Code)
	case "go":
		// Create temporary directory
		tmpDir, err := ioutil.TempDir("", "go-code")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer os.RemoveAll(tmpDir) // Clean up the temporary directory

		// Create temporary Go file
		tmpFile, err := ioutil.TempFile(tmpDir, "example-*.go")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer tmpFile.Close()

		// Write code to temporary Go file
		if _, err := tmpFile.WriteString(req.Code); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Set the working directory to the temporary directory
		cmd = exec.Command("go", "run", tmpFile.Name())
		cmd.Dir = tmpDir
	case "javascript":
		cmd = exec.Command("node", "-e", req.Code)
	default:
		http.Error(w, "Unsupported language", http.StatusBadRequest)
		return
	}

	if cmd == nil {
		http.Error(w, "Failed to create command", http.StatusInternalServerError)
		return
	}

	cmd.Stdout = &output
	cmd.Stderr = &output
	err := cmd.Run()

	res := model.CodeResponse{
		Output: output.String(),
	}

	if err != nil {
		res.Error = err.Error()
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
