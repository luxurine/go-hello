package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "go-hello",
		Short: "A simple HTTP service using Cobra",
		Run: func(cmd *cobra.Command, args []string) {
			http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "text/plain")
				fmt.Fprint(w, "pong")
			})
			fmt.Println("Serving on :8080...")
			if err := http.ListenAndServe(":8080", nil); err != nil {
				fmt.Fprintf(os.Stderr, "Server failed: %v\n", err)
				os.Exit(1)
			}
		},
	}
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
