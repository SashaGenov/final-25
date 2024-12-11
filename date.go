package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const TimeFormat = "20060102"

// NextDate вычисляет следующую дату для задачи с правилами

func NextDate(now time.Time, date string, repeat string) (string, error) {
	if repeat == "" {
		return "", errors.New("no repeat role")
	}

	// Проверка формата даты
	taskDate, err := time.Parse(TimeFormat, date)
	if err != nil {
		return "", fmt.Errorf("incorrect date: %v", err)
	}

	if repeat == "d 1" && !taskDate.After(now) {
		return now.Format(TimeFormat), nil
	}

	for {
		if repeat == "y" {
			// Добавляем 1 год к дате

			taskDate = taskDate.AddDate(1, 0, 0)
		} else if strings.HasPrefix(repeat, "d ") {
			// Извлекаем количество дней из правила повторения

			daysStr := strings.TrimPrefix(repeat, "d ")
			days, err := strconv.Atoi(daysStr)
			if err != nil || days < 1 || days > 400 {
				return "", fmt.Errorf("incorrect role: d %s", daysStr)
			}
			// Добавляем указанное количество дней к дате
			
			taskDate = taskDate.AddDate(0, 0, days)
		} else {
			return "", fmt.Errorf("incorrect format role: %s", repeat)
		}

		// Проверка, превышает ли новая дата текущую дату

		if taskDate.After(now) {
			return taskDate.Format(TimeFormat), nil
		}
	}
}
