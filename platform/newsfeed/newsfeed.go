package newsfeed

//news feed list
type Item struct {
	Title string `json:"title"`
	Post  string `json:"post"`
}

type Repo struct {
	//slice of item
	Items []Item
}

//return pointer to Repo
func New() *Repo {
	return &Repo{
		Items: []Item{},
	}
}

//Add item
func (r *Repo) Add(item Item) {
	r.Items = append(r.Items, item)
}

func (r *Repo) GetAll() []Item {
	return r.Items
}
