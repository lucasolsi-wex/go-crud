package service

func (ud *userDomainService) ExistsByFirstNameAndLastName(firstName, lastName string) bool {
	return ud.repository.ExistsByFirstNameAndLastName(firstName, lastName)
}
