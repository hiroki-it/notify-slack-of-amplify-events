package file

type File struct {
	path *Path
}

// NewFile コンストラクタ
func NewFile(path *Path) *File {

	return &File{
		path: path,
	}
}

// FilePath Pathを返却します．
func (f *File) FilePath() *Path {
	return f.path
}
