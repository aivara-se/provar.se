package feedback

import (
	"strconv"
)

// SetFeedbackData sets additional feedback data fields from available data
func SetFeedbackData(m *map[string]string, feedbackType FeedbackType) {
	questionType := (*m)["question-type"]
	feedbackData := (*m)["response-data"]
	if feedbackType == CNPSFeedback {
		if questionType == "rating-11p" {
			rating, err := strconv.Atoi(feedbackData)
			if err != nil {
				return
			}
			if rating >= 0 && rating <= 6 {
				(*m)["response-type"] = "detractor"
			} else if rating > 6 && rating <= 8 {
				(*m)["response-type"] = "passive"
			} else if rating > 8 && rating <= 10 {
				(*m)["response-type"] = "promoter"
			}
		}
	}
	if feedbackType == CSATFeedback {
		if questionType == "rating-11p" {
			rating, err := strconv.Atoi(feedbackData)
			if err != nil {
				return
			}
			if rating >= 0 && rating <= 6 {
				(*m)["response-type"] = "low"
			} else if rating > 6 && rating <= 8 {
				(*m)["response-type"] = "mid"
			} else if rating > 8 && rating <= 10 {
				(*m)["response-type"] = "high"
			}
		}
	}
}
