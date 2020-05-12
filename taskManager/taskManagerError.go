package taskManager

import (
	"fmt"
	"errors"
	"strconv"
	"strings"
)

func checkTaskNameConflict(task string) (isConflict bool) {
	isConflict = true

	rows, err := db.Query(`SELECT * FROM TASKS WHERE TASK=? LIMIT 1`, task)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		return
	}

	isConflict = false
	return
}

func checkDatePatarn(date string) (days map[string]int, err error) {
	dateStrings := strings.Split(date, "/")

	if len(dateStrings) < 1 {
		err = errors.New("invalid patarn")
		return
	}

	rawMonth := dateStrings[0]
	rawDay := dateStrings[1]

	month, err := strconv.Atoi(rawMonth)
	if err != nil || month < 1 || month > 12 {
		err = errors.New("mouth cannot convert to int")
		return
	}

	day, err := strconv.Atoi(rawDay)
	if err != nil || day < 1 || day > 31 {
		err = errors.New("day cannot convert to int")
		return
	}

	days = map[string]int{"month": month, "day": day}
	return
}
