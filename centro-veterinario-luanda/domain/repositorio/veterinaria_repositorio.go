package domain

type VeterinariaRepositorio interface{
	
}

type InternamentoRepositorio struct{

}

func NovoInternamentoRepositorio() VeterinariaRepositorio{
	return InternamentoRepositorio{}
}