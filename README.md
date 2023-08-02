# moviesBacked
**Used ::** Golang , GorilaMUX , GORM , Postgres(database)

## //Databse Models

    
    type Movie struct {
    	gorm.Model
    	ID          string    `gorm:"primaryKey"`
    	ReleaseDate string    `json:"release_date" validate:"required"`
    	Title       string    `json:"title"`
    	IsWatched    bool     `json:"is_seen"  validate:"required"`
    	Director    Director  `json:"director" gorm:"foreignKey:ID"`
    	Created_at  time.Time `json:"created_at"`
    	Updated_at  time.Time `json:"updated_at"`
    }


    type Director struct {
        ID        string
        FirstName string `json:"firstname"`
        LastName  string `json:"lastname"`
      }


//routes 
<img width="1440" alt="Screenshot 2023-08-02 at 4 16 01 PM" src="https://github.com/siddharthsingh025/moviesBacked/assets/87073574/68e04e65-2829-414c-ab2c-080222a6dcb5">

<img width="1440" alt="Screenshot 2023-08-02 at 4 16 44 PM 1" src="https://github.com/siddharthsingh025/moviesBacked/assets/87073574/966e602a-8d9e-47f6-bcb5-64ab736a8712">



<img width="1001" alt="Screenshot 2023-08-02 at 2 16 47 PM" src="https://github.com/siddharthsingh025/moviesBacked/assets/87073574/60bbb686-e0d2-41f2-b909-8ddf12d28ef4">
<img width="1440" alt="Screenshot 2023-08-02 at 2 17 37 PM" src="https://github.com/siddharthsingh025/moviesBacked/assets/87073574/3352f460-98a1-40ee-9a33-b5231e8bb954">
<img width="1440" alt="Screenshot 2023-08-02 at 2 17 45 PM" src="https://github.com/siddharthsingh025/moviesBacked/assets/87073574/8ad62145-4570-4346-a249-15a19fb87269">
