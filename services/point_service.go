package services

import (
	"log"
	"strings"
	"time"
	"math"
	"math/big"
	"receipt-processor/models"
)

func GetAdjNumOfItems(items []models.Item) int {
	numOfItems := len(items)
	adj := numOfItems % 2
	adjNumOfItems := (numOfItems - adj)
	pairsOfItems := adjNumOfItems / 2
	log.Printf("Pairs of Items: %d\n",pairsOfItems)
	return pairsOfItems
}

func GetPointsForRetailer(retailer string) int {
	points := 0
	for _, char := range retailer {
		if IsAlphaNumChar(char) {
			points += 1
		}
	}
	log.Printf("Points for Retailer: %d\n", points)
	return points
}

func GetPointsForTotal(amount *big.Rat) int {
	points := 0
	amountFloat,_ := amount.Float64()
	if IsRoundDollar(amount) {
		log.Printf("%f is round dollar amount\n", amountFloat)
		points += 50
	}
	if IsMultipleOfQuarter(amount) {
		log.Printf("%f is multiple of 0.25\n", amountFloat)
		points += 25
	}
	return points
}

func GetPointsForEveryTwoItems(items []models.Item) int {
	adjNumOfItems := GetAdjNumOfItems(items)
	points := adjNumOfItems * 5
	log.Printf("Points for Pairs of Items: %d\n", points)
	return points
}

func GetPointsForItemDescription(description string, price *big.Rat) int {
	trimmedDescription := strings.TrimSpace(description)
	trimmedDescriptionLength := len(trimmedDescription)
	evenlyDivides := EvenlyDivides(trimmedDescriptionLength, 3)
	if evenlyDivides {
		priceFloat,_ := price.Float64()
		rawPoints := priceFloat * 0.2
		points := int(math.Ceil(rawPoints))
		log.Printf("Points for Item with Description `%s` and Price `%f`: %d\n", description, priceFloat, points)
		return points
	}
	return 0
}

func GetPointsForItemDescriptions(items []models.Item) int {
	points := 0
	for _, item := range items {
		points += GetPointsForItemDescription(item.ShortDescription, &item.Price)
	}
	return points
}

func GetPointsForPurchaseDate(purchaseDate time.Time) int {
	day := purchaseDate.Day()
	isOdd := !EvenlyDivides(day, 2)
	if isOdd {
		log.Printf("Points for Purchase Date with Odd Day: %d\n", 6)
		return 6
	}
	return 0
}

func GetPointsForPurchaseTime(purchaseTime time.Time) int {
	layout := "15:04"
	startTime, _ := time.Parse(layout, "14:00")
	endTime, _ := time.Parse(layout, "16:00")
	timeIsInRange := purchaseTime.After(startTime) && purchaseTime.Before(endTime)
	if timeIsInRange {
		log.Printf("Points for Purchase Time between 2 pm and 4 pm: %d\n", 6)
		return 10
	}
	return 0
}