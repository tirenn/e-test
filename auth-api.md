---
title: Auth v1.0
language_tabs:
  - shell: Shell
  - http: HTTP
  - javascript: JavaScript
  - ruby: Ruby
  - python: Python
  - php: PHP
  - java: Java
  - go: Go
toc_footers: []
includes: []
search: true
highlight_theme: darkula
headingLevel: 2
---

<!-- Generator: Widdershins v4.0.1 -->

<h1 id="auth">Auth v1.0</h1>

> Scroll down for code samples, example requests and responses. Select a language for code samples from the tabs above or the mobile navigation menu.

Base URLs:

- <a href="http://localhost:8080/api">http://localhost:8080/api</a>

<h1 id="auth-default">Default</h1>

## post-auth-signup

<a id="opIdpost-auth-signup"></a>

> Code samples

```shell
# You can also use wget
curl -X POST http://localhost:8080/api/auth/signup \
  -H 'Content-Type: application/json' \
  -H 'Accept: application/json'

```

```http
POST http://localhost:8080/api/auth/signup HTTP/1.1
Host: localhost:8080
Content-Type: application/json
Accept: application/json

```

```javascript
const inputBody = '{
  "phone": "081234567890",
  "name": "tio",
  "role": "staff"
}';
const headers = {
  'Content-Type':'application/json',
  'Accept':'application/json'
};

fetch('http://localhost:8080/api/auth/signup',
{
  method: 'POST',
  body: inputBody,
  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Content-Type' => 'application/json',
  'Accept' => 'application/json'
}

result = RestClient.post 'http://localhost:8080/api/auth/signup',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Content-Type': 'application/json',
  'Accept': 'application/json'
}

r = requests.post('http://localhost:8080/api/auth/signup', headers = headers)

print(r.json())

```

```php
<?php

require 'vendor/autoload.php';

$headers = array(
    'Content-Type' => 'application/json',
    'Accept' => 'application/json',
);

$client = new \GuzzleHttp\Client();

// Define array of request body.
$request_body = array();

try {
    $response = $client->request('POST','http://localhost:8080/api/auth/signup', array(
        'headers' => $headers,
        'json' => $request_body,
       )
    );
    print_r($response->getBody()->getContents());
 }
 catch (\GuzzleHttp\Exception\BadResponseException $e) {
    // handle exception or api errors.
    print_r($e->getMessage());
 }

 // ...

```

```java
URL obj = new URL("http://localhost:8080/api/auth/signup");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("POST");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Content-Type": []string{"application/json"},
        "Accept": []string{"application/json"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("POST", "http://localhost:8080/api/auth/signup", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`POST /auth/signup`

Signup Endpoint

> Body parameter

```json
{
  "phone": "081234567890",
  "name": "tio",
  "role": "staff"
}
```

<h3 id="post-auth-signup-parameters">Parameters</h3>

| Name    | In   | Type   | Required | Description |
| ------- | ---- | ------ | -------- | ----------- |
| body    | body | object | false    | none        |
| » phone | body | string | true     | none        |
| » name  | body | string | true     | none        |
| » role  | body | string | true     | none        |

> Example responses

> Example response

```json
{
  "message": "Success",
  "data": {
    "phone": "081234567890",
    "name": "tio",
    "role": "staff",
    "password": "BpLn",
    "timestamp": "2022-02-23T06:23:18.429462916Z"
  }
}
```

```json
{
  "message": "string",
  "data": null
}
```

```json
{
  "message": "string",
  "data": null
}
```

<h3 id="post-auth-signup-responses">Responses</h3>

| Status | Meaning                                                                    | Description      | Schema |
| ------ | -------------------------------------------------------------------------- | ---------------- | ------ |
| 201    | [Created](https://tools.ietf.org/html/rfc7231#section-6.3.2)               | Example response | Inline |
| 400    | [Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)           | Example response | Inline |
| 500    | [Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1) | Example response | Inline |

<h3 id="post-auth-signup-responseschema">Response Schema</h3>

Status Code **201**

| Name         | Type   | Required | Restrictions | Description |
| ------------ | ------ | -------- | ------------ | ----------- |
| » message    | string | true     | none         | none        |
| » data       | object | true     | none         | none        |
| »» phone     | string | true     | none         | none        |
| »» name      | string | true     | none         | none        |
| »» role      | string | true     | none         | none        |
| »» password  | string | true     | none         | none        |
| »» timestamp | string | true     | none         | none        |

Status Code **400**

| Name      | Type   | Required | Restrictions | Description |
| --------- | ------ | -------- | ------------ | ----------- |
| » message | string | true     | none         | none        |
| » data    | any    | false    | none         | none        |

Status Code **500**

| Name      | Type   | Required | Restrictions | Description |
| --------- | ------ | -------- | ------------ | ----------- |
| » message | string | true     | none         | none        |
| » data    | any    | false    | none         | none        |

<aside class="success">
This operation does not require authentication
</aside>

## post-auth-login

<a id="opIdpost-auth-login"></a>

> Code samples

```shell
# You can also use wget
curl -X POST http://localhost:8080/api/auth/login \
  -H 'Content-Type: application/json' \
  -H 'Accept: application/json'

