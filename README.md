# Numbers API
## Just playing around HNG tasks

FUNCTIONS

func classifyNumberHandler(w http.ResponseWriter, r *http.Request)
    classifyNumberHandler is a handle function for the api route to to write an
    http response

func computeProperties(isArmstrong bool, isOdd bool) []string
    compute properties returns a slice for the json response property field.

    Returns a slice of string with allowed values ["armstrong", "odd", "even"]

func digitSum(n int) int
    digitSum calculates the sum of the all the digits in the number.

    returns int.

func getFunFact(n int) string
    getFunFact is a function that gets a random math funfact using
    http://numbersapi.com/{n}/math where n is the number.

    Returns a string.

func isArmstrong(n int) bool
    isArmstrong checks if a number is an armstrong number.

    A number is regarded as an armstrong number if the sum of its digits each
    raised to the power of the number of digits equals itself.

    returns a bool.

func isOdd(n int) bool
    isOdd checks if a number is odd.

    A number is odd if it is not diviisble by 2 without having remainders.

    Retuns a bool.

func isPerfect(n int) bool
    isPerfect checks if a number is perfect.

    A number is perfect if the sum of its positive proper divisor excluding
    itself.

    Retuns a bool.

func isPrime(n int) bool
    isPrime checks if a number is prime.

    A number is prime if it is only divisble by 1 and itself.

    Retuns a bool.

func main()
func sendResponse(w http.ResponseWriter, statusCode int, message interface{})
    sendResponse is an helper function to write certain headers into a response.


TYPES

type errorResponse struct {
	Number string `json:"number"`
	Error  bool   `json:"error"`
}
    successResponse contains fields for a 400 BadRequest response

type successResponse struct {
	Number     int      `json:"number"`
	IsPrime    bool     `json:"is_prime"`
	IsPerfect  bool     `json:"is_perfect"`
	Properties []string `json:"properties"`
	DigitSum   int      `json:"digit_sum"`
	FunFact    string   `json:"fun_fact"`
}
    successResponse contains fields for a 200 OK response

func newSuccessResponse(n int) successResponse
    newSuccessResponse creates an instance of successResponse.

    It makes use of go routines to run all functions asyncronously to improve
    the speed of generating a response.

    Returns successResponse

