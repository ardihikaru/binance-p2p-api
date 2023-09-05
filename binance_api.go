package binance_p2p_api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

type BinanceP2PApi struct {
	p2pOriginUrl string // e.g. https://p2p.binance.me or https://p2p.binance.com
}

// NewBinanceP2PApi creates new binance P2P service
// source: https://p2p.binance.me/en/trade/sell/USDT?fiat=IDR&payment=all-payments
func NewBinanceP2PApi(p2pOriginUrl string) *BinanceP2PApi {
	return &BinanceP2PApi{
		p2pOriginUrl: p2pOriginUrl,
	}
}

// GetExchange fetches exchange information
func (b *BinanceP2PApi) GetExchange(assets string, fiat string, row int, payTypes []string,
	tradeType string, transAmount float64, countries []string, proMerchantAds, shieldMerchantAds, ignoreZeroOrder bool,
	publisherType, orderBy *string) (*ExchangeDataReport, error) {

	var exchangeData []ExchangeData
	var cheapAdvPro ExchangeData
	var cheapAdvGeneral ExchangeData
	edReport := ExchangeDataReport{
		ExchangeData:              exchangeData,
		CheapestAdvertiserPro:     cheapAdvPro,
		CheapestAdvertiserGeneral: cheapAdvGeneral,
	}

	// infinite loop until no record in the data
	page := 1
	for {
		rawExchange, err := b.GetExchangesRaw(assets, fiat, page, payTypes, row, tradeType, transAmount, countries,
			proMerchantAds, shieldMerchantAds, publisherType, orderBy)
		if err != nil {
			return nil, err
		}

		if len(rawExchange.Data) == 0 {
			break // break here
		}

		// adds to report
		for _, data := range rawExchange.Data {
			thisExData := extractExchangeData(data, b.p2pOriginUrl)
			edReport.ExchangeData = append(edReport.ExchangeData, thisExData)

			// if empty, set default
			if edReport.CheapestAdvertiserPro.AdvertiserName == "" && thisExData.ProMerchant {
				edReport.CheapestAdvertiserPro = thisExData
			}
			if edReport.CheapestAdvertiserGeneral.AdvertiserName == "" && !thisExData.ProMerchant {
				// if ignoreZeroOrder enabled, verify first
				if ignoreZeroOrder && thisExData.TotalOrder == 0 {
					// ignores
				} else {
					edReport.CheapestAdvertiserGeneral = thisExData
				}
			}

			currentPrice := thisExData.Price
			proCurPrice := edReport.CheapestAdvertiserPro.Price
			GeneralCurPrice := edReport.CheapestAdvertiserGeneral.Price
			// adds cheapest one (pro merchant)
			if currentPrice < proCurPrice && thisExData.ProMerchant {
				edReport.CheapestAdvertiserPro = thisExData
			}

			// adds cheapest one (normal merchant)
			if currentPrice < GeneralCurPrice && !thisExData.ProMerchant {
				if ignoreZeroOrder && thisExData.TotalOrder == 0 {
					// ignores
				} else {
					edReport.CheapestAdvertiserGeneral = thisExData
				}
			}

		}

		// if records < rows, break now
		if len(rawExchange.Data) < row {
			break
		}

		// go to next page
		page = page + 1
	}

	return &edReport, nil
}

// toFloat casts string value to float
func toFloat(val string) float64 {
	price, _ := strconv.ParseFloat(val, 32)

	return price
}

// extractPaymentMethods extracts payment methods
func extractPaymentMethods(pmList []TradeMethods) []PaymentMethods {
	var pms []PaymentMethods

	for _, method := range pmList {
		pm := PaymentMethods{
			Identifier: method.Identifier,
			Name:       method.TradeMethodName,
			ShortName:  method.TradeMethodShortName,
		}
		pms = append(pms, pm)
	}

	return pms
}

// toProMerchant casts to merchant status
func toProMerchant(userType string) bool {
	if userType == Merchant {
		return true
	} else {
		return false
	}
}

// extractExchangeData extracts exchange data
func extractExchangeData(data Data, p2pOriginUrl string) ExchangeData {
	profileUrl := p2pOriginUrl + getAdvProfile + data.Advertiser.UserNo

	exchangeData := ExchangeData{
		AdvertiserProfileUrl: profileUrl,
		AdvertiserUserNo:     data.Advertiser.UserNo,
		AdvertiserName:       data.Advertiser.NickName,
		ProMerchant:          toProMerchant(data.Advertiser.UserType),
		TotalOrder:           data.Advertiser.MonthOrderCount,
		CompletionRate:       data.Advertiser.MonthFinishRate * 100,
		CommisionRate:        toFloat(data.Adv.CommissionRate),
		Price:                toFloat(data.Adv.Price),
		Stock:                toFloat(data.Adv.SurplusAmount),
		PaymentMethods:       extractPaymentMethods(data.Adv.TradeMethods),
		MinSingleTransAmount: toFloat(data.Adv.MinSingleTransAmount),
		MaxSingleTransAmount: toFloat(data.Adv.MaxSingleTransAmount),
	}

	return exchangeData
}

// GetExchangesRaw extracts RAW exchange data
func (b *BinanceP2PApi) GetExchangesRaw(assets string, fiat string, page int, payTypes []string,
	rows int, tradeType string, transAmount float64, countries []string, proMerchantAds, shieldMerchantAds bool,
	publisherType, orderBy *string) (Response, error) {

	body := Request{
		Asset:             assets,            // e.g. USDT
		Fiat:              fiat,              // e.g. IDR
		Page:              page,              // e.g. 1
		PayTypes:          payTypes,          // e.g. [] (to show all available trade types
		Rows:              rows,              // e.g. 10
		TradeType:         tradeType,         // e.g. SELL or BUY
		TransAmount:       transAmount,       // e.g. 750000
		Countries:         countries,         // e.g. ["ID"]
		ProMerchantAds:    proMerchantAds,    // e.g. false
		ShieldMerchantAds: shieldMerchantAds, // e.g. false
		OrderBy:           orderBy,           // e.g. completion_rate or trade_count
	}

	if publisherType != nil {
		body.PublisherType = publisherType // e.g. merchant
	}

	bodyJson, err := json.Marshal(body)
	if err != nil {
		return Response{}, err
	}

	bodyReader := bytes.NewReader(bodyJson)
	request, err := http.NewRequest("POST", b.p2pOriginUrl+bapi+getExchange, bodyReader)
	if err != nil {
		return Response{}, err
	}
	request.Header.Set(HeaderContentType, ApplicationJsonContentType)
	request.Header.Set(HeaderOrigin, b.p2pOriginUrl)
	request.Header.Set(HeaderPragma, NoCashPragma)
	request.Header.Set(HeaderTE, TrailersTE)
	request.Header.Set(HeaderUserAgent, MozillaUserAgent)

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	responseRaw, err := client.Do(request)
	if err != nil {
		return Response{}, err
	}
	defer responseRaw.Body.Close()

	var response Response
	err = json.NewDecoder(responseRaw.Body).Decode(&response)
	if err != nil {
		return Response{}, err
	}
	return response, nil
}
