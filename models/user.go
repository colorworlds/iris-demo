package models

type UserModel struct {
	ID                 int    `gorm:"column:id"`
	UserName           string `gorm:"column:user_name"`
	AuthKey            string `gorm:"column:auth_key"`
	PasswordHash       string `gorm:"column:password_hash"`
	PasswordResetToken string `gorm:"column:password_reset_token"`
	Email              string `gorm:"column:email"`
	Status             int    `gorm:"column:status"`
	CreatedAt          int    `gorm:"column:created_at"`
	UpdatedAt          int    `gorm:"column:updated_at"`
	VerificationToken  string `gorm:"column:verification_token"`
}

/*
CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `auth_key` varchar(32) COLLATE utf8_unicode_ci NOT NULL,
  `password_hash` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `password_reset_token` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `email` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `status` smallint(6) NOT NULL DEFAULT '10',
  `created_at` int(11) NOT NULL,
  `updated_at` int(11) NOT NULL,
  `verification_token` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`),
  UNIQUE KEY `email` (`email`),
  UNIQUE KEY `password_reset_token` (`password_reset_token`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
*/

func (UserModel) TableName() string {
	return "user"
}
