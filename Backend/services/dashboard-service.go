package services

import "github.com/TusharChoudharykp/Email-Campaign-Tool/repositories"

func FetchDashboardStats() (map[string]int, error) {
	return repositories.GetDashboardStats()
}
