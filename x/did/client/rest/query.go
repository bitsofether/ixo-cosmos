package rest

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/gorilla/mux"

	rest "github.com/ixofoundation/ixo-cosmos/client"
	"github.com/ixofoundation/ixo-cosmos/x/did/internal/keeper"
	"github.com/ixofoundation/ixo-cosmos/x/did/internal/types"
	"github.com/ixofoundation/ixo-cosmos/x/ixo"
)

func registerQueryRoutes(cliCtx context.CLIContext, r *mux.Router) {
	r.HandleFunc("/didToAddr/{did}", queryAddressFromDidRequestHandler(cliCtx)).Methods("GET")
	r.HandleFunc("/did/{did}", queryDidDocRequestHandler(cliCtx)).Methods("GET")
	r.HandleFunc("/did", queryAllDidsRequestHandler(cliCtx)).Methods("GET")
	r.HandleFunc("/allDidDocs", queryAllDidDocsRequestHandler(cliCtx)).Methods("GET")
}

func queryAddressFromDidRequestHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		accAddress := sdk.AccAddress(crypto.AddressHash([]byte(vars["did"])))

		rest.PostProcessResponse(w, cliCtx.Codec, accAddress, true)
	}
}

func queryDidDocRequestHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		didAddr := vars["did"]
		key := ixo.Did(didAddr)
		res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", types.QuerierRoute,
			keeper.QueryDidDoc, key), nil)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(fmt.Sprintf("Could't query did. Error: %s", err.Error())))

			return
		}
		if len(res) == 0 {
			w.WriteHeader(http.StatusNoContent)
			_, _ = w.Write([]byte("No data for respected did address."))

			return
		}

		var didDoc types.BaseDidDoc
		cliCtx.Codec.MustUnmarshalJSON(res, &didDoc)

		rest.PostProcessResponse(w, cliCtx.Codec, didDoc, true)
	}
}

func queryAllDidsRequestHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", types.QuerierRoute,
			keeper.QueryAllDids), nil)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(fmt.Sprintf("Could't query did. Error: %s", err.Error())))

			return
		}

		if len(res) == 0 {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		dids := []ixo.Did{}
		cliCtx.Codec.MustUnmarshalJSON(res, &dids)

		rest.PostProcessResponse(w, cliCtx.Codec, dids, true)
	}
}

func queryAllDidDocsRequestHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", types.QuerierRoute,
			keeper.QueryAllDidDocs), nil)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(fmt.Sprintf("Could't query did. Error: %s", err.Error())))

			return
		}

		if len(res) == 0 {
			w.WriteHeader(http.StatusNoContent)
			_, _ = w.Write([]byte("No data present."))

			return
		}

		didDocs := []types.BaseDidDoc{}
		cliCtx.Codec.MustUnmarshalJSON(res, &didDocs)

		rest.PostProcessResponse(w, cliCtx.Codec, didDocs, true)
	}
}
