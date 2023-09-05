package binance_p2p_api

const (
	bapi        = "/bapi"
	getExchange = "/c2c/v2/friendly/c2c/adv/search"
)

// user type
const (
	User     = "user"
	Merchant = "merchant" // pro merchant
)

const (
	HeaderContentType = "content-type"
	HeaderOrigin      = "Origin"
	HeaderPragma      = "Pragma"
	HeaderTE          = "TE"
	HeaderUserAgent   = "User-Agent"

	ApplicationJsonContentType = "application/json"
	NoCashPragma               = "no-cache"
	TrailersTE                 = "Trailers"
	MozillaUserAgent           = "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:88.0) Gecko/20100101 Firefox/88.0"
)

const (
	OperationBuy  = "BUY"
	OperationSell = "SELL"
)

type Request struct {
	Asset             string   `json:"asset"`
	Fiat              string   `json:"fiat"`
	MerchantCheck     bool     `json:"merchantCheck"`
	Page              int      `json:"page"`
	PayTypes          []string `json:"payTypes"`
	PublisherType     *string  `json:"publisherType"` // Only show Merchant Ads (if nil, it will show all data)
	ProMerchantAds    bool     `json:"proMerchantAds"`
	ShieldMerchantAds bool     `json:"shieldMerchantAds"`
	Rows              int      `json:"rows"`
	TradeType         string   `json:"tradeType"` // use OperationBuy or OperationSell
	TransAmount       float64  `json:"transAmount"`
	Countries         []string `json:"countries"`
	OrderBy           *string  `json:"order"` // use TradeCount or CompletionRate
}

type Response struct {
	Code          string  `json:"code"`
	Message       *string `json:"message"`
	MessageDetail *string `json:"messageDetail"`
	Data          []Data  `json:"data"`
	Total         int     `json:"total"`
	Success       bool    `json:"success"`
}

type Data struct {
	Adv        Adv        `json:"adv"`
	Advertiser Advertiser `json:"advertiser"`
}

type ExchangeDataReport struct {
	ExchangeData              []ExchangeData `json:"exchange_data"`
	CheapestAdvertiserPro     ExchangeData   `json:"cheapest_advertiser_pro"` // verified merchant
	CheapestAdvertiserGeneral ExchangeData   `json:"cheapest_advertiser_general"`
}

type ExchangeData struct {
	AdvertiserUserNo     string           `json:"advertiser_user_no"`
	AdvertiserName       string           `json:"advertiser_name"`
	ProMerchant          bool             `json:"pro_merchant"`
	TotalOrder           int              `json:"total_order"`
	CompletionRate       float64          `json:"completion_rate"`
	CommisionRate        float64          `json:"commision_rate"`
	Price                float64          `json:"price"`
	Stock                float64          `json:"stock"`
	PaymentMethods       []PaymentMethods `json:"payment_methods"`
	MinSingleTransAmount float64          `json:"min_single_trans_amount"`
	MaxSingleTransAmount float64          `json:"max_single_trans_amount"`
}

