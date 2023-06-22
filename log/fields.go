package log

type Fields map[string]interface{}

func (f Fields) Merge(other Fields){
	for k, v := range other{
		f[k] = v
	}
}