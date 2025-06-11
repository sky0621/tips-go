package main

import (
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/sky0621/tips-go/library_selection/converter/copier/dto"
	"github.com/sky0621/tips-go/library_selection/converter/copier/model"
	"time"
)

func main() {
	schoolModel := createModel()
	var schoolDTO dto.SchoolDTO
	err := copier.Copy(&schoolDTO, &schoolModel)
	if err != nil {
		panic(err)
	}
	fmt.Println(schoolDTO)
}

func createModel() model.School {
	return model.School{
		ID:        "school1",
		Name:      "テスト学校",
		No:        1,
		CreatedAt: time.Now(),
		Grades: []model.Grade{
			{
				ID:        "grade1",
				GradeName: "1年",
				No:        1,
				CreatedAt: time.Now(),
				Classes: []model.Class{
					{
						ID:   "class1",
						Name: "1組",
						Students: []model.Student{
							{
								ID:       "student1",
								Name:     "生徒1",
								IsReader: true,
							},
						},
					},
					{
						ID:   "class2",
						Name: "2組",
						Students: []model.Student{
							{
								ID:       "student2",
								Name:     "生徒2",
								IsReader: false,
							},
							{
								ID:       "student3",
								Name:     "生徒3",
								IsReader: true,
							},
						},
					},
				},
			},
		},
	}
}
