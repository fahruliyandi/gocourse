//dependecy inversion
//harus bergantung ke interface(abstrak)

package SOLID

import "fmt"

type UserRepository interface {
	save()
}

type postgreRepo struct{}

func (p *postgreRepo) save() {
	fmt.Println("Save ke postgre")
}

type fileRepo struct{}

func (f *fileRepo) save() {
	fmt.Println("Save ke file")
}

type UserService struct {
	repo UserRepository
}

func (us *UserService) register() {
	us.repo.save()
}

func main() {
	pg := &postgreRepo{}
	us := &UserService{
		repo: pg,
	}
	us.register()

	us.repo = &fileRepo{}
	us.register()
}
