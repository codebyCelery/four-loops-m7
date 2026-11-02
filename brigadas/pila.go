package brigadas

import "errors"

type Pila struct {
	elementos []string
}

func NuevaPila() *Pila {
	return &Pila{
		elementos: []string{},
	}
}

func (p *Pila) Apilar(v string) {
	p.elementos = append(p.elementos, v)
}

func (p *Pila) Desapilar() (string, error) {
	if p.Len() == 0 {
		return "", errors.New("la pila está vacía")
	}
	ultimoIndex := len(p.elementos) - 1
	valor := p.elementos[ultimoIndex]

	p.elementos = p.elementos[:ultimoIndex]
	return valor, nil
}

func (p *Pila) Tope() (string, error) {
	if p.Len() == 0 {
		return "", errors.New("la pila está vacía")
	}
	return p.elementos[len(p.elementos)-1], nil
}

func (p *Pila) Len() int {
	return len(p.elementos)
}
