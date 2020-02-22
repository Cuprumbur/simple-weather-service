package storage

type Detector struct {
	ID   int
	Name string
}

func (s *Storage) FindDetector(nId int) (d *Detector, err error) {

	selDB, err := s.db.Query("SELECT * FROM detector Where id=?", nId)
	if err != nil {
		return
	}

	for selDB.Next() {
		var id int
		var name string
		err = selDB.Scan(&id, &name)
		if err != nil {
			return
		}

		d = &Detector{
			ID:   id,
			Name: name,
		}
	}

	return
}

func (s *Storage) FindAllDetectors() (detectors []*Detector, err error) {
	selDB, err := s.db.Query("SELECT * FROM detector")
	if err != nil {
		return
	}
	for selDB.Next() {
		var id int
		var name string
		err = selDB.Scan(&id, &name)
		if err != nil {
			return
		}
		d := &Detector{
			ID:   id,
			Name: name,
		}

		detectors = append(detectors, d)
	}

	return
}

func (s *Storage) Update(d *Detector) (err error) {
	upd, err := s.db.Prepare("UPDATE detector SET name=? WHERE  id=?")
	if err != nil {
		return
	}

	_, err = upd.Exec(d.Name, d.ID)
	return
}

func (s *Storage) Delete(id int) (err error) {
	del, err := s.db.Prepare("DELETE FROM detector WHERE id=?")
	if err != nil {
		return
	}
	_, err = del.Exec(id)
	return
}

func (s *Storage) Store(d *Detector) (id int, err error) {
	ins, err := s.db.Prepare("INSERT Into detector(name) VALUES(?)")
	if err != nil {
		return
	}

	_, err = ins.Exec(d.Name)
	return
}
