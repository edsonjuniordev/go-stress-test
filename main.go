package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Report struct {
	TotalRequests      int
	SuccessRequests    int
	StatusDistribution map[int]int
	TotalTimeInSeconds float64
}

func worker(url string, requests int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < requests; i++ {
		resp, err := http.Get(url)
		if err != nil {
			results <- 0
			continue
		}
		results <- resp.StatusCode
		resp.Body.Close()
	}
}

func main() {
	url := flag.String("url", "", "URL do serviço a ser testado")
	totalRequests := flag.Int("requests", 100, "Número total de requests")
	concurrency := flag.Int("concurrency", 10, "Número de chamadas simultâneas")
	flag.Parse()

	if *url == "" {
		fmt.Println("A URL é obrigatória. Use --url para especificar.")
		return
	}

	startTime := time.Now()
	numRequestsPerWorker := *totalRequests / *concurrency
	remainingRequests := *totalRequests % *concurrency

	results := make(chan int, *totalRequests)
	var wg sync.WaitGroup

	for i := 0; i < *concurrency; i++ {
		requests := numRequestsPerWorker
		if i < remainingRequests {
			requests++
		}
		wg.Add(1)
		go worker(*url, requests, results, &wg)
	}

	wg.Wait()
	close(results)

	report := Report{
		TotalRequests:      *totalRequests,
		StatusDistribution: make(map[int]int),
	}

	for status := range results {
		if status == 200 {
			report.SuccessRequests++
		}
		report.StatusDistribution[status]++
	}

	report.TotalTimeInSeconds = time.Since(startTime).Seconds()

	fmt.Println("--- Resultado dos Testes ---")
	fmt.Printf("Tempo Total Gasto: %.2f segundos\n", report.TotalTimeInSeconds)
	fmt.Printf("Total de Requests: %d\n", report.TotalRequests)
	fmt.Printf("Requests com Status 200: %d\n", report.SuccessRequests)
	fmt.Println("Distribuição de Status HTTP:")
	for status, count := range report.StatusDistribution {
		if status == 0 {
			fmt.Printf("Erros: %d\n", count)
		} else {
			fmt.Printf("%d: %d\n", status, count)
		}
	}
}
