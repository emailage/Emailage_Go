package emailage

// ErrorCodeLookup provies the ability to look up an error code
// returned from the API
func ErrorCodeLookup(code int) string {
	switch code {
	case 400:
		return "Invalid input parameter. Error message should indicate which one"
	case 401:
		return "Invalid or expired token. This can happen if an access token was either revoked or has expired. This can be fixed by re-authenticating the user"
	case 403:
		return "Invalid oAuth request (wrong consumer key, bad nonce, expired timestamp...)"
	case 404:
		return "File or folder was not found at the specified path"
	case 405:
		return "Request method not expected (generally should be GET or POST)"
	case 503:
		return "Calls are exceeding the defined throttle limit."
	case 3001:
		return "Authentication Error: The signature doesn't match or the user/consumer key file wasn't found"
	case 3002:
		return "Authentication Error: Your account status is inactive or disabled. Please contact us at support@Emailage.com to activate your account"
	case 3003:
		return "Authentication Error: Your account is currently expired. The free trial access period has ended. Please contact support@Emailage.com to upgrade your account"
	case 3004:
		return "Authentication Error: You do not have enough query credits available. Please contact support@Emailage.com to upgrade your account"
	case 3005:
		return "Authentication Error: A general error has occurred which prevented the proper authorization by our API in response to your request. Please contact us at apisupport@Emailage.com"
	case 3006:
		return "Invalid Parameter: Parameter not provided or empty"
	case 3007:
		return "Invalid Parameter: Malformed or wrong parameter provided"
	case 3008:
		return "Authentication Error: Invalid timestamp provided"
	case 3009:
		return "Authentication Error: Invalid nonce provided"
	case 3010:
		return "Invalid Parameter: Invalid partnerid provided"
	case 3011:
		return "Invalid Parameter: Invalid CustomerIdentifierID provided"
	case 3012:
		return "Invalid Parameter: The user_email is not valid or is inactive"
	case 3013:
		return "Invalid Parameter: The user_email and departmentid do not match"
	case 3014:
		return "Invalid Parameter: Consumer Key not found"
	case 3015:
		return "Invalid Parameter: No account found for your consumer Key"
	case 3016:
		return "Authentication Error: Invalid Signature Method"
	default:
		return ""
	}
}
