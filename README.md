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

**Description:** Performs a refresh operation using a cookie refresh_token. Only guid is needed to refresh tokens.

**Request Parameters:**
- `guid`: The unique identifier (GUID) of the user.

**Response:**
- `200 OK`: Returns the access token in the response body. And refresh in the cookie
