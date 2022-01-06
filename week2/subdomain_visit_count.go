package week2

import (
	"fmt"
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	countMap := map[string]int{}

	for i := range cpdomains {
		spaceIdx := strings.Index(cpdomains[i], " ")
		count, _ := strconv.Atoi(cpdomains[i][:spaceIdx])

		domain := cpdomains[i][spaceIdx+1:]
		for {
			countMap[domain] += count

			dotIdx := strings.Index(domain, ".")
			if dotIdx == -1 {
				break
			}
			domain = domain[dotIdx+1:]
		}
	}

	ans := make([]string, 0, len(countMap))
	for domain, count := range countMap {
		ans = append(ans, fmt.Sprintf("%d %s", count, domain))
	}
	return ans
}
