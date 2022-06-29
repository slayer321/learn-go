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

type ProfileMaker interface {
	UpdateProfilePicture(ProfilePicture)
	CheckDuplicateProfile(Profile) bool
}

func (c *ProfilePicture) UpdateProfilePicture(cProfile ProfilePicture) {

	c.ImageName = cProfile.ImageName
	c.ImagePath = cProfile.ImagePath

}

func CheckDuplicateProfile(dProfile Profile) bool {

	if "slayer321" == dProfile.Username {
		return true
	}
	return false

}

func main() {

	sachin := Profile{
		Name:          "sachin",
		Username:      "slayer321",
		Designation:   "Associate Product Engineer",
		ContactNumber: "7666454860",
		MyProfilePic: ProfilePicture{
			ImageName: "sachin_maurya.jpg",
			ImagePath: "https://imugar.io/sachin_maurya.jpg",
		},
	}

	//fmt.Printf("Old profile : %+v \n", sachin)

	// Updating Profile Picture

	sachin.MyProfilePic.ImageName = "new_sachin_maurya.jpg"
	sachin.MyProfilePic.ImagePath = "https://imugar.io/new_sachin_maurya.jpg"

	sachinNewProfilePic := ProfilePicture{
		ImageName: sachin.MyProfilePic.ImageName,
		ImagePath: sachin.MyProfilePic.ImagePath,
	}

	sachinNewProfilePic.UpdateProfilePicture(sachinNewProfilePic)

	//fmt.Printf("New profile pic: %+v", sachin)

	// Check Duplicate Profile

	chandu := Profile{
		Name:          "chandu",
		Username:      "robinhood",
		Designation:   "Data scientist",
		ContactNumber: "7666454860",
		MyProfilePic: ProfilePicture{
			ImageName: "chandu.jpg",
			ImagePath: "https://imugar.io/chandu.jpg",
		},
	}

	val := CheckDuplicateProfile(chandu)

	if val == true {
		fmt.Printf("%s, is the Duplicate profile\n", chandu.Name)
	} else {
		fmt.Printf("%s, is the New Profile\n", chandu.Name)
	}

}
