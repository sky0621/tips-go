package crud

import "github.com/sky0621/tips-go/library_selection/ormapper/sqlc/infra"

// 入れ子構造の定義
type StudentRow struct {
	StudentID int64
	Name      string
}

type ClassRow struct {
	ClassID   int64
	ClassName string
	Students  []StudentRow
}

type GradeRow struct {
	GradeID   int64
	GradeName string
	Classes   []ClassRow
}

type SchoolRow struct {
	SchoolID   int64
	SchoolName string
	Grades     []GradeRow
}

// aggregateRows はフラットな行データから入れ子構造の School のスライスへ変換します。
func aggregateRows(rows []infra.ListStudentsWithClassWithGradeWithSchoolRow) []SchoolRow {
	// 学校ごとに集約するための一時マップ
	schoolMap := make(map[int64]*SchoolRow)

	for _, row := range rows {
		// 学校が既に存在するかチェック。なければ新規作成。
		school, exists := schoolMap[row.SchoolID]
		if !exists {
			school = &SchoolRow{
				SchoolID:   row.SchoolID,
				SchoolName: row.SchoolName,
				Grades:     []GradeRow{},
			}
			schoolMap[row.SchoolID] = school
		}

		// 指定の GradeRow が学校内に既に存在するか検索
		gradeIndex := -1
		for i, grade := range school.Grades {
			if grade.GradeID == row.GradeID {
				gradeIndex = i
				break
			}
		}
		// 存在しなければ新たに GradeRow を追加
		if gradeIndex == -1 {
			school.Grades = append(school.Grades, GradeRow{
				GradeID:   row.GradeID,
				GradeName: row.GradeName,
				Classes:   []ClassRow{},
			})
			gradeIndex = len(school.Grades) - 1
		}

		// 同様に、指定の ClassRow が GradeRow 内に既に存在するか検索
		classIndex := -1
		for i, class := range school.Grades[gradeIndex].Classes {
			if class.ClassID == row.ClassID {
				classIndex = i
				break
			}
		}
		// 存在しなければ新たに ClassRow を追加
		if classIndex == -1 {
			school.Grades[gradeIndex].Classes = append(school.Grades[gradeIndex].Classes, ClassRow{
				ClassID:   row.ClassID,
				ClassName: row.ClassName,
				Students:  []StudentRow{},
			})
			classIndex = len(school.Grades[gradeIndex].Classes) - 1
		}

		// 最後に StudentRow を ClassRow に追加
		student := StudentRow{
			StudentID: row.StudentID,
			Name:      row.Name,
		}
		school.Grades[gradeIndex].Classes[classIndex].Students = append(
			school.Grades[gradeIndex].Classes[classIndex].Students,
			student,
		)
	}

	// マップから最終結果のスライスに変換
	result := make([]SchoolRow, 0, len(schoolMap))
	for _, s := range schoolMap {
		result = append(result, *s)
	}
	return result
}

func extract[T any, R any](items []T, selector func(T) R) []R {
	result := make([]R, len(items))
	for i, item := range items {
		result[i] = selector(item)
	}
	return result
}
