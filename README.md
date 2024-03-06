## Routes

### Generate Token

**Endpoint:** `POST /generate-token/:guid` 

**Description:** Generates a pair of Access and Refresh tokens for the user identified by the GUID provided in the request parameter. Inject refresh token into cookie http-only

**Request Parameters:**
- `guid`: The unique identifier (GUID) of the user in the url 

**Response:**
- `200 OK`: Returns the generated access token in the response body and refresh in cookie

### Refresh Tokens

**Endpoint:** `POST /refresh/:guid`
### Token will be generated only if user exist

**Description:** Performs a refresh operation using a cookie refresh_token. Only guid is needed to refresh tokens.

**Request Parameters:**
- `guid`: The unique identifier (GUID) of the user.

**Response:**
- `200 OK`: Returns the access token in the response body. And refresh in the cookie


  ## Screenshots
  ### Postman screenshots cookie updated and created also
  
![Screenshot 2024-03-06 231216](https://github.com/MamushevArup/medods/assets/93328884/6df4d3e0-bac4-41ee-916e-3666ad654c12)
![Screenshot 2024-03-06 231739](https://github.com/MamushevArup/medods/assets/93328884/bdc59c1a-2e4d-4432-bbd4-214ffc01801c)
![Screenshot 2024-03-06 231858](https://github.com/MamushevArup/medods/assets/93328884/4c32597d-347f-488f-9454-21613606d63f)
