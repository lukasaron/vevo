package vevo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	jsonURL = "https://www.ecom.immi.gov.au/evo/ws/first-party/json"
)

// VEVO defines the unite information package needed to get a Visa.
type VEVO struct {
	n   string    // visa number
	p   string    // passport number
	c   string    // country
	dob time.Time // date of birth
}

// NewVEVO function creates the instance of VEVO. All parameters are mandatory to get a valid response from the server.
// Date of Birth does not have to be precise in the manner of hours or minute - the date is enough.
// Country code is the 3 letter code of the visa issuer country.
// VisaNumber could be one of:
// 		Visa Grant Number
//		Transaction Reference Number
func NewVEVO(dob time.Time, passport, countryCode, visaNumber string) VEVO {
	return VEVO{
		n:   visaNumber,
		p:   passport,
		c:   countryCode,
		dob: dob,
	}
}

// Visa method handles the connection to the immigration department (government) and returns the Visa and/or error.
func (v VEVO) Visa() (Visa, error) {
	visa := Visa{}
	resp, err := http.Get(v.prepareURL())
	if err != nil {
		return visa, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	_ = resp.Body.Close()
	if err != nil {
		return visa, err
	}

	return visa, json.Unmarshal(body, &visa)
}

// Visa is the response expected from the immigration server.
type Visa struct {
	ClientEvoID        string `json:"clientEvoId"`
	DateOfEnquiry      string `json:"dateOfEnquiry"`
	VisaTypeCode       string `json:"visaTypeCode"`
	Success            bool   `json:"wsSuccess"`
	Error              string `json:"wsErrorMessage"`
	VisaDescription    string `json:"visaDescription"`
	PassportNumber     string `json:"passportNumber"`
	VisaClassSubclass  string `json:"visaClassSubclass"`
	VisaSubclass       string `json:"visaSubclass"`
	VisaStream         string `json:"visaStream"`
	VisaApplicant      string `json:"visaApplicant"`
	VisaGrantDate      string `json:"visaGrantDate"`
	VisaExpiryDate     string `json:"visaExpiryDate"`
	Location           string `json:"location"`
	VisaStatus         string `json:"visaStatus"`
	VisaGrantNumber    string `json:"visaGrantNumber"`
	EntriesAllowed     string `json:"entriesAllowed"`
	InitialStayDate    string `json:"initialStayDate"`
	MustNotArriveAfter string `json:"mustNotArriveAfter"`
	PeriodOfStay       string `json:"periodOfStay"`
	EnterBeforeDate    string `json:"enterBeforeDate"`
	VisaConditions     string `json:"visaConditions"`
}

// ------------------------------------------------ PRIVATE METHODS ------------------------------------------------

func (v VEVO) prepareURL() string {
	u, _ := url.Parse(jsonURL)
	q := u.Query()

	q.Add("passport", strings.ToUpper(v.p))
	q.Add("country", strings.ToUpper(v.c))
	q.Add("dateofbirth", v.dob.Format("20060102"))
	vn := strings.ToUpper(v.n)

	if strings.HasPrefix(vn, "E") { // with letter E starts Transaction Reference Number
		q.Add("trn", vn)
	} else { // otherwise it's a Visa Grant Number
		q.Add("visagrant", vn)
	}

	u.RawQuery = q.Encode()
	return u.String()
}
