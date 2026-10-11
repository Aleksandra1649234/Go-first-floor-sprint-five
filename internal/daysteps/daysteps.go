package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	dataSl := strings.Split(datastring, ",")
	if len(dataSl) != 2 {
		return errors.New("некорректная строка")
	}
	ds.Steps, err = strconv.Atoi(dataSl[0])
	if err != nil {
		return err
	}
	if ds.Steps <= 0 {
		return errors.New("колличество шагов <= 0")
	}
	ds.Duration, err = time.ParseDuration(dataSl[1])
	if err != nil {
		return err
	}
	if ds.Duration.Minutes() <= 0 {
		return errors.New("время <= 0")
	}
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Personal.Height)
	walkingSpentCalories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Personal.Weight, ds.Personal.Height, ds.Duration)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Количество шагов: %d.\n"+
		"Дистанция составила %.2f км.\n"+
		"Вы сожгли %.2f ккал.\n", ds.Steps, distance, walkingSpentCalories), nil
}
