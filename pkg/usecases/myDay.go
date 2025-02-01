package usecases

import (
	"time"
	"toDoApp/pkg/model"
)

type MyDayUseCase struct {
	MyDayRepository model.MyDayRepository
}

func NewMyDayUseCase(myDayRepository model.MyDayRepository) *MyDayUseCase {
	return &MyDayUseCase{
		MyDayRepository: myDayRepository,
	}
}

func (uc *MyDayUseCase) Get(userID int64, date string) (*model.MyDay, error) {

	formatedDate, err := formattingDate(date)
	if err != nil {
		return nil, err
	}

	myDay, err := uc.MyDayRepository.Get(userID, formatedDate)
	return myDay, err
}

func formattingDate(inputDate string) (string, error) {
	parseDate, err := time.Parse("01-02-2006", inputDate)
	if err != nil {
		return "", err
	}

	fornatedDate := parseDate.Format(time.RFC3339)
	return fornatedDate, nil
}
