package billing

import (
	"math"
	"time"

	"github.com/Brilhante29/kiri-aws/internal/clock"
)

// PriceCalculator is a generic function that calculates cost.
type PriceCalculator func(res *Resource, usage []UsageEvent, start, end time.Time) float64

var awsCatalog = map[string]PriceCalculator{
	"s3": func(res *Resource, usage []UsageEvent, start, end time.Time) float64 {
		cost := 0.0
		
		// S3 Pricing (mocked): $0.023 per GB-month. Let's say we just charge a flat $0.01 per day per bucket.
		// And $0.005 per 1000 requests (UsageEvent).
		
		// 1. Calculate uptime within the window
		resStart := res.CreatedAt
		if resStart.Before(start) {
			resStart = start
		}
		resEnd := end
		if res.DeletedAt != nil && res.DeletedAt.Before(end) {
			resEnd = *res.DeletedAt
		}
		
		if resEnd.After(resStart) {
			days := resEnd.Sub(resStart).Hours() / 24.0
			cost += days * 0.01 // flat 1 cent per day just for having a bucket
		}

		// 2. Add usage costs
		for _, u := range usage {
			if u.Timestamp.After(start) && u.Timestamp.Before(end) {
				if u.MetricName == "PutObject" || u.MetricName == "GetObject" {
					cost += (u.Value / 1000.0) * 0.005
				}
			}
		}

		return cost
	},
	"dynamodb": func(res *Resource, usage []UsageEvent, start, end time.Time) float64 {
		cost := 0.0
		// DynamoDB Pricing (mocked): $0.25 per GB-month (flat $0.05/day/table)
		resStart := res.CreatedAt
		if resStart.Before(start) {
			resStart = start
		}
		resEnd := end
		if res.DeletedAt != nil && res.DeletedAt.Before(end) {
			resEnd = *res.DeletedAt
		}
		
		if resEnd.After(resStart) {
			days := resEnd.Sub(resStart).Hours() / 24.0
			cost += days * 0.05 
		}

		// Usage costs
		for _, u := range usage {
			if u.Timestamp.After(start) && u.Timestamp.Before(end) {
				if u.MetricName == "WCU" || u.MetricName == "RCU" {
					cost += u.Value * 0.0001
				}
			}
		}

		return cost
	},
	"sqs": func(res *Resource, usage []UsageEvent, start, end time.Time) float64 {
		cost := 0.0
		// SQS is typically only charged per request, no flat hourly fee.
		// Mock pricing: $0.40 per 1 million requests.
		for _, u := range usage {
			if u.Timestamp.After(start) && u.Timestamp.Before(end) {
				if u.MetricName == "SendMessage" || u.MetricName == "ReceiveMessage" {
					cost += (u.Value / 1000000.0) * 0.40
				}
			}
		}
		return cost
	},
	"kms": func(res *Resource, usage []UsageEvent, start, end time.Time) float64 {
		cost := 0.0
		// KMS: $1 per month per key. We do $0.033 per day.
		resStart := res.CreatedAt
		if resStart.Before(start) {
			resStart = start
		}
		resEnd := end
		if res.DeletedAt != nil && res.DeletedAt.Before(end) {
			resEnd = *res.DeletedAt
		}
		
		if resEnd.After(resStart) {
			days := resEnd.Sub(resStart).Hours() / 24.0
			cost += days * 0.033 
		}
		return cost
	},
}

// CalculateCost calculates the cost per service for a given time window.
// It returns a map of ServiceName -> Cost.
func (l *Ledger) CalculateCost(start, end time.Time) map[string]float64 {
	l.mu.RLock()
	defer l.mu.RUnlock()

	costs := make(map[string]float64)

	// Group usage by resource ARN for faster processing
	usageByRes := make(map[string][]UsageEvent)
	for _, u := range l.usage {
		usageByRes[u.ResourceARN] = append(usageByRes[u.ResourceARN], u)
	}

	for _, res := range l.resources {
		// Currently only AWS pricing logic is hardcoded.
		if res.Provider == "AWS" {
			if calc, ok := awsCatalog[res.Service]; ok {
				svcCost := calc(res, usageByRes[res.ARN], start, end)
				costs[res.Service] += svcCost
			}
		}
	}

	// For neatness, let's round to 4 decimals.
	for k, v := range costs {
		costs[k] = math.Round(v*10000) / 10000
	}

	return costs
}

// Ensure "start" and "end" default to logical bounds if not provided.
func (l *Ledger) EvaluateTotalCost() map[string]float64 {
	return l.CalculateCost(time.Time{}, clock.Now())
}
