package hairstyles

type HairStyleService struct {
	repo *HairStyleRepository
}

func NewHairStyleService(repo *HairStyleRepository) *HairStyleService {
	return &HairStyleService{repo: repo}
}

func (s *HairStyleService) GetAllHairStyles() ([]HairStyle, error) {
	return s.repo.GetAll()
}

func (s *HairStyleService) CreateHairStyle(hairStyle *HairStyle) error {
	return s.repo.Create(hairStyle)
}

func (s *HairStyleService) GetHairStyleByID(id uint) (*HairStyle, error) {
	return s.repo.FindById(id)
}

func (s *HairStyleService) DeleteHairStyleByID(id uint) error {
	return s.repo.DeleteByID(id)
}

func (s *HairStyleService) UpdateHairStyle(hairStyle *HairStyle) error {
	return s.repo.Update(hairStyle)
}
