package models

type Comment struct {
	Model
	Key     string `gorm:"unique_index;notnull" json:"key"`
	NoteKey string `json:"note_key"`
	User    User   `json:"user"`
	UserId  int    `json:"user_id"`
	Content string `json:"content"`
	Praise  int    `gorm:"default:0" json:"praise"`
}

func SaveComment(comment *Comment) error {
	return db.Save(comment).Error
}
func QuerryByNotekey(notekey string) (comments []*Comment, e error) {
	return comments, db.Preload("User").Where("note_key = ?", notekey).Order("updated_at desc").Find(&comments).Error
}
func QuerryCountByNoteKey(notekey string) (count int, err error) {
	return count, db.Model(&Comment{}).Where("note_key=?", notekey).Count(&count).Error
}
func QuerrByPageCommentByNoteKey(page, limit int) (comment []*Comment, err error) {
	return comment, db.Preload("User").Where("note_key=?", "").Offset((page - 1) * limit).Limit(limit).Order("updated_at desc").Find(&comment).Error
}
