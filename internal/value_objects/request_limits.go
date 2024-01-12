package value_objects

type RequestLimits struct {
	IPLimit  uint
	APILimit uint
}

func NewRequestLimit(IPLimit uint, APILimit uint) RequestLimits {
	return RequestLimits{
		IPLimit:  IPLimit,
		APILimit: APILimit,
	}
}
