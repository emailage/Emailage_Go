package emailage

// Response is the overall payload returned from
// the Classic API
type Response struct {
	Query *Query `json:"query"`
}

// Query holds the relevant data for the given request
type Query struct {
	Email          string          `json:"email,omitempty"`
	QueryType      string          `json:"queryType,omitempty"`
	Count          int64           `json:"count,omitempty"`
	Created        string          `json:"created,omitempty"`
	Lang           string          `json:"lang,omitempty"`
	ResponseCount  int64           `json:"responseCount,omitempty"`
	Results        []Result        `json:"results,omitempty"`
	ResponseStatus *ResponseStatus `json:"responseStatus,omitempty"`
}

// ResponseStatus contains the status for the given request
type ResponseStatus struct {
	Status      string `json:"status,omitempty"`
	ErrorCode   string `json:"errorCode,omitempty"`
	Description string `json:"description,omitempty"`
}

// SMLink contains links to social media
type SMLink struct {
	Source string `json:"source,omitempty"`
	Link   string `json:"link,omitempty"`
}

// Result represents a result given back from the Classic API
type Result struct {
	Email                   string   `json:"email,omitempty"`
	EName                   string   `json:"eName,omitempty"`
	EmailAge                string   `json:"emailAge,omitempty"`
	EmailCreationDays       string   `json:"email_creation_days,omitempty"`
	DomainAge               string   `json:"domainAge,omitempty"`
	DomainCreationDays      string   `json:"domain_creation_days,omitempty"`
	FirstVerificationDate   string   `json:"firstVerificationDate,omitempty"`
	FirstSeenDays           string   `json:"first_seen_days,omitempty"`
	LastVerificationDate    string   `json:"lastVerificationDate,omitempty"`
	Status                  string   `json:"status,omitempty"`
	Company                 string   `json:"company,omitempty"`
	Count                   string   `json:"count,omitempty"`
	Country                 string   `json:"country,omitempty"`
	Created                 string   `json:"created,omitempty"`
	FraudRisk               string   `json:"fraudRisk,omitempty"`
	EAScore                 string   `json:"EAScore,omitempty"`
	EAReason                string   `json:"EAReason,omitempty"`
	EAStatusID              string   `json:"EAStatusID,omitempty"`
	EAReasonID              string   `json:"EAReasonID,omitempty"`
	EAAdviceID              string   `json:"EAAdviceID,omitempty"`
	EAAdvice                string   `json:"EAAdvice,omitempty"`
	EARiskBandID            string   `json:"EARiskBandID,omitempty"`
	EARiskBand              string   `json:"EARiskBand,omitempty"`
	SourceIndustry          string   `json:"source_industry,omitempty"`
	FraudType               string   `json:"fraud_type,omitempty"`
	LastflaggedOn           string   `json:"lastflaggedon,omitempty"`
	DOB                     string   `json:"dob,omitempty"`
	Gender                  string   `json:"gender,omitempty"`
	Location                string   `json:"location,omitempty"`
	Lang                    string   `json:"lang,omitempty"`
	SMFriends               string   `json:"smfriends,omitempty"`
	TotalHits               string   `json:"totalhits,omitempty"`
	UniqueHits              string   `json:"uniquehits,omitempty"`
	ImageURL                string   `json:"imageurl,omitempty"`
	EmailExists             string   `json:"emailExists,omitempty"`
	DomainExists            string   `json:"domainExists,omitempty"`
	Title                   string   `json:"title,omitempty"`
	DomainName              string   `json:"domainname,omitempty"`
	DomainCompany           string   `json:"domaincompany,omitempty"`
	DomainCountryName       string   `json:"domaincountryname,omitempty"`
	DomainCategory          string   `json:"domaincategory,omitempty"`
	DomainCorporate         string   `json:"domaincorporate,omitempty"`
	DomainRiskLevel         string   `json:"domainrisklevel,omitempty"`
	DomainRelevantInfo      string   `json:"domainrelevantinfo,omitempty"`
	DomainRisklevelID       string   `json:"domainrisklevelID,omitempty"`
	DomainRelevantInfoID    string   `json:"domainrelevantinfoID,omitempty"`
	DomainRiskCountry       string   `json:"domainriskcountry,omitempty"`
	SMLinks                 []SMLink `json:"smlinks,omitempty"`
	PhoneStatus             string   `json:"phone_status,omitempty"`
	ShipForward             string   `json:"shipforward,omitempty"`
	UserdefinedRecordID     string   `json:"userdefinedrecordid,omitempty"`
	PhoneOwner              string   `json:"phoneowner,omitempty"`
	PhoneOwnerType          string   `json:"phoneownertype,omitempty"`
	PhoneOwnerCarrierType   string   `json:"phoneownercarriertype,omitempty"`
	PhoneCarrierNetworkCode string   `json:"phonecarriernetworkcode,omitempty"`
	PhoneCarrierName        string   `json:"phonecarriername,omitempty"`
	PhoneOwnerMatch         string   `json:"phoneownermatch,omitempty"`
	IssuerBank              string   `json:"issuerBank,omitempty"`
	IssuerBrand             string   `json:"issuerBrand,omitempty"`
	IssuerCountry           string   `json:"issuerCountry,omitempty"`
	CardCategory            string   `json:"cardCategory,omitempty"`
	CardType                string   `json:"cardType,omitempty"`
	NameMatch               string   `json:"namematch,omitempty"`
	CustomerIdentifierMatch string   `json:"customeridentifiermatch,omitempty"`
	IPRiskLevelID           string   `json:"ip_risklevelid,omitempty"`
	IPRiskLevel             string   `json:"ip_risklevel,omitempty"`
	IPRiskLevelReasonID     string   `json:"ip_risklevelReasonid,omitempty"`
	IPRiskLevelReason       string   `json:"ip_risklevelreason,omitempty"`
	IPReputation            string   `json:"ip_reputation,omitempty"`
	IPAnonymousDetected     string   `json:"ip_anonymousdetected,omitempty"`
	IPProxyType             string   `json:"ip_proxytype,omitempty"`
	IPProxyDescription      string   `json:"ip_proxydescription,omitempty"`
	IPISP                   string   `json:"ip_isp,omitempty"`
	IPOrg                   string   `json:"ip_org,omitempty"`
	IPUserType              string   `json:"ip_usertype,omitempty"`
	IPNetSpeedCell          string   `json:"ip_netSpeedCell,omitempty"`
	IPCorporateProxy        string   `json:"ip_corporateProxy,omitempty"`
	IPContinentCode         string   `json:"ip_continentCode,omitempty"`
	IPCountry               string   `json:"ip_country,omitempty"`
	IPCountryCode           string   `json:"ip_countryCode,omitempty"`
	IPRegion                string   `json:"ip_region,omitempty"`
	IPCity                  string   `json:"ip_city,omitempty"`
	IPCallingCode           string   `json:"ip_callingcode,omitempty"`
	IPMetroCode             string   `json:"ip_metroCode,omitempty"`
	IPLatitude              string   `json:"ip_latitude,omitempty"`
	IPLongitude             string   `json:"ip_longitude,omitempty"`
	IPMap                   string   `json:"ip_map,omitempty"`
	IPCountryMatch          string   `json:"ipcountrymatch,omitempty"`
	IPRiskCountry           string   `json:"ipriskcountry,omitempty"`
	IPDistanceKM            string   `json:"ipdistancekm,omitempty"`
	IPDistanceMil           string   `json:"ipdistancemil,omitempty"`
	IPAccuracyRadius        string   `json:"ipaccuracyradius,omitempty"`
	IPTimezone              string   `json:"iptimezone,omitempty"`
	IPasNum                 string   `json:"ipasnum,omitempty"`
	IPDomain                string   `json:"ipdomain,omitempty"`
	IPCountryConf           string   `json:"ip_countryconf,omitempty"`
	IPRegionConf            string   `json:"ip_regionconf,omitempty"`
	IPCityConf              string   `json:"ip_cityconf,omitempty"`
	IPPostalCode            string   `json:"ip_postalcode,omitempty"`
	IPRiskScore             string   `json:"ip_riskscore,omitempty"`
	CustPhoneInBillingLoc   string   `json:"custphoneInbillingloc,omitempty"`
	CityPostalMatch         string   `json:"citypostalmatch,omitempty"`
	ShipCityPostalMatch     string   `json:"shipcitypostalmatch,omitempty"`
}

