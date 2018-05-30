package emailage

// Response
type Response struct {
	Query *Query `json:"query"`
}

// Query
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

// ResponseStatus
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

// Result
type Result struct {
	Userdefinedrecordid   string   `json:"userdefinedrecordid,omitempty"`
	Email                 string   `json:"email,omitempty"`
	EName                 string   `json:"eName,omitempty"`
	EmailAge              string   `json:"emailAge,omitempty"`
	EmailCreationDays     string   `json:"email_creation_days,omitempty"`
	DomainAge             string   `json:"domainAge,omitempty"`
	DomainCreationDays    string   `json:"domain_creation_days,omitempty"`
	FirstVerificationDate string   `json:"firstVerificationDate,omitempty"`
	FirstSeenDays         string   `json:"first_seen_days,omitempty"`
	LastVerificationDate  string   `json:"lastVerificationDate,omitempty"`
	Status                string   `json:"status,omitempty"`
	Country               string   `json:"country,omitempty"`
	FraudRisk             string   `json:"fraudRisk,omitempty"`
	EAScore               string   `json:"EAScore,omitempty"`
	EAReason              string   `json:"EAReason,omitempty"`
	EAStatusID            string   `json:"EAStatusID,omitempty"`
	EAReasonID            string   `json:"EAReasonID,omitempty"`
	EAAdviceID            string   `json:"EAAdviceID,omitempty"`
	EAAdvice              string   `json:"EAAdvice,omitempty"`
	EARiskBandID          string   `json:"EARiskBandID,omitempty"`
	EARiskBand            string   `json:"EARiskBand,omitempty"`
	SourceIndustry        string   `json:"source_industry,omitempty"`
	FraudType             string   `json:"fraud_type,omitempty"`
	Lastflaggedon         string   `json:"lastflaggedon,omitempty"`
	Dob                   string   `json:"dob,omitempty"`
	Gender                string   `json:"gender,omitempty"`
	Location              string   `json:"location,omitempty"`
	Smfriends             string   `json:"smfriends,omitempty"`
	Totalhits             string   `json:"totalhits,omitempty"`
	Uniquehits            string   `json:"uniquehits,omitempty"`
	Imageurl              string   `json:"imageurl,omitempty"`
	EmailExists           string   `json:"emailExists,omitempty"`
	DomainExists          string   `json:"domainExists,omitempty"`
	Company               string   `json:"company,omitempty"`
	Title                 string   `json:"title,omitempty"`
	Domainname            string   `json:"domainname,omitempty"`
	Domaincompany         string   `json:"domaincompany,omitempty"`
	Domaincountryname     string   `json:"domaincountryname,omitempty"`
	Domaincategory        string   `json:"domaincategory,omitempty"`
	Domaincorporate       string   `json:"domaincorporate,omitempty"`
	Domainrisklevel       string   `json:"domainrisklevel,omitempty"`
	Domainrelevantinfo    string   `json:"domainrelevantinfo,omitempty"`
	DomainrisklevelID     string   `json:"domainrisklevelID,omitempty"`
	DomainrelevantinfoID  string   `json:"domainrelevantinfoID,omitempty"`
	Domainriskcountry     string   `json:"domainriskcountry,omitempty"`
	Smlinks               []SMLink `json:"smlinks,omitempty"`
	PhoneStatus           string   `json:"phone_status,omitempty"`
	Shipforward           string   `json:"shipforward,omitempty"`
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
