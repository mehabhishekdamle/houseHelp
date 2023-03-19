package app

import (
	//"github.com/rsj-rishabh/urbanClapClone/server/app/model"
)

//DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func (a *App) DBMigrate() {
	// Drop the table if it exists
	/*a.DB.AutoMigrate().DropTable(&model.User{})
	a.DB.AutoMigrate().DropTable(&model.Service{})
	a.DB.AutoMigrate().DropTable(&model.Booking{})
	a.DB.AutoMigrate().DropTable(&model.CityServiceMapping{})

	// Migrate the schema
	a.DB.AutoMigrate(&model.User{}, &model.Service{}, &model.Booking{}, &model.CityServiceMapping{})

	// Create users table
	a.DB.Create(&model.User{
		Id:       1,
		Name:     "Abhishek Damle",
		Username: "abhishek",
		Password: "abhishek10",
		Email:    "abhishekdamle@gmail.com",
		Gender:   "M",
	})
	a.DB.Create(&model.User{
		Id:       2,
		Name:     "Siddharth Dubey",
		Username: "siddharth",
		Password: "siddharth11",
		Email:    "siddharthdubey@gmail.com",
		Gender:   "M",
	})
	a.DB.Create(&model.User{
		Id:       3,
		Name:     "Tarun Surname",
		Username: "tarun",
		Password: "tarun12",
		Email:    "tarunsurname@gmail.com",
		Gender:   "M",
	})

	// Create services table
	a.DB.Create(&model.Service{
		Id:          1,
		Name:        "Home Cook",
		Description: "Welcome to our home cook services, where we provide a unique combination of cooking services to simplify your life. Our team of experienced chefs will keep your fridge stocked with delicious meals.",
		Category:    "Household",
		ImageName:   "Cook_Image.png",
		Price:       200,
	})
	a.DB.Create(&model.Service{
		Id:          2,
		Name:        "Cleaning Service",
		Description: "We provide top-notch cleaning services for both residential and commercial spaces. Our team of professional cleaners is dedicated to ensuring a comfortable and healthy environment for you.",
		Category:    "Household",
		ImageName:   "Cleaning_Service.png",
		Price:       350,
	})
	a.DB.Create(&model.Service{
		Id:          3,
		Name:        "Gardner",
		Description: "Welcome to our gardening service,we provide exceptional gardening services to enhance the beauty and functionality of your outdoor spaces.",
		Category:    "Household",
		ImageName:   "Gardner.png",
		Price:       500,
	})
	// Create booking table
	a.DB.Create((&model.Booking{
		UserId:      1,
		ServiceId:   1,
		Date:        "2022-02-15",
		StartTime:   "12:30",
		EndTime:     "13:30",
		IsCancelled: false,
	}))
	a.DB.Create((&model.Booking{
		UserId:      1,
		ServiceId:   2,
		Date:        "2022-02-15",
		StartTime:   "16:30",
		EndTime:     "17:30",
		IsCancelled: false,
	}))
	a.DB.Create((&model.Booking{
		UserId:      2,
		ServiceId:   3,
		Date:        "2022-02-15",
		StartTime:   "16:30",
		EndTime:     "17:30",
		IsCancelled: false,
	}))

	// Create CityServiceMapping table
	a.DB.Create((&model.CityServiceMapping{
		CityName:  "Newyork",
		ServiceId: 3,
	}))
	a.DB.Create((&model.CityServiceMapping{
		CityName:  "Newyork",
		ServiceId: 2,
	}))
	a.DB.Create((&model.CityServiceMapping{
		CityName:  "LA",
		ServiceId: 2,
	}))
	a.DB.Create((&model.CityServiceMapping{
		CityName:  "LA",
		ServiceId: 3,
	}))
	a.DB.Create((&model.CityServiceMapping{
		CityName:  "Boston",
		ServiceId: 1,
	}))
	a.DB.Create((&model.CityServiceMapping{
		CityName:  "Boston",
		ServiceId: 2,
	}))*/
}
