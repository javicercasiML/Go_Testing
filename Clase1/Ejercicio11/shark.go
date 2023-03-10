package hunt

import "fmt"

type Shark struct {
	hungry bool
	tired  bool
	speed  int
}

type Prey struct {
	name  string
	speed int
}

func (s *Shark) Hunt(p *Prey) error {

	if err := p.Validate(); err != nil {
		return err
	}

	if s.tired {
		return fmt.Errorf("cannot hunt, i am really tired")
	}
	if !s.hungry {
		return fmt.Errorf("cannot hunt, i am not hungry")
	}
	if p.speed >= s.speed {
		s.tired = true
		return fmt.Errorf("could not catch it")
	}

	s.hungry = false
	s.tired = true
	return nil
}

func (p *Prey) Validate() error {
	if p.speed == 0 || p.name == "" {
		return fmt.Errorf("invalid Prey")
	}
	return nil
}