```

```http
POST http://localhost:8080/api/auth/login HTTP/1.1
Host: localhost:8080
Content-Type: application/json
Accept: application/json

```

```javascript
const inputBody = '{
  "phone": "string",
  "password": "string"
}';
const headers = {
  'Content-Type':'application/json',
  'Accept':'application/json'
};

fetch('http://localhost:8080/api/auth/login',
{
  method: 'POST',
  body: inputBody,
  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Content-Type' => 'application/json',
  'Accept' => 'application/json'
}

result = RestClient.post 'http://localhost:8080/api/auth/login',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Content-Type': 'application/json',
  'Accept': 'application/json'
}

r = requests.post('http://localhost:8080/api/auth/login', headers = headers)

print(r.json())

```

```php
<?php

require 'vendor/autoload.php';

$headers = array(
    'Content-Type' => 'application/json',
    'Accept' => 'application/json',
);

$client = new \GuzzleHttp\Client();

// Define array of request body.
$request_body = array();

try {
    $response = $client->request('POST','http://localhost:8080/api/auth/login', array(
        'headers' => $headers,
        'json' => $request_body,
       )
    );
    print_r($response->getBody()->getContents());
 }
 catch (\GuzzleHttp\Exception\BadResponseException $e) {
    // handle exception or api errors.
    print_r($e->getMessage());
 }

 // ...

```

```java
URL obj = new URL("http://localhost:8080/api/auth/login");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("POST");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Content-Type": []string{"application/json"},
        "Accept": []string{"application/json"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("POST", "http://localhost:8080/api/auth/login", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`POST /auth/login`

Login

> Body parameter

```json
{
  "phone": "string",
  "password": "string"
}
```

<h3 id="post-auth-login-parameters">Parameters</h3>

| Name       | In   | Type   | Required | Description |
| ---------- | ---- | ------ | -------- | ----------- |
| body       | body | object | false    | none        |
| » phone    | body | string | true     | none        |
| » password | body | string | true     | none        |

> Example responses

> Example response

```json
{
  "message": "string",
  "data": {
    "token": "string"
  }
}
```

```json
{
  "message": "string",
  "data": null
}
```

```json
{
  "message": "string",
  "data": null
}
```

<h3 id="post-auth-login-responses">Responses</h3>

| Status | Meaning                                                                    | Description      | Schema |
| ------ | -------------------------------------------------------------------------- | ---------------- | ------ |
| 200    | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)                    | Example response | Inline |
| 400    | [Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)           | Example response | Inline |
| 500    | [Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1) | Example response | Inline |

<h3 id="post-auth-login-responseschema">Response Schema</h3>

Status Code **200**

| Name      | Type   | Required | Restrictions | Description |
| --------- | ------ | -------- | ------------ | ----------- |
| » message | string | true     | none         | none        |
| » data    | object | true     | none         | none        |
| »» token  | string | true     | none         | none        |

Status Code **400**

| Name      | Type   | Required | Restrictions | Description |
| --------- | ------ | -------- | ------------ | ----------- |
| » message | string | true     | none         | none        |
| » data    | any    | false    | none         | none        |

Status Code **500**

| Name      | Type   | Required | Restrictions | Description |
| --------- | ------ | -------- | ------------ | ----------- |
| » message | string | true     | none         | none        |
| » data    | any    | false    | none         | none        |

<aside class="success">
This operation does not require authentication
</aside>

## get-auth-token

<a id="opIdget-auth-token"></a>

> Code samples

```shell
# You can also use wget
curl -X GET http://localhost:8080/api/auth/token/{token} \
  -H 'Accept: application/json'

```

```http
GET http://localhost:8080/api/auth/token/{token} HTTP/1.1
Host: localhost:8080
Accept: application/json

```

```javascript
const headers = {
  Accept: "application/json",
};

fetch("http://localhost:8080/api/auth/token/{token}", {
  method: "GET",

  headers: headers,
})
  .then(function (res) {
    return res.json();
  })
  .then(function (body) {
    console.log(body);
  });
```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Accept' => 'application/json'
}

