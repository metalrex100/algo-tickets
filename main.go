package main

import (
	"github.com/metalrex100/algo-tester"
	"log"
	"strconv"
)

func getStrLenTask() algo_tester.Task {
	return func(data []string) string {
		return strconv.Itoa(len(data[0]))
	}
}

func getLuckyTicketsTask() algo_tester.Task {
	return func(data []string) string {
		const maxDigit = 9

		digitsInTicket, err := strconv.Atoi(data[0])
		if err != nil {
			log.Fatalln(err)
		}

		lastRow := getRowForSingleDigit()

		offset := 0
		for i := 1; i < digitsInTicket; i++ {
			currentRow := make(map[int]int64, i*maxDigit)
			for j := 0; j <= maxDigit; j++ {
				for k, v := range lastRow {
					currentRow[k+offset] += v
				}
				offset++
			}
			lastRow = currentRow
		}

		var luckyTicketsCount uint64
		for _, sum := range lastRow {
			luckyTicketsCount += uint64(sum * sum)
		}

		return strconv.FormatUint(luckyTicketsCount, 10)
	}
}

func getRowForSingleDigit() map[int]int64 {
	return map[int]int64{0:1, 1:1, 2:1, 3:1, 4:1, 5:1, 6:1, 7:1, 8:1, 9:1}
}