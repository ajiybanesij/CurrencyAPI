# Currency API Doc

## Running

This API run on Docker.
- Docker Image  https://hub.docker.com/repository/docker/1864/currencyapi
- Postman Collection  [Here](https://github.com/ajiybanesij/CurrencyAPI/blob/main/Currency%20API.postman_collection.json)
- Swagger YAML  [Here](https://github.com/ajiybanesij/CurrencyAPI/blob/main/swagger.yaml)

### Requirement and Dependencies

- Docker

### Start API

- The following command runs **PostgreSQL**, **Redis**, and **API**  on Docker.
  ```bash
      make start
  ```

### Stop API

- The following command stops **PostgreSQL**, **Redis**, and **API**  on Docker.
  ```bash
      make stop
  ```

---

## Flow of Currency Converting

In the following section you see how to use apis.

1. Register user
2. Login user
3. Create wallet of currencies
4. Get offer for converting
5. Accept offer

---

## API Requests

You can find Postman Collection as a JSON and also Swagger yaml file in the repository.

- Requests accept a JSON body and return a JSON structure with "**status**", "**message**", and "**data**".
- "**status**" is a boolean outcome of the process.
- "**message**" gives two distinct types of results. The string "success" is returned if the result is successful;
  otherwise, an error message is given.
- "**data**" is result of requests. If the result is successful, the required data is returned; otherwise, it returns
  null.
  ```
    {
      "status" : <BOOLEAN>,
      "message": <STRING>,
      "data"   : <ANY_TYPE>
    }
  ```

### User APIs

- **Register** [POST]
    - This endpoint enables user registration with the system.
        - Sample Body
      ```
      {
        "username": <STRING>,
        "password": <STRING>
      }
      ```
- **Login** [POST]
    - This endpoint is utilized for user authentication. This endpoint accepts the username and password credentials in
      that body.
    - If the request is successful, the response contains a jwt token.
        - Sample Body
      ```
      {
        "username": <STRING>,
        "password": <STRING>
      }
      ```

### Currency APIs

> **Note:** User must use JWT token in header.

- **Read Currency List** [GET]
    - This endpoint provides currency list.
        - Sample Response
      ```
      {
        "status": true,
        "message": "success",
        "data": [
            "TRY",
            "EURO",
            "USD"
         ]
      }
      ```

- **Read Currency List** [GET]
    - This endpoint provides rate of currencies.
        - Sample Response
      ```
      {
        "status": true,
        "message": "success",
        "data": {
            "EURO-TRY": 19.63,
            "EURO-USD": 1.05,
            "TRY-EURO": 0.051,
            "TRY-USD": 0.054,
            "USD-EURO": 0.95,
            "USD-TRY": 18.55
        }
      }
      ```

### Wallet APIs

> **Note**: User must use JWT token in header.

- **Create Wallet** [POST]
    - With this endpoint User can create new wallet. The currency parameter must be string type and one of then currency
      list. 1000 balance units are added to the newly constructed wallet.
        - Sample Body
      ```
      {
       "currency":<STRING>
      }
      ```
        - Sample Response
      ```
      {
        "status": true,
        "message": "success",
        "data": {
            "ID": 1,
            "Currency": "TRY",
            "Balance": 1000,
            "Transaction": null
        }
      }
      ```
- **Read All Wallets** [GET]
    - User can see own wallets.
        - Sample Response
      ```
      {
        "status": true,
        "message": "success",
        "data": [
            {
                "ID": 1,
                "Currency": "TRY",
                "Balance": 1000,
                "Transaction": null
            },
            {
                "ID": 2,
                "Currency": "USD",
                "Balance": 1000,
                "Transaction": null
            }
        ]
      }
      ```
    - **Read Wallet** [GET]
        - User can see own wallet by wallet id. If the user has converted foreign currency before, user can see the
          transactions history of this wallet.
            - Sample Response
          ```
          {
            "status": true,
            "message": "success",
            "data": {
                "ID": 1,
                "Currency": "TRY",
                "Balance": 814.8148,
                "Transaction": [
                    {
                        "ID": 1,
                        "CreatedAt": "2022-12-04T18:51:42.428025Z",
                        "UpdatedAt": "2022-12-04T18:51:42.428025Z",
                        "DeletedAt": null,
                        "FromCurrency": "TRY",
                        "FromAmount": -185.1851,
                        "FromWalletID": 1,
                        "ToCurrency": "USD",
                        "ToAmount": 10,
                        "ToWalletID": 2,
                        "Rate": 0.054,
                        "UserID": 1
                    }
                ]
             }
          }
          ```

### Offer APIs

> **Note**: User must use JWT token in header.

- **Create Offer** [GET]
    - User can start currency convert session using this api. User must select currencies and request returns currency
      rate, expire and offerID.
    - rate shows the current exchange rate.
    - expire is a time and this time indicates when the offer will expire. Offer expires 3 minutes after creation.
        - Sample Body : User wants give TRY and take USD.
      ```
      {
        "fromCurrency":"TRY",
        "toCurrency":"USD"
      }
      ```
        - Sample Response
      ```
      {
        "status": true,
        "message": "success",
        "data": {
            "expire": 1670180074,
            "offerId": "c1448c56-f33c-48c3-b6e6-a7c9f95efae7",
            "rate": 0.054
         }
      }
      ```
- **Accept Offer** [GET]
    - User can accept offer using this api. User must enter offerId and amount.
        - Sample Body : User created offer and wants take USD. In the following body user entered 10 USD as a amount.
          ````
          {
            "offerId":"c1448c56-f33c-48c3-b6e6-a7c9f95efae7",
            "amount":10
          }
          ````
        - Sample Response
          ```
          {
            "status": true,
            "message": "success",
            "data": {
                "ID": 1,
                "CreatedAt": "2022-12-04T18:51:42.428025888Z",
                "UpdatedAt": "2022-12-04T18:51:42.428025888Z",
                "DeletedAt": null,
                "FromCurrency": "TRY",
                "FromAmount": -185.1851851851852,
                "FromWalletID": 1,
                "ToCurrency": "USD",
                "ToAmount": 10,
                "ToWalletID": 2,
                "Rate": 0.054,
                "UserID": 1
            }
          }
          ```
