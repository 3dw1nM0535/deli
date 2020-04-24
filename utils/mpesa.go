package utils

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/joho/godotenv"
)

const (
	// MpesaTimeoutRequestMessage : Mpesa STK timeout message response
	MpesaTimeoutRequestMessage = "Request cancelled by user"

	// MpesaSuccessfulRequestMessage : Mpesa confirmation message response
	MpesaSuccessfulRequestMessage = "The service request is processed successfully."

	// MpesaInsufficientMessage : Mpesa insufficient message response
	MpesaInsufficientMessage = "The balance is insufficient for the transaction"

	// MpesaErrorMessage : Mpesa pending transaction error response message
	MpesaErrorMessage = "The transaction is being processed"

	// MpesaWrongPinMessage : Mpesa response message if initiator info is invalid
	MpesaWrongPinMessage = "The initiator information is invalid."

	transcType = "CustomerPayBillOnline"
)

var consumerKey,
	consumerSecret,
	authURL, paymentURL,
	shortCode,
	passKey,
	paymentConfirmationURL,
	encodedString,
	callbackURL string

// AuthResp : response
type AuthResp struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   string `json:"expires_in"`
}

// ErrorResp : error response
type ErrorResp struct {
	RequestID    string `json:"requestId"`
	ErrorCode    string `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
}

// ConfirmationResponse : payment response
type ConfirmationResponse struct {
	ResponseCode        string `json:"ResponseCode"`
	ResponseDescription string `json:"ResponseDescription"`
	MerchantRequestID   string `json:"MerchantRequestID"`
	CheckoutRequestID   string `json:"CheckoutRequestID"`
	ResultCode          string `json:"ResultCode"`
	ResultDesc          string `json:"ResultDesc"`
}

// PaymentResponse : payment confirmation respose
type PaymentResponse struct {
	ResponseCode        string `json:"ResponseCode"`
	ResponseDescription string `json:"ResponseDescription"`
	MerchantRequestID   string `json:"MerchantRequestID"`
	CheckoutRequestID   string `json:"CheckoutRequestID"`
	ResultCode          string `json:"ResultCode"`
	ResultDesc          string `json:"ResultDesc"`
}

// PaymentInput : payment processing data
type PaymentInput struct {
	BusinessShortCode string `json:"BusinessShortCode"`
	Password          string `json:"Password"`
	Timestamp         string `json:"Timestamp"`
	TransactionType   string `json:"TransactionType"`
	Amount            string `json:"Amount"`
	PartyA            string `json:"PartyA"`
	PartyB            string `json:"PartyB"`
	PhoneNumber       string `json:"PhoneNumber"`
	CallBackURL       string `json:"CallBackURL"`
	AccountReference  string `json:"AccountReference"`
	TransactionDesc   string `json:"TransactionDesc"`
}

// ConfirmPaymentInput : payment confirmation data input
type ConfirmPaymentInput struct {
	BusinessShortCode string `json:"BusinessShortCode"`
	Password          string `json:"Password"`
	Timestamp         string `json:"Timestamp"`
	CheckoutRequestID string `json:"CheckoutRequestID"`
}

func init() {
	godotenv.Load()
	consumerSecret = MustGetEnv("MPESA_CONSUMER_SECRET")
	consumerKey = MustGetEnv("MPESA_CONSUMER_KEY")
	authURL = MustGetEnv("AUTH_TOKEN_URL")
	paymentURL = MustGetEnv("MPESA_PAYMENT_URL")
	shortCode = MustGetEnv("BUSINESS_SHORT_CODE")
	passKey = MustGetEnv("MPESA_PASS_KEY")
	paymentConfirmationURL = MustGetEnv("PAYMENT_CONFIRMATION_URL")
	encodedString = base64.StdEncoding.EncodeToString([]byte(shortCode + passKey + formatted))
	callbackURL = MustGetEnv("MPESA_CALLBACK_URL")
}

var t time.Time = time.Now()
var formatted string = fmt.Sprintf("%d%02d%02d%02d%02d%02d",
	t.Year(), t.Month(), t.Day(),
	t.Hour(), t.Minute(), t.Second())

// GetToken : get oauth token from Safaricom Mpesa API
func GetToken() string {
	msg := consumerKey + ":" + consumerSecret
	enc := base64.StdEncoding.EncodeToString([]byte(msg))

	client := &http.Client{}
	req, err := http.NewRequest("GET", authURL, nil)
	if err != nil {
		panic("err when creating request " + err.Error())
	}

	req.Header.Set("Authorization", "Basic "+enc)

	res, err := client.Do(req)
	if err != nil {
		panic("error during request " + err.Error())
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic("err reading body content" + err.Error())
	}
	data := &AuthResp{}
	if err := json.Unmarshal(body, data); err != nil {
		panic(err)
	}
	return data.AccessToken
}

// MakePayment : process payment
func MakePayment(amount, phoneNumber string) string {
	jsonString := PaymentInput{
		BusinessShortCode: shortCode,
		Password:          encodedString,
		Timestamp:         formatted,
		TransactionType:   transcType,
		Amount:            amount,
		PartyA:            phoneNumber,
		PartyB:            shortCode,
		PhoneNumber:       phoneNumber,
		CallBackURL:       callbackURL,
		AccountReference:  "test",
		TransactionDesc:   "Food payment and delivery",
	}
	out, err := json.Marshal(jsonString)

	client := &http.Client{}
	req, err := http.NewRequest("POST", paymentURL, bytes.NewBuffer(out))
	if err != nil {
		panic("err making new request" + err.Error())
	}

	// Get auth token
	accessToken := GetToken()

	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		panic("err processing payment" + err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic("err reading body" + err.Error())
	}
	defer res.Body.Close()
	data := &PaymentResponse{}
	if err := json.Unmarshal(body, data); err != nil {
		panic(err)
	}
	return data.CheckoutRequestID
}

func delay(n time.Duration) {
	time.Sleep(n * time.Second)
}

// ConfirmPayment : confirms payment
func ConfirmPayment(checkOutID string) (*ConfirmationResponse, *ErrorResp, int) {
	jsonString := ConfirmPaymentInput{
		BusinessShortCode: shortCode,
		Password:          encodedString,
		Timestamp:         formatted,
		CheckoutRequestID: checkOutID,
	}

	out, _ := json.Marshal(jsonString)
	req, err := http.NewRequest("POST", paymentConfirmationURL, bytes.NewBuffer(out))
	if err != nil {
		panic(err)
	}

	// Get auth token
	accessToken := GetToken()

	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	// delay(20)
	res, err := client.Do(req)
	// check error code if 500
	if res.StatusCode == 500 {
		errorMessage := &ErrorResp{}
		body, _ := ioutil.ReadAll(res.Body)
		if err := json.Unmarshal(body, errorMessage); err != nil { // marshal body string
			panic(err)
		}
		if errorMessage.ErrorMessage == MpesaErrorMessage { // check if request is pending
			// delay 20 second before making a request
			delay(20)
			// call function again; will double delay time to 40seconds
			c, _, statusCode := ConfirmPayment(errorMessage.RequestID)
			return &ConfirmationResponse{
				ResponseCode:        c.ResponseCode,
				ResponseDescription: c.ResponseDescription,
				MerchantRequestID:   c.MerchantRequestID,
				CheckoutRequestID:   c.CheckoutRequestID,
				ResultCode:          c.ResultCode,
				ResultDesc:          c.ResultDesc,
			}, &ErrorResp{}, statusCode
		}
		return &ConfirmationResponse{}, &ErrorResp{
			RequestID:    errorMessage.RequestID,
			ErrorCode:    errorMessage.ErrorCode,
			ErrorMessage: errorMessage.ErrorMessage,
		}, res.StatusCode
	}
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	data := &ConfirmationResponse{}
	if err := json.Unmarshal(body, data); err != nil {
		panic(err)
	}
	defer res.Body.Close()
	return &ConfirmationResponse{
		ResponseCode:        data.ResponseCode,
		ResponseDescription: data.ResponseDescription,
		MerchantRequestID:   data.MerchantRequestID,
		CheckoutRequestID:   data.CheckoutRequestID,
		ResultCode:          data.ResultCode,
		ResultDesc:          data.ResultDesc,
	}, &ErrorResp{}, res.StatusCode
}
