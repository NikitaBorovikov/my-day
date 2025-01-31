package usecases

import "toDoApp/pkg/model"

type MyDayUseCase struct {
	MyDayRepository model.MyDayRepository
}

func NewMyDayUseCase(myDayRepository model.MyDayRepository) *MyDayUseCase {
	return &MyDayUseCase{
		MyDayRepository: myDayRepository,
	}
}

func (uc *MyDayUseCase) Get(userID int64, date string) (*model.MyDay, error) {
	return nil, nil
}