// FraudCodeLookup provides code to fraud message lookup
func FraudCodeLookup(code int) string {
	switch code {
	case 1:
		return "Card Not Present Fraud"
	case 2:
		return "Customer Dispute (Chargeback)"
	case 3:
		return "First Party Fraud"
	case 4:
		return "First Payment Default"
	case 5:
		return "Identify Theft (Fraud Application)"
	case 6:
		return "Identify Theft (Account Take Over)"
	case 7:
		return "Suspected Fraud (Not Confirmed)"
	case 8:
		return "Synthetic ID"
	case 9:
		return "Other"
	default:
		return ""
	}
}

// PhoneOwnerMatchLookup provides a lookup translation
// from a received response
func PhoneOwnerMatchLookup(s string) string {
	switch s {
	case "Y":
		return "Full Match"
	case "P":
		return "Partial Match"
	case "N":
		return "No Match"
	case "U":
		return "Owner Unknown"
	default:
		return ""
	}
}

// RiskLevelLookup takes a level integer and
// converts it to a string representation
func RiskLevelLookup(level int) string {
	switch level {
	case 1:
		return "Very High"
	case 2:
		return "High"
	case 3:
		return "Moderate"
	case 4:
		return "Low"
	case 5:
		return "Very Low"
	case 6:
		return "Review"
	default:
		return ""
	}
}

