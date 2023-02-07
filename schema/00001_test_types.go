package schema

type users struct {
	id           int
	email        string
	username     string
	is_activated bool
	password     string
}

type tokens struct {
	token_id     int
	refreshToken string
}

type user_api struct {
	id         int
	owner_id   int
	api_name   string
	api_key    string
	createdAt  int64
	expiryTime int64
	bearer     string
}
