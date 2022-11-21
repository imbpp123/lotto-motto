package rows_loader

type RowDataLoader interface {
	Load(filename string) ([][]string, error)
}
