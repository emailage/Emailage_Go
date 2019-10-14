package emailage

// Response is the overall payload returned from
// the Classic API
type Response struct {
	Query          *Query          `json:"query"`
	ResponseStatus *ResponseStatus `json:"responseStatus,omitempty"`
}

// Query holds the relevant data for the given request
type Query struct {
	Email         string   `json:"email,omitempty"`
	QueryType     string   `json:"queryType,omitempty"`
	Count         int64    `json:"count,omitempty"`
	Created       string   `json:"created,omitempty"`
	Lang          string   `json:"lang,omitempty"`
	ResponseCount int64    `json:"responseCount,omitempty"`
	Results       []Result `json:"results,omitempty"`
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
	CardCategory                       string   `json:"cardCategory"`
	CardType                           string   `json:"cardType"`
	CityPostalMatch                    string   `json:"citypostalmatch"`
	Company                            string   `json:"company"`
	Count                              string   `json:"count"`
	Country                            string   `json:"country"`
	Created                            string   `json:"created"`
	CustomerIdentifierMatch            string   `json:"customeridentifiermatch"`
	CustPhoneInBillingLoc              string   `json:"custphoneInbillingloc"`
	DOB                                string   `json:"dob"`
	DomainAge                          string   `json:"domainAge"`
	DomainCategory                     string   `json:"domaincategory"`
	DomainCompany                      string   `json:"domaincompany"`
	DomainCorporate                    string   `json:"domaincorporate"`
	DomainCountryName                  string   `json:"domaincountryname"`
	DomainCreationDays                 string   `json:"domain_creation_days"`
	DomainExists                       string   `json:"domainExists"`
	DomainName                         string   `json:"domainname"`
	DomainRelevantInfo                 string   `json:"domainrelevantinfo"`
	DomainRelevantInfoID               string   `json:"domainrelevantinfoID"`
	DomainRiskCountry                  string   `json:"domainriskcountry"`
	DomainRiskLevel                    string   `json:"domainrisklevel"`
	DomainRisklevelID                  string   `json:"domainrisklevelID"`
	EAAdvice                           string   `json:"EAAdvice"`
	EAAdviceID                         string   `json:"EAAdviceID"`
	EAReason                           string   `json:"EAReason"`
	EAReasonID                         string   `json:"EAReasonID"`
	EARiskBand                         string   `json:"EARiskBand"`
	EARiskBandID                       string   `json:"EARiskBandID"`
	EAScore                            string   `json:"EAScore"`
	EAStatusID                         string   `json:"EAStatusID"`
	Email                              string   `json:"email"`
	EmailAge                           string   `json:"emailAge"`
	EmailCreationDays                  string   `json:"email_creation_days"`
	EmailExists                        string   `json:"emailExists"`
	EName                              string   `json:"eName"`
	FirstSeenDays                      string   `json:"first_seen_days"`
	FirstVerificationDate              string   `json:"firstVerificationDate"`
	FraudRisk                          string   `json:"fraudRisk"`
	FraudType                          string   `json:"fraud_type"`
	Gender                             string   `json:"gender"`
	ImageURL                           string   `json:"imageurl"`
	IPAccuracyRadius                   string   `json:"ipaccuracyradius"`
	IPAnonymousDetected                string   `json:"ip_anonymousdetected"`
	IPAsNum                            string   `json:"ipasnum"`
	IPCallingCode                      string   `json:"ip_callingcode"`
	IPCity                             string   `json:"ip_city"`
	IPCityConf                         string   `json:"ip_cityconf"`
	IPContinentCode                    string   `json:"ip_continentCode"`
	IPCorporateProxy                   string   `json:"ip_corporateProxy"`
	IPCountry                          string   `json:"ip_country"`
	IPCountryCode                      string   `json:"ip_countryCode"`
	IPCountryConf                      string   `json:"ip_countryconf"`
	IPCountryMatch                     string   `json:"ipcountrymatch"`
	IPDistanceKM                       string   `json:"ipdistancekm"`
	IPDistanceMil                      string   `json:"ipdistancemil"`
	IPDomain                           string   `json:"ipdomain"`
	IPISP                              string   `json:"ip_isp"`
	IPLatitude                         string   `json:"ip_latitude"`
	IPLongitude                        string   `json:"ip_longitude"`
	IPMap                              string   `json:"ip_map"`
	IPMetroCode                        string   `json:"ip_metroCode"`
	IPNetSpeedCell                     string   `json:"ip_netSpeedCell"`
	IPOrg                              string   `json:"ip_org"`
	IPPostalCode                       string   `json:"ip_postalcode"`
	IPProxyDescription                 string   `json:"ip_proxydescription"`
	IPProxyType                        string   `json:"ip_proxytype"`
	IPRegion                           string   `json:"ip_region"`
	IPRegionConf                       string   `json:"ip_regionconf"`
	IPReputation                       string   `json:"ip_reputation"`
	IPRiskCountry                      string   `json:"ipriskcountry"`
	IPRiskLevel                        string   `json:"ip_risklevel"`
	IPRiskLevelID                      string   `json:"ip_risklevelid"`
	IPRiskLevelReason                  string   `json:"ip_risklevelreason"`
	IPRiskLevelReasonID                string   `json:"ip_risklevelReasonid"`
	IPRiskScore                        string   `json:"ip_riskscore"`
	IPTimezone                         string   `json:"iptimezone"`
	IPUserType                         string   `json:"ip_usertype"`
	IssuerBank                         string   `json:"issuerBank"`
	IssuerBrand                        string   `json:"issuerBrand"`
	IssuerCountry                      string   `json:"issuerCountry"`
	Lang                               string   `json:"lang"`
	LastFlaggedOn                      string   `json:"lastflaggedon"`
	LastVerificationDate               string   `json:"lastVerificationDate"`
	Location                           string   `json:"location"`
	NameMatch                          string   `json:"namematch"`
	PhoneCarrierName                   string   `json:"phonecarriername"`
	PhoneCarrierNetworkCode            string   `json:"phonecarriernetworkcode"`
	PhoneOwner                         string   `json:"phoneowner"`
	PhoneOwnerCarrierType              string   `json:"phoneownercarriertype"`
	PhoneOwnerMatch                    string   `json:"phoneownermatch"`
	PhoneOwnerType                     string   `json:"phoneownertype"`
	PhoneStatus                        string   `json:"phone_status"`
	ShipCityPostalMatch                string   `json:"shipcitypostalmatch"`
	ShipForward                        string   `json:"shipforward"`
	SMFriends                          string   `json:"smfriends"`
	SMLinks                            []SMLink `json:"smlinks"`
	SourceIndustry                     string   `json:"source_industry"`
	Status                             string   `json:"status"`
	Title                              string   `json:"title"`
	TotalHits                          string   `json:"totalhits"`
	UniqueHits                         string   `json:"uniquehits"`
	UserdefinedRecordID                string   `json:"userdefinedrecordid"`
	OverallDigitalIdentityScore        string   `json:"overallDigitalIdentityScore"`
	EmailToIPConfidence                string   `json:"emailToIpConfidence"`
	EmailToPhoneConfidence             string   `json:"emailToPhoneConfidence"`
	EmailToBillAddressConfidence       string   `json:"emailToBillAddressConfidence"`
	EmailToShipAddressConfidence       string   `json:"emailToShipAddressConfidence"`
	EmailToFullNameConfidence          string   `json:"emailToFullNameConfidence"`
	EmailToLastNameConfidence          string   `json:"emailToLastNameConfidence"`
	IPToPhoneConfidence                string   `json:"ipToPhoneConfidence"`
	IPToBillAddressConfidence          string   `json:"ipToBillAddressConfidence"`
	IPToShipAddressConfidence          string   `json:"ipToShipAddressConfidence"`
	IPToFullNameConfidence             string   `json:"ipToFullNameConfidence"`
	IPToLastNameConfidence             string   `json:"ipToLastNameConfidence"`
	PhoneToBillAddressConfidence       string   `json:"phoneToBillAddressConfidence"`
	PhoneToShipAddressConfidence       string   `json:"phoneToShipAddressConfidence"`
	PhoneToFullNameConfidence          string   `json:"phoneToFullNameConfidence"`
	PhoneToLastNameConfidence          string   `json:"phoneToLastNameConfidence"`
	BillAddressToFullNameConfidence    string   `json:"billAddressToFullNameConfidence"`
	BillAddressToLastNameConfidence    string   `json:"billAddressToLastNameConfidence"`
	ShipAddressToBillAddressConfidence string   `json:"shipAddressToBillAddressConfidence"`
	ShipAddressToFullNameConfidence    string   `json:"shipAddressToFullNameConfidence"`
	ShipAddressToLastNameConfidence    string   `json:"shipAddressToLastNameConfidence"`
	StandardizedBillAddress            string   `json:"standardizedbillingaddress"`
	StandardizedShipAddress            string   `json:"standardizedshippingaddress"`
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
