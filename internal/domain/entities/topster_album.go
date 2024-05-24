package entities

type TopsterAlbum struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	TopsterID uint   `gorm:"index"`
	AlbumID   uint   `gorm:"index"`
	Position  string `gorm:"type:json"`
}
