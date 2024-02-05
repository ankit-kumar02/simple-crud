package request

type CreateBureauRequest struct {
	UserID            int                    `validate:"required,min=10,max=20" json:"user_id"`
	DiffNumM2M8       int                    `validate:"required,min=10,max=20" json:"diff_num_m2m8"`
	CountPlLoan       int                    `validate:"required,min=10,max=20" json:"count_pl_loan"`
	NewBureauVariable map[string]interface{} `validate:"required,min=10,max=20" json:"new_bureau_variable"`
}

type UpdateBureauRequest struct {
	DiffNumM2M8           int `validate:"required,min=10,max=20" json:"diff_num_m2m8"`
	CountPlLoan           int `validate:"required,min=10,max=20" json:"count_pl_loan"`
	BureauReferenceNumber string
}
