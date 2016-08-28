package api

import (
	"errors"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/skycoin/skycoin-exchange/src/coin"
	"github.com/skycoin/skycoin-exchange/src/pp"
	"github.com/skycoin/skycoin-exchange/src/wallet"
)

// CreateWallet api for creating local wallet.
// mode: POST
// url: /api/v1/wallet?type=[:type]&seed=[:seed]
// params:
// 		type: bitcoin or skycoin
// 		seed: wallet seed.
func CreateWallet(se Servicer) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		rlt := &pp.EmptyRes{}
		for {
			// get coin type
			cp, err := coin.TypeFromStr(r.FormValue("type"))
			if err != nil {
				rlt = pp.MakeErrRes(err)
				break
			}

			// get seed
			sd := r.FormValue("seed")
			if sd == "" {
				rlt = pp.MakeErrRes(errors.New("seed is required"))
				break
			}

			wlt, err := wallet.New(cp, sd)
			if err != nil {
				logger.Error(err.Error())
				rlt = pp.MakeErrResWithCode(pp.ErrCode_ServerError)
				break
			}
			res := struct {
				Result *pp.Result `json:"result"`
				ID     string     `json:"id"`
			}{
				Result: pp.MakeResultWithCode(pp.ErrCode_Success),
				ID:     wlt.GetID(),
			}
			sendJSON(w, &res)
			return
		}
		sendJSON(w, rlt)
	}
}

// NewAddress create address in wallet.
// mode: POST
// url: /api/v1/wallet/:id/address?&id=[:id]
// params:
// 		id: wallet id.
func NewAddress(se Servicer) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var rlt *pp.EmptyRes
		for {
			// get wallet id
			wltID := r.FormValue("id")
			if wltID == "" {
				rlt = pp.MakeErrRes(errors.New("id is required"))
				break
			}

			addrEntries, err := wallet.NewAddresses(wltID, 1)
			if err != nil {
				rlt = pp.MakeErrResWithCode(pp.ErrCode_ServerError)
				break
			}
			res := struct {
				Result  *pp.Result
				Address string `json:"address"`
			}{
				Result:  pp.MakeResultWithCode(pp.ErrCode_Success),
				Address: addrEntries[0].Address,
			}
			sendJSON(w, &res)
			return
		}
		sendJSON(w, rlt)
	}
}

// GetAddresses get all addresses in wallet.
// mode: GET
// url: /api/v1/wallet/addresses?id=[:id]
// params:
// 		id: wallet id.
func GetAddresses(se Servicer) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var rlt *pp.EmptyRes
		for {
			id := r.FormValue("id")
			if id == "" {
				rlt = pp.MakeErrRes(errors.New("id is required"))
				break
			}

			addrs, err := wallet.GetAddresses(id)
			if err != nil {
				logger.Error(err.Error())
				rlt = pp.MakeErrResWithCode(pp.ErrCode_WrongRequest)
				break
			}

			res := struct {
				Result    *pp.Result `json:"result"`
				Addresses []string   `json:"addresses"`
			}{
				Result:    pp.MakeResultWithCode(pp.ErrCode_Success),
				Addresses: addrs,
			}
			sendJSON(w, &res)
			return
		}
		sendJSON(w, rlt)
	}
}

// GetKeys get keys of specific address in wallet.
// mode: GET
// url: /api/v1/wallet/address/keys?id=[:id]&address=[:address]
func GetKeys(se Servicer) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var rlt *pp.EmptyRes
		for {
			// get wallet id
			wltID := r.FormValue("id")
			if wltID == "" {
				rlt = pp.MakeErrRes(errors.New("no id"))
				break
			}

			// get address
			addr := r.FormValue("address")
			if addr == "" {
				rlt = pp.MakeErrRes(errors.New("no address"))
				break
			}
			p, s, err := wallet.GetKeypair(wltID, addr)
			if err != nil {
				logger.Error(err.Error())
				rlt = pp.MakeErrResWithCode(pp.ErrCode_ServerError)
				break
			}

			res := struct {
				Result *pp.Result
				Pubkey string `json:"pubkey"`
				Seckey string `json:"seckey"`
			}{
				Result: pp.MakeResultWithCode(pp.ErrCode_Success),
				Pubkey: p,
				Seckey: s,
			}
			sendJSON(w, &res)
			return
		}
		sendJSON(w, rlt)
	}
}