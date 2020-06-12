<p align="center">
  <a href="emailage"><img src="https://www.emailage.com/wp-content/uploads/2018/01/logo-dark.svg" width="200" height="200" border="0" alt="Emailage"></a>
</p>
<p align="center">
  <a href="https://github.com/emailage/Emailage_Go/releases"><img src="https://img.shields.io/badge/version-0.1.6-green.svg?" alt="Version"></a>
  <a href="https://goreportcard.com/report/github.com/emailage/Emailage_Go"><img src="https://goreportcard.com/badge/Emailage/Emailage_Go"></a>
</p>

The Emailage&#8482; API is organized around REST (Representational State Transfer). The API was built to help companies integrate with our highly efficient fraud risk and scoring system. By calling our API endpoints and simply passing us an email and/or IP Address, companies will be provided with real-time risk scoring assessments based around machine learning and proprietary algorithms that evolve with new fraud trends.

The package no longer supports `HMAC-SHA1`, we recommend use of `HMAC-SHA256`, `HMAC-SHA384`, or `HMAC-SHA512`.

## Getting Started

### Requirements

```sh
git clone https://github.com/emailage/Emailage_Go.git 
```

This package can be imported with:

```Go
import github.com/emailage/emailage
```

## Usage

### Settings

### Email Only Validation

```Go
params := map[string]string{
    "first_seen_days": 30,
    "phone": "8675309",
    "transorigin": "O",
}
res, err := client.EmailOnlyScore("nigerian.prince@legit.ru", params)
if err != nil {
    log.Fatalln(err)
}
fmt.Printf("Result: %+v\n", res.Query)
```

### IP Only Validation

```Go
res, err := client.IPOnlyScore("192.168.0.1", nil)
if err != nil {
    log.Fatalln(err)
}
fmt.Printf("Result: %+v\n", res.Query)
```

### Email and IP Validation

```Go
params := map[string]string{
    "billcity": "Phoenix",
    "responseCount": 10,
}
res, err := client.EmailAndIPScore("nigerian.prince@legit.ru", "192.168.0.1", params)
if err != nil {
    log.Fatalln(err)
}
fmt.Printf("Result: %+v\n", res.Query)
```

