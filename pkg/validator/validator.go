package validator

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"time"
	"xyz-finance-api/internal/payment/domain"
)

func IsDataEmpty(fields []string, data ...interface{}) error {
	if len(fields) != len(data) {
		return errors.New("column names and data length mismatch")
	}

	for i, value := range data {
		switch v := value.(type) {
		case string:
			if v == "" {
				return fmt.Errorf("%s is empty", fields[i])
			}
		case int:
			if v == 0 {
				return fmt.Errorf("%s is empty", fields[i])
			}
		case time.Time:
			if v.IsZero() {
				return fmt.Errorf("%s is empty", fields[i])
			}
		case []interface{}:
			if len(v) == 0 {
				return fmt.Errorf("%s is empty", fields[i])
			}
		case []string:
			if len(v) == 0 {
				return fmt.Errorf("%s is empty", fields[i])
			}
		case []int:
			if len(v) == 0 {
				return fmt.Errorf("%s is empty", fields[i])
			}
		default:
			if reflect.TypeOf(v).Kind() == reflect.Slice {
				slice := reflect.ValueOf(v)
				if slice.Len() == 0 {
					return fmt.Errorf("%s is empty", fields[i])
				}
			} else {
				return fmt.Errorf("unsupported data type for %s: %T", fields[i], v)
			}
		}
	}
	return nil
}

func IsEmailValid(email string) error {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	if !re.MatchString(email) {
		return errors.New("invalid email format. example: emailname@gmail.com")
	}
	return nil
}

func IsMinLengthValid(minLength int, fields map[string]string) error {
	for fieldName, fieldValue := range fields {
		if len(fieldValue) < minLength {
			return fmt.Errorf("minimum length for field %s is %d characters", fieldName, minLength)
		}
	}
	return nil
}

func IsMaxLengthValid(maxLength int, fields map[string]string) error {
	for fieldName, fieldValue := range fields {
		if len(fieldValue) > maxLength {
			return fmt.Errorf("maximum length for field %s is %d characters", fieldName, maxLength)
		}
	}
	return nil
}

func IsDataValid(data interface{}, validData []interface{}, caseSensitive bool) error {
	dataStr := fmt.Sprintf("%v", data)
	validDataStr := make([]string, len(validData))
	for i, v := range validData {
		validDataStr[i] = fmt.Sprintf("%v", v)
	}

	if !caseSensitive {
		dataStr = strings.ToLower(dataStr)
		for i, v := range validDataStr {
			validDataStr[i] = strings.ToLower(v)
		}
	}

	for _, validValue := range validDataStr {
		if dataStr == validValue {
			return nil
		}
	}

	return errors.New("invalid data. allowed data: " + strings.Join(validDataStr, ", "))
}

func IsDateValid(date string) error {
	if date == "" {
		return nil
	}

	dateRegex := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
	if !dateRegex.MatchString(date) {
		return errors.New("invalid date format. expected format: '2000-12-30'")
	}

	return nil
}

func CalculateUsedAmount(payments []domain.Payment) int {
	totalUsed := 0
	for _, payment := range payments {
		totalUsed += payment.AmountPaid
	}
	return totalUsed
}

func CalculateLoanLimit(salary int, tenor int) int {

	limitMultiplier := map[int]float64{
		1: 0.5, // 50% salary
		2: 0.7, // 70% salary
		3: 0.9, // 90% salary
		6: 1.2, // 120% salary
	}

	multiplier, exists := limitMultiplier[tenor]
	if !exists {
		multiplier = 0.5
	}

	return int(float64(salary) * multiplier)
}
