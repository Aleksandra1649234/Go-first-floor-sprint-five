package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	dataSl := strings.Split(datastring, ",")
	if len(dataSl) != 3 {
		return errors.New("некорректная строка")
	}
	t.Steps, err = strconv.Atoi(dataSl[0])
	if err != nil {
		return err
	}
	if t.Steps <= 0 {
		return fmt.Errorf("колличество шагов <= 0")
	}
	t.TrainingType = dataSl[1]
	t.Duration, err = time.ParseDuration(dataSl[2])
	if err != nil {
		return err
	}
	if t.Duration.Minutes() <= 0 {
		return fmt.Errorf("время <= 0")
	}
	return nil
}

func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps, t.Personal.Height)
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Personal.Height, t.Duration)
	spentCalories := 0.0
	var err error
	switch t.TrainingType {
	case "Бег":
		spentCalories, err = spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
		if err != nil {
			return "", err
		}
		err = nil
	case "Ходьба":
		spentCalories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
		if err != nil {
			return "", err
		}
		err = nil
	default:
		return "", errors.New("неизвестный тип тренировки")
	}
	return fmt.Sprintf("Тип тренировки: %s\n"+
		"Длительность: %.2f ч.\n"+
		"Дистанция: %.2f км.\n"+
		"Скорость: %.2f км/ч\n"+
		"Сожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), distance, meanSpeed, spentCalories), nil

}
