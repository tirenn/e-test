---
title: Fetch v1.0
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

<h1 id="fetch">Fetch v1.0</h1>

> Scroll down for code samples, example requests and responses. Select a language for code samples from the tabs above or the mobile navigation menu.

Base URLs:

* <a href="http://localhost:8081/api">http://localhost:8081/api</a>

<h1 id="fetch-default">Default</h1>

## get-user

<a id="opIdget-user"></a>

> Code samples

```shell
# You can also use wget
curl -X GET http://localhost:8081/api/fetch \
  -H 'Accept: application/json' \
  -H 'Authorization: string'

```

```http
GET http://localhost:8081/api/fetch HTTP/1.1
Host: localhost:8081
Accept: application/json
Authorization: string

```

```javascript

const headers = {
  'Accept':'application/json',
  'Authorization':'string'
};

fetch('http://localhost:8081/api/fetch',
{
  method: 'GET',

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
  'Accept' => 'application/json',
  'Authorization' => 'string'
}

result = RestClient.get 'http://localhost:8081/api/fetch',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Accept': 'application/json',
  'Authorization': 'string'
}

r = requests.get('http://localhost:8081/api/fetch', headers = headers)

print(r.json())

```

```php
<?php

require 'vendor/autoload.php';

$headers = array(
    'Accept' => 'application/json',
    'Authorization' => 'string',
);

$client = new \GuzzleHttp\Client();

// Define array of request body.
$request_body = array();

try {
    $response = $client->request('GET','http://localhost:8081/api/fetch', array(
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
URL obj = new URL("http://localhost:8081/api/fetch");
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
        "Authorization": []string{"string"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("GET", "http://localhost:8081/api/fetch", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`GET /fetch`

<h3 id="get-user-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|Authorization|header|string|true|Bearer|

> Example responses

> Example response

```json
{
  "message": "string",
  "data": [
    {
      "uuid": "string",
      "komoditas": "string",
      "area_provinsi": "string",
      "area_kota": "string",
      "size": "string",
      "price": "string",
      "price_usd": "string",
      "tgl_parsed": "string",
      "timestamp": "string"
    }
  ]
}
```

```json
{
  "message": "string",
  "data": null
}
```

<h3 id="get-user-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Example response|Inline|
|401|[Unauthorized](https://tools.ietf.org/html/rfc7235#section-3.1)|Example response|Inline|

<h3 id="get-user-responseschema">Response Schema</h3>

Status Code **200**

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|» message|string|true|none|none|
|» data|[[Komoditas](#schemakomoditas)]|true|none|none|
|»» uuid|string|true|none|none|
|»» komoditas|string|true|none|none|
|»» area_provinsi|string|true|none|none|
|»» area_kota|string|true|none|none|
|»» size|string|true|none|none|
|»» price|string|true|none|none|
|»» price_usd|string|true|none|none|
|»» tgl_parsed|string|true|none|none|
|»» timestamp|string|true|none|none|

Status Code **401**

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|» message|string|true|none|none|
|» data|any|false|none|none|

<aside class="success">
This operation does not require authentication
</aside>

## get-fetch-aggregate

<a id="opIdget-fetch-aggregate"></a>

> Code samples

```shell
# You can also use wget
curl -X GET http://localhost:8081/api/fetch/aggregate \
  -H 'Accept: application/json' \
  -H 'Authorization: string'

```

```http
GET http://localhost:8081/api/fetch/aggregate HTTP/1.1
Host: localhost:8081
Accept: application/json
Authorization: string

```

```javascript

const headers = {
  'Accept':'application/json',
  'Authorization':'string'
};

fetch('http://localhost:8081/api/fetch/aggregate',
{
  method: 'GET',

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
  'Accept' => 'application/json',
  'Authorization' => 'string'
}

result = RestClient.get 'http://localhost:8081/api/fetch/aggregate',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Accept': 'application/json',
  'Authorization': 'string'
}

r = requests.get('http://localhost:8081/api/fetch/aggregate', headers = headers)

print(r.json())

```

```php
<?php

require 'vendor/autoload.php';

$headers = array(
    'Accept' => 'application/json',
    'Authorization' => 'string',
);

$client = new \GuzzleHttp\Client();

// Define array of request body.
$request_body = array();

