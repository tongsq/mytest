package main

import (
	"log"
)

func MaxProfit(prices []int) (int, int, int) {
	n := len(prices)
	if n <= 1 {
		return 0, 0, 0
	}
	buy := -prices[0]
	sell := 0
	buy_day := 0
	sell_day := 0
	buy_day_tmp := 0
	for i := 0; i < n; i++ {
		buy = max(buy, -prices[i])
		if buy == -prices[i] {
			buy_day_tmp = i
		}
		sell = max(sell, buy+prices[i])
		if sell == (buy + prices[i]) {
			sell_day = i
			buy_day = buy_day_tmp
		}
	}
	return sell, buy_day, sell_day
}

func SecondMaxProfit(prices []int) int {
	n := len(prices)
	if n <= 1 {
		return 0
	}
	firstBuy := -prices[0]
	firstSell := 0
	secondBuy := -prices[0]
	secondSell := 0
	for i := 0; i < n; i++ {
		firstBuy = max(firstBuy, -prices[i])

		firstSell = max(firstSell, firstBuy+prices[i])

		secondBuy = max(secondBuy, firstSell-prices[i])

		secondSell = max(secondSell, secondBuy+prices[i])
	}
	return secondSell
}

func max(args ...int) (m int) {
	m = args[0]
	for _, v := range args {
		if v > m {
			m = v
		}
	}
	return
}

func main() {
	prices := []int{1, 2, 10, 5, 3, 5, 7, 50, 1}
	log.Println(MaxProfit(prices))
	log.Println(SecondMaxProfit(prices))
}
