# Introduction

The Remington API is a set of HTTP endpoints that adhere to RESTful design
principles and CRUD actions with predictable URIs. It uses standard HTTP
response codes, authentication, and verbs. The API has consistent and
well-formed JSON requests and responses with pagination to simplify list
handling. Error messages are descriptive and easy to understand. All functions
of both the "Admin Panel" and the "Storefront" services are accessible via the
API.

## Requests

Communicate with the API by making an HTTP request at the correct endpoint. The
chosen method determines the action taken.

| Method | Usage                                                                                                                                                           |
| ------ | --------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| DELETE | Use the DELETE method to destroy a resource from the database. If it is not found, the operation will return a 4xx error and an appropriate message.            |
| GET    | To retrieve information about a resource, use the GET method. The data is returned as a JSON object. GET methods are read-only and do not affect any resources. |
| PATCH  | Some resources support partial modification with PATCH, which modifies specific attributes without updating the entire object representation.                   |
| POST   | Issue a POST method to create a new object. Include all needed attributes in the request body encoded as JSON.                                                  |
| PUT    | Use the PUT method to update information about a resource. PUT will set new values on the item without regard to their current values.                          |

## Response Codes

We use standard HTTP response codes to show the success or failure of requests.
Response codes in the 2xx range indicate success, while codes in the 4xx range
indicate an error, such as an authorization failure or a malformed request. All
4xx errors will return a JSON response object with an error attribute explaining
the error. Codes in the 5xx range indicate a server-side problem preventing the
Remington API from fulfilling your request.

| Response                  | Description                                                          |
| ------------------------- | -------------------------------------------------------------------- |
| 200 OK                    | The response contains your requested information.                    |
| 201 Created               | Your request was accepted. The resource was created.                 |
| 202 Accepted              | Your request was accepted. The resource was created or updated.      |
| 204 No Content            | Your request succeeded, there is no additional information returned. |
| 400 Bad Request           | Your request was malformed.                                          |
| 401 Unauthorized          | You did not supply valid authentication credentials.                 |
| 403 Forbidden             | You are not allowed to perform that action.                          |
| 404 Not Found             | No results were found for your request.                              |
| 429 Too Many Requests     | Your request exceeded the API rate limit.                            |
| 500 Internal Server Error | We were unable to perform the request due to server-side problems.   |

## Pagination

TODO: Add documentation about pagination

## Parameters

It is possible to pass information to the API with three different types of
parameters.

### Path Parameters

Some API calls require variable parameters as part of the endpoint path. For
example, to retrieve information about a product, supply the `product-id` in the
endpoint.

```console
curl "http://localhost:8000/users/{product-id}" \
  -X GET \
  -H "Authorization: Bearer ${API_KEY}"
```

### Query Parameters

Some API calls allow filtering with query parameters.

TODO: Add documentation about the query parameters.

### Request Body

PUT, POST, and PATCH methods may include an object in the request body with a
content type of `application/json` (unless otherwise specified). The
documentation for each endpoint below has more information about the expected
object.

## API Example Conventions

The examples in this documentation use `curl`, a command-line HTTP client, to
demonstrate usage. Linux and macOS computers usually have `curl` installed by
default, and it's [available for download](https://curl.se/download.html) on all
popular platforms including Windows.

Each of the protected endpoints require an `API_KEY` to be passed as a "bearer
token" along with the `Authorizaton` header. You can get a token for development
needs by making a GET request to the `/http://localhost:8000/auth/login` route
by invoking the following command:

```console
curl -X POST "http://localhost:8000/auth/login" \
 -H 'accept: application/json'\
 -H 'content-type: application/x-www-form-urlencoded' \
 -d email=dev.mode%40remington.bg&password=dev-mode
```

If the development PostgreSQL database was spun up and seeded appropriately then
a "Staff" (a user with access to the Admin Panel **ONLY**) with the email -
`dev.mode@remington.bg` and password - `dev-mode` will already exist!

Each example is split across multiple lines with the `\` character, which is
compatible with a bash terminal. A typical example looks like this:

```console
curl "http://localhost:8000/settings" \
  -X POST \
  -H "Authorization: Bearer ${API_KEY}" \
  -H "Content-Type: application/json" \
  --data '{
    "token" : "example-token",
  }'
```

- The `-X` parameter sets the request method. For consistency, we show the
  method on all examples, even though it's not explicitly required for GET
  methods.
- The `-H` lines set required HTTP headers. These examples are formatted to
  expand the `API_KEY` environment variable for your convenience.
- Examples that require a JSON object in the request body pass the required data
  via the `--data` parameter.
