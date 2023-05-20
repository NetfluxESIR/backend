package models

func Models() []interface{} {
	return []interface{}{
		&Video{},
		&Account{},
		&Token{},
		&ProcessingStep{},
		&Processing{},
	}
}
