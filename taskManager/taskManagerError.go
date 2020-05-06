package taskManager

import (
	"errors"
	"strconv"
	"strings"
)

func checkTaskNameConflict(task string) bool {
	rows, err := db.Query(`SELECT * FROM TASKS WHERE TASK=?`, task)
	defer rows.Close()
	if err != nil {
		return true
	}

	for rows.Next() {
		return true
	}
	return false
}

func checkDatePatarn(date string) (map[string]int, error) {
	dateStrings := strings.Split(date, "/")

	if len(dateStrings) < 1 {
		return nil, errors.New("invalid patarn")
	}

	rawMonth := dateStrings[0]
	rawDay := dateStrings[1]

	month, err := strconv.Atoi(rawMonth)
	if err != nil || month < 1 || month > 12 {
		return nil, errors.New("mouth cannot convert to int")
	}

	day, err := strconv.Atoi(rawDay)
	if err != nil || day < 1 || day > 31 {
		return nil, errors.New("day cannot convert to int")
	}

	return map[string]int{"month": month, "day": day}, nil
}
