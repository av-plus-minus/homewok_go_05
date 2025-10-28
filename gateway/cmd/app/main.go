package main

import (
	"log/slog"
	"net/http"

	// Импортируем необходимые пакеты из ledger
	"github.com/av-plus-minus/homewok_go_05/ledger/internal/store"
	"github.com/av-plus-minus/homewok_go_05/ledger/internal/usecase"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("pong"))
	})

	addr := ":8081"
	slog.Info("Gateway listening", slog.Any("port", addr))
	if err := http.ListenAndServe(addr, mux); err != nil {
		slog.Error("gateway server error", slog.Any("error", err))
	}
}
