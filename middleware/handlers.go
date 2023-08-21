package middleware

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

// GetUser will return a single artist by its id
//func GetArtist(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
//	w.Header().Set("Access-Control-Allow-Origin", "*")
//
//	// get the userid from the request params, key is "id"
//	params := mux.Vars(r)
//
//	// convert the id type from string to int
//	id, err := strconv.Atoi(params["id"])
//
//	if err != nil {
//		log.Fatalf("Unable to convert the string into int.  %v", err)
//	}
//
//	// call the getUser function with user id to retrieve a single user
//	user, err := getUser(int64(id))
//
//	if err != nil {
//		log.Fatalf("Unable to get user. %v", err)
//	}
//
//	// send the response
//	json.NewEncoder(w).Encode(user)
//}

// GetAllArtists will return all the users
//func GetAllArtists(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
//	w.Header().Set("Access-Control-Allow-Origin", "*")
//	// get all the users in the db
//	users, err := getAllUsers()
//
//	if err != nil {
//		log.Fatalf("Unable to get all user. %v", err)
//	}
//
//	// send all the users as response
//	json.NewEncoder(w).Encode(users)
//}

//func (artist models.Artist) getAllArtists() ([]models.Artist, error) {
//	db, err := initializers.ConnectDB()
//	if err != nil {
//		return nil, err
//	} else {
//		var artists []models.Artist
//		db.Preload("Users").Find(&artists)
//		return artists, nil
//	}
//}

// CreateArtist create a artist in the postgres db
//func CreateArtist(w http.ResponseWriter, r *http.Request) {
//	// set the header to content type x-www-form-urlencoded
//	// Allow all origin to handle cors issue
//	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
//	w.Header().Set("Access-Control-Allow-Origin", "*")
//	w.Header().Set("Access-Control-Allow-Methods", "POST")
//	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
//
//	// create an empty artist of type models.Artist
//	var artist models.Artist
//
//	// decode the json request to artist
//	err := json.NewDecoder(r.Body).Decode(&artist)
//
//	if err != nil {
//		log.Fatalf("Unable to decode the request body.  %v", err)
//	}
//
//	// call insert user function and pass the user
//	insertID := insertUser(user)
//
//	// format a response object
//	res := response{
//		ID:      insertID,
//		Message: "User created successfully",
//	}
//
//	// send the response
//	json.NewEncoder(w).Encode(res)
//}
