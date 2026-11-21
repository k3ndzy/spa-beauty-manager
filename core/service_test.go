package spa

import "testing"

func TestSpaManager_CheckInSession(t *testing.T) {
	mgr := NewSpaManager()
	c, err := mgr.CheckInSession("0909123456")
	if err != nil {
		t.Fatalf("lỗi check-in: %v", err)
	}

	if c.Sessions != 4 {
		t.Errorf("kỳ vọng còn 4 buổi nhưng nhận %d", c.Sessions)
	}
	if c.Points != 400 {
		t.Errorf("kỳ vọng 400 điểm nhưng nhận %d", c.Points)
	}
}
