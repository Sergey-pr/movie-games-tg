package serializers

type JwtToken struct {
	Token   string `json:"token"`
	ExpTime int64  `json:"exp_time"`
}