type Adv struct {
	AdvNo                           string         `json:"advNo"`
	Classify                        string         `json:"classify"`
	TradeType                       string         `json:"tradeType"`
	Asset                           string         `json:"asset"`
	FiatUnit                        string         `json:"fiatUnit"`
	AdvStatus                       interface{}    `json:"advStatus"`
	PriceType                       interface{}    `json:"priceType"`
	PriceFloatingRatio              interface{}    `json:"priceFloatingRatio"`
	RateFloatingRatio               interface{}    `json:"rateFloatingRatio"`
	CurrencyRate                    interface{}    `json:"currencyRate"`
	Price                           string         `json:"price"`
	InitAmount                      interface{}    `json:"initAmount"`
	SurplusAmount                   string         `json:"surplusAmount"`
	AmountAfterEditing              interface{}    `json:"amountAfterEditing"`
	MaxSingleTransAmount            string         `json:"maxSingleTransAmount"`
	MinSingleTransAmount            string         `json:"minSingleTransAmount"`
	BuyerKycLimit                   interface{}    `json:"buyerKycLimit"`
	BuyerRegDaysLimit               interface{}    `json:"buyerRegDaysLimit"`
	BuyerBtcPositionLimit           interface{}    `json:"buyerBtcPositionLimit"`
	Remarks                         interface{}    `json:"remarks"`
	AutoReplyMsg                    string         `json:"autoReplyMsg"`
	PayTimeLimit                    interface{}    `json:"payTimeLimit"`
	TradeMethods                    []TradeMethods `json:"tradeMethods"`
	UserTradeCountFilterTime        interface{}    `json:"userTradeCountFilterTime"`
	UserBuyTradeCountMin            interface{}    `json:"userBuyTradeCountMin"`
	UserBuyTradeCountMax            interface{}    `json:"userBuyTradeCountMax"`
	UserSellTradeCountMin           interface{}    `json:"userSellTradeCountMin"`
	UserSellTradeCountMax           interface{}    `json:"userSellTradeCountMax"`
	UserAllTradeCountMin            interface{}    `json:"userAllTradeCountMin"`
	UserAllTradeCountMax            interface{}    `json:"userAllTradeCountMax"`
	UserTradeCompleteRateFilterTime interface{}    `json:"userTradeCompleteRateFilterTime"`
	UserTradeCompleteCountMin       interface{}    `json:"userTradeCompleteCountMin"`
	UserTradeCompleteRateMin        interface{}    `json:"userTradeCompleteRateMin"`
	UserTradeVolumeFilterTime       interface{}    `json:"userTradeVolumeFilterTime"`
	UserTradeType                   interface{}    `json:"userTradeType"`
	UserTradeVolumeMin              interface{}    `json:"userTradeVolumeMin"`
	UserTradeVolumeMax              interface{}    `json:"userTradeVolumeMax"`
	UserTradeVolumeAsset            interface{}    `json:"userTradeVolumeAsset"`
	CreateTime                      interface{}    `json:"createTime"`
	AdvUpdateTime                   interface{}    `json:"advUpdateTime"`
	FiatVo                          interface{}    `json:"fiatVo"`
	AssetVo                         interface{}    `json:"assetVo"`
	AdvVisibleRet                   interface{}    `json:"advVisibleRet"`
	AssetLogo                       interface{}    `json:"assetLogo"`
	AssetScale                      int            `json:"assetScale"`
	FiatScale                       int            `json:"fiatScale"`
	PriceScale                      int            `json:"priceScale"`
	FiatSymbol                      string         `json:"fiatSymbol"`
	IsTradable                      bool           `json:"isTradable"`
	DynamicMaxSingleTransAmount     string         `json:"dynamicMaxSingleTransAmount"`
	MinSingleTransQuantity          string         `json:"minSingleTransQuantity"`
	MaxSingleTransQuantity          string         `json:"maxSingleTransQuantity"`
	DynamicMaxSingleTransQuantity   string         `json:"dynamicMaxSingleTransQuantity"`
	TradableQuantity                string         `json:"tradableQuantity"`
	CommissionRate                  string         `json:"commissionRate"`
	TradeMethodCommissionRates      []interface{}  `json:"tradeMethodCommissionRates"`
	LaunchCountry                   interface{}    `json:"launchCountry"`
	AbnormalStatusList              interface{}    `json:"abnormalStatusList"`
	CloseReason                     interface{}    `json:"closeReason"`
}

type TradeMethods struct {
	PayId                interface{} `json:"payId"`
	PayMethodId          string      `json:"payMethodId"`
	PayType              interface{} `json:"payType"`
	PayAccount           interface{} `json:"payAccount"`
	PayBank              interface{} `json:"payBank"`
	PaySubBank           interface{} `json:"paySubBank"`
	Identifier           string      `json:"identifier"`
	IconUrlColor         interface{} `json:"iconUrlColor"`
	TradeMethodName      string      `json:"tradeMethodName"`
	TradeMethodShortName string      `json:"tradeMethodShortName"`
	TradeMethodBgColor   string      `json:"tradeMethodBgColor"`
}

type PaymentMethods struct {
	Identifier string `json:"identifier"`
	Name       string `json:"name"`
	ShortName  string `json:"short_name"`
}

type Advertiser struct {
	UserNo           string        `json:"userNo"`
	RealName         interface{}   `json:"realName"`
	NickName         string        `json:"nickName"`
	Margin           interface{}   `json:"margin"`
	MarginUnit       interface{}   `json:"marginUnit"`
	OrderCount       interface{}   `json:"orderCount"`
	MonthOrderCount  int           `json:"monthOrderCount"`
	MonthFinishRate  float64       `json:"monthFinishRate"`
	AdvConfirmTime   interface{}   `json:"advConfirmTime"`
	Email            interface{}   `json:"email"`
	RegistrationTime interface{}   `json:"registrationTime"`
	Mobile           interface{}   `json:"mobile"`
	UserType         string        `json:"userType"`
	TagIconUrls      []interface{} `json:"tagIconUrls"`
	UserGrade        int           `json:"userGrade"`
	UserIdentity     string        `json:"userIdentity"`
	ProMerchant      interface{}   `json:"proMerchant"`
	IsBlocked        interface{}   `json:"isBlocked"`
}