// RiskReasonLookup provides the code to reason mapping for IP risk
func RiskReasonLookup(reason int) string {
	switch reason {
	case 301:
		return "Moderate Risk"
	case 307:
		return "Risk Country"
	case 308:
		return "Anonymous Proxy"
	case 309:
		return "Risk Proxy"
	case 310:
		return "IP Not Found"
	case 311:
		return "Moderate By Proxy Reputation And Country Code"
	case 312:
		return "Invalid IP Syntax"
	case 313:
		return "TOR Network IP"
	case 321:
		return "Low Risk IP for Company"
	case 322:
		return "Low Risk IP Geolocation for Company"
	case 323:
		return "Low Risk IP for Industry"
	case 324:
		return "Low Risk IP Geolocation for Industry"
	case 325:
		return "Low Risk IP for Network"
	case 326:
		return "Low Risk IP Geolocation for Network"
	case 327:
		return "Very Low Risk IP for Company"
	case 328:
		return "Very Low Risk IP Geolocation for Company"
	case 329:
		return "Very Low Risk IP for Industry"
	case 330:
		return "Very Low Risk IP Geolocation for Industry"
	case 331:
		return "Very Low Risk IP for Network"
	case 332:
		return "Very Low Risk IP Geolocation for Network"
	case 333:
		return "High Risk IP for Company"
	case 334:
		return "High Risk IP Geolocation for Company"
	case 335:
		return "High Risk IP for Industry"
	case 336:
		return "High Risk IP Geolocation for Industry"
	case 337:
		return "High Risk IP for Network"
	case 338:
		return "High Risk IP Geolocation for Network"
	case 339:
		return "Very High Risk IP for Company"
	case 340:
		return "Very High Risk IP Geolocation for Company"
	case 341:
		return "Very High Risk IP for Industry"
	case 342:
		return "Very High Risk IP Geolocation for Industry"
	case 343:
		return "Very High Risk IP for Network"
	case 344:
		return "Very High Risk IP Geolocation for Network"
	default:
		return ""
	}
}
