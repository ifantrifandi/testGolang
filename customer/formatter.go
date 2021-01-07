package customer

type CustomerFormatter struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Occupation string `json:"occupation"`
	Role string `json:"role"`
	Token string `json:"token"`
}

type DataCustomer  struct {
	Cust []CustomerDataFormatter `json:"cust"`
}
type CustomerDataFormatter struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Occupation string `json:"occupation"`
	Role string `json:"role"`
}
func FormatCustomer (customer Customer, token string) CustomerFormatter {
	formatter := CustomerFormatter{
		ID : customer.ID,
		Name : customer.Name,
		Occupation : customer.Occupation,
		Role : customer.Role,
		Token : token,
	}

	return  formatter
}

func DataFormatCustomer (customers []Customer) []CustomerDataFormatter {
	var newCustomer []CustomerDataFormatter

	for _, value := range customers{
		newFormat := CustomerDataFormatter{
			ID : value.ID,
			Name : value.Name,
			Occupation : value.Occupation,
			Role : value.Role,
		}

		newCustomer = append(newCustomer, newFormat)
	}
	return newCustomer
}