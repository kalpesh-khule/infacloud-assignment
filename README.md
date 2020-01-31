# REST API Backend for URL shortner

This project is meant to build the REST API backend for the URL shortner website. Two APIs are built for this purpose. First for accepting URL and shortening it and second for redirecting users to the URL associated for the shortURL

## Consumption by front end client

Front end client can consume first API. This API will accept the URL, generate and store a short code for that URL and send back the json response with URL and short code.

### URL: http://localhost:8080/shorten

Request Body will consist of a single parameter "url", a valid URL string.

```json
{
  "url": "http://google.com/"
}
```

If request body does not contain "url" or if its value is invalid URL, status code 400 will be sent back. On success, API will respond with status code 201 and json response body as:

```json
{
  "url": "http://google.com/",
  "shortCode": "BpLnfgDs"
}
```

Frontend application should take the shortCode value from json response body and append it to the base URL (http://localhost:8080/ if you are running it on local). Frontend should display URL and associated short url on the page.

### URL: http://localhost:8080/{shorCode}

When user clicks on this API, if it is a valid short url, then user will be redirected to the original URL that was used to create the short code.
If it is an invalid shortcode, user will be shown "Invalid URL" message. (Can be redirected to custom page on front end).

## Production Version

This version of REST APIs make use of string to string map for storing short codes. This data will be lost when the server restarts.
In production version, Redis database can be used to retain key-value pair of short codes and URLs.
