[![Go](https://github.com/nmalinowski/time-go/actions/workflows/go.yml/badge.svg)](https://github.com/nmalinowski/time-go/actions/workflows/go.yml)
# time-go API

This is a simple API written in Go that allows you to fetch the current time in any timezone.

## Endpoints

### GET /time

Fetches the current time in the specified timezone.

#### Query Parameters

- `location`: The timezone to fetch the time for. This should be a valid IANA Time Zone database name.
- See: https://en.wikipedia.org/wiki/List_of_tz_database_time_zones for a full list.

#### Example

```
GET /time?location=America/New_York
```

#### Response

The API will return a JSON object with the following structure:

```json
{
  "time": "2022-01-01T00:00:00Z"
}
```

If there's an error, the response will have an `error` field:

```json
{
  "error": "Invalid timezone"
}
```

## Running the API

1. Make sure you have Go installed on your machine.
2. Clone this repository.
3. Navigate to the project directory.
4. Set `CERT_FILE_PATH` and `KEY_FILE_PATH` in your environment variables.
5. Run `go run main.go` to start the server.
6. The server will be running on `http://localhost:8080`.

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

[MIT](https://choosealicense.com/licenses/mit/)
```

Remember to replace the example URL and response with the actual URL and response structure you expect from your API. Also, ensure that you have the correct license for your project.