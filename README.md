<p align="center">
  <a href="emailage"><img src="https://www.emailage.com/wp-content/uploads/2018/01/logo-dark.svg" width="200" height="200" border="0" alt="Emailage"></a>
</p>
<p align="center">
  <a href="https://godoc.org/github.com/emailage/Emailage_Go"><img src="https://godoc.org/github.com/emailage/emailage?status.svg" alt="GoDoc"></a>
  <a href="https://github.com/emailage/Emailage_Go/releases"><img src="https://img.shields.io/badge/version-0.1.0-green.svg?" alt="Version"></a>
</p>

The Emailage&#8482; API is organized around REST (Representational State Transfer). The API was built to help companies integrate with our highly efficient fraud risk and scoring system. By calling our API endpoints and simply passing us an email and/or IP Address, companies will be provided with real-time risk scoring assessments based around machine learning and proprietary algorithms that evolve with new fraud trends.

The package only supports `HMAC-SHA1` at this time however it will be extended to include `HMAC-SHA256`, `HMAC-SHA384`, and `HMAC-SHA512`.

## Getting Started

### Requirements

```sh
git clone https://github.com/emailage/Emailage_Go.git
```

This package can be imported with 

```Go
import github.com/emailage/emailage
```

## Usage

### Settings

### Email Only Validation

```Go
res, err := client.EmailOnlyScore("nigerian.prince@legit.ru")
if err != nil {
    log.Fatalln(err)
}
fmt.Printf("Result: %+v\n", res.Query)
```

### IP Only Validation

```Go
res, err := client.IPOnlyScore("192.168.0.1")
if err != nil {
    log.Fatalln(err)
}
fmt.Printf("Result: %+v\n", res.Query)
```

### Email and IP Validation

### Email and IP Validation with extra input parameters

### Mark email as fraud/good

## Frequent asked integration problems