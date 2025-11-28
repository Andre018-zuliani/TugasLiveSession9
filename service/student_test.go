package service

import (
	"session-9/model"
	"session-9/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

// InMemoryRepo adalah implementasi sederhana dari StudentRepositoryInterface
// yang menyimpan data dalam memori (slice). Ini menggantikan Mock yang kompleks.
type InMemoryRepo struct {
	Students []model.Student
	ErrGet   error // Simulasi error saat GetAll dipanggil
	ErrSave  error // Simulasi error saat SaveAll dipanggil
}

func (r *InMemoryRepo) GetAll() ([]model.Student, error) {
	return r.Students, r.ErrGet
}

func (r *InMemoryRepo) SaveAll(students []model.Student) error {
	if r.ErrSave != nil {
		return r.ErrSave
	}
	// Menyimpan data yang diperbarui ke memori
	r.Students = students
	return nil
}

// newTestService adalah helper function untuk membuat service baru dengan in-memory repo.
func newTestService(initial []model.Student) (*StudentService, *InMemoryRepo) {
	repo := &InMemoryRepo{Students: initial}
	svc := NewStudentService(repo)
	return svc, repo
}

func TestStudentService_GetAll_Success(t *testing.T) {
	initial := []model.Student{
		{ID: 1, Name: "Andi", Age: 21},
	}
	svc, _ := newTestService(initial)

	students, err := svc.GetAll()

	assert.Nil(t, err)
	assert.Equal(t, len(initial), len(students))
	assert.Equal(t, initial[0].Name, students[0].Name)
}

func TestStudentService_GetAll_Error(t *testing.T) {
	svc, repo := newTestService([]model.Student{})
	// Set error untuk simulasi kegagalan repository
	repo.ErrGet = utils.ErrFile

	_, err := svc.GetAll()

	assert.Error(t, err)
	assert.Equal(t, utils.ErrFile, err)
}

func TestStudentService_GetByID_Found(t *testing.T) {
	initial := []model.Student{
		{ID: 1, Name: "Andi", Age: 21},
		{ID: 2, Name: "Siti", Age: 22},
	}
	svc, _ := newTestService(initial)

	st, err := svc.GetByID(2)

	assert.Nil(t, err)
	assert.Equal(t, "Siti", st.Name)
	assert.Equal(t, 2, st.ID)
}

func TestStudentService_GetByID_NotFound(t *testing.T) {
	initial := []model.Student{
		{ID: 1, Name: "Andi", Age: 21},
	}
	svc, _ := newTestService(initial)

	_, err := svc.GetByID(999)

	assert.Error(t, err)
	assert.Equal(t, utils.ErrNotFound, err)
}

func TestStudentService_GetByID_GetAllError(t *testing.T) {
	svc, repo := newTestService([]model.Student{})
	// Set error untuk simulasi kegagalan repository
	repo.ErrGet = utils.ErrFile

	_, err := svc.GetByID(1)

	assert.Error(t, err)
	assert.Equal(t, utils.ErrFile, err)
}

func TestStudentService_Create_Success(t *testing.T) {
	initial := []model.Student{
		{ID: 10, Name: "MaxID", Age: 30}, // ID awal tertinggi = 10
	}
	newStudent := model.Student{Name: "Rudi", Age: 20}
	expectedID := 11

	svc, repo := newTestService(initial)

	created, err := svc.Create(newStudent)

	assert.Nil(t, err)
	assert.Equal(t, expectedID, created.ID)
	assert.Equal(t, "Rudi", created.Name)
	// Pastikan data tersimpan di repo
	assert.Equal(t, 2, len(repo.Students))
}

func TestStudentService_Create_GetAllError(t *testing.T) {
	svc, repo := newTestService([]model.Student{})
	repo.ErrGet = utils.ErrFile

	_, err := svc.Create(model.Student{Name: "Test", Age: 10})

	assert.Error(t, err)
	assert.Equal(t, utils.ErrFile, err)
}

func TestStudentService_Create_SaveAllError(t *testing.T) {
	initial := []model.Student{}
	svc, repo := newTestService(initial)
	repo.ErrSave = utils.ErrFile

	_, err := svc.Create(model.Student{Name: "Test", Age: 10})

	assert.Error(t, err)
	assert.Equal(t, utils.ErrFile, err)
}

func TestStudentService_Update_Success(t *testing.T) {
	initial := []model.Student{
		{ID: 1, Name: "Andi", Age: 21},
	}
	updateInput := model.Student{Name: "Andi Updated", Age: 23}
	studentID := 1

	svc, repo := newTestService(initial)

	updated, err := svc.Update(studentID, updateInput)

	assert.Nil(t, err)
	assert.Equal(t, "Andi Updated", updated.Name)
	assert.Equal(t, 23, updated.Age)
	// Pastikan ID tidak berubah
	assert.Equal(t, 1, updated.ID)
	// Pastikan data di repo juga terupdate
	assert.Equal(t, "Andi Updated", repo.Students[0].Name)
}

func TestStudentService_Update_NotFound(t *testing.T) {
	initial := []model.Student{
		{ID: 1, Name: "Andi", Age: 21},
	}
	updateInput := model.Student{Name: "Missing", Age: 99}
	studentID := 999

	svc, _ := newTestService(initial)

	_, err := svc.Update(studentID, updateInput)

	assert.Error(t, err)
	assert.Equal(t, utils.ErrNotFound, err)
}

func TestStudentService_Update_GetAllError(t *testing.T) {
	svc, repo := newTestService([]model.Student{})
	repo.ErrGet = utils.ErrFile

	_, err := svc.Update(1, model.Student{})

	assert.Error(t, err)
	assert.Equal(t, utils.ErrFile, err)
}

func TestStudentService_Update_SaveAllError(t *testing.T) {
	initial := []model.Student{
		{ID: 1, Name: "Andi", Age: 21},
	}
	svc, repo := newTestService(initial)
	repo.ErrSave = utils.ErrFile

	_, err := svc.Update(1, model.Student{Name: "Andi Updated", Age: 23})

	assert.Error(t, err)
	assert.Equal(t, utils.ErrFile, err)
}

func TestStudentService_Delete_Success(t *testing.T) {
	initial := []model.Student{
		{ID: 1, Name: "Andi", Age: 21},
		{ID: 2, Name: "Siti", Age: 22},
	}
	studentID := 1 // Akan dihapus

	svc, repo := newTestService(initial)

	err := svc.Delete(studentID)

	assert.Nil(t, err)
	// Pastikan hanya tersisa 1 student di repo (Siti)
	assert.Equal(t, 1, len(repo.Students))
	assert.Equal(t, 2, repo.Students[0].ID)
}

func TestStudentService_Delete_NotFound(t *testing.T) {
	initial := []model.Student{
		{ID: 1, Name: "Andi", Age: 21},
	}
	studentID := 999 // Tidak ditemukan

	svc, _ := newTestService(initial)

	err := svc.Delete(studentID)

	assert.Error(t, err)
	assert.Equal(t, utils.ErrNotFound, err)
}

func TestStudentService_Delete_GetAllError(t *testing.T) {
	svc, repo := newTestService([]model.Student{})
	repo.ErrGet = utils.ErrFile

	err := svc.Delete(1)

	assert.Error(t, err)
	assert.Equal(t, utils.ErrFile, err)
}

func TestStudentService_Delete_SaveAllError(t *testing.T) {
	initial := []model.Student{
		{ID: 1, Name: "Andi", Age: 21},
	}
	svc, repo := newTestService(initial)
	repo.ErrSave = utils.ErrFile

	err := svc.Delete(1)

	assert.Error(t, err)
	assert.Equal(t, utils.ErrFile, err)
}