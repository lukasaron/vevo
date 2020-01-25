// Package vevo implements Visa Entitlement Verification Online system that
// is the Australian Visa check online system:
// https://immi.homeaffairs.gov.au/visas/already-have-a-visa/check-visa-details-and-conditions/overview
//
// This library allows users to check their Australian Visa without browser and facilitate the whole
// process of filling forms in any internet browser.
//
// To to get valid result basic information is needed such as Passport number,
// Country code of the specific country in which the passport has been issued, date of birth of the
// Visa applicant and lastly the number of the visa.
//
// There are two types of identification of the visa: Visa Grant Number or
// Transaction Reference Number. Any of them can be used to get the Visa status.
//
// Installation
//		go get github.com/lukasaron/vevo
//
// Example of usage
//		package main
//
//		import (
//			"fmt"
//			"github.com/lukasaron/vevo"
//			"log"
//			"time"
//		)
//
//		func main() {
//			// date of birth
//			dob := time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC)
//			// passport number
//			passport := "123456P"
//			// country code
//			cc := "AUS"
//			// visa grant number or transaction reference number
//			vgn := "1234567891011"
//			v := vevo.NewVEVO(dob, passport, cc, vgn)
//			visa, err := v.Visa()
//			if err != nil {
//				log.Fatal(err)
//			}
//
//			fmt.Printf("%+v\n", visa)
//		}
package vevo
