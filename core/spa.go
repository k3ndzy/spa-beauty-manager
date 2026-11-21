package spa

type MemberTier string
const (
	TierSilver MemberTier = "SILVER"
	TierGold   MemberTier = "GOLD"
	TierVIP    MemberTier = "VIP"
)

type Customer struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	Phone     string     `json:"phone"`
	Points    int        `json:"points"`
	Tier      MemberTier `json:"tier"`
	Sessions  int        `json:"sessions_remaining"`
}
