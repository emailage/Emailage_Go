package emailage

import "time"

// Response
type Response struct {
	Query *Query `json:"query"`
}

// Query
type Query struct {
	Email          string          `json:"email"`
	QueryType      string          `json:"queryType"`
	Count          int64           `json:"count"`
	Created        time.Time       `json:"created"`
	Lang           string          `json:"lang"`
	ResponseCount  int64           `json:"responseCount"`
	Results        []Result        `json:"results"`
	ResponseStatus *ResponseStatus `json:"responseStatus"`
}

// ResponseStatus
type ResponseStatus struct {
	Status      string `json:"status"`
	ErrorCode   string `json:"errorCode"`
	Description string `json:"description"`
}

// SMLink
type SMLink struct {
	Source string `json:"source"`
	Link   string `json:"link"`
}

// Result
type Result struct {
	Userdefinedrecordid   string    `json:"userdefinedrecordid"`
	Email                 string    `json:"email"`
	EName                 string    `json:"eName"`
	EmailAge              string    `json:"emailAge"`
	EmailCreationDays     string    `json:"email_creation_days"`
	DomainAge             time.Time `json:"domainAge"`
	DomainCreationDays    string    `json:"domain_creation_days"`
	FirstVerificationDate time.Time `json:"firstVerificationDate"`
	FirstSeenDays         string    `json:"first_seen_days"`
	LastVerificationDate  time.Time `json:"lastVerificationDate"`
	Status                string    `json:"status"`
	Country               string    `json:"country"`
	FraudRisk             string    `json:"fraudRisk"`
	EAScore               string    `json:"EAScore"`
	EAReason              string    `json:"EAReason"`
	EAStatusID            string    `json:"EAStatusID"`
	EAReasonID            string    `json:"EAReasonID"`
	EAAdviceID            string    `json:"EAAdviceID"`
	EAAdvice              string    `json:"EAAdvice"`
	EARiskBandID          string    `json:"EARiskBandID"`
	EARiskBand            string    `json:"EARiskBand"`
	SourceIndustry        string    `json:"source_industry"`
	FraudType             string    `json:"fraud_type"`
	Lastflaggedon         string    `json:"lastflaggedon"`
	Dob                   string    `json:"dob"`
	Gender                string    `json:"gender"`
	Location              string    `json:"location"`
	Smfriends             string    `json:"smfriends"`
	Totalhits             string    `json:"totalhits"`
	Uniquehits            string    `json:"uniquehits"`
	Imageurl              string    `json:"imageurl"`
	EmailExists           string    `json:"emailExists"`
	DomainExists          string    `json:"domainExists"`
	Company               string    `json:"company"`
	Title                 string    `json:"title"`
	Domainname            string    `json:"domainname"`
	Domaincompany         string    `json:"domaincompany"`
	Domaincountryname     string    `json:"domaincountryname"`
	Domaincategory        string    `json:"domaincategory"`
	Domaincorporate       string    `json:"domaincorporate"`
	Domainrisklevel       string    `json:"domainrisklevel"`
	Domainrelevantinfo    string    `json:"domainrelevantinfo"`
	DomainrisklevelID     string    `json:"domainrisklevelID"`
	DomainrelevantinfoID  string    `json:"domainrelevantinfoID"`
	Domainriskcountry     string    `json:"domainriskcountry"`
	Smlinks               []SMLink  `json:"smlinks"`
	PhoneStatus           string    `json:"phone_status"`
	Shipforward           string    `json:"shipforward"`
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

// RiskLevelLookup
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
