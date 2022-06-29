package main

import "fmt"

type ProfilePicture struct {
	ImageName string
	ImagePath string
}

type Profile struct {
	Name          string
	Username      string
	Designation   string
	ContactNumber string
	MyProfilePic  ProfilePicture
}

func (p *Profile) UpdateProfile(UpdateField Profile) *Profile {

	p.Name = UpdateField.Name
	p.Username = UpdateField.Username
	p.Designation = UpdateField.Designation
	p.ContactNumber = UpdateField.ContactNumber
	p.MyProfilePic.ImageName = UpdateField.MyProfilePic.ImageName
	p.MyProfilePic.ImagePath = UpdateField.MyProfilePic.ImagePath
	return p
}

func main() {

	sachin := Profile{
		Name:          "sachin",
		Username:      "slayer321",
		Designation:   "Associate Product Engineer",
		ContactNumber: "7666454860",
		MyProfilePic: ProfilePicture{
			ImageName: "sachin_maurya.jpg",
			ImagePath: "https://imugar.io",
		},
	}

	fmt.Printf("%+v \n", sachin)

	updateSachin := Profile{
		Name:        "Sachin Maurya",
		Designation: "Sr Product Engineer",
	}

	sachin.UpdateProfile(updateSachin)

	fmt.Printf("Updated: %+v \n", sachin)
}
