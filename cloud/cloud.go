package cloud

type CloudDb struct {
	url string
}

func NewJsonDb(name string) *CloudDb {
	return &CloudDb{
		url: name,
	}
}

func (db *CloudDb) Write(content []byte) {

}

func (db *CloudDb) Read() ([]byte, error) {
	return []byte{}, nil
}
