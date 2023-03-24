package service

import (
	"backend/data"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func GetNetworkInfo() (int, data.NetworkInfo) {
	var returnNetworkInfo data.NetworkInfo
	returnNetworkInfo.Status = "fail"

	resp, err := http.Get("https://xmrchain.net/api/networkinfo")
	if err != nil {
		return resp.StatusCode, returnNetworkInfo
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, returnNetworkInfo
	}

	var networkInfo data.NetworkInfo
	if err := json.Unmarshal(body, &networkInfo); err != nil {
		return resp.StatusCode, returnNetworkInfo
	}

	if resp.StatusCode == 200 {
		return resp.StatusCode, networkInfo
	} else {
		return resp.StatusCode, returnNetworkInfo
	}
}

func GetPrice(r *http.Request) (int, data.Price) {
	var returnPrice data.Price

	resp, err := http.Get("https://api.coingecko.com/api/v3/simple/price?ids=monero&vs_currencies=usd")
	if err != nil {
		return resp.StatusCode, returnPrice
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, returnPrice
	}

	var price data.Price
	if err := json.Unmarshal(body, &price); err != nil {
		return resp.StatusCode, returnPrice
	} else {
		return resp.StatusCode, price
	}
}

func GetBlockByNumber(number int) (int, data.BlockInfo) {
	var returnBlock data.BlockInfo
	returnBlock.Status = "fail"

	url := fmt.Sprintf("https://xmrchain.net/api/block/%s", strconv.Itoa(number))
	resp, err := http.Get(url)
	if err != nil {
		return resp.StatusCode, returnBlock
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, returnBlock
	}

	var block data.BlockInfo
	if err := json.Unmarshal(body, &block); err != nil {
		return resp.StatusCode, returnBlock
	} else if block.Status == "fail" {
		return resp.StatusCode, returnBlock
	} else {
		return resp.StatusCode, block
	}

}