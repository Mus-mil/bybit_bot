package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Функция для анализа цен
func analyzePrices(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var prices []float64
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			// Извлекаем цену из строки
			priceStr := line[len(""):]
			price, err := strconv.ParseFloat(priceStr, 64)
			if err == nil {
				prices = append(prices, price)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("No prices available for analysis.")
}

// Функция для вычисления средней цены
func calculateAverage(prices []float64) float64 {
	total := 0.0
	for _, price := range prices {
		total += price
	}
	return total / float64(len(prices))
}
