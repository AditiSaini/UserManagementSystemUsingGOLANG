package main

type profile struct {
	username       string
	nickname       string
	profilePicture string
	password       string
}

func getProfile(username string) *profile {
	return &profile{
		username:       username,
		nickname:       "Adris",
		profilePicture: "Dummy_picture",
		password:       "Dummy_password",
	}
}

func updateProfile(username string, nickname string,
	profilePicture string, password string) *profile {
	return &profile{
		username:       username,
		nickname:       nickname,
		profilePicture: profilePicture,
		password:       password,
	}
}
