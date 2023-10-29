package tmdb

type GetMovieListResponse struct {
	Movies   []MovieList `json:"movies" validate:"dive,required"`
	PageInfo string      `json:"page_info"`
	Error    string      `json:"error"`
}

type MovieList struct {}

func GetMovieList(client *ProxyClient) (movies []MovieList, err error) {
	return movies, nil
}