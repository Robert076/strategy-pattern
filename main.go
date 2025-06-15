package main

import "fmt"

type ILanguage interface {
	sayHi(name string)
}

type English struct {
}

func (e *English) sayHi(name string) {
	fmt.Printf("Hi in English, %s", name)
}

type Romanian struct {
}

func (r *Romanian) sayHi(name string) {
	fmt.Printf("Salut in Romana, %s", name)
}

type User struct {
	name             string
	preferedLanguage ILanguage
}

func (u *User) setPreferedLanguage(language ILanguage) {
	u.preferedLanguage = language
}

func (u *User) setName(name string) {
	u.name = name
}

func main() {
	u := &User{}
	u.setPreferedLanguage(&English{})
	u.setName("Robert")

	u.preferedLanguage.sayHi(u.name)

	/* if we didnt have those, the code would look something like this(pseudocode):

	u := &User{}
	u.setPreferedLanguage("English")
	u.setName("Robert")

	if u.preferedLanguage == "English" {
		sayHiInEnglish(u.name)
	}
	if u.preferedLanguage == "Romanian" {
		sayHiInRomanian(u.name)
	}

	...

	and any new language we have we would add to this never ending if

	If we use the strategy pattern as above however, we are using the open/closed principle

	and so our code is closed for modification but open for extension

	we never have to modify existing code, just define the new strategy and it will work flawlessly

	this might just be my favorite pattern

	*/
}
