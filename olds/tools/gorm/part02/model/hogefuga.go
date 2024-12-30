package model

type HogeFuga struct {
	HogeID             int    `gorm:"primary_key"`           // プライマリーキーの指示
	Name               string `gorm:"column:hoge_fuga_name"` // プロパティとは別の名前でカラム定義
	Field              string `gorm:"type:varchar(100)"`     // 型・桁を指示
	UniqueField        string `gorm:"unique"`                // 重複不可を指示
	NotNullField       string `gorm:"not null"`              // Not Nullを指示
	IgnoreField        string `gorm:"-"`                     // カラム化しない指示
	AutoIncField       int    `gorm:"AUTO_INCREMENT"`        // オートインクリメント指示
	IndexField         int    `gorm:"index"`                 // インデックスを貼る
	UniqueIndexField   int    `gorm:"unique_index"`          // ユニークインデックスを貼る
	NamedIndexField    int    `gorm:"index:other_name_idx"`  // インデックスの名前を指示
	NoInstructionField string // 指定なし
}
