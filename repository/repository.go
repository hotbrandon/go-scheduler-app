package repository

// 預約組數摘要
type ReservationGroupSummary struct {
	// 這筆資料的名稱
	Name string `json:"name"`
	// 統計值(日)
	AmtD int `json:"amt_d"`
	// 統計值(月)
	AmtM int `json:"amt_m"`
	// 統計值(年)
	AmtY int `json:"amt_y"`
}

type GolfRepository interface {
	GetReservationGroupSummary() ([]ReservationGroupSummary, error)
}

// 取得預約組數摘要

type golfRepository struct {
}

func NewGolfRepository() GolfRepository {
	return &golfRepository{}
}

func (r *golfRepository) GetReservationGroupSummary() ([]ReservationGroupSummary, error) {
	return []ReservationGroupSummary{}, nil
}
