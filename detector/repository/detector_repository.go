package detector

import (
	"database/sql"

	"github.com/Cuprumbur/weather-service/detector"
	"github.com/Cuprumbur/weather-service/model"
)

type mySqlDetectorRepository struct {
	db *sql.DB
}

func NewMySqlDetectorRepository(db *sql.DB) detector.Repository {
	return &mySqlDetectorRepository{db}
}

func (r *mySqlDetectorRepository) FindDetector(ID int) (d *model.Detector, err error) {

	selDB, err := r.db.Query("SELECT * FROM detector Where id=?", ID)
	if err != nil {
		return
	}

	for selDB.Next() {
		var id int
		var name, serial string
		err = selDB.Scan(&id, &name, &serial)
		if err != nil {
			return
		}

		d = &model.Detector{
			ID:     id,
			Name:   name,
			Serial: serial,
		}
	}

	return
}

func (r *mySqlDetectorRepository) FindAllDetectors() (detectors []*model.Detector, err error) {
	selDB, err := r.db.Query("SELECT * FROM detector")
	if err != nil {
		return
	}

	for selDB.Next() {
		var id int
		var name, serial string
		err = selDB.Scan(&id, &name, &serial)
		if err != nil {
			return
		}

		d := &model.Detector{
			ID:     id,
			Name:   name,
			Serial: serial,
		}

		detectors = append(detectors, d)
	}

	return
}

func (r *mySqlDetectorRepository) Update(d *model.Detector) (err error) {
	upd, err := r.db.Prepare("UPDATE detector SET name=?, serial=? WHERE  id=?")
	if err != nil {
		return
	}

	_, err = upd.Exec(d.Name, d.Serial, d.ID)
	return
}

func (r *mySqlDetectorRepository) Delete(id int) (err error) {
	del, err := r.db.Prepare("DELETE FROM detector WHERE id=?")
	if err != nil {
		return
	}
	_, err = del.Exec(id)

	return
}

func (r *mySqlDetectorRepository) Store(d *model.Detector) (id int64, err error) {
	ins, err := r.db.Prepare("INSERT Into detector(name, serial) VALUES(?, ?)")
	if err != nil {
		return
	}

	res, err := ins.Exec(d.Name, d.Serial)
	if err != nil {
		return
	}

	id, err = res.LastInsertId()

	return
}
