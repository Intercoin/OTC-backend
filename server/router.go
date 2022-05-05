package server

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"net/http"
	"otc-backend/types"
)

func SetupRouter(logger *zap.Logger, app SwapApp) http.Handler {
	r := &router{
		app:    app,
		logger: logger,
	}
	return r.Handler()
}

type SwapApp interface {
	GetTrade(ctx context.Context, tradehash string) (*types.Trade, error)
	GetLocks(ctx context.Context, tradehash string) ([]*types.Lock, error)
	GetEngages(ctx context.Context, tradehash string) ([]*types.Engage, error)
	GetClaims(ctx context.Context, tradehash string) ([]*types.Claim, error)
}

type router struct {
	app    SwapApp
	logger *zap.Logger
}

func (ro *router) Handler() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/trade/{tradehash}", ro.tradeByTradehashHandler).Methods("GET")
	return r
}

func (ro *router) tradeByTradehashHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tradehash := vars["tradehash"]
	if tradehash == "" {
		ro.httpJSONError(w, "empty tradehash", http.StatusBadRequest)
		return
	}
	trade, err := ro.app.GetTrade(r.Context(), tradehash)
	if err != nil {
		ro.handlerError(w, err)
		return
	}
	var res types.GetTradeByTradehashResponse

	res.Address1 = trade.Address1
	res.Address2 = trade.Address2

	res.Locks, err = ro.app.GetLocks(r.Context(), tradehash)
	if err != nil && err != sql.ErrNoRows {
		ro.handlerError(w, err)
		return
	}

	res.Engages, err = ro.app.GetEngages(r.Context(), tradehash)
	if err != nil && err != sql.ErrNoRows {
		ro.handlerError(w, err)
		return
	}

	res.Claims, err = ro.app.GetClaims(r.Context(), tradehash)
	if err != nil && err != sql.ErrNoRows {
		ro.handlerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(res); err != nil {
		ro.logger.Warn("error writing json response", zap.Error(err))
	}
}

func (ro *router) handlerError(w http.ResponseWriter, err error) {
	if errors.Cause(err) == types.ErrNoObject {
		ro.httpJSONError(w, err.Error(), http.StatusNotFound)
		return
	}

	ro.httpJSONError(w, err.Error(), http.StatusInternalServerError)
}

func (ro *router) httpJSONError(w http.ResponseWriter, error string, code int) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(map[string]string{
		"error": error,
	})
	if err != nil {
		ro.logger.Warn("error writing json response", zap.Error(err))
	}
}
