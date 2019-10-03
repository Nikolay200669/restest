# RESTEST

## REST FUll api all HTTP methods (GET, POST, UPDATE, ...) & stattic page

Write ./static/index.html and test.json

```bash
touch test.json && mkdir static && cd static && touch index.html
```

Write ./test.json

```javascript
[
    {
        "route": "/api/v1/current",
        "method": "POST",
        "request": "",
        "response": "[{\"shname\":\"BTCC\",\"loname\":\"Bitcoin\"},{\"shname\":\"ETH\",\"loname\":\"Ethereum\"},{\"shname\":\"BCH\",\"loname\":\"Bitcoin Cash\"},{\"shname\":\"XRP\",\"loname\":\"Ripple\"}]"
    },
    {
        "route": "/api/v1/test",
        "method": "GET",
        "request": "",
        "response": "[{\"shname\":\"BTCC\",\"loname\":\"Bitcoin\"},{\"shname\":\"ETH\",\"loname\":\"Ethereum\"},{\"shname\":\"BCH\",\"loname\":\"Bitcoin Cash\"},{\"shname\":\"XRP\",\"loname\":\"Ripple\"}]"
    },
    {
        "route": "/api/v1/{key}/{amount}/{currency}/{orderid}",
        "method": "GET",
        "request": "",
        "response": "[{\"shname\":\"BTCC\",\"loname\":\"Bitcoin\"},{\"shname\":\"ETH\",\"loname\":\"Ethereum\"},{\"shname\":\"BCH\",\"loname\":\"Bitcoin Cash\"},{\"shname\":\"XRP\",\"loname\":\"Ripple\"}]"
    }
]
```

## Use

Download file "w", crete file "test.json" & dir "static" with file "index.html"

Run file "w"

```bash
http://0.0.0.0:8181/api/v1/asdf366as4df654adf6/199.65/USD/2578-52415-965855
```

P.S. - CORS protection OFF
