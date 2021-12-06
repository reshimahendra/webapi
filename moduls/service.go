package moduls

import "time"


type Service interface {
    Create(bookRequest BookRequest) (Book, error)
    Update(ID int, bookRequest BookRequest) (Book, error)
    First() (Book, error)
    Last() (Book, error)
    FindByID(ID int) (Book, error)
    Delete(ID int) error
    All() ([]Book, error)
    // Filter(Rating int) ([]Book, error)
}

type service struct {
    repo Repo
}

func NewService(repo Repo) *service {
    return &service{repo}
}

func (s *service) All() ([]Book, error) {
    return s.repo.All() 
}

func (s *service) First() (Book, error) {
    return s.repo.First() 
}

func (s *service) Last() (Book, error) {
    return s.repo.Last()
}

func (s *service) FindByID(ID int) (Book, error){
    return s.repo.FindByID(ID)
}

func (s *service) Delete(ID int) error{
    return s.repo.Delete(ID)
}

func (s *service) Create(bookRequest BookRequest) (Book, error) {
    price, _ := bookRequest.Price.Int64()
    rating, _ := bookRequest.Rating.Int64()
    book:= Book{
        Title:       bookRequest.Title,
        Description: bookRequest.Description,
        Price:       int(price),
        Rating:      int(rating),
    }

    return s.repo.Create(book)
}

func (s *service) Update(id int,bookRequest BookRequest) (Book, error) {
    price, _ := bookRequest.Price.Int64()
    rating, _ := bookRequest.Rating.Int64()
    book := Book{
        Title:       bookRequest.Title,
        Description: bookRequest.Description,
        Price:       int(price),
        Rating:      int(rating),
        UpdatedAt:   time.Now(),
    }
    b, err := s.repo.Update(id, book)
    return b, err
}

// func DB_Conn() {
//     // Connect to SQLITE DB server
//     var err error
//     DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
//     if err != nil {
//         panic("Failed to connect to DB")
//     }
// }
