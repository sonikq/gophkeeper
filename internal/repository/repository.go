package repository

func NewGophKeeperRepository() Repository {
	return newInMemoryRepo()
}
