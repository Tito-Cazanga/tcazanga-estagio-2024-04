package domain

type VeterinaryRepository struct{
	
}

type NewVeterinaryRepository struct{
	paciente map[string]int
}

func (value NewVeterinaryRepository) Size() int{
	return len(value.paciente)
}

func  PatientHospitalization() *NewVeterinaryRepository{
	return &NewVeterinaryRepository{}

}