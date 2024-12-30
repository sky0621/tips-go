package crud

type crud struct {
	value string
}

var (
	Create = crud{"create"}
	Read   = crud{"read"}
	Update = crud{"update"}
	Delete = crud{"delete"}
	Other  = crud{"other"}
)
