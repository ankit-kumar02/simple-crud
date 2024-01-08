package request

type CreateBureauRequest struct {
	UserID                int    `validate:"required,min=10,max=20" json:"user_id"`
	BureauReferenceNumber string `validate:"required,min=10,max=20" json:"bureau_reference_number"`
	UserReferenceNumber   string `validate:"required,min=10,max=20" json:"user_reference_number"`
	DiffNumM2M8           int    `validate:"required,min=10,max=20" json:"diff_num_m2m8"`
	CountPlLoan           int    `validate:"required,min=10,max=20" json:"count_pl_loan"`
	NewBureauVariable     string `validate:"required,min=10,max=20" json:"new_bureau_variable"`
}

type UpdateBureauRequest struct {
	UserID                int    `validate:"required"`
	BureauReferenceNumber string `validate:"required,max=200,min=1" json:"bureau_reference_number"`
}