result = RestClient.get 'http://localhost:8080/api/auth/token/{token}',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Accept': 'application/json'
}

r = requests.get('http://localhost:8080/api/auth/token/{token}', headers = headers)

print(r.json())

```

```php
<?php

require 'vendor/autoload.php';

$headers = array(
    'Accept' => 'application/json',
);

$client = new \GuzzleHttp\Client();

// Define array of request body.
$request_body = array();

try {
    $response = $client->request('GET','http://localhost:8080/api/auth/token/{token}', array(
        'headers' => $headers,
        'json' => $request_body,
       )
    );
    print_r($response->getBody()->getContents());
 }
 catch (\GuzzleHttp\Exception\BadResponseException $e) {
    // handle exception or api errors.
    print_r($e->getMessage());
 }

 // ...

```

```java
URL obj = new URL("http://localhost:8080/api/auth/token/{token}");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("GET");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Accept": []string{"application/json"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("GET", "http://localhost:8080/api/auth/token/{token}", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`GET /auth/token/{token}`

_Your GET endpoint_

<h3 id="get-auth-token-parameters">Parameters</h3>

| Name  | In   | Type   | Required | Description |
| ----- | ---- | ------ | -------- | ----------- |
| token | path | string | true     | none        |

> Example responses

> Example response

```json
{
  "message": "string",
  "data": {
    "phone": "string",
    "name": "string",
    "role": "string"
  }
}
```

```json
{
  "message": "string",
  "data": null
}
```

```json
{
  "message": "string",
  "data": null
}
```

<h3 id="get-auth-token-responses">Responses</h3>

| Status | Meaning                                                                    | Description      | Schema |
| ------ | -------------------------------------------------------------------------- | ---------------- | ------ |
| 200    | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)                    | Example response | Inline |
| 400    | [Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)           | Example response | Inline |
| 500    | [Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1) | Example response | Inline |

<h3 id="get-auth-token-responseschema">Response Schema</h3>

Status Code **200**

| Name      | Type   | Required | Restrictions | Description |
| --------- | ------ | -------- | ------------ | ----------- |
| » message | string | true     | none         | none        |
| » data    | object | true     | none         | none        |
| »» phone  | string | true     | none         | none        |
| »» name   | string | true     | none         | none        |
| »» role   | string | true     | none         | none        |

Status Code **400**

| Name      | Type   | Required | Restrictions | Description |
| --------- | ------ | -------- | ------------ | ----------- |
| » message | string | true     | none         | none        |
| » data    | any    | false    | none         | none        |

Status Code **500**

| Name      | Type   | Required | Restrictions | Description |
| --------- | ------ | -------- | ------------ | ----------- |
| » message | string | true     | none         | none        |
| » data    | any    | false    | none         | none        |

<aside class="success">
This operation does not require authentication
</aside>

# Schemas

<h2 id="tocS_Auth">Auth</h2>
<!-- backwards compatibility -->
<a id="schemaauth"></a>
<a id="schema_Auth"></a>
<a id="tocSauth"></a>
<a id="tocsauth"></a>

```json
{
  "phone": "string",
  "name": "string",
  "role": "string",
  "password": "string",
  "timestamp": "string"
}
```

Auth

### Properties

| Name      | Type   | Required | Restrictions | Description |
| --------- | ------ | -------- | ------------ | ----------- |
| phone     | string | true     | none         | none        |
| name      | string | true     | none         | none        |
| role      | string | true     | none         | none        |
| password  | string | false    | none         | none        |
| timestamp | string | false    | none         | none        |
