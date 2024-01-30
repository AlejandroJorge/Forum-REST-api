package logging

import "log"

func LogDomainError(err error) {
	msg := `
	[DOMAIN ERROR] %s
	`

	log.Printf(msg, err)
}

func LogUnexpectedDomainError(err error) {
	msg := `
	[UNEXPECTED DOMAIN ERROR] %s
	`

	log.Printf(msg, err)
}

func LogRepositoryError(err error) {
	msg := `
	[REPOSITORY ERROR] %s
	`

	log.Printf(msg, err)
}

func LogUnexpectedRepositoryError(err error) {
	msg := `
	[UNEXPECTED REPOSITORY ERROR] %s
	`

	log.Printf(msg, err)
}
