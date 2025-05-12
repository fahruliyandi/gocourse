// single responsibility
package SOLID

import "fmt"

type UserRepository interface {
	save(nama string)
}

type EmailSender interface {
	sendEmail(to string)
}

type UserService struct {
	repo  UserRepository
	email EmailSender
}

func newUserService(user UserRepository, email EmailSender) *UserService {
	us := &UserService{repo: user, email: email}
	return us
}

func (us *UserService) register(name string) {
	us.repo.save(name)
	us.email.sendEmail(name)
}

type fileRepo struct{}

func (u *fileRepo) save(nama string) {
	fmt.Printf("User %s berhasil register\n", nama)
}

type gmail struct{}

func (g *gmail) sendEmail(nama string) {
	fmt.Printf("Send dari gmail to %s\n", nama)
}

func main() {

	usr := &fileRepo{}
	eml := &gmail{}
	userService := newUserService(usr, eml)
	userService.register("Fahrul")

}
