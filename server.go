package main

import (
  "os"
  "log"
  "net/http"
)

func showUsage() {
  log.Println("Usage: ./server <folder> [port]")
}

func serve(port string, folder string) {
  http.Handle("/", http.FileServer(http.Dir(folder)))
  err := http.ListenAndServe(":" + port, nil)
  if err != nil {
    log.Fatal(err)
  }
}

func getPort() string {
  port := "80"
  if len(os.Args) > 2 {
    port = os.Args[2]
  }
  return port
}

func getFolder() string {
  return os.Args[1]
}

func main() {
  
  if len(os.Args) == 1 {
    showUsage()
    os.Exit(0)
  }

  folder := getFolder()
  port := getPort()

  log.Println("Server running on port " + port + " serving " + folder)
  serve(port, folder)
}