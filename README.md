# Jed-NCA: Jed's Number Classification API
▶️ This API classifies a given number and returns its mathematical properties along with a fun fact
▶️ This API is written in Go

## Setup Instructions
### Prerequisites
- Go 1.20 or higher

### Running locally
1. Clone this repo:
```bash
git clone https://github.com/Jedway/jed-nca.git
```
2. Navigate to the project directory:
```bash
cd jed-nca
```
3. Run the server:
```bash
go run main.go
```
4. Access the API at port 8080, appending "/api/classify-number?number=x" at the end by either using curl:
```bash
curl "http://localhost:8080/api/classify-number?number=22"
```
Or, going to `http://localhost:8080/api/classify-number?number=22` with your browser of choice
Take note: `x` can be replaced with any number of your choice

### Endpoint
`GET /api/classify-number?number=<number>`

## Example request

```bash
curl "https://your-app.herokuapp.com/api/classify-number?number=371"
```
## Example response

```json
{
    "number": 371,
    "is_prime": false,
    "is_perfect": false,
    "properties": ["armstrong", "odd"],
    "digit_sum": 11,
    "fun_fact": "371 is an Armstrong number because 3^3 + 7^3 + 1^3 = 371"
}
```
