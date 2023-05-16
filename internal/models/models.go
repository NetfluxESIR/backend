package models

func Models() []interface{} {
	return []interface{}{
		&Processing{},
		&ProcessingStep{},
		&Video{},
		&User{},
		&RobotAccount{},
		&Token{},
	}
}
