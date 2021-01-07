package customer

type RegisterCustomerInput struct {
	Name string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	Occupation string `json:"occupation" binding: "required"`
}