package services

import (
	"encoding/csv"
  "log"
  "os"
	"fmt"
	"time"
	"strings"
	"strconv"
	"github.com/juliocesar1235/bank-api/models"
	"github.com/juliocesar1235/bank-api/db"
	"github.com/juliocesar1235/bank-api/utilities"
)

type ITransactionService interface {
	GetTransactionsByAccountId(accountId string, db db.Database) (*models.TransactionList, error)
}

type TransactionService struct {
	ITransactionService
}

type TransactionResume struct {
	TotalBalance 	float64 					`json:"totalBalance"`
	Report				TransactionReport	`json:"MonthlyReport"`
}

type TransactionReport struct {
	TransactionQuantities 	[]TransactionQuantity	 	`json:"transactionQuantities"`
	AverageAmounts 					[]AverageAmount					`json:"averageAmounts"`
}

type TransactionQuantity struct {
	Quantity 				int64 	`json:"quantity"`
	DateFrequency		string	`json:"dateFrequency"`
	DateAtFrequency string	`json:"dateRegistered`
}

type AverageAmount struct {
	Amount						float64		`json:"amount"`
	TrnType						string		`json:"trnType"`
	DateAtFrequency 	string		`json:"dateRegistered"`
}


var dbInstance db.Database

func (service TransactionService) GetTransactionsByAccountId(accountId string, db db.Database) (*models.TransactionList, error) {
	dbInstance = db
	transactions, err := dbInstance.GetTransactionsByAccountId(accountId)
	if err != nil {
			return nil, err
	}
	filep := fmt.Sprintf("./%s-txns.csv",accountId)
	csvRecords := readCSVFile(filep)
	log.Printf("CSV Records %s", csvRecords)

	//report := &TransactionReport{}
	//resume := &TransactionResume{}
	reportsCount := map[string]int{}
	amountByType := map[string]float64{
		"credit": 0.0,
		"debit": 	0.0,
	} 
	var totalAmount float64


	monthMap := utilities.GetMonthMap()

	for i, txnsArr := range csvRecords {
		if i < 1 {
			continue
		}
		dateArr := strings.Split(txnsArr[1],"-")
		amountInTrn, _ := strconv.ParseFloat(txnsArr[2], 8) 

		currentYear := time.Now().Year()

		// Convert the month name to a numeric value
		month, missing := monthMap[dateArr[1]]
		if !missing {
			fmt.Println("Invalid month name")
			return nil, err
		}

		day, err := strconv.Atoi(dateArr[0])
    if err != nil {
        return nil, err
    }

		dateConverted := time.Date(currentYear, month, day, 0, 0, 0, 0, time.UTC)
		fmt.Println(dateConverted.Format("2006-01-02"))
		
		monthNameStr := dateConverted.Month().String()

		_, ok := reportsCount[monthNameStr]
		// If the key exists
		if !ok {
			reportsCount[monthNameStr] = 0
		}
		reportsCount[monthNameStr]++

		totalAmount = totalAmount + amountInTrn
		if(amountInTrn < 0) {
			amountByType["debit"] += amountInTrn
		} else {
			amountByType["credit"] += amountInTrn
		}
	}
	fmt.Println(reportsCount)
	fmt.Println(totalAmount)
	fmt.Println(amountByType)

	return transactions, nil
}

func readCSVFile(filePath string) [][] string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file " + filePath, err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
			log.Fatal("Unable to parse file as CSV for " + filePath, err)
	}

  return records
}


func mapTransactions(reportsCount map[string]int, amountByType map[string]float64, totalAmount float64) (*TransactionResume, error) {
	resume := &TransactionResume{}
	report := &TransactionReport{}
	resume.TotalBalance = totalAmount
	txnsQuantities := []TransactionQuantity{}
	for key := range reportsCount{
		txnsQuantity := TransactionQuantity{}
		txnsQuantities :=
	}
}