try {
    $response = $client->request('GET','http://localhost:8081/api/fetch/aggregate', array(
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
URL obj = new URL("http://localhost:8081/api/fetch/aggregate");
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
        "Authorization": []string{"string"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("GET", "http://localhost:8081/api/fetch/aggregate", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`GET /fetch/aggregate`

*Your GET endpoint*

<h3 id="get-fetch-aggregate-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|Authorization|header|string|true|Bearer|

> Example responses

> Example response

```json
{
  "message": "string",
  "data": [
    {
      "area_provinsi": "string",
      "max_price": "string",
      "min_price": "string",
      "average_price": "string",
      "median_price": "string",
      "data": [
        {
          "uuid": "string",
          "komoditas": "string",
          "area_provinsi": "string",
          "area_kota": "string",
          "size": "string",
          "price": "string",
          "price_usd": "string",
          "tgl_parsed": "string",
          "timestamp": "string"
        }
      ]
    }
  ]
}
```

```json
{
  "message": "string",
  "data": null
}
```

<h3 id="get-fetch-aggregate-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Example response|Inline|
|401|[Unauthorized](https://tools.ietf.org/html/rfc7235#section-3.1)|Example response|Inline|

<h3 id="get-fetch-aggregate-responseschema">Response Schema</h3>

Status Code **200**

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|» message|string|true|none|none|
|» data|[object]|true|none|none|
|»» area_provinsi|string|true|none|none|
|»» max_price|string|true|none|none|
|»» min_price|string|true|none|none|
|»» average_price|string|true|none|none|
|»» median_price|string|true|none|none|
|»» data|[[Komoditas](#schemakomoditas)]|false|none|none|
|»»» uuid|string|true|none|none|
|»»» komoditas|string|true|none|none|
|»»» area_provinsi|string|true|none|none|
|»»» area_kota|string|true|none|none|
|»»» size|string|true|none|none|
|»»» price|string|true|none|none|
|»»» price_usd|string|true|none|none|
|»»» tgl_parsed|string|true|none|none|
|»»» timestamp|string|true|none|none|

Status Code **401**

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|» message|string|true|none|none|
|» data|any|false|none|none|

<aside class="success">
This operation does not require authentication
</aside>

## get-fetch-token-token

<a id="opIdget-fetch-token-token"></a>

> Code samples

```shell
# You can also use wget
curl -X GET http://localhost:8081/api/fetch/token/{token} \
  -H 'Accept: application/json'

```

```http
GET http://localhost:8081/api/fetch/token/{token} HTTP/1.1
Host: localhost:8081
Accept: application/json

```

```javascript

const headers = {
  'Accept':'application/json'
};

fetch('http://localhost:8081/api/fetch/token/{token}',
{
  method: 'GET',

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
  'Accept' => 'application/json'
}

result = RestClient.get 'http://localhost:8081/api/fetch/token/{token}',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Accept': 'application/json'
}

r = requests.get('http://localhost:8081/api/fetch/token/{token}', headers = headers)

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
    $response = $client->request('GET','http://localhost:8081/api/fetch/token/{token}', array(
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
URL obj = new URL("http://localhost:8081/api/fetch/token/{token}");
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
    req, err := http.NewRequest("GET", "http://localhost:8081/api/fetch/token/{token}", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`GET /fetch/token/{token}`

*Your GET endpoint*

<h3 id="get-fetch-token-token-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|token|path|string|true|none|

> Example responses

> 200 Response

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

> Example response

```json
{
  "message": "string",
  "data": null
}
```

<h3 id="get-fetch-token-token-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Example response|Inline|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|Example response|Inline|

<h3 id="get-fetch-token-token-responseschema">Response Schema</h3>

Status Code **200**

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|» message|string|true|none|none|
|» data|object|true|none|none|
|»» phone|string|true|none|none|
|»» name|string|true|none|none|
|»» role|string|true|none|none|

Status Code **500**

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|» message|string|true|none|none|
|» data|any|false|none|none|

<aside class="success">
This operation does not require authentication
</aside>

# Schemas

<h2 id="tocS_Komoditas">Komoditas</h2>
<!-- backwards compatibility -->
<a id="schemakomoditas"></a>
<a id="schema_Komoditas"></a>
<a id="tocSkomoditas"></a>
<a id="tocskomoditas"></a>

```json
{
  "uuid": "string",
  "komoditas": "string",
  "area_provinsi": "string",
  "area_kota": "string",
  "size": "string",
  "price": "string",
  "price_usd": "string",
  "tgl_parsed": "string",
  "timestamp": "string"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|uuid|string|true|none|none|
|komoditas|string|true|none|none|
|area_provinsi|string|true|none|none|
|area_kota|string|true|none|none|
|size|string|true|none|none|
|price|string|true|none|none|
|price_usd|string|true|none|none|
|tgl_parsed|string|true|none|none|
|timestamp|string|true|none|none|

