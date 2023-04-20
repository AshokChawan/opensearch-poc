# Go application integration with Open Search. 
This application sends the data object to endpoint configured. It uses http client to send the data. 

1. Golang Gin boilerplate 
2. REST API /upload  

#How to run application 
1. Make 
2. Run ./dashboard 
3. curl command to upload the data 

```
curl --location 'http://localhost/upload' \
--header 'Content-Type: application/json' \
--data '[
    {
  "promotion_id": "promotion-01",
  "name": "promotion-01 ",
  "description": "Promotional Rewards",
  "image_url": "",
  "link_url": "",
  "total": 500000,
  "remaining": 400000,
  "start_date": 1651363200000,
  "end_date": 1675054800000,
  "pre_generated_codes": true,
  "require_input_code": true,
  "claims_per_user": 5,
  "show_code": true,
  "promotion_type": "external_code_redemption",
  "entitle_items": [
    "32031020",
    "31040030"
  ],
  "creator_code_required": false,
  "eligible_countries": null
},
{
  "promotion_id": "promotion-02",
  "name": "promotion-02 ",
  "description": "Promotional Rewards",
  "image_url": "",
  "link_url": "",
  "total": 500000,
  "remaining": 400000,
  "start_date": 1651363200000,
  "end_date": 1675054800000,
  "pre_generated_codes": true,
  "require_input_code": true,
  "claims_per_user": 5,
  "show_code": true,
  "promotion_type": "external_code_redemption",
  "entitle_items": [
    "32031020",
    "31040030"
  ],
  "creator_code_required": false,
  "eligible_countries": null
},
]
'
```


### Boilerplate structure

```
.
├── Makefile
├── Procfile
├── README.md
├── config
│   ├── config.go
│   ├── development.yaml
│   ├── production.yaml
│   └── test.yaml
├── controllers
│   └── user.go
├── db
│   └── db.go
├── forms
│   └── user.go
├── header.jpg
├── main.go
├── middlewares
│   └── auth.go
├── models
│   └── user.go
└── server
    ├── router.go
    └── server.go
```