package external

import (
	"bytes"
	"context"
	"encoding/json"
	"ewallet-ums/helpers"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

type Wallet struct {
	ID      int     `json:"id"`
	UserID  int     `json:"user_id"`
	Balance float64 `json:"balance"`
}

type ExtWallet struct {

}

func (e *ExtWallet) CreateWallet(ctx context.Context, userID int) (*Wallet, error) {
	req := Wallet{UserID: userID}
	payload, err := json.Marshal(req)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to marhsal json")
	}

	url := helpers.GetEnv("WALLET_HOST", "") + helpers.GetEnv("WALLET_ENDPOINT_CREATE", "")
	httpReq, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create wallet http request")
	}

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to connect wallet service")
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("got error response from wallet service: %d", resp.StatusCode)
	}

	result := &Wallet{}
	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to read response body")
	}
	defer resp.Body.Close()

	return result, nil
}
