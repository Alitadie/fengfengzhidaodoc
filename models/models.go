package models

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique"`
	Password string
}

type Book struct {
	ID       uint `gorm:"primaryKey"`
	Title    string
	Chapters []Chapter
}

type Chapter struct {
	ID        uint `gorm:"primaryKey"`
	BookID    uint
	Title     string
	Documents []Document
}

type Document struct {
	ID        uint `gorm:"primaryKey"`
	ChapterID uint
	Title     string
	Content   string `gorm:"type:text"`
}
